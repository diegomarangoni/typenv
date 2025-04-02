# Typenv

[![go.dev](https://img.shields.io/static/v1?label=go.dev&message=reference&color=00add8)](https://pkg.go.dev/diegomarangoni.dev/typenv)
[![go report](https://goreportcard.com/badge/diegomarangoni.dev/typenv)](https://goreportcard.com/report/diegomarangoni.dev/typenv)
[![codecov](https://codecov.io/gh/diegomarangoni/typenv/branch/master/graph/badge.svg)](https://codecov.io/gh/diegomarangoni/typenv)

`typenv` is a minimalistic, zero-dependency Go library for typed environment variables.  
It eliminates the need for manual type conversion and allows default values for missing environment variables.

## Supported Types

```golang
var _ time.Duration = typenv.Duration("MY_ENV")
var _ string = typenv.String("MY_ENV")
var _ float64 = typenv.Float64("MY_ENV")
var _ float32 = typenv.Float32("MY_ENV")
var _ int64 = typenv.Int64("MY_ENV")
var _ int32 = typenv.Int32("MY_ENV")
var _ int16 = typenv.Int16("MY_ENV")
var _ int8 = typenv.Int8("MY_ENV")
var _ int = typenv.Int("MY_ENV")
var _ bool = typenv.Bool("MY_ENV")
```

And their slice variants:

```golang
var _ []time.Duration = typenv.Slice[[]time.Duration]("MY_ENV", ",")
var _ []string = typenv.Slice[[]string]("MY_ENV", ",")
var _ []float64 = typenv.Slice[[]float64]("MY_ENV", ",")
var _ []float32 = typenv.Slice[[]float32]("MY_ENV", ",")
var _ []int64 = typenv.Slice[[]int64]("MY_ENV", ",")
var _ []int32 = typenv.Slice[[]int32]("MY_ENV", ",")
var _ []int16 = typenv.Slice[[]int16]("MY_ENV", ",")
var _ []int8 = typenv.Slice[[]int8]("MY_ENV", ",")
var _ []int = typenv.Slice[[]int]("MY_ENV", ",")
var _ []bool = typenv.Slice[[]bool]("MY_ENV", ",")
```

# Why `typenv`?

Handling environment variables with types other than strings can be cumbersome.
`typenv` simplifies this process by:

- ✅ Eliminating manual type conversions
- ✅ Supporting slices out of the box
- ✅ Allowing fallback values for missing environment variables
- ✅ Having zero external dependencies

# How to use it

## Basic usage

Go get the library:

```shell
go get diegomarangoni.dev/typenv
```

And use it:

```go
if typenv.Bool("DEBUG") {
	// do something
}
```

If the environment is **NOT** set, the zero value of the type will return, in this case `false`.

## With default value

You can fallback to a default value in case the environment variable is not set.

```go
if typenv.Bool("DEBUG", true) {
	// do something
}
```

If the environment is **NOT** set, the return value will be `true`.
