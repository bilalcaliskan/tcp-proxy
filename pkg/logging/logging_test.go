package logging

import (
	"go.uber.org/zap"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	_, err := zap.NewProduction()
	if err != nil {
		t.Errorf("%v\n", err.Error())
	}
}
