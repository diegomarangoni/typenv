package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestString(t *testing.T) {
	var expected, got string

	got = typenv.String("TEST_STRING")
	if got != expected {
		t.Errorf("String failed, expected `%s` but got `%s`.", expected, got)
	}

	os.Setenv("TEST_STRING", "foobar")
	expected = "foobar"
	got = typenv.String("TEST_STRING", "overrided foobar")
	if got != expected {
		t.Errorf("String with default failed, expected `%s` but got `%s`.", expected, got)
	}
}
