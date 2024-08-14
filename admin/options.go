package admin

import (
	"crypto/tls"
)

//go:generate options-gen -from-struct=Options -defaults-from=var
type Options struct {
	bootstrapServers []string `validate:"required,min=1"`
	tlsConfig        *tls.Config
}

var defaultOptions = Options{ //nolint:gochecknoglobals
	bootstrapServers: []string{"127.0.0.1:9092"},
	tlsConfig:        nil,
}
