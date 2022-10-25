package command

type ConnectOptions struct {
	ClusterDomain string
	ProxyPort     int
	DnsCacheTtl   int
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
