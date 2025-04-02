package typenv

import (
	"log"
	"os"
	"time"
)

type types interface {
	int64 | int32 | int16 | int8 | int | float64 | float32 | bool | string | time.Duration
}

func parse[T types](name string, read func(env string) (T, error), defaults ...T) T {
	var zero T

	raw, set := os.LookupEnv(name)
	if !set {
		if len(defaults) > 0 {
			return defaults[0]
		}
		return zero
	}

	val, err := read(raw)
	if err != nil {
		log.Printf(`unable to read environment variable "%s" as type "%T": %s`, name, zero, err)
		if len(defaults) > 0 {
			return defaults[0]
		}
		return zero
	}

	return val
}
