package typenv

import (
	"time"
)

// Duration returns given registered environment variable and mark as used
func Duration(name string, defaults ...time.Duration) time.Duration {
	return parse(name, func(env string) (time.Duration, error) {
		return time.ParseDuration(env)
	}, defaults...)
}
