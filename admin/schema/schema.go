package schema

import (
	"errors"
	"fmt"
)

var ErrBadConfiguration = errors.New("bad configuration")

// New creates a new configuration from options with validation.
func New(opts ...option) (*Schema, error) {
	scm := Schema{kv: make(map[string]string, len(opts))}
	for i := range opts {
		if err := opts[i](&scm); err != nil {
			return nil, fmt.Errorf("bad option (by index #%d): %w", i, err)
		}
	}

	return &scm, nil
}

// ServerDefault returns default server-defined schema.
// With this schema you will create a topic with default configuration like with command:
func ServerDefault() *Schema {
	return FromRawMap(nil)
}

// FromRawMap create a config from raw map without any validation.
func FromRawMap(m map[string]string) *Schema {
	return &Schema{kv: m}
}
