package command

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
)

func CmdTransfer(c *cli.Context) {
	issueNumer, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		cli.ShowCommandHelp(c, "transfer")
	}
	pipelineName := c.Args().Get(1)
	if pipelineName == "" {
		cli.ShowCommandHelp(c, "transfer")
	}
	client := NewZenhubClient(c)

	_, err = client.UpdatePipeline(issueNumer, pipelineName)
	DieIfError(err)

	os.Exit(0)
}
