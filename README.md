# env [![Build Status](https://travis-ci.org/themccallister/env.svg?branch=master)](https://travis-ci.org/themccallister/env)
A Go package to make application environments a little more sane.

Golang has a lot of awesome ways to get environment variables from the operating systems. However, one of the languages that I used often (Laravel) has some excellent helpers for grabbing or setting default environment variables. This package is a small helper to reduce the amount of code I write in other applications. Overtime this package will add more options, such as `AppMode`, to set some common practices that I see in the applications I build.

Pull requests are welcome and if you would like to have a discussion please feel free to open an issue!

## Installation

Install env using the command `go get themccallister/env`.

## Common Helpers

There are a few helper methods on Env that are used often, these are `AppMode` and `AppKey`.

`AppMode` is a string that defines the applications mode, such as production, development or staging.

`AppKey` is a random string that can be used for encrypting within the application.

## Documentation

You can create a new env instance in your app like so:

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{}
        // more boilerplate
    }

### Getting the applications "mode"

By default, `env` sets an `AppMode` which allows you to quickly define the application mode (e.g production, staging or development). You can access the `AppMode` like so:

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{}
        if e.AppMode == "development" {
            // do something only for development
        }
    }

### Setting a default app mode for the application

By default, `AppMode` will look for an OS environment variable `APP_MODE`. If `APP_MODE` is not set, it will return a string "development".

However, if you wish to change the default `AppMode`, when creating the env, you can set a field on the Env struct `DefaultMode`. If this is set, calling `AppMode` will return the `DefaultMode` instead of "development".

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{DefaultMode: "production"}
        // the rest of your application
    }

### Getting the applications key

Similar to `AppMode` By default Env sets an `AppKey` which allows you to quickly access the applications key for encryption purposes:

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{}
        k := e.AppKey()
        // do something awesome with the key
    }

You can also set a default, using the following code:

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{DefaultKey: "D035ABB6AC62A811D462D9BD572396E5C52E383699737A9D4B022E3C3B2618CF"}
        // the rest of your application
    }

### Getting an environment variable but setting a default (fallback)

Env has a helper to get an environment variable or fall back to a default you specify. This helper is really just a wrapper around the `env.LookupEnv` to remove some extra code in your application.

Lets say you need to get the applications URL but want to set a default at the same time, you would write something like this:

    package main

    import "github.com/themccallister/env"

    func main() {
        e := env.Env{DefaultMode: "production"}
        url := e.GetOr("APP_URL", "https://mccallister.io")
        // the rest of your application
    }
