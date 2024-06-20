package consumer

import (
	"context"
	"fmt"
	"github.com/segmentio/kafka-go"
)

type Consumer[T any] struct {
	opts Options

	ch chan Msg[T]
}

func New[T any](opts Options) (*Consumer[T], error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("validate options: %w", err)
	}

	return &Consumer[T]{
		opts: opts,
		ch:   make(chan Msg[T]),
	}, nil
}

func (c *Consumer[T]) Run(ctx context.Context) {
	config := kafka.ReaderConfig{ //nolint:exhaustruct
		Brokers:               c.opts.brokers,
		Topic:                 c.opts.topic,
		GroupID:               c.opts.consGroup,
		MinBytes:              c.opts.minBytes,
		MaxBytes:              c.opts.maxBytes,
		MaxWait:               c.opts.maxWait,
		RetentionTime:         c.opts.retentionTime,
		Dialer:                &kafka.Dialer{TLS: c.opts.tlsConfig}, //nolint:exhaustruct
		CommitInterval:        c.opts.commitInterval,
		WatchPartitionChanges: true,
	}

	kfkReader := kafka.NewReader(config)

	go c.runConsumer(ctx, kfkReader)
}

// C return a chan of messages from kafka.
func (c *Consumer[T]) C() <-chan Msg[T] {
	return c.ch
}
