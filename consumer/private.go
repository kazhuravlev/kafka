package consumer

import (
	"context"
	"encoding/json"
	"github.com/segmentio/kafka-go"
	"log/slog"
)

func (c *Consumer[T]) runConsumer(ctx context.Context, kfkReader *kafka.Reader) {
	defer close(c.ch)
	defer func() {
		if err := kfkReader.Close(); err != nil {
			c.opts.logger.WarnContext(ctx, "close kafka reader", slog.String("error", err.Error()))
		}
	}()

	for {
		select {
		case <-ctx.Done():
			return
		default:
			message, err := kfkReader.ReadMessage(ctx)
			if err != nil {
				c.opts.logger.ErrorContext(ctx, "read message", slog.String("error", err.Error()))

				break
			}

			var payload T
			if err := json.Unmarshal(message.Value, &payload); err != nil {
				c.opts.logger.ErrorContext(ctx, "unmarshal message",
					slog.String("error", err.Error()),
					slog.String("msg", string(message.Value)))

				continue
			}

			msg := Msg[T]{
				Partition:     message.Partition,
				Offset:        message.Offset,
				HighWaterMark: message.HighWaterMark,
				Key:           message.Key,
				Value:         payload,
				Headers:       message.Headers,
				Time:          message.Time,
			}

			select {
			case <-ctx.Done():
				return
			case c.ch <- msg:
			}
		}
	}
}
