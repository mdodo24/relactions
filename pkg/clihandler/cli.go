package clihandler

import (
	"log"
	"os"
	"time"
	"fmt"

	"github.com/urfave/cli/v2"
)

func Run(appname, version string) {

	app := &cli.App{
		Name: appname,
		Version: version,
		Compiled: time.Now(),
		Authors: []*cli.Author{
			{
				Name: "Madhur Dodeja",
				Email: "dodejamadhur@gmail.com",
			},
		},
		Usage: "execute release pipeline actions",
		Commands: []*cli.Command{rcCmd, jiraCmd},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func before(cCtx *cli.Context) error {
	fmt.Fprintf(cCtx.App.Writer, "Before function\n")
	return nil
}