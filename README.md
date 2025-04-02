# Typenv

[![go.dev](https://img.shields.io/static/v1?label=go.dev&message=reference&color=00add8)](https://pkg.go.dev/diegomarangoni.dev/typenv)
[![go report](https://goreportcard.com/badge/diegomarangoni.dev/typenv)](https://goreportcard.com/report/diegomarangoni.dev/typenv)
[![codecov](https://codecov.io/gh/diegomarangoni/typenv/branch/master/graph/badge.svg)](https://codecov.io/gh/diegomarangoni/typenv)

Typenv is a minimalistic, zero dependency, typed environment variables library.
It does support the following types:

* time.Duration: `typenv.Duration()`
* string: `typenv.String()`
* float64: `typenv.Float64()`
* float32: `typenv.Float32()`
* int64: `typenv.Int64()`
* int32: `typenv.Int32()`
* int16: `typenv.Int16()`
* int8: `typenv.Int8()`
* int: `typenv.Int()`
* bool: `typenv.Bool()`

# Why `typenv`?

Handling environment variables in types other than strings can be annoying, `typenv` makes it easier and also allows you to set default values for them.

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
