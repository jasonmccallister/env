package env_test

import (
	"os"
	"testing"
	. "github.com/themccallister/env"
)

func TestCanGetTheEnvAppModeDefault(t *testing.T) {
	e := Env{}

	if e.AppMode() != "development" {
		t.Fatalf("expected the environment to be development, got %v instead", e.AppMode)
	}
}

func TestCanDefineTheAppModeUsingEnvVars(t *testing.T) {
	e := Env{}
	os.Setenv("APP_MODE", "production")
	if e.AppMode() != "production" {
		t.Fatalf("expected the environment to be production, got %v instead", e.AppMode)
	}
	if e.AppMode() == "development" {
		t.Fatalf("expected the environment to be production, got %v instead", e.AppMode)
	}
}

func TestGetOrWillReturnTheRequestedEnvVarOrTheDefaultThatWasPassed(t *testing.T) {
	e := Env{}
	notSet := e.GetOr("APP_FAKE_VAR", "https://mccallister.io")
	if notSet != "https://mccallister.io" {
		t.Fatalf("expected the environment var to be development, got %v instead", e.AppMode)
	}

	os.Setenv("APP_FAKE_VAR", "1234")
	set := e.GetOr("APP_FAKE_VAR", "https://www.google.com")
	if set != "1234" {
		t.Fatalf("expected the environment var to be 1234, got %v instead", e.AppMode)
	}
	if set == "https://www.google.com" {
		t.Fatalf("expected the environment var to be 1234, got %v instead", e.AppMode)
	}
}

func TestEnvCanSetADefaultModeToOverrideThePackageDefaultAppMode(t *testing.T) {
	// clear the env to be sure
	os.Unsetenv("APP_MODE")
	e := Env{DefaultMode: "staging"}
	m := e.AppMode()
	if m != "staging" {
		t.Fatalf("expected the app mode to be staging, got %v instead", m)
	}
}
