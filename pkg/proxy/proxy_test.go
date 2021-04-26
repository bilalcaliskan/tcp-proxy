package proxy

import (
	"fmt"
	"io"
	"log"
	"net"
	"testing"
	"time"
)

// https://gist.github.com/mtilson/00f72d7cbd98e3d1b9cf2c8bb9ec39b7
// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html
func TestProxy(t *testing.T) {
	var cases = []struct{
		name, proxyProto, targetProto, targetDns string
		proxyPort, targetPort int
	}{
		// TODO: Different protos, like tcp4, ip4. Check https://golang.org/pkg/net/#Dial
		{"TCP3000", "tcp", "tcp", "en.wikipedia.org", 3000,
			443},
		{"TCP3001", "tcp", "tcp", "en.wikipedia.org", 3001,
			443},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			go func() {
				listener, err := net.Listen(tc.proxyProto, fmt.Sprintf(":%d", tc.proxyPort))
				if err != nil {
					t.Errorf("%v\n", err.Error())
					return
				}

				conn, err := listener.Accept()
				if err != nil {
					t.Errorf("%v\n", err.Error())
					return
				}

				go func(conn net.Conn) {
					dst, err := net.Dial(tc.targetProto, fmt.Sprintf("%s:%d", tc.targetDns, tc.targetPort))
					if err != nil {
						t.Errorf("%v\n", err.Error())
						return
					}

					if _, err := io.Copy(dst, conn); err != nil {
						t.Errorf("%v\n", err.Error())
						return
					}

					if _, err := io.Copy(conn, dst); err != nil {
						t.Errorf("%v\n", err.Error())
					}

					defer func() {
						err := conn.Close()
						if err != nil {
							t.Errorf("%v\n", err.Error())
							return
						}
					}()

					err = dst.Close()
					if err != nil {
						t.Errorf("%v\n", err.Error())
						return
					}
				}(conn)
			}()

			go func() {
				for i := 0; i < 5; i++ {
					_, err := net.Dial(tc.targetProto, fmt.Sprintf("127.0.0.1:%d", tc.proxyPort))
					if err == nil {
						log.Println("connection succeeded")
						break
					}
					time.Sleep(2 * time.Second)
					log.Println("connection could not succeeded, retrying...")
				}
			}()

			// TODO: wait for goroutines to complete their jobs using channel
			time.Sleep(10 * time.Second)
		})
	}
}