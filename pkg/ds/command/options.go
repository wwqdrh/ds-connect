package command

type ConnectOptions struct {
	Shadow      string
	SshUser     string
	SshPass     string
	ProxyPort   int
	DnsPort     int
	DnsCacheTtl int
}

// DaemonOptions cli options
type DaemonOptions struct {
	Connect *ConnectOptions
}

var opt *DaemonOptions = &DaemonOptions{
	Connect: &ConnectOptions{},
}

// Get fetch options instance
func Opts() *DaemonOptions {
	return opt
}
