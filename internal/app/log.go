package app

import (
	"log"

	"github.com/aligntzy/logcat/pkg/config"
)

type LogProcess struct {
	RC      chan []byte
	WC      chan []byte
	Readers []Reader
	Writer  Writer
}

func NewLogProcess() *LogProcess {
	var readers []Reader
	folder := config.C.Path
	for _, v := range folder {
		log.Println("listening log file directory:", v)
		readers = append(readers, NewReadFromFile(v))
	}

	return &LogProcess{
		RC:      make(chan []byte),
		WC:      make(chan []byte),
		Readers: readers,
		Writer:  NewWriteToIM(),
	}
}

func (l *LogProcess) Process() {
	go l.Writer.Write(l.WC)
	for _, reader := range l.Readers {
		go reader.Read(l.RC)
	}

	for data := range l.RC {
		l.WC <- data
	}
}
