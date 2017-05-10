package env_test

import (
	. "github.com/themccallister/env"
	"os"
	"testing"
)

func TestCanGetTheEnvAppModeDefault(t *testing.T) {
	e := Set{}
	if e.AppMode() != "development" {
		t.Fatalf("expected the environment to be development, got %v instead", e.AppMode)
	}
}

func TestCanDefineTheAppModeUsingEnvVars(t *testing.T) {
	e := Set{}
	os.Setenv("APP_MODE", "production")
	if e.AppMode() != "production" {
		t.Fatalf("expected the environment to be production, got %v instead", e.AppMode)
	}
	if e.AppMode() == "development" {
		t.Fatalf("expected the environment to be production, got %v instead", e.AppMode)
	}
}

func TestGetOrWillReturnTheRequestedEnvVarOrTheDefaultThatWasPassed(t *testing.T) {
	e := Set{}
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
	e := Set{DefaultMode: "staging"}
	m := e.AppMode()
	if m != "staging" {
		t.Fatalf("expected the app mode to be staging, got %v instead", m)
	}
}

func TestCanGetTheAppKeyFromEnv(t *testing.T) {
	os.Setenv("APP_KEY", "1234")
	e := Set{}
	k, err := e.AppKey()
	if err != nil {
		t.Fatalf("could not get the AppKey from the Env, got %v instead", err)
	}
	if k != "1234" {
		t.Fatalf("expected the key to not be empty, got %v instead", k)
	}
}

func TestGettingTheAppKeyWithoutTheEnvVarReturnsAnError(t *testing.T) {
	os.Unsetenv("APP_KEY")
	e := Set{}
	_, err := e.AppKey()
	if err == nil {
		t.Fatal("expected an error, but none was returned")
	}
}

func TestDefaultAppKeyCanBeSetAndReturnedIfTheAppKeyIsNotSetInEnv(t *testing.T) {
	os.Unsetenv("APP_KEY")
	e := Set{DefaultKey: "654321"}
	k, err := e.AppKey()
	if err != nil {
		t.Fatalf("an unexpected error occured, %v", err)
	}
	if k != "654321" {
		t.Fatalf("expected the app key to be 654321, got %v instead", k)
	}
}
