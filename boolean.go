package typenv

import (
	"os"
)

// Bool returns given registered environment variable and mark as used.
// `true` for those values: y, yes, true, 1;
// `false` for anything else.
func Bool(name string, defaults ...bool) bool {
	val, set := os.LookupEnv(name)

	if !set {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := booleans[name]; ok {
			return global
		}
	}

	switch val {
	case "y", "yes", "true", "1":
		return true

	default:
		return false
	}
}
