package typenv_test

import (
	"testing"

	"diegomarangoni.dev/typenv"
)

var expectedglobalstr string = "global foobar"
var expectedglobalint32 int32 = -2147483648

func init() {
	typenv.SetGlobalDefault(
		typenv.E(typenv.String, "TEST_GLOBAL_STRING", expectedglobalstr),
		typenv.E(typenv.Int32, "TEST_GLOBAL_INT32", expectedglobalint32),
	)
}

func TestUnset(t *testing.T) {
	var expected, got string

	expected = "foobar"
	got = typenv.String("TEST_UNSET", expected)
	if got != expected {
		t.Errorf("Unset failed, expected `%s` but got `%s`.", expected, got)
	}
}

func TestStringGlobalDefault(t *testing.T) {
	var expected, got string

	got = typenv.String("TEST_GLOBAL_STRING")
	if got != expectedglobalstr {
		t.Errorf("String global default failed, expected `%s` but got `%s`.", expectedglobalstr, got)
	}

	expected = "new global value"
	got = typenv.String("TEST_GLOBAL_STRING", expected)
	if got != expected {
		t.Errorf("String global default failed, expected `%s` but got `%s`.", expected, got)
	}
}

func TestInt32GlobalDefault(t *testing.T) {
	var expected, got int32

	got = typenv.Int32("TEST_GLOBAL_INT32")
	if got != expectedglobalint32 {
		t.Errorf("Int32 global default failed, expected `%v` but got `%v`.", expectedglobalint32, got)
	}

	expected = 99
	got = typenv.Int32("TEST_GLOBAL_INT32", expected)
	if got != expected {
		t.Errorf("Int32 global default failed, expected `%v` but got `%v`.", expected, got)
	}
}
