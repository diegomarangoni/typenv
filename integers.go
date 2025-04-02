package typenv

import (
	"strconv"
)

// Int64 returns given environment variable as int64
func Int64(name string, fallback ...int64) int64 {
	return parse(name, formatInt64, fallback)
}

func formatInt64(raw string) (int64, error) {
	return strconv.ParseInt(raw, 10, 64)
}

// Int32 returns given environment variable as int32
func Int32(name string, fallback ...int32) int32 {
	return parse(name, formatInt32, fallback)
}

func formatInt32(raw string) (int32, error) {
	val, err := strconv.ParseInt(raw, 10, 36)
	if err != nil {
		return 0, err
	}
	return int32(val), nil
}

// Int16 returns given environment variable as int16
func Int16(name string, fallback ...int16) int16 {
	return parse(name, formatInt16, fallback)
}

func formatInt16(raw string) (int16, error) {
	val, err := strconv.ParseInt(raw, 10, 36)
	if err != nil {
		return 0, err
	}
	return int16(val), nil
}

// Int8 returns given environment variable as int8
func Int8(name string, fallback ...int8) int8 {
	return parse(name, formatInt8, fallback)
}

func formatInt8(raw string) (int8, error) {
	val, err := strconv.ParseInt(raw, 10, 36)
	if err != nil {
		return 0, err
	}
	return int8(val), nil
}

// Int returns given environment variable as int
func Int(name string, fallback ...int) int {
	return parse(name, formatInt, fallback)
}

func formatInt(raw string) (int, error) {
	val, err := strconv.ParseInt(raw, 10, 36)
	if err != nil {
		return 0, err
	}
	return int(val), nil
}
