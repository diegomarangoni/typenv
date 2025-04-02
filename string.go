package typenv

// String returns given registered environment variable and mark as used
func String(name string, defaults ...string) string {
	return parse(name, func(env string) (string, error) {
		return env, nil
	}, defaults...)
}
