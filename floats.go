package typenv

import (
	"os"
	"strconv"
)

// Float64 returns given registered environment variable and mark as used
func Float64(name string, defaults ...float64) float64 {
	use(name)

	val, set := os.LookupEnv(name)

	if !set {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := floats[name]; ok {
			return global
		}
	}

	f64, err := strconv.ParseFloat(val, 0)

	if nil != err {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := floats[name]; ok {
			return global
		}
	}

	return f64
}

// Float32 returns given registered environment variable and mark as used
func Float32(name string, defaults ...float32) float32 {
	var subdefaults []float64

	if 1 == len(defaults) {
		subdefaults = []float64{float64(defaults[0])}
	}

	f64 := Float64(name, subdefaults...)

	return float32(f64)
}
