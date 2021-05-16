package main

import (
	"fmt"
	"github.com/bilalcaliskan/tcp-proxy/pkg/logging"
	"github.com/bilalcaliskan/tcp-proxy/pkg/proxy"
	_ "github.com/dimiro1/banner/autoload"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"net"
)

var (
	logger                             *zap.Logger
	proxyProto, targetProto, targetDns string
	proxyPort, targetPort              int
)

func init() {
	logger = logging.GetLogger()

	flag.StringVar(&proxyProto, "proxyProto", "tcp", "Provide a proxy server protocol")
	flag.IntVar(&proxyPort, "proxyPort", 3000, "Provide a port to run proxy server on")
	flag.StringVar(&targetProto, "targetProto", "tcp", "Provide a target protocol to proxy")
	flag.StringVar(&targetDns, "targetDns", "en.wikipedia.org", "Provide a target DNS to proxy")
	flag.IntVar(&targetPort, "targetPort", 443, "Provide a target port to proxy")
	flag.Parse()
}

func main() {
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	connectionStr := fmt.Sprintf("%s:%d", targetDns, targetPort)
	listener, err := net.Listen(proxyProto, fmt.Sprintf(":%d", proxyPort))
	if err != nil {
		logger.Fatal("fatal error occurred while listening", zap.String("proto", proxyProto),
			zap.Int("port", proxyPort), zap.Error(err))
	}

	logger.Info("server successfully started listening for requests", zap.String("proto", proxyProto),
		zap.Int("port", proxyPort))
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatal("fatal error occurred while accepting connection", zap.Error(err))
		}

		go proxy.Proxy(conn, targetProto, connectionStr)
	}
}
