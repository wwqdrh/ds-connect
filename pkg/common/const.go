package common

const (
	// Localhost ip of localhost
	Localhost = "127.0.0.1"
	// YyyyMmDdHhMmSs timestamp format
	YyyyMmDdHhMmSs = "2006-01-02 15:04:05"
	// StandardSshPort standard ssh port
	StandardSshPort = 22
	// StandardDnsPort standard dns port
	StandardDnsPort = 53

	// EnvVarLocalDomains environment variable for local domain config
	EnvVarLocalDomains = "DS_LOCAL_DOMAIN"
	// EnvVarDnsProtocol environment variable for shadow pod dns protocol
	EnvVarDnsProtocol = "DS_DNS_PROTOCOL"
	// EnvVarLogLevel environment variable for shadow pod log level
	EnvVarLogLevel = "DS_LOG_LEVEL"

	DnsModeLocalDns = "localDNS"
)
