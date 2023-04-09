package app

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/fsnotify/fsnotify"
)

// LatestFile returns the latest file name under the folder
func LatestFile(folder string) (path string, err error) {
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
			path = filepath.Join(folder, f.Name())
			latestTime = modTime
		}
	}

	return
}

func FileWatcher(folder string) {
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Has(fsnotify.Create) {
					log.Println("create file:", event.Name)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(folder)
	if err != nil {
		log.Fatal(err)
	}
}
