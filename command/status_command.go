package command

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
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

func execStatusCommandFunc(opts *CommandHandlerOptions) {
	fullName := opts.Manifest.FullName()
	var s []docker.APIContainers

	containers, _ := opts.Client.ListContainers(docker.ListContainersOptions{})

	for _, c := range containers {
		if fullName == c.Image {
			s = append(s, c)
		}
	}

	fmt.Println(len(s))
}
