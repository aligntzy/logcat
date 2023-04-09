package app

import "log"

type Notifier interface {
	Notify(data []byte)
}

type IMNotifier struct{}

func NewIMNotifier() *IMNotifier {
	return &IMNotifier{}
}

func (*IMNotifier) Notify(data []byte) {
	log.Println("sending data:", string(data))
}
