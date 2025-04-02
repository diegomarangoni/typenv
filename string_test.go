package typenv_test

import (
	"os"
	"testing"
	"time"

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

func init() {
	var _ time.Duration = typenv.Duration("MY_ENV")
	var _ string = typenv.String("MY_ENV")
	var _ float64 = typenv.Float64("MY_ENV")
	var _ float32 = typenv.Float32("MY_ENV")
	var _ int64 = typenv.Int64("MY_ENV")
	var _ int32 = typenv.Int32("MY_ENV")
	var _ int16 = typenv.Int16("MY_ENV")
	var _ int8 = typenv.Int8("MY_ENV")
	var _ int = typenv.Int("MY_ENV")
	var _ bool = typenv.Bool("MY_ENV")

	var _ []time.Duration = typenv.Slice[[]time.Duration]("MY_ENV", ",")
	var _ []string = typenv.Slice[[]string]("MY_ENV", ",")
	var _ []float64 = typenv.Slice[[]float64]("MY_ENV", ",")
	var _ []float32 = typenv.Slice[[]float32]("MY_ENV", ",")
	var _ []int64 = typenv.Slice[[]int64]("MY_ENV", ",")
	var _ []int32 = typenv.Slice[[]int32]("MY_ENV", ",")
	var _ []int16 = typenv.Slice[[]int16]("MY_ENV", ",")
	var _ []int8 = typenv.Slice[[]int8]("MY_ENV", ",")
	var _ []int = typenv.Slice[[]int]("MY_ENV", ",")
	var _ []bool = typenv.Slice[[]bool]("MY_ENV", ",")
}
