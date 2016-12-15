package main

import (
	"bufio"
	"bytes"
	"path/filepath"
	"strings"
)

type StackTrace struct {
	Items     []*StackTraceItem
	RealStack []byte
}

func ParseStackTrace(stack []byte) *StackTrace {
	stackTrace := &StackTrace{RealStack: stack}

	scanner := bufio.NewScanner(bytes.NewBuffer(stack))

	for i := 0; scanner.Scan(); i++ {
		if i == 0 {
			continue
		}

		methodLine := scanner.Text()
		if !scanner.Scan() {
			break
		}
		path, lineNo, ok := pathAndLineNoFromLine(scanner.Text())
		if !ok {
			continue
		}

		item := &StackTraceItem{
			Method:       methodLine[:strings.Index(methodLine, "(")],
			AbsolutePath: path,
			LineNo:       lineNo,
			Name:         filepath.Base(path),
		}
		stackTrace.Items = append(stackTrace.Items, item)
	}

	return stackTrace
}

func (s *StackTrace) List() (items []string) {
	for _, stackItem := range s.Items {
		items = append(items, stackItem.Line())
	}
	return items
}

func (s *StackTrace) ErrorLocation() string {
	defaultMessage := "no error location found"

	lastPanicIndex := -1
	for index, item := range s.Items {
		if item.isPanic() {
			lastPanicIndex = index
		}
	}

	if lastPanicIndex == -1 {
		return defaultMessage
	}

	var errorLocation *StackTraceItem
	for _, item := range s.Items[lastPanicIndex : len(s.Items)-1] {
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

func pathAndLineNoFromLine(line string) (string, string, bool) {
	parts := strings.Fields(line)
	if len(parts) < 2 {
		return "", "", false
	}
	parts = strings.SplitN(parts[0], ":", 2)
	if len(parts) != 2 || parts[1] == "" {
		return "", "", false
	}
	return parts[0], parts[1], true
}
