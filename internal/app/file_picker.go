package app

import (
	"errors"
	"os"
	"path/filepath"
	"time"
)

// LatestFile returns the latest file name under the folder
func LatestFile(folder string) (filename string, err error) {
	dirEntry, err := os.ReadDir(folder)
	if err != nil {
		return
	}
	if len(dirEntry) == 0 {
		return "", errors.New("no files")
	}

	var latestTime time.Time
	for _, f := range dirEntry {
		if f.Type().IsDir() {
			continue
		}

		fi, _ := f.Info()
		modTime := fi.ModTime()
		if latestTime.Before(modTime) {
			filename = filepath.Join(folder, f.Name())
			latestTime = modTime
		}
	}

	return
}
