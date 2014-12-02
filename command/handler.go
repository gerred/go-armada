package command

import (
	"fmt"

	"github.com/codegangsta/cli"
)

func handleWithManifest(c *cli.Context, handler func(c *cli.Context, m Manifest)) {
	filename := "armada.toml"

	if len(c.Args()) > 0 {
		filename = c.Args().First()
	}

	manifest, err := ParseManifest(filename)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	handler(c, manifest)
}
