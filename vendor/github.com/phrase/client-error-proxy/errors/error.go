package errors

import (
	"github.com/bugsnag/bugsnag-go/errors"
)

// Error is a serializable implementation of the
// errors.ErrorWithStackFrames interface.
type Error struct {
	ErrorMessage string      `json:"error_message"`
	Stack        StackFrames `json:"stack"`
	MetaData     MetaData    `json:"meta_data"`
}

// Error method to implement the standard error interface.
func (e *Error) Error() string {
	return e.ErrorMessage
}

// StackFrames returns bugsnag stack frames. With this method (and Error()),
// Error implements bugsnag's errors.ErrorWithStackFrames interface.
func (e *Error) StackFrames() []errors.StackFrame {
	return e.Stack.ToBugsnagStackFrames()
}

// ToBugsnagError returns a bugsnag *Error built from e.
func (e *Error) ToBugsnagError() *errors.Error {
	return errors.New(e, 0)
}

// NewFromBugsnagError creates a new (serializable) *Error from
// a bugsnag *Error.
func NewFromBugsnagError(err *errors.Error) *Error {
	return &Error{
		ErrorMessage: err.Error(),
		Stack:        stackFramesFromBugsnagError(err),
	}
}
