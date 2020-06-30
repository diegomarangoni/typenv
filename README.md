# Typenv

[![go.dev](https://img.shields.io/static/v1?label=go.dev&message=reference&color=00add8)](https://pkg.go.dev/diegomarangoni.dev/typenv) [![go report](https://goreportcard.com/badge/diegomarangoni.dev/typenv)](https://goreportcard.com/report/diegomarangoni.dev/typenv) [![Coverage Status](https://coveralls.io/repos/github/diegomarangoni/typenv/badge.svg?branch=master)](https://coveralls.io/github/diegomarangoni/typenv?branch=master)

Typenv is a minimalistic environment variables library with types.
You can simple get a value from environment variables as the type you need.
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

// if env not set, will be false
if typenv.Bool("DEBUG", false) {
	// do something
}
```

If environment is **NOT** set, the return value will be `false`, even if global default is telling that is `true`.

<!-- go test -coverprofile coverage.out && go tool cover -html=coverage.out -o coverage.html -->
