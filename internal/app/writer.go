package app

type Writer interface {
	Write(wc chan []byte)
}

type WriteToIM struct {
	Notifier Notifier
}

func NewWriteToIM() *WriteToIM {
	return &WriteToIM{
		Notifier: NewIMNotifier(),
	}
}

func (w *WriteToIM) Write(wc chan []byte) {
	for data := range wc {
		w.Notifier.Notify(data)
	}
}
