package clihandler

import (
	"github.com/urfave/cli/v2"
	"relactions/pkg/relhandler"
)

var rcCmd = &cli.Command{
	Name: "rc",
	Usage: "release checklist related commands",
	Subcommands: []*cli.Command{
		{
			Name: "validate",
			Usage: "rc validate command",
			Flags: []cli.Flag{ManifestFlag},
			Before: before,
			Action: rcValidate,
		},
	},
}

func rcValidate(ctx *cli.Context) error {
	manifest := ctx.String(FlagNameManifestID)
	res := relhandler.ExecuteRCValidate(manifest)
	if res.Error != "" {
		return cli.Exit(res.Error, 1)
	}
	return cli.Exit(res.Message, 0)
}