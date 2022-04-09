package main

import (
	"fmt"
	"io/ioutil"
	"net"
	"os"
	"strings"

	"github.com/bilalcaliskan/tcp-proxy/internal/logging"
	"github.com/bilalcaliskan/tcp-proxy/internal/options"
	"github.com/bilalcaliskan/tcp-proxy/internal/proxy"
	"github.com/dimiro1/banner"
	"go.uber.org/zap"
)

var (
	logger *zap.Logger
	tpo    *options.TcpProxyOptions
)

func init() {
	logger = logging.GetLogger()
	tpo = options.GetTcpProxyOptions()

	bannerBytes, _ := ioutil.ReadFile("banner.txt")
	banner.Init(os.Stdout, true, false, strings.NewReader(string(bannerBytes)))
}

func main() {
	defer func() {
		err := logger.Sync()
		if err != nil {
			panic(err)
		}
	}()

	connectionStr := fmt.Sprintf("%s:%d", tpo.TargetDns, tpo.TargetPort)
	listener, err := net.Listen(tpo.ProxyProto, fmt.Sprintf(":%d", tpo.ProxyPort))
	if err != nil {
		logger.Fatal("fatal error occurred while listening", zap.String("proto", tpo.ProxyProto),
			zap.Int("port", tpo.ProxyPort), zap.Error(err))
	}

	logger.Info("server successfully started listening for requests", zap.String("proto", tpo.ProxyProto),
		zap.Int("port", tpo.ProxyPort))
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatal("fatal error occurred while accepting connection", zap.Error(err))
		}

		go func() {
			if err := proxy.Proxy(conn, tpo.TargetProto, connectionStr); err != nil {
				logger.Fatal("fatal error occured while proxying", zap.Error(err))
			}
		}()
	}
}
