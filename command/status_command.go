package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func NewStatusCommand() cli.Command {
	return cli.Command{
		Name:      "status",
		ShortName: "s",
		Usage:     "Gets the status of hosts from an Armada manifest",
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:   "env, e",
				Value:  "",
				Usage:  "Filter hosts by environment",
				EnvVar: "ARMADA_ENVIRONMENT",
			},
		},
		Action: func(c *cli.Context) {
			handleWithManifest(c, execStatusCommandFunc)
		},
	}
}

func execStatusCommandFunc(c *cli.Context, m Manifest) {
	fmt.Println("Doing status stuff...")
	fmt.Printf("Manifest: %+v\n", m)
}
