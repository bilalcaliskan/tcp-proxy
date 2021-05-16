package proxy

import (
	"github.com/bilalcaliskan/tcp-proxy/pkg/logging"
	"go.uber.org/zap"
	"io"
	"net"
)

var logger *zap.Logger

func init() {
	logger = logging.GetLogger()
}

// Proxy proxies the requests to the remote
func Proxy(src net.Conn, targetProto, connectionStr string) {
	dst, err := net.Dial(targetProto, connectionStr)
	if err != nil {
		logger.Fatal("fatal error occurred while connecting to remote host", zap.String("remoteHost", connectionStr),
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
			logger.Fatal("fatal error occurred while proxying", zap.String("src", src.LocalAddr().String()),
				zap.String("dst", src.RemoteAddr().String()), zap.Error(err))
		}
	}()

	// Copy our destination's output back to our source
	if _, err := io.Copy(src, dst); err != nil {
		logger.Fatal("fatal error occurred while proxying", zap.String("src", src.LocalAddr().String()),
			zap.String("dst", src.RemoteAddr().String()), zap.Error(err))
	}
}
