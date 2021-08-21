package options

import (
	"github.com/spf13/pflag"
)

var tcpProxyOptions = &TcpProxyOptions{}

func init() {
	tcpProxyOptions.addFlags(pflag.CommandLine)
	pflag.Parse()
}

// GetTcpProxyOptions returns the pointer of TcpProxyOptions
func GetTcpProxyOptions() *TcpProxyOptions {
	return tcpProxyOptions
}

// TcpProxyOptions contains frequent command line and application options.
type TcpProxyOptions struct {
	// ProxyProto is the protocol of the proxy server
	ProxyProto string
	// ProxyPort is the port to run proxy server on
	ProxyPort int
	// TargetProto is the target protocol to proxy
	TargetProto string
	// TargetDns is the target DNS to proxy
	TargetDns string
	// TargetPort is the target port to proxy
	TargetPort int
}

func (tpo *TcpProxyOptions) addFlags(fs *pflag.FlagSet) {
	fs.StringVar(&tpo.ProxyProto, "proxyProto", "tcp", "Provide a proxy server protocol")
	fs.IntVar(&tpo.ProxyPort, "proxyPort", 3000, "Provide a port to run proxy server on")
	fs.StringVar(&tpo.TargetProto, "targetProto", "tcp", "Provide a target protocol to proxy")
	fs.StringVar(&tpo.TargetDns, "targetDns", "en.wikipedia.org", "Provide a target DNS to proxy")
	fs.IntVar(&tpo.TargetPort, "targetPort", 443, "Provide a target port to proxy")
}
