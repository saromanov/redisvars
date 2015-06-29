package redisvars

import (
	"testing"
)

func TestCreateClient(t *testing.T) {
	result := New("localhost:6379")
	if !result.Status {
		t.Errorf("Redis server is not ready")
	}
}
