package typenv

import (
	"os"
)

// String returns given registered environment variable and mark as used
func String(name string, defaults ...string) string {
	val, set := os.LookupEnv(name)

	if !set {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := strings[name]; ok {
			return global
		}
	}

	return val
}
