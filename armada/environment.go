package armada

type Environment struct {
	Hosts    []string
	HostsUri string `toml:"hosts_uri"`
}
