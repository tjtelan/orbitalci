package output

import (
	"bitbucket.org/level11consulting/ocelot/admin/models"
	"bitbucket.org/level11consulting/ocelot/client/commandhelper"
	pb "bitbucket.org/level11consulting/ocelot/werker/protobuf"
	"context"
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"google.golang.org/grpc"
	"io"
	"bitbucket.org/level11consulting/ocelot/util/cmd_table"
	"os"
	"os/exec"
	"strings"
)

const synopsis = "stream logs on running or completed build"
const help = `
Usage: ocelot logs --hash <git_hash>
  Will stream logs of a running or completed build identified by the hash. 
  If the build is: 
    - completed:   logs will stream from storage via the admin
    - in progress: logs will stream from the werker node running the build 
    - not found:   he cli will return an error and exit. 
  If there are multiple builds of the same hash, it will stream the latest found in storage.
  If you have the build id from running 'ocelot summary', you can also stream old logs by running: 
    ocelot logs -build-id <id>
`


func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui, config: commandhelper.Config}
	c.init()
	return c
}

type cmd struct {
	UI cli.Ui
	flags   *flag.FlagSet
	config *commandhelper.ClientConfig
	hash string
	buildId int
}


func (c *cmd) GetClient() models.GuideOcelotClient {
	return c.config.Client
}

func (c *cmd) GetUI() cli.Ui {
	return c.UI
}

func (c *cmd) GetConfig() *commandhelper.ClientConfig {
	return c.config
}


func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.hash, "hash", "ERROR",
		"*REQUIRED* hash to get build data of")
	c.flags.IntVar(&c.buildId, "build-id", 0, "build id from build_summary table, if you do not want to get the last")
}

func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}
	ctx := context.Background()
	if err := commandhelper.CheckConnection(c, ctx); err != nil {
		return 1
	}
	var build *models.Builds
	var err error

	if c.buildId != 0 {
		// todo: this is now broken. just fyi
		builds := make(map[string]*models.BuildRuntimeInfo)
		builds["br"] = &models.BuildRuntimeInfo{Done: true}
		build = &models.Builds{Builds:builds}
	} else {
		if c.hash == "ERROR" {

			var (
				cmdOut []byte
				err    error
			)

			cmdName := "git"

			getBranch := []string{"rev-parse", "--abbrev-ref",  "HEAD"}
			if cmdOut, err = exec.Command(cmdName, getBranch...).Output(); err != nil {
				fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command to find the current branch: ", err)
			}

			remoteBranch := fmt.Sprintf("origin/%s", string(cmdOut))
			if cmdOut, err = exec.Command(cmdName, "rev-parse", strings.TrimSpace(remoteBranch)).Output(); err != nil {
				fmt.Fprintln(os.Stderr, "There was an error running git rev-parse command to find the most recently pushed commit: ", err)
			}

			sha := string(cmdOut)

			if len(sha) > 0 {
				c.UI.Info(fmt.Sprintf("no -hash passed, using detected hash %s", sha))
				c.hash = sha
			} else {
				c.UI.Error("flag --hash is required, otherwise there is no build to tail")
				return 1
			}
		}

		build, err = c.config.Client.BuildRuntime(ctx, &models.BuildQuery{Hash: c.hash})
		if err != nil {
			c.UI.Error("unable to get build runtime! error: " + err.Error())
			return 1
		}
	}

	if len(build.Builds) > 1 {
		c.UI.Info(fmt.Sprintf("it's your lucky day, there's %d hashes matching that str. Please enter a more complete git hash", len(build.Builds)))
		c.UI.Output(cmd_table.SelectFromHashes(build))
		return 0
	} else if len(build.Builds) == 1 {
		for _, build := range build.Builds {
			if build.Done {
				return c.fromStorage(ctx, build.Hash)
			} else {
				return c.fromWerker(ctx, build)
			}
		}
	} else {
		c.UI.Info(fmt.Sprintf("no builds found for entry: %s", c.hash))
		return 0
	}
	return 0
}

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}


func (c *cmd) fromStorage(ctx context.Context, hash string) int {
	stream, err := c.config.Client.Logs(ctx, &models.BuildQuery{Hash: hash})
	if err != nil {
		commandhelper.UIErrFromGrpc(err, c.UI, "Unable to get stream from admin.")
		return 1
	}
	for {
		line, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			return 0
		} else if err != nil {
			//todo: implement this
			c.UI.Warn("Warning: partial matching is not yet supported for streaming from storage")
			commandhelper.UIErrFromGrpc(err, c.UI, "Error streaming from storage via admin.")
			return 1
		}
		c.UI.Info(line.GetOutputLine())
	}
}

func (c *cmd) fromWerker(ctx context.Context, build models.BuildRuntime) int {
	var opts []grpc.DialOption
	// right now werker is insecure
	opts = append(opts, grpc.WithInsecure())
	client, err := build.CreateBuildClient(opts)
	if err != nil {
		c.UI.Error(fmt.Sprintf("Error dialing the werker at %s:%s! Error: %s", build.GetIp(), build.GetGrpcPort(), err.Error()))
		return 1
	}

	stream, err := client.BuildInfo(ctx, &pb.Request{Hash: build.GetHash()})
	if err != nil {
		commandhelper.UIErrFromGrpc(err, c.UI, fmt.Sprintf("Unable to get build info stream from client at %s:%s!", build.GetIp(), build.GetGrpcPort()))
		return 1
	}
	for {
		line, err := stream.Recv()
		if err == io.EOF {
			stream.CloseSend()
			return 0
		} else if err != nil {
			commandhelper.UIErrFromGrpc(err, c.UI, "Error streaming from werker.")
			return 1
		}
		c.UI.Info(line.GetOutputLine())
	}
}


