package main

import (
	"log"

	"github.com/sailwith/logcat/internal/app"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	lp := app.NewLogProcess()
	go lp.Process()

	select {}
}
