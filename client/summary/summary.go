package summary

import (
	clr "github.com/fatih/color"

	models "bitbucket.org/level11consulting/ocelot/models/pb"
	"bitbucket.org/level11consulting/ocelot/client/commandhelper"
	"bytes"
	"context"
	"flag"
	"fmt"
	"github.com/mitchellh/cli"
	"github.com/olekukonko/tablewriter"
	"time"
)

const synopsis = "show summary table of specific repo"
const help = `
Usage: ocelot summary -acct-repo <acct>/<repo>
  Retrieve summary table of a specific repo (i.e. level11consulting/ocelot). If -limit is not specified, then the 
  limit will be set to 5, and only the last 5 runs will be shown.
  Full usage:
    $ ocelot summary -acct-repo mariannefeng/test-ocelot -limit 2

+----------+-------------+--------------------------+--------------------+--------+--------+------------------------------------------+
| BUILD ID |    REPO     |      BUILD DURATION      |     START TIME     | RESULT | BRANCH |                   HASH                   |
+----------+-------------+--------------------------+--------------------+--------+--------+------------------------------------------+
| 20       | test-ocelot | 2 minutes and 37 seconds | Thu Feb 8 10:58:36 | PASS   | master | 2f4117d4a38eede1d7c8db27d94253491bf2064d |
| 19       | test-ocelot | running                  | Thu Feb 8 10:54:06 | FAIL   | master | 2f4117d4a38eede1d7c8db27d94253491bf2064d |
+----------+-------------+--------------------------+--------------------+--------+--------+------------------------------------------+


` + commandhelper.AcctRepoHelp


func New(ui cli.Ui) *cmd {
	c := &cmd{UI: ui, config: commandhelper.Config, OcyHelper: &commandhelper.OcyHelper{}}
	c.init()
	return c
}

type cmd struct {
	UI cli.Ui
	flags   *flag.FlagSet
	config *commandhelper.ClientConfig
	*commandhelper.OcyHelper
	limit int
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

func (c *cmd) Synopsis() string {
	return synopsis
}

func (c *cmd) Help() string {
	return help
}

func (c *cmd) init() {
	c.flags = flag.NewFlagSet("", flag.ContinueOnError)
	c.flags.StringVar(&c.OcyHelper.AcctRepo, "acct-repo", "ERROR", "<account>/<repo> to display build summaries for ")
	c.flags.IntVar(&c.limit, "limit", 5, "number of rows to fetch")
}


func (c *cmd) Run(args []string) int {
	if err := c.flags.Parse(args); err != nil {
		return 1
	}

	if err := c.DetectAcctRepo(c.UI); err != nil {
		return 1
	}
	if err := c.SplitAndSetAcctRepo(c.UI); err != nil {
		return 1
	}
	ctx := context.Background()
	if err := commandhelper.CheckConnection(c, ctx); err != nil {
		return 1
	}
	commandhelper.Debuggit(c.UI, "getting last few summaries")
	commandhelper.Debuggit(c.UI, c.OcyHelper.Repo + ": " + c.OcyHelper.Account + "  " + c.OcyHelper.AcctRepo)
	commandhelper.Debuggit(c.UI, fmt.Sprintf("%v", &models.RepoAccount{Repo: c.OcyHelper.Repo, Account: c.OcyHelper.Account, Limit: int32(c.limit)}))
	summaries, err := c.config.Client.LastFewSummaries(ctx, &models.RepoAccount{Repo: c.OcyHelper.Repo, Account: c.OcyHelper.Account, Limit: int32(c.limit)})
	if err != nil {
		// todo: add more descriptive error
		c.LastFewSummariesErr(err, c.GetUI())
		return 1
	}
	commandhelper.Debuggit(c.UI, "found them!")
	// todo: need a check/error for when nothing is found, right now just generated an empty table
	writer := &bytes.Buffer{}
	writ := tablewriter.NewWriter(writer)
	writ.SetAlignment(tablewriter.ALIGN_LEFT)   // Set Alignment
	writ.SetHeader([]string{"Build ID", "Repo", "Build Duration", "Start Time", "Result", "Branch", "Hash"})
	if !c.config.Theme.NoColor {
		writ.SetHeaderColor(
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold},
			tablewriter.Colors{tablewriter.FgBlackColor, tablewriter.Bold})
	}

	for _, sum := range summaries.Sums {
		writ.Append(generateTableRow(sum, c.config.Theme))
	}
	writ.Render()
	her := writer.String()
	c.UI.Output("\n" + her)
	return 0
}

func generateTableRow(summary *models.BuildSummary, theme *commandhelper.ColorDefs) []string {
	tym := time.Unix(summary.BuildTime.Seconds, int64(summary.BuildTime.Nanos))
	var row []string
	var color *commandhelper.Color
	var status string
	failedValidation := summary.QueueTime.Seconds == 0
	isQueued := summary.BuildDuration < 0 && summary.BuildTime.Seconds == 0
	isRunning := summary.BuildDuration < 0
	//we color line output based on success/failure
	if isRunning || isQueued {
		status = "N/A"
		color = theme.Running
	} else if failedValidation {
		status = "FAILED PRESTART"
		color = theme.Failed
	} else if summary.Failed {
		status = "FAIL"
		color = theme.Failed
	} else {
		status = "PASS"
		color = theme.Passed
	}
	start, end := writeFirstAndLastColumns(summary, color)

	row = append(row,
		start,
		summary.Repo,
		commandhelper.PrettifyTime(summary.BuildDuration, isQueued),
		tym.Format("Mon Jan 2 15:04:05"),
		status,
		summary.Branch,
		end,
	)
	return row
}


func writeFirstAndLastColumns(summary *models.BuildSummary, color *commandhelper.Color) (start, end string) {
	buf := bytes.NewBuffer([]byte{})
	old := clr.Output
	clr.Output = buf
	defer func(){ clr.Output = old }()
	color.Set()
	fmt.Fprintf(buf, "%d", summary.BuildId)
	start = buf.String()
	buf.Reset()
	fmt.Fprint(buf, summary.Hash)
	clr.Unset()
	end = buf.String()
	return
}