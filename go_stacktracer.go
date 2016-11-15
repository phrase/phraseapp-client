package main

import (
	"errors"
	"fmt"
	"strings"
)

type StackTrace struct {
	Stack     []*StackItem
	RealStack string
}

func NewStackTrace(stack string) StackTrace {
	stackSlice := strings.Split(stack, "\n")
	stackSlice = stackSlice[1:len(stackSlice)]

	stackTrace := StackTrace{
		RealStack: stack,
	}

	for i := range stackSlice {
		stackMethod := strings.TrimSpace(stackSlice[i])

		if stackMethod == "" && i%2 == 1 {
			continue
		}

		var stackPath string
		if i+1 < len(stackSlice) {
			stackPath = strings.TrimSpace(stackSlice[i+1])
		}

		if newStackItem := NewStackItem(stackMethod, stackPath); newStackItem != nil {
			stackTrace.Stack = append(stackTrace.Stack, newStackItem)
		}
	}

	return stackTrace
}

func (s *StackTrace) ErrorList() (errorStrings []string) {
	for _, err := range s.Errors() {
		errorStrings = append(errorStrings, err.Error())
	}
	return errorStrings
}

func (s *StackTrace) ErrorContext() string {
	panicIndex := s.panicIndex()
	if panicIndex == -1 {
		return ""
	}

	if panicIndex+1 < len(s.Stack) {
		panicIndex = panicIndex + 1
	}
	if item := s.Stack[panicIndex]; item != nil {
		return item.ItemContext()
	}
	return ""
}

func (s *StackTrace) Errors() []error {
	var stackErrors []error

	panicIndex := s.panicIndex()
	if panicIndex == -1 {
		return stackErrors
	}

	for _, stackItem := range s.Stack[panicIndex:len(s.Stack)] {
		stackErrors = append(stackErrors, errors.New(stackItem.Line()))
	}

	for i, j := 0, len(stackErrors)-1; i < j; i, j = i+1, j-1 {
		stackErrors[i], stackErrors[j] = stackErrors[j], stackErrors[i]
	}

	return stackErrors
}

func (s *StackTrace) panicIndex() int {
	index := -1
	for i := range s.Stack {
		if strings.Contains(s.Stack[i].Raw, "panic.go") {
			index = i
		}
	}
	return index
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

	stackMethodParts := strings.Split(stackMethod, "(")
	if len(stackMethodParts) > 1 {
		stackMethod = fmt.Sprintf("%s()", stackMethodParts[0])
	}

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
