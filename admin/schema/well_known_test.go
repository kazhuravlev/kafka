package schema //nolint:testpackage

import (
	"testing"

	"github.com/kazhuravlev/just"
	"github.com/stretchr/testify/require"
)

func TestWellKnownIntegrity(t *testing.T) {
	t.Parallel()

	knownOptionsLen := len(knownOptions)
	knownOptionsUniqLen := len(just.SliceUniq(knownOptions))

	require.Equal(t, knownOptionsLen, knownOptionsUniqLen, "variable `knownOptions` contains not-uniq values")
}
