package command

import (
	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

func NewCleanCommand() cli.Command {
	return cli.Command{
		Name:      "clean",
		ShortName: "c",
		Usage:     "options for cleaning Docker artifacts",
		Subcommands: []cli.Command{
			{
				Name:  "images",
				Usage: "clean unused images",
				Action: func(c *cli.Context) {
					handleWithManifest(c, execCleanImagesCommandFunc)
				},
			},
			{
				Name:  "containers",
				Usage: "clean unused containers",
				Action: func(c *cli.Context) {
					handleWithManifest(c, execCleanContainersCommandFunc)
				},
			},
		},
	}
}

func execCleanImagesCommandFunc(opts *CommandHandlerOptions) {
	live := !opts.Context.GlobalBool("dry-run")

	images, _ := opts.Client.ListImages(true)

	for _, v := range images {
		if contains(v.RepoTags, "<none>:<none>") {
			hostLog(opts, "Removing orphaned image", v.ID[0:11])
			if live {
				err := opts.Client.RemoveImage(v.ID)
				if err != nil {
					hostLog(opts, "Unable to remove image", v.ID[0:11], "because of error:", err.Error())
				}
			}
		}
	}
}

func execCleanContainersCommandFunc(opts *CommandHandlerOptions) {
	live := !opts.Context.GlobalBool("dry-run")

	containers, _ := opts.Client.ListContainers(docker.ListContainersOptions{
		All: true,
	})

	for _, v := range containers {
		container, _ := opts.Client.InspectContainer(v.ID)
		if !(container.State.Running || container.State.Paused) {
			hostLog(opts, "Removing unused container", v.ID[0:11])
			if live {
				err := opts.Client.RemoveContainer(docker.RemoveContainerOptions{
					ID:            v.ID,
					RemoveVolumes: true,
				})
				if err != nil {
					hostLog(opts, "Unable to remove container", v.ID[0:11], "because of error:", err.Error())
				}
			}
		}
	}
}
