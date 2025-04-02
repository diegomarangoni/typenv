package typenv_test

import (
	"os"
	"testing"
	"time"

	"diegomarangoni.dev/typenv"
	"github.com/stretchr/testify/assert"
)

func TestSlice(t *testing.T) {
	t.Parallel()

	t.Run("string", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_STRING", "my env,another env")

		expected := []string{"my env", "another env"}

		got := typenv.Slice[[]string]("TEST_SLICE_STRING", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int64", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT64", "123,456")

		expected := []int64{123, 456}

		got := typenv.Slice[[]int64]("TEST_SLICE_INT64", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int64 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT64_ERROR", "123,not int")

		var expected []int64

		got := typenv.Slice[[]int64]("TEST_SLICE_INT64_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int32", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT32", "123,456")

		expected := []int32{123, 456}

		got := typenv.Slice[[]int32]("TEST_SLICE_INT32", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int32 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT32_ERROR", "123,not int")

		var expected []int32

		got := typenv.Slice[[]int32]("TEST_SLICE_INT32_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int16", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT16", "123,456")

		expected := []int16{123, 456}

		got := typenv.Slice[[]int16]("TEST_SLICE_INT16", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int16 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT16_ERROR", "123,not int")

		var expected []int16

		got := typenv.Slice[[]int16]("TEST_SLICE_INT16_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int8", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT8", "-128,127")

		expected := []int8{-128, 127}

		got := typenv.Slice[[]int8]("TEST_SLICE_INT8", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int8 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT8_ERROR", "123,not int")

		var expected []int8

		got := typenv.Slice[[]int8]("TEST_SLICE_INT8_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT", "123,456")

		expected := []int{123, 456}

		got := typenv.Slice[[]int]("TEST_SLICE_INT", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("int error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_INT_ERROR", "123,not int")

		var expected []int

		got := typenv.Slice[[]int]("TEST_SLICE_INT_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("float64", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_FLOAT64", "1.11,2.22")

		expected := []float64{1.11, 2.22}

		got := typenv.Slice[[]float64]("TEST_SLICE_FLOAT64", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("float64 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_FLOAT64_ERROR", "1.11,not float")

		var expected []float64

		got := typenv.Slice[[]float64]("TEST_SLICE_FLOAT64_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("float32", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_FLOAT32", "1.11,2.22")

		expected := []float32{1.11, 2.22}

		got := typenv.Slice[[]float32]("TEST_SLICE_FLOAT32", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("float32 error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_FLOAT32_ERROR", "1.11,not float")

		var expected []float32

		got := typenv.Slice[[]float32]("TEST_SLICE_FLOAT32_ERROR", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("bool", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_BOOL", "true,no,yes")

		expected := []bool{true, false, true}

		got := typenv.Slice[[]bool]("TEST_SLICE_BOOL", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("duration", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_DURATION", "10s,2h30m")

		expected := []time.Duration{
			10 * time.Second,
			2*time.Hour + 30*time.Minute,
		}

		got := typenv.Slice[[]time.Duration]("TEST_SLICE_DURATION", ",")
		assert.Equal(t, expected, got)
	})

	t.Run("duration error", func(t *testing.T) {
		t.Parallel()

		os.Setenv("TEST_SLICE_DURATION_ERROR", "10s,not duration")

		var expected []time.Duration

		got := typenv.Slice[[]time.Duration]("TEST_SLICE_DURATION_ERROR", ",")
		assert.Equal(t, expected, got)
	})
}
