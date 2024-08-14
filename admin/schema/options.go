package schema

import (
	"fmt"
	"strconv"
	"time"

	"github.com/kazhuravlev/just"
)

type option func(*Schema) error

// WithCleanupPolicy set cleanup policy.
func WithCleanupPolicy(policy string) option {
	return func(scm *Schema) error {
		switch policy {
		default:
			return fmt.Errorf("unknown cleanup policy: %w", ErrBadConfiguration)
		case "compact,delete", "compact", "delete":
		}

		return WithKV("cleanup.policy", policy)(scm)
	}
}

// WithMinCleanableDirtyRatio допускается минимум (ratio * 100)% дубликатов в логе.
// Влияет на частоту удаления сегментов. Пока порог не превышен - чистка не начнется.
func WithMinCleanableDirtyRatio(ratio float64) option {
	return func(scm *Schema) error {
		if ratio < 0 || ratio > 1 {
			return fmt.Errorf("ratio should be between 0 and 1: %w", ErrBadConfiguration)
		}

		return WithKV("min.cleanable.dirty.ratio", strconv.FormatFloat(ratio, 'f', 5, 64))(scm)
	}
}

// WithRetentionDur сообщение будет точно находиться в логе как минимум dur времени. Далее - может быть удалено.
func WithRetentionDur(dur time.Duration) option {
	return func(scm *Schema) error {
		if dur < 0 {
			return fmt.Errorf("renention duration should be positive: %w", ErrBadConfiguration)
		}

		return WithKV("retention.ms", toMsStr(dur))(scm)
	}
}

// WithDeleteRetentionDur надгробия удаленных сообщений будут находиться в топике как минимум dur.
// Далее - могут быть удалены.
func WithDeleteRetentionDur(dur time.Duration) option {
	return func(scm *Schema) error {
		if dur < 0 {
			return fmt.Errorf("delete renention duration should be positive: %w", ErrBadConfiguration)
		}

		return WithKV("delete.retention.ms", toMsStr(dur))(scm)
	}
}

// WithSegmentDur каждый dur кафка будет принудительно создавать новый сегмент на диске.
// Удаление может производиться только по-сегментно.
func WithSegmentDur(dur time.Duration) option {
	return func(scm *Schema) error {
		if dur < 0 {
			return fmt.Errorf("segment duration should be positive: %w", ErrBadConfiguration)
		}

		return WithKV("segment.ms", toMsStr(dur))(scm)
	}
}

func WithCompressionType(compression string) option {
	return func(scm *Schema) error {
		switch compression {
		default:
			return fmt.Errorf("unknown compression: %w", ErrBadConfiguration)
		case "uncompressed", "zstd", "lz4", "snappy", "gzip", "producer":
		}

		return WithKV("compression.type", compression)(scm)
	}
}

// WithKV проверит только ключ на его "известность" в документации. В случае если вы хотите добавить неизвестную
// опцию - используйте схему FromRawMap.
func WithKV(key, value string) option {
	return func(scm *Schema) error {
		if !just.MapContainsKey(knownOptionsMap, key) {
			return fmt.Errorf("unknown key '%s': %w", key, ErrBadConfiguration)
		}

		scm.set(key, value)

		return nil
	}
}
