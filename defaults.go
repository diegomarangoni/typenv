package typenv

import (
	"reflect"
	"time"
)

var (
	booleans  map[string]bool
	durations map[string]time.Duration
	floats    map[string]float64
	integers  map[string]int64
	strings   map[string]string
)

var called bool

// SetGlobalDefault register a environment variable global defaul value
func SetGlobalDefault(l ...e) {
	if called {
		panic("function `SetGlobalDefault` already called")
	}

	called = true

	booleans = map[string]bool{}
	durations = map[string]time.Duration{}
	floats = map[string]float64{}
	integers = map[string]int64{}
	strings = map[string]string{}

	for _, i := range l {
		switch reflect.TypeOf(i.fn) {
		case reflect.TypeOf(Bool):
			booleans[i.name] = i.val.(bool)

		case reflect.TypeOf(Duration):
			durations[i.name] = i.val.(time.Duration)

		case reflect.TypeOf(Float64), reflect.TypeOf(Float32):
			switch val := i.val.(type) {
			case float64:
				floats[i.name] = val
			case float32:
				floats[i.name] = float64(val)

			default:
				panic("given global default value must be a valid float type")
			}

		case reflect.TypeOf(Int64), reflect.TypeOf(Int32), reflect.TypeOf(Int8), reflect.TypeOf(Int):
			switch val := i.val.(type) {
			case int64:
				integers[i.name] = val
			case int32:
				integers[i.name] = int64(val)
			case int8:
				integers[i.name] = int64(val)
			case int:
				integers[i.name] = int64(val)

			default:
				panic("given global default value must be a valid integer type")
			}

		case reflect.TypeOf(String):
			strings[i.name] = i.val.(string)
		}
	}
}

type e struct {
	fn   interface{}
	name string
	val  interface{}
}

// E tests if given environment is registered and returns a `e{}`
func E(fn interface{}, name string, val interface{}) e {
	return e{fn: fn, name: name, val: val}
}
