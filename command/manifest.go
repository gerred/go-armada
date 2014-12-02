package command

import "github.com/BurntSushi/toml"

type Manifest struct {
	Container    `toml:"container"`
	Environments map[string]Environment `toml:"environments"`
}

func ParseManifest(filename string) (Manifest, error) {
	var manifest Manifest

	if _, err := toml.DecodeFile(filename, &manifest); err != nil {
		return manifest, err
	}

	return manifest, nil
}
