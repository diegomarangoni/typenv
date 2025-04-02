package typenv

import (
	"strconv"
)

// Float64 returns given environment variable as float64
func Float64(name string, fallback ...float64) float64 {
	return parse(name, formatFloat64, fallback)
}

func formatFloat64(raw string) (float64, error) {
	return strconv.ParseFloat(raw, 64)
}

// Float32 returns given environment variable as float32
func Float32(name string, fallback ...float32) float32 {
	return parse(name, formatFloat32, fallback)
}

func formatFloat32(raw string) (float32, error) {
	val, err := strconv.ParseFloat(raw, 32)
	if err != nil {
		return 0, err
	}
	return float32(val), nil
}
