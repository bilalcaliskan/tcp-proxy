package logging

import (
	"go.uber.org/zap"
	"testing"
)

func TestCreateLogger(t *testing.T) {
	_, err := zap.NewProduction()
	if err != nil {
		t.Fatalf("an error occured while creating logger, error is %v\n", err.Error())
	}
}