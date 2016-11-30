package main

import (
	"fmt"
	"strings"
)

type StackTrace struct {
	Stack     []*StackItem
	RealStack string
}

func NewStackTrace(in string) StackTrace {
	return parseStackTrace(in)
}

func (s *StackTrace) List() (items []string) {
	for _, stackItem := range s.Stack {
		items = append(items, stackItem.Line())
	}
	return items
}

func (s *StackTrace) ErrorLocation() string {
	defaultMessage := "no error location found"

	lastPanicIndex := -1
	for index, item := range s.Stack {
		if item.isPanic() {
			lastPanicIndex = index
		}
	}

	if lastPanicIndex == -1 {
		return defaultMessage
	}

	var errorLocation *StackItem
	for _, item := range s.Stack[lastPanicIndex : len(s.Stack)-1] {
		errorLocation = item

		if item.isGoLibFile() || item.isClientFile() {
			break
		}
	}

	if errorLocation != nil {
		return errorLocation.ItemContext()
	}

	return defaultMessage
}

type StackItem struct {
	Method       string
	Name         string
	AbsolutePath string
	LineNo       string
}

func (s *StackItem) Line() string {
	return fmt.Sprintf("%s:%s - %s", s.AbsolutePath, s.LineNo, s.Method)
}

func (s *StackItem) ItemContext() string {
	return fmt.Sprintf("%s:%s - %s", s.Name, s.LineNo, s.Method)
}

func (s *StackItem) isVendored() bool {
	return strings.Contains(s.AbsolutePath, "/vendor/") || strings.Contains(s.AbsolutePath, "/Godeps/")
}

func (s *StackItem) isPanic() bool {
	return strings.Contains(s.AbsolutePath, "panic.go")
}

func (s *StackItem) isGoLibFile() bool {
	return strings.Contains(s.AbsolutePath, "/github.com/phrase/phraseapp-go/")
}

func (s *StackItem) isClientFile() bool {
	return strings.Contains(s.AbsolutePath, "/github.com/phrase/phraseapp-client/") && !s.isVendored()
}
