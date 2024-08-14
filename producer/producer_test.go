package producer_test

import (
	"context"
	"github.com/kazhuravlev/kafka/producer"
)

type IProducer interface {
	Close(ctx context.Context) error
	Stats() producer.Stats
	WriteMessages(ctx context.Context, msgs ...producer.Message) error
}

var (
	_ IProducer = (*producer.Producer)(nil)
	_ IProducer = (*producer.DummyProducer)(nil)
)
