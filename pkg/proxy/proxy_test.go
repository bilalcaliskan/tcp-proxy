package proxy

import (
	"testing"
)

// https://gist.github.com/mtilson/00f72d7cbd98e3d1b9cf2c8bb9ec39b7
// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html
func TestProxy(t *testing.T) {
	var cases = []struct{
		name, proxyProto, targetProto, targetDns string
		proxyPort, targetPort int
		success bool
	}{
		{"TCP3000", "tcp", "tcp", "en.wikipedia.org", 3000,
			443, true},
		{"TCP3001", "tcp", "tcp", "en.wikipedia.org", 3001,
			443, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {

		})
	}

	/*for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			connChan := make(chan net.Conn, 1)

			// For every listener spawn the following routine
			go func() {
				listener, err := net.Listen(tc.proxyProto, fmt.Sprintf(":%d", tc.proxyPort))
				if err != nil {
					t.Errorf("error occured while listening on port %d. error=%v\n", tc.proxyPort, err.Error())
					return
				}
				c, err := listener.Accept()
				if err != nil {
					t.Errorf("error occured while acception connection. error=%v\n", err.Error())
					return
				}
				connChan <- c
			}()

			select {
			case conn := <-connChan:
				// new connection or nil if acceptor is down, in which case we should
				// do something (respawn, stop when everyone is down or just explode)
				connectionStr := fmt.Sprintf("%s:%d", tc.targetDns, tc.targetPort)
				dst, err := net.Dial(tc.targetProto, connectionStr)
				if err != nil {
					t.Errorf("error occured while creating connection. error=%v\n", err.Error())
					return
				}

				// Copy our source's output to the destination
				if _, err := io.Copy(dst, conn); err != nil {
					t.Errorf("error occured while proxying. error=%v\n", err.Error())
					return
				}

				// Copy our destination's output back to our source
				if _, err := io.Copy(conn, dst); err != nil {
					t.Errorf("error occured while proxying. error=%v\n", err.Error())
				}

				err = conn.Close()
				if err != nil {
					t.Errorf("error occured while closing listener. error=%v\n", err.Error())
					return
				}

				err = dst.Close()
				if err != nil {
					t.Errorf("error occured while closing connection. error=%v\n", err.Error())
					return
				}
			case <-time.After(time.Second * 10):
				// timeout branch, no connection for a minute
			}

		})
	}*/
}