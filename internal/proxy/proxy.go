package proxy

import (
	"io"
	"net"
)

// Proxy proxies the requests to the remote
func Proxy(src net.Conn, targetProto, connectionStr string) error {
	dst, err := net.Dial(targetProto, connectionStr)
	if err != nil {
		return err
	}

	defer dst.Close()

	// Run in goroutine to prevent io.Copy from blocking
	go func() {
		// Copy our source's output to the destination
		_, _ = io.Copy(dst, src)
	}()

	// Copy our destination's output back to our source
	_, _ = io.Copy(src, dst)

	return nil
}
