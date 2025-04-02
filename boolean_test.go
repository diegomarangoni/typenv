package typenv_test

import (
	"os"
	"reflect"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestBoolean(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		expected := false

		got := typenv.Bool("TEST_BOOL_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default true", func(t *testing.T) {
		t.Parallel()

		expected := true

		got := typenv.Bool("TEST_BOOL_UNSET_WITH_DEFAULT_TRUE", true)
		assert_equal(t, expected, got)
	})

	t.Run("set to true", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_TRUE", "true")
		require_no_error(t, err)

		expected := true

		got := typenv.Bool("TEST_BOOL_SET_TRUE")
		assert_equal(t, expected, got)
	})

	t.Run("set to true with default false", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_TRUE_DEFAULT_FALSE", "true")
		require_no_error(t, err)

		expected := true

		got := typenv.Bool("TEST_BOOL_SET_TRUE_DEFAULT_FALSE", false)
		assert_equal(t, expected, got)
	})

	t.Run("set to false", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_FALSE", "false")
		require_no_error(t, err)

		expected := false

		got := typenv.Bool("TEST_BOOL_SET_FALSE")
		assert_equal(t, expected, got)
	})

	t.Run("set to false with default true", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_BOOL_SET_FALSE_DEFAULT_FALSE", "false")
		require_no_error(t, err)

		expected := false

		got := typenv.Bool("TEST_BOOL_SET_FALSE_DEFAULT_FALSE", true)
		assert_equal(t, expected, got)
	})
}

func assert_equal[T comparable](t *testing.T, expected T, got T) {
	t.Helper()

	if !reflect.DeepEqual(expected, got) {
		t.Errorf("expected `%v` but got `%v`.", expected, got)
	}
}

func require_no_error(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("expected no error, but got: %s", err)
		t.FailNow()
	}
}
