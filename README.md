# Typenv

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

## With default value

You can fallback to a default value in case the environment variable is not set.

```go
// if env not set, will be true
if typenv.Bool("DEBUG", true) {
	// do something
}
```

In this example, if the environment `DEBUG` is not set, the default value will be `true`.

## With global default value

You can also set a global default value for a environment.
Useful when you use the same environment with same default value several times.

```go
func init() {
	typenv.SetGlobalDefault(
		typenv.E(env.Bool, "DEBUG", true),
	)
}
```

You must declare this function only once, it will panic if you declare more than once.

## Overriding global default value

You still can override the global default value by setting a local default.

```go
func init() {
	typenv.SetGlobalDefault(
		typenv.E(env.Bool, "DEBUG", true),
	)
}

...

// if env not set, will be false
if typenv.Bool("DEBUG", false) {
	// do something
}
```

In this case, if environment is **NOT** set, it will **NOT** get inside the if block even if global default is telling that is `true`.
