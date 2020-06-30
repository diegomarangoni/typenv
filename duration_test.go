package typenv_test

import (
	"os"
	"testing"
	"time"

	"diegomarangoni.dev/typenv"
)

func TestDuration(t *testing.T) {
	var expected, got time.Duration

	got = typenv.Duration("TEST_DURATION")
	if got != expected {
		t.Errorf("Duration failed, expected `%v` but got `%v`.", expected, got)
	}

	expected = 5 * time.Second
	got = typenv.Duration("TEST_DURATION", 5*time.Second)
	if got != expected {
		t.Errorf("Duration failed, expected `%v` but got `%v`.", expected, got)
	}

	os.Setenv("TEST_DURATION", "10h")
	expected = 10 * time.Hour
	got = typenv.Duration("TEST_DURATION", 10*time.Hour)
	if got != expected {
		t.Errorf("Duration failed, expected `%v` but got `%v`.", expected, got)
	}
}
