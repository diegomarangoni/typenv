package typenv

import (
	"os"
	"strconv"
)

// Int64 returns given registered environment variable and mark as used
func Int64(name string, defaults ...int64) int64 {
	val, set := os.LookupEnv(name)

	if !set {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := integers[name]; ok {
			return global
		}
	}

	i64, err := strconv.ParseInt(val, 10, 0)

	if nil != err {
		if 1 == len(defaults) {
			return defaults[0]
		}

		if global, ok := integers[name]; ok {
			return global
		}
	}

	return i64
}

// Int32 returns given registered environment variable and mark as used
func Int32(name string, defaults ...int32) int32 {
	var subdefaults []int64

	if 1 == len(defaults) {
		subdefaults = []int64{int64(defaults[0])}
	}

	i64 := Int64(name, subdefaults...)

	return int32(i64)
}

// Int8 returns given registered environment variable and mark as used
func Int8(name string, defaults ...int8) int8 {
	var subdefaults []int64

	if 1 == len(defaults) {
		subdefaults = []int64{int64(defaults[0])}
	}

	i64 := Int64(name, subdefaults...)

	return int8(i64)
}

// Int returns given registered environment variable and mark as used
func Int(name string, defaults ...int) int {
	var subdefaults []int64

	if 1 == len(defaults) {
		subdefaults = []int64{int64(defaults[0])}
	}

	i64 := Int64(name, subdefaults...)

	return int(i64)
}
