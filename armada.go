package main

import (
	"os"

	"github.com/codegangsta/cli"
	"github.com/gerred/armada/command"
)

func main() {

	app := cli.NewApp()
	app.Version = releaseVersion
	app.Name = "armada"
	app.Usage = "Control Docker deployments and hosts across platforms"

	app.Commands = []cli.Command{
		command.NewStatusCommand(),
		command.NewCleanCommand(),
	}
	app.Flags = []cli.Flag{
		cli.BoolFlag{
			Name:  "dry-run, d",
			Usage: "show the results of an action without performing it.",
		},
	}
	app.Run(os.Args)

}
