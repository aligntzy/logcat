package app

type LogProcess struct {
	RC     chan []byte
	WC     chan []byte
	Reader Reader
	Writer Writer
}

func NewLogProcess() *LogProcess {
	return &LogProcess{
		RC:     make(chan []byte),
		WC:     make(chan []byte),
		Reader: NewReadFromFile(),
		Writer: NewWriteToIM(),
	}
}

func (l *LogProcess) Process() {
	for data := range l.RC {
		l.WC <- data
	}
}
