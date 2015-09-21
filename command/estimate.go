package command

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
)

func CmdEstimate(c *cli.Context) {
	issueNumer, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		cli.ShowCommandHelp(c, "estimate")
	}
	estimateValue, err := strconv.Atoi(c.Args().Get(1))
	if err != nil {
		cli.ShowCommandHelp(c, "estimate")
	}
	client := NewZenhubClient(c)

	_, err = client.UpdateEstimate(issueNumer, estimateValue)
	DieIfError(err)

	os.Exit(0)
}
