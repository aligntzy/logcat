package config

import (
	"testing"
)

func TestGet(t *testing.T) {
	if len(C.Path) == 0 {
		t.Error("C.Path is empty")
	}
}
