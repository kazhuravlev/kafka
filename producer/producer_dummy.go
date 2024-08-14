package producer

import (
	"context"
	"github.com/segmentio/kafka-go"
)

type DummyProducer struct{}

func NewDummy() DummyProducer {
	return DummyProducer{}
}

func (DummyProducer) Stats() Stats { //nolint:funlen
	return Stats{
		Writes:   100, //nolint:gomnd
		Messages: 2,   //nolint:gomnd
		Bytes:    3,   //nolint:gomnd
		Errors:   4,   //nolint:gomnd
		BatchQueueTime: kafka.DurationStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		BatchTime: kafka.DurationStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		WriteTime: kafka.DurationStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		WaitTime: kafka.DurationStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		Retries: 5, //nolint:gomnd
		BatchSize: kafka.SummaryStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		BatchBytes: kafka.SummaryStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		MaxAttempts:     6,  //nolint:gomnd
		WriteBackoffMin: 7,  //nolint:gomnd
		WriteBackoffMax: 8,  //nolint:gomnd
		MaxBatchSize:    9,  //nolint:gomnd
		BatchTimeout:    10, //nolint:gomnd
		ReadTimeout:     11, //nolint:gomnd
		WriteTimeout:    12, //nolint:gomnd
		RequiredAcks:    13, //nolint:gomnd
		Async:           false,
		Topic:           "dummy-1",
		Dials:           14, //nolint:gomnd
		DialTime: kafka.DurationStats{
			Avg:   100, //nolint:gomnd
			Min:   200, //nolint:gomnd
			Max:   300, //nolint:gomnd
			Count: 400, //nolint:gomnd
			Sum:   500, //nolint:gomnd
		},
		Rebalances:        15, //nolint:gomnd
		RebalanceInterval: 16, //nolint:gomnd
		QueueLength:       17, //nolint:gomnd
		QueueCapacity:     18, //nolint:gomnd
		ClientID:          "dummy-2",
	}
}

func (DummyProducer) WriteMessages(ctx context.Context, msgs ...Message) error {
	return nil
}

func (DummyProducer) Close(ctx context.Context) error {
	return nil
}
