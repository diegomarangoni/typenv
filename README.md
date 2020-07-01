# Typenv

[![go.dev](https://img.shields.io/static/v1?label=go.dev&message=reference&color=00add8)](https://pkg.go.dev/diegomarangoni.dev/typenv) [![go report](https://goreportcard.com/badge/diegomarangoni.dev/typenv)](https://goreportcard.com/report/diegomarangoni.dev/typenv) [![codecov](https://codecov.io/gh/diegomarangoni/typenv/branch/master/graph/badge.svg)](https://codecov.io/gh/diegomarangoni/typenv)

Typenv is a minimalistic, zero dependency, typed environment variables library.
It does support the following types:

* Bool
* Duration
* Float64
* Float32
* Int64
* Int32
* Int8
* Int
* String

# Why typenv?

Handling environment variables in other type than strings or even handling fallback values is not a simple task, typenv makes easy to handle string, boolean, integer, float and time duration environment variables and also allows you to easily set default values for then.

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

## With global default value

You can also set a global default value for a environment.
Useful when you use the same environment with same default value several times.

```go
func init() {
	typenv.SetGlobalDefault(
		typenv.E(typenv.Bool, "DEBUG", true),
	)
}

...

if typenv.Bool("DEBUG") {
	// do something
}
```

If the environment is **NOT** set, the return value will be `true`.

Be aware, you must declare the `SetGlobalDefault` function only once, it will panic if you declare more than once.

## Overriding global default value

You still can override the global default value by setting a local default.

```go
func init() {
	typenv.SetGlobalDefault(
		typenv.E(typenv.Bool, "DEBUG", true),
	)
}

...

if typenv.Bool("DEBUG", false) {
	// do something
}
```

If environment is **NOT** set, the return value will be `false`, even if global default is telling that is `true`.

<!-- go test -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html -->
