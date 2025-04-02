package typenv

import (
	"strconv"
)

// Float64 returns given registered environment variable and mark as used
func Float64(name string, defaults ...float64) float64 {
	return parse(name, func(env string) (float64, error) {
		return strconv.ParseFloat(env, 64)
	}, defaults...)
}

// Float32 returns given registered environment variable and mark as used
func Float32(name string, defaults ...float32) float32 {
	return parse(name, func(env string) (float32, error) {
		val, err := strconv.ParseFloat(env, 32)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	}, defaults...)
}
