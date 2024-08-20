package clihandler

import (
	"github.com/urfave/cli/v2"
)

const (
	FlagNameManifestID = "manifest"
	FlagNameJiraID = "jira"
)

var (
	ManifestFlag = &cli.StringFlag{
		Name: FlagNameManifestID,
		Usage: "artifact ID",
		Aliases: []string{"m"},
		EnvVars: []string{"MANIFEST_ID"},
	}

	JiraFlag = &cli.StringFlag{
		Name: FlagNameJiraID,
		Usage: "release ticket id",
		Aliases: []string{"j"},
		EnvVars: []string{"JIRA_ID"},
	}
)