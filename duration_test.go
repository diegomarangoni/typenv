package typenv_test

import (
	"os"
	"testing"
	"time"

	"diegomarangoni.dev/typenv"
)

func TestDuration(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected time.Duration

		got := typenv.Duration("TEST_DURATION_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := 3 * time.Hour

		got := typenv.Duration("TEST_DURATION_UNSET_WITH_DEFAULT", expected)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_DURATION_SET", "3h")
		require_no_error(t, err)

		expected := 3 * time.Hour

		got := typenv.Duration("TEST_DURATION_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_DURATION_SET_WITH_DEFAULT", "3h")
		require_no_error(t, err)

		expected := 3 * time.Hour

		got := typenv.Duration("TEST_DURATION_SET_WITH_DEFAULT", 10*time.Minute)
		assert_equal(t, expected, got)
	})
}
