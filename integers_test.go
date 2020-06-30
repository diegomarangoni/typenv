package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestInt8(t *testing.T) {
	var expected, got int8

	got = typenv.Int8("TEST_INT8")
	if got != expected {
		t.Errorf("Int8 failed, expected `%v` but got `%v`.", expected, got)
	}

	expected = 127
	got = typenv.Int8("TEST_INT8", 127)
	if got != expected {
		t.Errorf("Int8 failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_INT8", "-128")
	expected = -128
	got = typenv.Int8("TEST_INT8", 99)
	if got != expected {
		t.Errorf("Int8 failed, expected `%v` but got `%v`.", expected, got)
	}
}

func TestInt32(t *testing.T) {
	var expected, got int32

	got = typenv.Int32("TEST_INT32")
	if got != expected {
		t.Errorf("Int32 failed, expected `%v` but got `%v`.", expected, got)
	}

	expected = 2147483647
	got = typenv.Int32("TEST_INT32", 2147483647)
	if got != expected {
		t.Errorf("Int32 failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_INT32", "-2147483648")
	expected = -2147483648
	got = typenv.Int32("TEST_INT32", 99)
	if got != expected {
		t.Errorf("Int32 failed, expected `%v` but got `%v`.", expected, got)
	}
}

func TestInt(t *testing.T) {
	var expected, got int

	got = typenv.Int("TEST_INT")
	if got != expected {
		t.Errorf("Int failed, expected `%v` but got `%v`.", expected, got)
	}

	expected = 2147483647
	got = typenv.Int("TEST_INT", 2147483647)
	if got != expected {
		t.Errorf("Int failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_INT", "-2147483648")
	expected = -2147483648
	got = typenv.Int("TEST_INT", 99)
	if got != expected {
		t.Errorf("Int failed, expected `%v` but got `%v`.", expected, got)
	}
}
