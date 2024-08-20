package main

import (
	"relactions/pkg/clihandler"
)

func main() {

	//run CLI server
	appName := "relactions"
	version := relactionsVersion
	clihandler.Run(appName, version)
}
