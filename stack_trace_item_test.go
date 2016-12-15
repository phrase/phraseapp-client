package main

import "testing"

func TestStackItem(t *testing.T) {
	stackItem := &StackTraceItem{
		LineNo:       "458",
		Method:       "panic",
		Name:         "runtime/panic.go",
		AbsolutePath: "/usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go",
	}

	expected := "runtime/panic.go:458 - panic"
	if stackItem.ItemContext() != expected {
		t.Fatalf("expected item context to eq\n%q, but was:\n%q", expected, stackItem.ItemContext())
	}

	expected = "/usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go"
	if stackItem.AbsolutePath != expected {
		t.Fatalf("expected absolute path to eq\n%q, but was:\n%q", expected, stackItem.AbsolutePath)
	}

	expected = "458"
	if stackItem.LineNo != expected {
		t.Fatalf("expected line no to eq\n%q, but was:\n%q", expected, stackItem.LineNo)
	}

	expected = "runtime/panic.go"
	if stackItem.Name != expected {
		t.Fatalf("expected name to eq\n%q, but was:\n%q", expected, stackItem.Name)
	}

	expected = "panic"
	if stackItem.Method != expected {
		t.Fatalf("expected method to eq\n%q, but was:\n%q", expected, stackItem.Method)
	}

	expected = "/usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go:458 - panic"
	if stackItem.Line() != expected {
		t.Fatalf("expected line to eq\n%q, but was:\n%q", expected, stackItem.Line())
	}
}

func TestItemContextPathEmpty(t *testing.T) {
	stackItem := &StackTraceItem{Name: "panic"}
	expected := "panic: - "
	if stackItem.ItemContext() != expected {
		t.Fatalf("expected item context to eq\n%q, but was:\n%q", expected, stackItem.ItemContext())
	}
}

func TestItemContextMethodAndPathEmpty(t *testing.T) {
	stackItem := &StackTraceItem{Name: ""}
	expected := ": - "
	if stackItem.ItemContext() != expected {
		t.Fatalf("expected item context to eq\n%q, but was:\n%q", expected, stackItem.ItemContext())
	}
}
