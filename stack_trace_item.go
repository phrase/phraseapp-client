package main

import (
	"fmt"
	"strings"
)

type StackTraceItem struct {
	Method       string
	Name         string
	AbsolutePath string
	LineNo       string
}

func (s *StackTraceItem) Line() string {
	return fmt.Sprintf("%s:%s - %s", s.AbsolutePath, s.LineNo, s.Method)
}

func (s *StackTraceItem) ItemContext() string {
	return fmt.Sprintf("%s:%s - %s", s.Name, s.LineNo, s.Method)
}

func (s *StackTraceItem) isVendored() bool {
	return strings.Contains(s.AbsolutePath, "/vendor/") || strings.Contains(s.AbsolutePath, "/Godeps/")
}

func (s *StackTraceItem) isPanic() bool {
	return strings.Contains(s.AbsolutePath, "panic.go")
}

func (s *StackTraceItem) isGoLibFile() bool {
	return strings.Contains(s.AbsolutePath, "/github.com/phrase/phraseapp-go/")
}

func (s *StackTraceItem) isClientFile() bool {
	return strings.Contains(s.AbsolutePath, "/github.com/phrase/phraseapp-client/") && !s.isVendored()
}
