package main

import (
	"fmt"
	flag "github.com/spf13/pflag"
	"go.uber.org/zap"
	"io"
	"net"
)

var (
	logger *zap.Logger
	proxyProto, targetProto, targetDns string
	proxyPort, targetPort int
	err error
)

func init() {
	logger, err = zap.NewProduction()
	if err != nil {
		panic(err)
	}

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
		logger.Fatal("fatal error occured while listening", zap.String("proto", proxyProto),
			zap.Int("port", proxyPort), zap.Error(err))
	}

	logger.Info("server successfully started listening for requests", zap.String("proto", proxyProto),
		zap.Int("port", proxyPort))
	for {
		conn, err := listener.Accept()
		if err != nil {
			logger.Fatal("fatal error occured while accepting connection", zap.Error(err))
		}

		go handle(conn, targetProto, connectionStr)
	}
}

func handle(src net.Conn, targetProto, connectionStr string) {
	dst, err := net.Dial(targetProto, connectionStr)
	if err != nil {
		logger.Fatal("fatal error occured while connecting to remote host", zap.String("remoteHost", connectionStr),
			zap.Error(err))
	}

	defer func() {
		err := dst.Close()
		if err != nil {
			panic(err)
		}
	}()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		if _, err := io.Copy(dst, src); err != nil {
			logger.Fatal("fatal error occured while proxying", zap.String("src", src.LocalAddr().String()),
				zap.String("dst", src.RemoteAddr().String()), zap.Error(err))
		}
	}()

	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		logger.Fatal("fatal error occured while proxying", zap.String("src", src.LocalAddr().String()),
			zap.String("dst", src.RemoteAddr().String()), zap.Error(err))
	}
}