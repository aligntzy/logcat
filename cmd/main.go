package main

import (
	"github.com/aligntzy/logcat/internal/app"
)

func main() {
	lp := app.NewLogProcess()

	go lp.Reader.Read(lp.RC)
	go lp.Process()
	go lp.Writer.Write(lp.WC)

	for {
	}
}
