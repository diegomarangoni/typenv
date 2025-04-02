package typenv

import (
	"strconv"
)

// Int64 returns given registered environment variable and mark as used
func Int64(name string, defaults ...int64) int64 {
	return parse(name, func(env string) (int64, error) {
		return strconv.ParseInt(env, 10, 64)
	}, defaults...)
}

// Int32 returns given registered environment variable and mark as used
func Int32(name string, defaults ...int32) int32 {
	return parse(name, func(env string) (int32, error) {
		val, err := strconv.ParseInt(env, 10, 36)
		if err != nil {
			return 0, err
		}
		return int32(val), nil
	}, defaults...)
}

// Int16 returns given registered environment variable and mark as used
func Int16(name string, defaults ...int16) int16 {
	return parse(name, func(env string) (int16, error) {
		val, err := strconv.ParseInt(env, 10, 36)
		if err != nil {
			return 0, err
		}
		return int16(val), nil
	}, defaults...)
}

// Int8 returns given registered environment variable and mark as used
func Int8(name string, defaults ...int8) int8 {
	return parse(name, func(env string) (int8, error) {
		val, err := strconv.ParseInt(env, 10, 36)
		if err != nil {
			return 0, err
		}
		return int8(val), nil
	}, defaults...)
}

// Int returns given registered environment variable and mark as used
func Int(name string, defaults ...int) int {
	return parse(name, func(env string) (int, error) {
		val, err := strconv.ParseInt(env, 10, 36)
		if err != nil {
			return 0, err
		}
		return int(val), nil
	}, defaults...)
}
