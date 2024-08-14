package admin

import (
	"fmt"
)

type Client struct {
	opts Options
}

func New(opts Options) (*Client, error) {
	if err := opts.Validate(); err != nil {
		return nil, fmt.Errorf("bad configuration: %w", err)
	}

	return &Client{opts: opts}, nil
}
