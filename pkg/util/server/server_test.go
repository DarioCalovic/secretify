package server_test

import (
	"testing"

	"github.com/DarioCalovic/secretify/pkg/util/server"
)

// Improve tests
func TestNew(t *testing.T) {
	e := server.New()
	if e == nil {
		t.Errorf("Server should not be nil")
	}
}
