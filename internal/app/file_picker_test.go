package app

import "testing"

func TestLatestFile(t *testing.T) {
	filename, err := LatestFile("/tmp/logs")
	if err != nil {
		t.Error(err)
	}
	t.Log(filename)
}
