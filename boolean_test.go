package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestBoolean(t *testing.T) {
	var expected, got bool

	got = typenv.Bool("TEST_BOOL")
	if got != expected {
		t.Errorf("Bool failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_BOOL", "true")
	expected = true
	got = typenv.Bool("TEST_BOOL", false)
	if got != expected {
		t.Errorf("Bool with default failed, expected `%v` but got `%v`.", expected, got)
	}
}
