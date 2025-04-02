package typenv

import (
	"strings"
	"time"
)

// Slice returns given environment variable as a slice
func Slice[T slices](name string, separator string, fallback ...T) T {
	return parse(name, func(raw string) (T, error) {
		return formatSlice[T](strings.Split(raw, separator))
	}, fallback)
}

func formatSlice[T slices](raw []string) (T, error) {
	var result T

	switch any(result).(type) {
	case []string:
		result = any(raw).(T)

	case []int64:
		var vals []int64
		for _, str := range raw {
			val, err := formatInt64(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []int32:
		var vals []int32
		for _, str := range raw {
			val, err := formatInt32(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []int16:
		var vals []int16
		for _, str := range raw {
			val, err := formatInt16(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []int8:
		var vals []int8
		for _, str := range raw {
			val, err := formatInt8(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []int:
		var vals []int
		for _, str := range raw {
			val, err := formatInt(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []float64:
		var vals []float64
		for _, str := range raw {
			val, err := formatFloat64(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []float32:
		var vals []float32
		for _, str := range raw {
			val, err := formatFloat32(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []bool:
		var vals []bool
		for _, str := range raw {
			val, _ := formatBool(str)
			vals = append(vals, val)
		}
		result = any(vals).(T)

	case []time.Duration:
		var vals []time.Duration
		for _, str := range raw {
			val, err := formatDuration(str)
			if err != nil {
				return T(nil), err
			}
			vals = append(vals, val)
		}
		result = any(vals).(T)
	}

	return result, nil
}
