package proxy

/*func TestProxy(t *testing.T) {
	var cases = []struct{
		name, proxyProto, targetProto, targetDns string
		proxyPort, targetPort int
		success bool
	}{
		{"TCP", "tcp", "tcp", "en.wikipedia.org", 3000,
			443, true},
		{"UDP", "udp", "udp", "en.wikipedia.org", 3000,
			443, false},
	}

	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			ch := make(chan net.Conn, 1)

			go func() {
				listener, err := net.Listen(tc.proxyProto, fmt.Sprintf(":%d", tc.proxyPort))
				if err != nil {
					t.Errorf("error occured while listening on port %d. error=%v\n", tc.proxyPort, err.Error())
					return
				}

				conn, err := listener.Accept()
				if err != nil {
					t.Errorf("error occured while accepting connections on listener. error=%v\n", err.Error())
					return
				}
				ch <- conn
			}()

			for {
				select {
				case <-time.After(time.Minute):
					// timeout branch, no connection for a minute
				case conn := <- ch:
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
				}
			}
		})
	}
}*/