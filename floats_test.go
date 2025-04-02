package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestFloat64(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected float64

		got := typenv.Float64("TEST_FLOAT64_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := float64(111.11)

		got := typenv.Float64("TEST_FLOAT64_UNSET_WITH_DEFAULT", 111.11)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT64_SET", "111.11")
		require_no_error(t, err)

		expected := float64(111.11)

		got := typenv.Float64("TEST_FLOAT64_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT64_SET_WITH_DEFAULT", "111.11")
		require_no_error(t, err)

		expected := float64(111.11)

		got := typenv.Float64("TEST_FLOAT64_SET_WITH_DEFAULT", 222.22)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT64_INVALID", "-------")
		require_no_error(t, err)

		var expected float64

		got := typenv.Float64("TEST_FLOAT64_INVALID")
		assert_equal(t, expected, got)
	})
}

func TestFloat32(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected float32

		got := typenv.Float32("TEST_FLOAT32_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := float32(111.11)

		got := typenv.Float32("TEST_FLOAT32_UNSET_WITH_DEFAULT", 111.11)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT32_SET", "111.11")
		require_no_error(t, err)

		expected := float32(111.11)

		got := typenv.Float32("TEST_FLOAT32_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT32_SET_WITH_DEFAULT", "111.11")
		require_no_error(t, err)

		expected := float32(111.11)

		got := typenv.Float32("TEST_FLOAT32_SET_WITH_DEFAULT", 222.22)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_FLOAT32_INVALID", "-------")
		require_no_error(t, err)

		var expected float32

		got := typenv.Float32("TEST_FLOAT32_INVALID")
		assert_equal(t, expected, got)
	})
}
