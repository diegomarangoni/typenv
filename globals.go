package typenv

// SetGlobalDefault was used to set a global default for a given environment variable.
//
// Deprecated: Global defaults is a bad idea
func SetGlobalDefault(...any) {}

// E was used to set a global default for a given environment variable.
//
// Deprecated: Global defaults is a bad idea
func E(fn any, name string, val any) any {
	return nil
}
