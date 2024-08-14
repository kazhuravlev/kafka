package producer

import (
	"context"
	"crypto/tls"
	"github.com/kazhuravlev/just"
	"github.com/pkg/errors"
	"github.com/segmentio/kafka-go"
	"time"
)

type Stats = kafka.WriterStats

type Producer struct {
	writer *kafka.Writer
}

func New(opts Options) (*Producer, error) {
	if err := opts.Validate(); err != nil {
		return nil, errors.Wrap(err, "bad configuration")
	}

	return &Producer{
		writer: &kafka.Writer{ //nolint:exhaustruct
			Addr:                   kafka.TCP(opts.addrs...),
			Topic:                  opts.topic,
			Balancer:               &kafka.Murmur2Balancer{Consistent: true},
			MaxAttempts:            3,    //nolint:gomnd // it is obvious
			BatchSize:              1024, //nolint:gomnd // it is obvious
			Async:                  true,
			BatchBytes:             1024 * 1024, //nolint:gomnd // it is obvious
			BatchTimeout:           time.Second,
			WriteTimeout:           time.Second,
			RequiredAcks:           kafka.RequireOne,
			AllowAutoTopicCreation: false,
			Transport: &kafka.Transport{ //nolint:exhaustruct
				ClientID: opts.clientID,
				TLS:      just.If(opts.tls, &tls.Config{InsecureSkipVerify: true}, nil), //nolint:exhaustruct,gosec
			},
		},
	}, nil
}

func (p Producer) Close(_ context.Context) error {
	if err := p.writer.Close(); err != nil {
		return errors.Wrap(err, "close kafka writer")
	}

	return nil
}

func (p Producer) Stats() Stats {
	return p.writer.Stats()
}

func (p Producer) WriteMessages(ctx context.Context, msgs ...Message) error {
	kafkaMessages := just.SliceMap(msgs, func(msg Message) kafka.Message {
		return kafka.Message{
			Offset:        msg.Offset,
			HighWaterMark: msg.HighWaterMark,
			Key:           msg.Key,
			Value:         msg.Value,
			Headers: just.SliceMap(msg.Headers, func(hdr Header) kafka.Header {
				return kafka.Header{
					Key:   string(hdr.Key),
					Value: hdr.Value,
				}
			}),
		}
	})
	if err := p.writer.WriteMessages(ctx, kafkaMessages...); err != nil {
		return errors.Wrap(err, "write message")
	}

	return nil
}
