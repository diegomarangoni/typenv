package typenv

// String returns given environment variable as string
func String(name string, fallback ...string) string {
	return parse(name, func(raw string) (string, error) {
		return raw, nil
	}, fallback)
}
