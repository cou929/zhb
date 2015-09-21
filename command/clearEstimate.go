package command

import (
	"github.com/codegangsta/cli"
	"os"
	"strconv"
)

func CmdClearEstimate(c *cli.Context) {
	issueNumer, err := strconv.Atoi(c.Args().Get(0))
	if err != nil {
		cli.ShowCommandHelp(c, "clearEstimate")
	}
	client := NewZenhubClient(c)

	err = client.ClearEstimate(issueNumer)
	DieIfError(err)

	os.Exit(0)
}
