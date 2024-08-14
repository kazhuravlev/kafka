package schema_test

import (
	"testing"

	"github.com/kazhuravlev/kafka/admin/schema"
	"github.com/stretchr/testify/require"
)

func apply(opt func(schema *schema.Schema) error) (*schema.Schema, error) {
	return schema.New(opt) //nolint:wrapcheck
}

func TestWithCleanupPolicy(t *testing.T) {
	t.Parallel()

	_, err := apply(schema.WithCleanupPolicy("hello"))
	require.Error(t, err)

	_, err = apply(schema.WithCleanupPolicy(""))
	require.Error(t, err)

	_, err = apply(schema.WithCleanupPolicy("compact,delete"))
	require.NoError(t, err)
}

func TestWithKV(t *testing.T) {
	t.Parallel()

	_, err := apply(schema.WithKV("cleanup.policy", "hello"))
	require.NoError(t, err)

	_, err = apply(schema.WithKV("unknown.policy", "unknown-value"))
	require.Error(t, err)

	_, err = apply(schema.WithKV("unknown.policy", ""))
	require.Error(t, err)

	_, err = apply(schema.WithKV("", ""))
	require.Error(t, err)
}
