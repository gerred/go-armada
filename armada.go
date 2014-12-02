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

	app.Commands = []cli.Command{
		command.NewStatusCommand(),
	}
	app.Run(os.Args)

}
