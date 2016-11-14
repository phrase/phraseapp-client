package main

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
)

type StackTrace struct {
	Stack     []*StackItem
	RealStack string
}

func NewStackTrace() StackTrace {
	stack := string(debug.Stack())
	stackSlice := strings.Split(stack, "\n")

	stackTrace := StackTrace{
		RealStack: stack,
	}

	for i := 0; i < len(stackSlice); i += 2 {
		stackItem := strings.TrimSpace(stackSlice[i])
		if stackItem == "" {
			continue
		}

		var stackMethod string
		if i+1 < len(stackSlice) {
			stackMethod = strings.TrimSpace(stackSlice[i+1])
		}

		if newStackItem := NewStackItem(stackItem, stackMethod); newStackItem != nil {
			stackTrace.Stack = append(stackTrace.Stack, newStackItem)
		}
	}

	return stackTrace
}

func (s *StackTrace) ErrorStrings() (errorStrings []string) {
	for _, err := range s.Errors() {
		errorStrings = append(errorStrings, err.Error())
	}
	return errorStrings
}

func (s *StackTrace) ErrorContext() string {
	for i := len(s.Stack) - 1; i >= 0; i-- {
		stackItem := s.Stack[i]
		if strings.Contains(stackItem.Raw, "panic.go") {
			stackItemBefore := s.Stack[i+1]
			return stackItemBefore.ItemContext()
		}
	}
	return ""
}

func (s *StackTrace) Errors() []error {
	var stackErrors []error

	for i := len(s.Stack) - 1; i >= 0; i-- {
		stackItem := s.Stack[i]
		if strings.Contains(stackItem.Raw, "panic.go") {
			stackItemBefore := s.Stack[i+1]
			stackErrors = append(stackErrors, errors.New(stackItemBefore.Line()))
			for j := i; j > 1; j-- {
				stackErrors = append(stackErrors, errors.New(s.Stack[j].Line()))
			}
		}
	}

	return stackErrors
}

type StackItem struct {
	Raw          string
	Method       string
	Name         string
	AbsolutePath string
	LineNo       string
}

func NewStackItem(rawItem, stackMethod string) *StackItem {
	rawItem = strings.TrimSpace(rawItem)

	var absolutePath string
	var fileName string
	var lineNo string

	tokens := strings.Split(rawItem, " ")
	pathWithLineNo := strings.Split(tokens[0], ":")

	if len(pathWithLineNo) > 0 {
		absolutePath = strings.TrimSpace(pathWithLineNo[0])
		pathTokens := strings.Split(absolutePath, separator)
		fileName = strings.TrimSpace(pathTokens[len(pathTokens)-1])

		if !strings.Contains(fileName, ".go") {
			return nil
		}
	}

	if len(pathWithLineNo) > 1 {
		lineNo = strings.TrimSpace(pathWithLineNo[1])
	}

	stackItem := &StackItem{
		Raw:          rawItem,
		Method:       strings.TrimSpace(stackMethod),
		Name:         fileName,
		AbsolutePath: absolutePath,
		LineNo:       lineNo,
	}

	return stackItem
}

func (s *StackItem) Line() string {
	return fmt.Sprintf("%s:%s - %s", s.AbsolutePath, s.LineNo, s.Method)
}

func (s *StackItem) ItemContext() string {
	return fmt.Sprintf("%s:%s - %s", s.Name, s.LineNo, s.Method)
}
