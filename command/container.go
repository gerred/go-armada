package command

type Container struct {
	Name    string
	Image   string
	Tag     string
	EnvVars map[string]string `toml:"env_vars"`
	Ports   []Port
}

type Port struct {
	Host      int16
	Container int16
	Type      string
}
