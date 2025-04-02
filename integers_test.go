package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestInt64(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected int64

		got := typenv.Int64("TEST_INT64_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := int64(111)

		got := typenv.Int64("TEST_INT64_UNSET_WITH_DEFAULT", 111)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT64_SET", "111")
		require_no_error(t, err)

		expected := int64(111)

		got := typenv.Int64("TEST_INT64_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT64_SET_WITH_DEFAULT", "111")
		require_no_error(t, err)

		expected := int64(111)

		got := typenv.Int64("TEST_INT64_SET_WITH_DEFAULT", 222)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT64_INVALID", "-------")
		require_no_error(t, err)

		var expected int64

		got := typenv.Int64("TEST_INT64_INVALID")
		assert_equal(t, expected, got)
	})
}

func TestInt32(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected int32

		got := typenv.Int32("TEST_INT32_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := int32(111)

		got := typenv.Int32("TEST_INT32_UNSET_WITH_DEFAULT", 111)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT32_SET", "111")
		require_no_error(t, err)

		expected := int32(111)

		got := typenv.Int32("TEST_INT32_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT32_SET_WITH_DEFAULT", "111")
		require_no_error(t, err)

		expected := int32(111)

		got := typenv.Int32("TEST_INT32_SET_WITH_DEFAULT", 222)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT32_INVALID", "-------")
		require_no_error(t, err)

		var expected int32

		got := typenv.Int32("TEST_INT32_INVALID")
		assert_equal(t, expected, got)
	})
}

func TestInt16(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected int16

		got := typenv.Int16("TEST_INT16_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := int16(111)

		got := typenv.Int16("TEST_INT16_UNSET_WITH_DEFAULT", 111)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT16_SET", "111")
		require_no_error(t, err)

		expected := int16(111)

		got := typenv.Int16("TEST_INT16_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT16_SET_WITH_DEFAULT", "111")
		require_no_error(t, err)

		expected := int16(111)

		got := typenv.Int16("TEST_INT16_SET_WITH_DEFAULT", 222)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT16_INVALID", "-------")
		require_no_error(t, err)

		var expected int16

		got := typenv.Int16("TEST_INT16_INVALID")
		assert_equal(t, expected, got)
	})
}

func TestInt8(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected int8

		got := typenv.Int8("TEST_INT8_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := int8(111)

		got := typenv.Int8("TEST_INT8_UNSET_WITH_DEFAULT", 111)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT8_SET", "111")
		require_no_error(t, err)

		expected := int8(111)

		got := typenv.Int8("TEST_INT8_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT8_SET_WITH_DEFAULT", "111")
		require_no_error(t, err)

		expected := int8(111)

		got := typenv.Int8("TEST_INT8_SET_WITH_DEFAULT", 127)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT8_INVALID", "-------")
		require_no_error(t, err)

		var expected int8

		got := typenv.Int8("TEST_INT8_INVALID")
		assert_equal(t, expected, got)
	})

	t.Run("overflow", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT8_OVERFLOW", "128")
		require_no_error(t, err)

		expected := int8(-128)

		got := typenv.Int8("TEST_INT8_OVERFLOW")
		assert_equal(t, expected, got)
	})
}

func TestInt(t *testing.T) {
	t.Parallel()

	t.Run("unset", func(t *testing.T) {
		t.Parallel()

		var expected int

		got := typenv.Int("TEST_INT_UNSET")
		assert_equal(t, expected, got)
	})

	t.Run("unset with default", func(t *testing.T) {
		t.Parallel()

		expected := int(111)

		got := typenv.Int("TEST_INT_UNSET_WITH_DEFAULT", 111)
		assert_equal(t, expected, got)
	})

	t.Run("set", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT_SET", "111")
		require_no_error(t, err)

		expected := int(111)

		got := typenv.Int("TEST_INT_SET")
		assert_equal(t, expected, got)
	})

	t.Run("set with default", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT_SET_WITH_DEFAULT", "111")
		require_no_error(t, err)

		expected := int(111)

		got := typenv.Int("TEST_INT_SET_WITH_DEFAULT", 222)
		assert_equal(t, expected, got)
	})

	t.Run("invalid", func(t *testing.T) {
		t.Parallel()

		err := os.Setenv("TEST_INT_INVALID", "-------")
		require_no_error(t, err)

		var expected int

		got := typenv.Int("TEST_INT_INVALID")
		assert_equal(t, expected, got)
	})
}
