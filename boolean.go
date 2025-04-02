package typenv

import (
	"strings"
)

// Bool returns given registered environment variable and mark as used.
// `true` for those values: y, yes, true, 1;
// `false` for anything else.
func Bool(name string, defaults ...bool) bool {
	return parse(name, func(env string) (bool, error) {
		switch strings.ToLower(env) {
		case "y", "yes", "true", "1":
			return true, nil

		default:
			return false, nil
		}
	}, defaults...)
}
