package logging

import (
	"os"
)

// NewEnvironmentError returns an *EnvironmentError related to
// the given environment variable.
//
// The EnvironmentError wraps os.ErrNotExist for use with errors.Is()
func NewEnvironmentError(name string) *EnvironmentError {
	return &EnvironmentError{
		VarName: name,
		Err:     os.ErrNotExist,
	}
}

// EnvironmentError records an error related to an environment variable.
type EnvironmentError struct {

	// name of the environment variable that was requested
	VarName string

	// using wrapped error to work with errors.Is()
	Err error `default:"os.ErrNotExist"`
}

func (m *EnvironmentError) Is(target error) bool { return target == os.ErrNotExist }

func (e *EnvironmentError) Error() string { return "environment variable not found" + ": " + e.VarName }

func (e *EnvironmentError) Unwrap() error { return e.Err }

// Timeout reports whether this error represents a timeout.
// Likely not used in EnvironmentError, but available
// for completeness and testing.
func (e *EnvironmentError) Timeout() bool {
	t, ok := e.Err.(timeout)
	return ok && t.Timeout()
}
