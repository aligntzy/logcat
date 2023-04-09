package app

import (
	"bufio"
	"io"
	"log"
	"os"
	"time"
)

type Reader interface {
	Read(rc chan []byte)
}

type ReadFromFile struct {
	path string
	f    *os.File
}

func NewReadFromFile(folder string) *ReadFromFile {
	path, err := LatestFile(folder)
	if err != nil {
		panic(err)
	}
	log.Println("listening log file:", path)
	// FileWatcher(folder)

	return &ReadFromFile{
		path: path,
	}
}

func (r *ReadFromFile) Read(rc chan []byte) {
	f, err := os.Open(r.path)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// read line by line starting from the end of the file
	f.Seek(0, io.SeekEnd)
	br := bufio.NewReader(f)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			time.Sleep(time.Second)
			continue
		} else if err != nil {
			panic(err)
		}

		rc <- line
	}
}
