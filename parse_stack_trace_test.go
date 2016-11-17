package main

import "testing"

func TestParseStackTrace(t *testing.T) {
	s := parseStackTrace(stackInMainFile)
	if v, x := len(s.Stack), 20; x != v {
		t.Fatalf(`expected len(s.Stack) to be %d, was %d`, x, v)
	}
	tests := []struct{ Has, Want interface{} }{
		{len(s.Stack), 20},
		{s.Stack[0].Name, "stack.go"},
		{s.Stack[0].Method, "runtime/debug.Stack"},
		{s.Stack[0].AbsolutePath, "/usr/local/go/src/runtime/debug/stack.go"},
		{s.Stack[0].LineNo, "24"},
		{s.Stack[1].Name, "error_handler.go"},
		{s.Stack[1].Method, "main.createBody"},
		{s.Stack[1].AbsolutePath, "/go/src/github.com/phrase/phraseapp-client/error_handler.go"},
		{s.Stack[1].LineNo, "73"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i+1, tc.Want, tc.Has)
		}
	}
}
