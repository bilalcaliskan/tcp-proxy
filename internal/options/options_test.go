package options

import "testing"

func TestGetTcpProxyOptions(t *testing.T) {
	t.Log("fetching default options.TcpProxyOptions")
	opts := GetTcpProxyOptions()
	t.Logf("fetched default options.TcpProxyOptions, %v\n", opts)
}
