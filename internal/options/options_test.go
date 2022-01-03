package options

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetTcpProxyOptions(t *testing.T) {
	t.Log("fetching default options.TcpProxyOptions")
	opts := GetTcpProxyOptions()
	assert.NotNil(t, opts)
	t.Logf("fetched default options.TcpProxyOptions, %v\n", opts)
}
