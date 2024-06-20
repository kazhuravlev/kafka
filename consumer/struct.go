package consumer

import (
	"github.com/segmentio/kafka-go"
	"time"
)

type Header = kafka.Header

type Msg[T any] struct {
	Partition     int
	Offset        int64
	HighWaterMark int64
	Key           []byte
	Value         T
	Headers       []Header
	Time          time.Time
}
