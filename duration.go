package typenv

import (
	"time"
)

// Duration returns given environment variable as time.Duration
func Duration(name string, fallback ...time.Duration) time.Duration {
	return parse(name, formatDuration, fallback)
}

func formatDuration(raw string) (time.Duration, error) {
	return time.ParseDuration(raw)
}
