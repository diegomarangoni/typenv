package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestBoolean(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		expected := false

		got := typenv.Bool("TEST_BOOL_UNSET")
		assert.Equal(t, expected, got)
	})

	t.Run("unset with default true", func(t *testing.T) {
		t.Parallel()

		expected := true

		got := typenv.Bool("TEST_BOOL_UNSET_WITH_DEFAULT_TRUE", true)
		assert.Equal(t, expected, got)
	})

	t.Run("set to true", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_TRUE", "true")
		require.NoError(t, err)

		expected := true

		got := typenv.Bool("TEST_BOOL_SET_TRUE")
		assert.Equal(t, expected, got)
	})

	t.Run("set to true with default false", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_TRUE_DEFAULT_FALSE", "true")
		require.NoError(t, err)

		expected := true

		got := typenv.Bool("TEST_BOOL_SET_TRUE_DEFAULT_FALSE", false)
		assert.Equal(t, expected, got)
	})

	t.Run("set to false", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_FALSE", "false")
		require.NoError(t, err)

		expected := false

		got := typenv.Bool("TEST_BOOL_SET_FALSE")
		assert.Equal(t, expected, got)
	})

	t.Run("set to false with default true", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_FALSE_DEFAULT_FALSE", "false")
		require.NoError(t, err)

		expected := false

		got := typenv.Bool("TEST_BOOL_SET_FALSE_DEFAULT_FALSE", true)
		assert.Equal(t, expected, got)
	})
}
