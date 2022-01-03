package utils

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestShouldPanic(t *testing.T) {
	defer func() { recover() }()

	PanicWithError(errors.New("calm down"))

	t.Errorf("should have panicked")
}

func TestShouldNotPanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Errorf("should not panic!")
		}
	}()

	PanicWithError(nil)
}

func TestShouldExitWithMessage(t *testing.T) {
	if os.Getenv("EXIT_TEST") == "true" {
		ExitWithMessage(errors.New("nice error"), "im running on subprocess")
		return
	}

	c := exec.Command(os.Args[0], "-test.run=TestShouldExitWithMessage")
	c.Env = append(os.Environ(), "EXIT_TEST=true")
	err := c.Run()

	_, ok := err.(*exec.ExitError)
	if !ok {
		t.Errorf("OS Exit not performed")
	}
}

func TestShouldNotExit(t *testing.T) {
	ExitWithMessage(nil, "")
}
