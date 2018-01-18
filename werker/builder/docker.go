package builder

import (
	ocelog "bitbucket.org/level11consulting/go-til/log"
	pb "bitbucket.org/level11consulting/ocelot/protos"
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/client"
	"github.com/mitchellh/go-homedir"
	"io"
	//"os/exec"
	"bitbucket.org/level11consulting/ocelot/util/cred"
	"strings"
)

type Docker struct{
	Log	io.ReadCloser
	ContainerId	string
	DockerClient *client.Client
	RemoteConfig cred.CVRemoteConfig
	*Basher
}

func NewDockerBuilder(b *Basher) Builder {
	remoteCred, _ := cred.New()
	return &Docker{nil, "", nil, remoteCred, b}
}

func (d *Docker) Setup(logout chan []byte, werk *pb.WerkerTask) *Result {
	//TODO: do some sort of util for the stage name + formatting
	stage := "setup"
	stagePrintln := "SETUP | "

	logout <- []byte(stagePrintln + "Setting up...")

	ctx := context.Background()
	cli, err := client.NewEnvClient()
	d.DockerClient = cli

	if err != nil {
		return &Result{
			Stage:  stage,
			Status: FAIL,
			Error:  err,
		}
	}

	imageName := werk.BuildConf.Image

	out, err := cli.ImagePull(ctx, imageName, types.ImagePullOptions{})
	if err != nil {
		ocelog.IncludeErrField(err).Error("couldn't pull image!")
		return &Result{
			Stage:  stage,
			Status: FAIL,
			Error:  err,
		}
	}
	var byt []byte
	buf := bufio.NewReader(out)
	buf.Read(byt)
	fmt.Println(string(byt))
	defer out.Close()

	bufReader := bufio.NewReader(out)
	d.writeToInfo(stagePrintln, bufReader, logout)

	logout <- []byte(stagePrintln + "Creating container...")


	//container configurations
	containerConfig := &container.Config{

		Image: imageName,
		Env: werk.BuildConf.Env,
		Cmd: d.DownloadCodebase(werk),
		AttachStderr: true,
		AttachStdout: true,
		AttachStdin:true,
		Tty:true,
	}

	homeDirectory, _ := homedir.Expand("~/.ocelot")
	//host configs like mount points
	hostConfig := &container.HostConfig{
		//TODO: have it be overridable via env variable
		Binds: []string{ homeDirectory + ":/.ocelot"},
		NetworkMode: "host",
	}

	resp, err := cli.ContainerCreate(ctx, containerConfig , hostConfig, nil, "")

	if err != nil {
		return &Result{
			Stage:  stage,
			Status: FAIL,
			Error:  err,
		}
	}

	for _, warning := range resp.Warnings {
		logout <- []byte(warning)
	}

	logout <- []byte(stagePrintln + "Container created with ID " + resp.ID)

	d.ContainerId = resp.ID

	if err := cli.ContainerStart(ctx, resp.ID, types.ContainerStartOptions{}); err != nil {
		return &Result{
			Stage:  stage,
			Status: FAIL,
			Error:  err,
		}
	}

	logout <- []byte(stagePrintln + "Container " + resp.ID + " started")

	//since container is created in setup, log tailing via container is also kicked off in setup
	containerLog, err := cli.ContainerLogs(ctx, resp.ID, types.ContainerLogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow: true,
	})

	if err != nil {
		return &Result{
			Stage: stage,
			Status: FAIL,
			Error:  err,
		}
	}

	d.Log = containerLog
	bufReader = bufio.NewReader(containerLog)
	d.writeToInfo(stagePrintln, bufReader, logout)

	return &Result{
		Stage:  stage,
		Status: PASS,
		Error:  nil,
	}
}

func (d *Docker) Cleanup() {
	//TODO: review, should we be creating new contexts for every stage?
	cleanupCtx := context.Background()

	d.Log.Close()
	if err := d.DockerClient.ContainerKill(cleanupCtx, d.ContainerId, "SIGKILL"); err != nil {
		ocelog.IncludeErrField(err).WithField("containerId", d.ContainerId).Error("couldn't kill")
	} else {
		if err := d.DockerClient.ContainerRemove(cleanupCtx, d.ContainerId, types.ContainerRemoveOptions{}); err != nil {
			ocelog.IncludeErrField(err).WithField("containerId", d.ContainerId).Error("couldn't rm")
		}
	}

	d.DockerClient.ContainerRemove(cleanupCtx, d.ContainerId, types.ContainerRemoveOptions{})
	d.DockerClient.Close()
}

