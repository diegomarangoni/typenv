package typenv_test

import (
	"testing"
	"time"

	"diegomarangoni.dev/typenv"
)

var (
	expectedstr     string        = "foobar"
	expectedint     int           = 2147483648
	expectedint8    int8          = 127
	expectedint32   int32         = -2147483648
	expectedint64   int64         = 2147483648
	expectedfloat32 float32       = 27837.3457
	expectedfloat64 float64       = 4327837.3457
	expectedbool    bool          = true
	expecteddur     time.Duration = 5 * time.Second
)

func init() {
	typenv.SetGlobalDefault(
		typenv.E(typenv.String, "TEST_GLOBAL_STRING", expectedstr),
		typenv.E(typenv.Int, "TEST_GLOBAL_INT", expectedint),
		typenv.E(typenv.Int8, "TEST_GLOBAL_INT8", expectedint8),
		typenv.E(typenv.Int32, "TEST_GLOBAL_INT32", expectedint32),
		typenv.E(typenv.Int64, "TEST_GLOBAL_INT64", expectedint64),
		typenv.E(typenv.Float32, "TEST_GLOBAL_FLOAT32", expectedfloat32),
		typenv.E(typenv.Float64, "TEST_GLOBAL_FLOAT64", expectedfloat64),
		typenv.E(typenv.Bool, "TEST_GLOBAL_BOOL", expectedbool),
		typenv.E(typenv.Duration, "TEST_GLOBAL_DURATION", expecteddur),
	)
}

func TestUnset(t *testing.T) {
	var expected, got string

	expected = "foobar"
	got = typenv.String("TEST_UNSET", "foobar")
	if got != expected {
		t.Errorf("Unset failed, expected `%s` but got `%s`.", expected, got)
	}
}

func TestStringGlobalDefault(t *testing.T) {
	var expected, got string

	got = typenv.String("TEST_GLOBAL_STRING")
	if got != expectedstr {
		t.Errorf("String global default failed, expected `%s` but got `%s`.", expectedstr, got)
	}

	expected = "new value"
	got = typenv.String("TEST_GLOBAL_STRING", expected)
	if got != expected {
		t.Errorf("String global default failed, expected `%s` but got `%s`.", expected, got)
	}
}

func TestInt32GlobalDefault(t *testing.T) {
	var expected, got int32

	got = typenv.Int32("TEST_GLOBAL_INT32")
	if got != expectedint32 {
		t.Errorf("Int32 global default failed, expected `%v` but got `%v`.", expectedint32, got)
	}

	expected = 99
	got = typenv.Int32("TEST_GLOBAL_INT32", expected)
	if got != expected {
		t.Errorf("Int32 global default failed, expected `%v` but got `%v`.", expected, got)
	}
}
