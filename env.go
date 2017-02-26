// Package env was created to make application environments a little more sane by allowing passing defaults when asking for environment variables.
package env

import (
	"os"
	"github.com/pkg/errors"
)

// Env is a helper struct to manage application environment configurations.
// It provides a few default helpers for working with the environment vars as
// well as a few helper methods to ask or set default config options.
type Env struct {
	DefaultMode string
	DefaultKey  string
}

// AppMode is the application mode, such as production, staging or development.
// You can define the AppMode by setting the `APP_MODE` environment variable. If you want to set a default mode,
// when you create the new Env struct you can set the DefaultMode field to a string of your choice.
// If the DefaultMode field is not set or if `APP_MODE` is not defined in the environment, the default
// response will be `development`.
func (e *Env) AppMode() string {
	v, ex := os.LookupEnv("APP_MODE")
	if ex {
		return v
	}
	if e.DefaultMode != "" {
		return e.DefaultMode
	}
	return "development"
}

// AppKey is the application key, which should be a random string used for encryption.
// You can define the AppKey by setting the `APP_KEY` environment variable. If you want to set a default key,
// when you create the new Env struct you can set the DefaultKey field to a random string of your choice.
// If the DefaultKey field is not set or if `APP_KEY` is not defined in the environment, an error will be returned.
func (e *Env) AppKey() (string, error) {
	v, ex := os.LookupEnv("APP_KEY")
	if ex {
		return v, nil
	}
	if e.DefaultKey != "" {
		return e.DefaultKey, nil
	}
	return "", errors.New("APP_KEY envrionment variable must be set to run the application")
}

// Helpers

// GetOr is a helper to get a variable from the environment, but specify a fall back in case the env var is not set.
func (e *Env) GetOr(val string, def string) string {
	variable, exists := os.LookupEnv(val)
	if exists {
		return variable
	}
	return def
}
