package main

import (
	"log"

	"github.com/aligntzy/logcat/internal/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lp := app.NewLogProcess()
	go lp.Process()

	for {
	}
}
