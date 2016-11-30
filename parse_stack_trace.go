package main

import (
	"path/filepath"
	"strings"
)

func parseStackTrace(in string) StackTrace {
	stackTrace := StackTrace{RealStack: in}
	var current *StackItem
	for i, line := range strings.Split(in, "\n") {
		if i == 0 {
			continue
		}
		switch i % 2 {
		case 1:
			// function
			current = &StackItem{Method: strings.Split(line, "(")[0]}
		case 0:
			// location
			if current == nil {
				continue
			}
			path, line, ok := pathAndLineNoFromLine(line)
			if !ok {
				continue
			}
			current.AbsolutePath = path
			current.LineNo = line
			current.Name = filepath.Base(path)
			stackTrace.Stack = append(stackTrace.Stack, current)
		}
	}
	return stackTrace
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
