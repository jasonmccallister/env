package env

import "os"

// Env is a helper struct to manage application environment configurations.
// It provides a few default helpers for working with the environment vars as
// well as a few helper methods to ask or set default config options.
type Env struct {
	DefaultMode string
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

// GetOr is a helper to get a variable from the environment, but specify a fall back in case the env var is not set.
func (e *Env) GetOr(val string, def string) string {
	variable, exists := os.LookupEnv(val)
	if exists {
		return variable
	}
	return def
}
