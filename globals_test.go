package typenv_test

import (
	"testing"

	"diegomarangoni.dev/typenv"
)

func TestSetGlobalDefault(t *testing.T) {
	typenv.SetGlobalDefault()
}

func TestE(t *testing.T) {
	typenv.E(nil, "", nil)
}
