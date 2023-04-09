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
	ch   chan string
}

func NewReadFromFile(folder string) *ReadFromFile {
	path, err := LatestFile(folder)
	if err != nil {
		panic(err)
	}
	log.Println("listening log file:", path)

	ch := make(chan string)
	FileCreateWatcher(folder, ch)

	return &ReadFromFile{
		path: path,
		ch:   ch,
	}
}

func (r *ReadFromFile) Read(rc chan []byte) {
	f, err := os.Open(r.path)
	if err != nil {
		panic(err)
	}
	r.f = f

	// read line by line starting from the end of the file
	r.f.Seek(0, io.SeekEnd)
	br := bufio.NewReader(r.f)

	for {
		select {
		case newFile := <-r.ch:
			log.Println("new file create:", newFile)
			log.Println("close file:", r.path)
			r.f.Close()
			log.Println("listening new log file:", newFile)
			f, err := os.Open(newFile)
			if err != nil {
				panic(err)
			}
			r.path = newFile
			r.f = f
			br = bufio.NewReader(r.f)
		default:
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
}
