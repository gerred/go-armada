package armada

type Manifest struct {
	Container    `toml:"container"`
	Environments map[string]Environment `toml:"environments"`
}
