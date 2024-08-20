package clihandler

import (
	"github.com/urfave/cli/v2"
	"relactions/pkg/relhandler"
)

var jiraCmd = &cli.Command{
	Name: "jira",
	Usage: "release ticket related commands",
	Subcommands: []*cli.Command{
		{
			Name: "create",
			Usage: "moves release ticket to closed status",
			Flags: []cli.Flag{ManifestFlag},
			Before: before,
			Action: jiraCreate,
		},
		{
			Name: "close",
			Usage: "moves release ticket to closed status",
			Flags: []cli.Flag{ManifestFlag, JiraFlag},
			Before: before,
			Action: jiraClose,
		},
	},
}

func jiraCreate(ctx *cli.Context) error {
	manifest := ctx.String(FlagNameManifestID)

	res := relhandler.ExecuteJiraCreate(manifest)
	if res.Error != "" {
		return cli.Exit(res.Error, 1)
	}
	return cli.Exit(res, 0)
}

func jiraClose(ctx *cli.Context) error {
	manifest := ctx.String(FlagNameManifestID)
	jira := ctx.String(FlagNameJiraID)

	res := relhandler.ExecuteJiraClose(manifest, jira)
	if res.Error != "" {
		return cli.Exit(res.Error, 1)
	}
	return cli.Exit(res, 0)
}