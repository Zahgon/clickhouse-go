package tests

import (
	"testing"
)

// AssertIsTimeoutError ensures that the error provided is a timeout error.
// It recursively unwraps the provided error and ensures the core error
// implements the Timeout() method and it returns true.
// context deadline error, os deadline error and poll deadline error each
// implement this and return true.
func AssertIsTimeoutError(t *testing.T, err error) { _ = "STUB: not implemented"; return }

type timeout interface {
	Timeout() bool
}

func isDeadlineExceededError(err error) bool { _ = "STUB: not implemented"; return false }

// unwrap recursively unwraps the error until it gets the core error
func unwrap(err error) error { _ = "STUB: not implemented"; return nil }
