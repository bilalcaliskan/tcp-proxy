package proxy

import (
	"crypto/tls"
	"fmt"
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net"
	"net/http"
	"sync"
	"testing"
	"time"
)

// https://gist.github.com/mtilson/00f72d7cbd98e3d1b9cf2c8bb9ec39b7
// https://ipfs.io/ipfs/QmfYeDhGH9bZzihBUDEQbCbTc5k5FZKURMUoUvfmc27BwL/socket/tcp_sockets.html
func TestProxy(t *testing.T) {
	var cases = []struct {
		name, proxyProto, targetProto, targetDns, contentUrl string
		proxyPort, targetPort                                int
	}{
		{"TCP3000", "tcp", "tcp", "en.wikipedia.org",
			"/wiki/OSI_model", 3000, 443},
		{"TCP3001", "tcp", "tcp", "en.wikipedia.org",
			"/wiki/Tesla,_Inc.", 3001, 443},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			go func() {
				var conn net.Conn
				var err error

				defer func(conn net.Conn) {
					err := conn.Close()
					if err != nil {
						t.Error("an error occured while closing current connection")
						return
					}
				}(conn)

				listener, err := net.Listen(tc.proxyProto, fmt.Sprintf(":%d", tc.proxyPort))
				if err != nil {
					t.Errorf("%v\n", err.Error())
					return
				}

				for {
					conn, err = listener.Accept()
					if err != nil {
						t.Errorf("%v\n", err.Error())
						return
					}

					connectionStr := fmt.Sprintf("%s:%d", tc.targetDns, tc.targetPort)
					go Proxy(conn, tc.targetProto, connectionStr)
				}
			}()

			var wg sync.WaitGroup
			wg.Add(1)

			go func() {
				for i := 0; i <= 5; i++ {
					if i == 5 {
						// * Error()  = Fail()    + Log()
						// * Errorf() = Fail()    + Logf()
						// * Fatal()  = FailNow() + Log()
						// * Fatalf() = FailNow() + Logf()
						t.Errorf("connection to port %d could not succeeded, not retrying!\n", tc.proxyPort)
						break
					}

					_, err := net.Dial(tc.targetProto, fmt.Sprintf("127.0.0.1:%d", tc.proxyPort))
					if err != nil {
						t.Logf("connection to port %d could not succeeded, retrying...\n", tc.proxyPort)
						time.Sleep(2 * time.Second)
						continue
					}

					t.Logf("connection to port %d succeeded!\n", tc.proxyPort)

					// make http request and see response
					client := &http.Client{
						Timeout: 5 * time.Second,
						Transport: &http.Transport{
							TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
						},
					}
					req, err := http.NewRequest("GET", fmt.Sprintf("https://127.0.0.1:%d%s", tc.proxyPort,
						tc.contentUrl), nil)
					assert.Nil(t, err)

					req.Host = tc.targetDns
					resp, err := client.Do(req)
					assert.Nil(t, err)

					body, err := ioutil.ReadAll(resp.Body)
					assert.Nil(t, err)
					assert.NotEmpty(t, string(body))

					break
				}

				wg.Done()
			}()

			wg.Wait()
		})
	}
}
