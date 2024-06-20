package consumer

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"time"
)

//go:generate options-gen -from-struct=Options -defaults-from=var
type Options struct {
	logger    ILogger `validate:"required"`
	name      string  `validate:"required,min=1"`
	tlsConfig *tls.Config
	brokers   []string `validate:"required,min=1"`
	topic     string   `validate:"required,min=1"`
	consGroup string
	minBytes  int
	maxBytes  int
	maxWait   time.Duration
	// consumer group retention time. consumer group will be keep in kafka for this period.
	retentionTime time.Duration
	// 0 - means sync mode.
	commitInterval time.Duration
	decoder        IDecoder
}

var defaultOptions = Options{ //nolint:gochecknoglobals
	logger:         slog.New(slog.NewTextHandler(io.Discard, nil)),
	name:           "unknown-consumer",
	tlsConfig:      nil,
	brokers:        []string{"127.0.0.1:9092"},
	topic:          "unknown-topic",
	consGroup:      "",
	minBytes:       1,
	maxBytes:       1,
	maxWait:        10 * time.Second,
	retentionTime:  24 * time.Hour,
	commitInterval: 0,
	decoder:        new(JSONDecoder),
}

type ILogger interface {
	WarnContext(ctx context.Context, msg string, attrs ...any)
	ErrorContext(ctx context.Context, msg string, attrs ...any)
}

type IDecoder interface {
	Decode(headers []Header, value []byte, dst any) error
}

type JSONDecoder struct{}

func (d *JSONDecoder) Decode(_ []Header, src []byte, res any) error {
	if err := json.Unmarshal(src, res); err != nil {
		return fmt.Errorf("unmarshal json message: %w", err)
	}

	return nil
}
