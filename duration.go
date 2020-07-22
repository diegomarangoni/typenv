package typenv

import (
	"os"
	"time"
)

// Duration returns given registered environment variable and mark as used
func Duration(name string, defaults ...time.Duration) time.Duration {
	val, set := os.LookupEnv(name)

	if !set && 1 == len(defaults) {
		return defaults[0]
	}

	dur, err := time.ParseDuration(val)

	if nil != err && 1 == len(defaults) {
		return defaults[0]
	}

	return dur
}
