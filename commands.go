package main

import (
	"fmt"
	"os"

	"github.com/codegangsta/cli"
	"github.com/cou929/zhb/command"
)

var GlobalFlags = []cli.Flag{
	cli.StringFlag{
		Name:   "authtoken, a",
		Value:  "",
		Usage:  "Auth token of zenhub",
		EnvVar: "ZHB_AUTH_TOKEN",
	},
	cli.StringFlag{
		Name:   "org, o",
		Value:  "",
		Usage:  "Name of github organization",
		EnvVar: "ZHB_ORG",
	},
	cli.StringFlag{
		Name:   "repo, r",
		Value:  "",
		Usage:  "Name of github repository",
		EnvVar: "ZHB_REPO",
	},
	cli.StringFlag{
		Name:   "repo_id",
		Value:  "",
		Usage:  "Repository Id of github",
		EnvVar: "ZHB_REPOID",
	},
	cli.BoolFlag{
		Name:  "verbose",
		Usage: "Show verbose logs",
	},
}

var Commands = []cli.Command{
	{
		Name:   "issue",
		Usage:  "Show details of the issue",
		Action: command.CmdIssue,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "transfer",
		Usage:  "Transfer the issue to the pipeline",
		Action: command.CmdTransfer,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "estimate",
		Usage:  "Set the estimate value to the issue",
		Action: command.CmdEstimate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "clearEstimate",
		Usage:  "Clear estimate value of the issue",
		Action: command.CmdClearEstimate,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "pipelines",
		Usage:  "Show all pipelines",
		Action: command.CmdPipelines,
		Flags:  []cli.Flag{},
	},
	{
		Name:   "events",
		Usage:  "Show events of ZenHub activities",
		Action: command.CmdEvents,
		Flags: []cli.Flag{
			cli.IntFlag{Name: "page, p", Value: 1, Usage: "Page number of events"},
		},
	},
}

func CommandNotFound(c *cli.Context, command string) {
	fmt.Fprintf(os.Stderr, "%s: '%s' is not a %s command. See '%s --help'.", c.App.Name, command, c.App.Name, c.App.Name)
	os.Exit(2)
}
