package typenv

import (
	"strings"
)

// Bool returns given environment variable as boolean.
// Returns `true` for if any of options: y, yes, true, 1;
// Returns `false` for anything else.
func Bool(name string, fallback ...bool) bool {
	return parse(name, formatBool, fallback)
}

func formatBool(raw string) (bool, error) {
	switch strings.ToLower(raw) {
	case "y", "yes", "true", "1":
		return true, nil

	default:
		return false, nil
	}
}
