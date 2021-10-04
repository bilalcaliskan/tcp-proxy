package options

import "testing"

func TestGetTcpProxyOptions(t *testing.T) {
	t.Logf("fetched default TcpProxyOptions, %v\n", GetTcpProxyOptions())
}
