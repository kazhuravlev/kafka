package schema

import (
	"strconv"
	"time"
)

// Schema describe a topic schema. Just keep all topic parameters.
type Schema struct {
	kv map[string]string
}

func (s Schema) KvPairs() map[string]string {
	return s.kv
}

func (s *Schema) set(key string, value string) {
	s.kv[key] = value
}

func toMsStr(dur time.Duration) string {
	return strconv.Itoa(int(dur.Milliseconds()))
}
