package typenv

import (
	"log"
	"os"
	"time"
)

type types interface {
	int64 | int32 | int16 | int8 | int | float64 | float32 | bool | string | time.Duration | slices
}

type slices interface {
	[]int64 | []int32 | []int16 | []int8 | []int | []float64 | []float32 | []bool | []string | []time.Duration
}

func parse[T types](name string, format func(raw string) (T, error), fallback []T) T {
	var zero T

	raw, set := os.LookupEnv(name)
	if !set {
		if len(fallback) > 0 {
			return fallback[0]
		}
		return zero
	}

	val, err := format(raw)
	if err != nil {
		log.Printf(`unable to read environment variable "%s" as type "%T": %s`, name, zero, err)
		if len(fallback) > 0 {
			return fallback[0]
		}
		return zero
	}

	return val
}
