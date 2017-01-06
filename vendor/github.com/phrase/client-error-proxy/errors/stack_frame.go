package errors

import "github.com/bugsnag/bugsnag-go/errors"

// StackFrames is a slice of StackFrame.
type StackFrames []StackFrame

// StackFrame is a JSON-serializable version of
// github.com/bugsnag/bugsnag-go/errors.StackFrame.
type StackFrame struct {
	File           string  `json:"file"`
	LineNumber     int     `json:"line_number"`
	Name           string  `json:"name"`
	Package        string  `json:"package"`
	ProgramCounter uintptr `json:"program_counter"`
}

// ToBugsnagStackFrames converts stackFrames to bugsnag stack frames.
func (stackFrames StackFrames) ToBugsnagStackFrames() []errors.StackFrame {
	stack := make([]errors.StackFrame, len(stackFrames))

	for i, frame := range stackFrames {
		stack[i] = errors.StackFrame{
			File:           frame.File,
			LineNumber:     frame.LineNumber,
			Name:           frame.Name,
			Package:        frame.Package,
			ProgramCounter: frame.ProgramCounter,
		}
	}

	return stack
}

// stackFramesFromBugsnagError is the inverse of StackFrames.ToBugsnagStackFrames().
func stackFramesFromBugsnagError(err *errors.Error) StackFrames {
	stackFrames := err.StackFrames()
	stack := make([]StackFrame, len(stackFrames))

	for i, frame := range stackFrames {
		stack[i] = StackFrame{
			File:           frame.File,
			LineNumber:     frame.LineNumber,
			Name:           frame.Name,
			Package:        frame.Package,
			ProgramCounter: frame.ProgramCounter,
		}
	}

	return stack
}
