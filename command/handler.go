package command

import (
	"fmt"
	"os/user"
	"path"

	"github.com/codegangsta/cli"
	"github.com/fsouza/go-dockerclient"
)

type DockerConnectionOptions struct {
	Endpoint  string
	TlsVerify bool
	CertPath  string
}

type CommandHandlerOptions struct {
	cli.Context
	Manifest
	docker.Client
	EnvName  string
	Endpoint string
}

func handleWithManifest(c *cli.Context, handler func(opts *CommandHandlerOptions)) {
	dry := c.GlobalBool("dry-run")

	if dry {
		fmt.Println("******* DRY RUN *******")
	}

	filename := "armada.toml"

	if len(c.Args()) > 0 {
		filename = c.Args().First()
	}

	manifest, err := ParseManifest(filename)

	if err != nil {
		fmt.Printf("Error: %s", err)
		return
	}

	for envName, env := range manifest.Environments {
		for _, host := range env.Hosts {
			client, _ := newDockerClient(&DockerConnectionOptions{
				Endpoint:  host,
				TlsVerify: true,
				CertPath:  "/Users/gdillon/.boot2docker/certs/boot2docker-vm/",
			})

			handler(&CommandHandlerOptions{*c, manifest, *client, envName, host})
		}
	}

}

func newDockerClient(d *DockerConnectionOptions) (*docker.Client, error) {

	if d.TlsVerify || d.CertPath != "" {
		if d.CertPath == "" {
			user, err := user.Current()
			if err != nil {
				return nil, err
			}

			d.CertPath = path.Join(user.HomeDir, ".docker")
		}

		cert := path.Join(d.CertPath, "cert.pem")
		key := path.Join(d.CertPath, "key.pem")
		ca := ""
		if d.TlsVerify {
			ca = path.Join(d.CertPath, "ca.pem")
		}

		return docker.NewTLSClient(d.Endpoint, cert, key, ca)
	} else {
		return docker.NewClient(d.Endpoint)
	}
}