func (d *Docker) Build(logout chan []byte, stage *pb.Stage, commitHash string) *Result {
	currStage := "build"
	currStageStr := "BUILD | "

	logout <- []byte(currStageStr + "Building...")

	if len(d.ContainerId) == 0 {
		return &Result {
			Stage: currStage,
			Status: FAIL,
			Error: errors.New("no container exists, setup before executing"),
		}
	}

	return d.Exec(currStage, currStageStr, stage.Env, d.BuildScript(stage.Script, commitHash), logout)
}

//uses the repo creds from admin to store artifact - keyed by acctname
func (d *Docker) SaveArtifact(logout chan []byte, task *pb.WerkerTask, commitHash string) *Result {
	currStage := "SaveArtifact"
	currStageStr := "SAVE_ARTIFACT | "

	logout <- []byte(currStageStr + "Saving artifact to ...")

	if len(d.ContainerId) == 0 {
		return &Result {
			Stage: currStage,
			Status: FAIL,
			Error: errors.New("no container exists, setup before executing"),
		}
	}

	//check if build tool if set to maven (cause that's the only thing that we use to push to nexus right now)
	if strings.Compare(task.BuildConf.BuildTool, "maven") != 0 {
		logout <- []byte(fmt.Sprintf(currStageStr + "build tool %s not part of accepted values: %s...", task.BuildConf.BuildTool, "maven"))
		return &Result {
			Stage: currStage,
			Status: FAIL,
			Error: errors.New(fmt.Sprintf("build tool %s not part of accepted values: %s...", task.BuildConf.BuildTool, "maven")),
		}
	}

	//check if nexus creds exist
	err := d.RemoteConfig.CheckExists(cred.BuildCredPath("nexus", task.AcctName, cred.Repo))

	if err != nil {
		logout <- []byte(currStageStr + err.Error())
		return &Result {
			Stage: currStage,
			Status: FAIL,
			Error: errors.New("nexus credentials don't exist for " + task.AcctName),
		}
	}

	return d.Exec(currStage, currStageStr, nil, d.PushToNexus(commitHash), logout)
}

func (d *Docker) Execute(stage string, actions *pb.Stage, logout chan []byte) *Result {
	if len(d.ContainerId) == 0 {
		return &Result {
			Stage: stage,
			Status: FAIL,
			Error: errors.New("No container exists, setup before executing"),
		}
	}

	//TODO: call d.Exec function
	return &Result{

	}
}

func (d *Docker) Exec(currStage string, currStageStr string, env []string, cmds []string, logout chan []byte) *Result {
	ctx := context.Background()

	resp, err := d.DockerClient.ContainerExecCreate(ctx, d.ContainerId, types.ExecConfig{
		Tty: true,
		AttachStdin: true,
		AttachStderr: true,
		AttachStdout: true,
		Env: env,
		Cmd: cmds,
	})

	if err != nil {
		return &Result{
			Stage:  currStage,
			Status: FAIL,
			Error:  err,
		}
	}

	attachedExec, err := d.DockerClient.ContainerExecAttach(ctx, resp.ID, types.ExecConfig{
		Tty: true,
		AttachStdin: true,
		AttachStderr: true,
		AttachStdout: true,
		Env: env,
		Cmd: cmds,
	})

	defer attachedExec.Conn.Close()

	d.writeToInfo(currStageStr, attachedExec.Reader, logout)

	if err != nil {
		return &Result{
			Stage:  currStage,
			Status: FAIL,
			Error:  err,
		}
	}

	return &Result{
		Stage:  currStage,
		Status: PASS,
		Error:  nil,
	}
}


func (d *Docker) writeToInfo(stage string, rd *bufio.Reader, infochan chan []byte) {
	for {
		//TODO: if we swap to scanner will it outputs nicer?
		str, err := rd.ReadString('\n')

		if err != nil {
			if err != io.EOF {
				ocelog.Log().Error("Read Error:", err)
			} else {
				ocelog.Log().Debug("EOF, finished writing to Info")
			}
			return
		}

		infochan <- []byte(stage + str)

		//our setup script will echo this to stdout, telling us script is finished downloading. This is HACK for keeping container alive
		if str == "Finished with downloading source code\r\n" {
			return
		}
	}
}
