package main

import (
	"fmt"
	"github.com/bilalcaliskan/tcp-proxy/pkg/logging"
	"github.com/bilalcaliskan/tcp-proxy/pkg/options"
	"github.com/bilalcaliskan/tcp-proxy/pkg/proxy"
	"github.com/dimiro1/banner"
	"go.uber.org/zap"
	"io/ioutil"
	"net"
	"os"
	"strings"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()

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

	tpo := options.GetTcpProxyOptions()
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

		go proxy.Proxy(conn, tpo.TargetProto, connectionStr)
	}
}
