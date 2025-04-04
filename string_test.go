package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestString(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected string

		got := typenv.String("TEST_STRING_UNSET")
		assert.Equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := "something else"

		got := typenv.String("TEST_STRING_UNSET_WITH_DEFAULT", "something else")
		assert.Equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_STRING_SET", "this is my env")
		require.NoError(t, err)

		expected := "this is my env"

		got := typenv.String("TEST_STRING_SET")
		assert.Equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_STRING_SET_WITH_DEFAULT", "this is my env")
		require.NoError(t, err)

		expected := "this is my env"

		got := typenv.String("TEST_STRING_SET_WITH_DEFAULT", "something else")
		assert.Equal(t, expected, got)
	})
}
