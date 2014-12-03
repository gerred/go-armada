package command

type Environment struct {
	Hosts    []Host `toml:"hosts"`
	HostsUri string `toml:"hosts_uri"`
}

type Host struct {
	Endpoint  string
	CertPath  string `toml:"cert_path"`
	TlsVerify bool   `toml:"tls_verify"`
}
