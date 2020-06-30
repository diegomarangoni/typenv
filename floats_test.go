package typenv_test

import (
	"os"
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestFloat32(t *testing.T) {
	var expected, got float32

	got = typenv.Float32("TEST_FLOAT32")
	if got != expected {
		t.Errorf("Float32 failed, expected `%v` but got `%v`.", expected, got)
	}

	expected = 172948.72918
	got = typenv.Float32("TEST_FLOAT32", 172948.72918)
	if got != expected {
		t.Errorf("Float32 failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_FLOAT32", "-2147483647.873920")
	expected = -2147483647.873920
	got = typenv.Float32("TEST_FLOAT32", 99)
	if got != expected {
		t.Errorf("Float32 failed, expected `%v` but got `%v`.", expected, got)
	}
}
