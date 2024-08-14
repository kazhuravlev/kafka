package producer

type Header struct {
	Key   []byte
	Value []byte
}

type Message struct {
	Offset        int64
	HighWaterMark int64
	Key           []byte
	Value         []byte
	Headers       []Header
}
