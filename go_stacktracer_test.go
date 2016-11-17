package main

import "testing"

func TestPanicWithEmptyStack(t *testing.T) {
	stackTrace := NewStackTrace("")
	if stackTrace.RealStack != "" {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "", stackTrace.RealStack)
	}

	expected := ""
	if stackTrace.ErrorLocation() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorLocation())
	}

	if len(stackTrace.List()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.List()))
	}
}

func TestPanicWithoutTrace(t *testing.T) {
	stackTrace := NewStackTrace("panic.go")
	if stackTrace.RealStack != "panic.go" {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "panic.go", stackTrace.RealStack)
	}

	expected := ""
	if stackTrace.ErrorLocation() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorLocation())
	}

	if len(stackTrace.List()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.List()))
	}
}

func TestNewStackItem(t *testing.T) {
	stackItem := &StackItem{
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

func TestPanicPathEmpty(t *testing.T) {
	stackItem := &StackItem{Name: "panic"}
	expected := "panic: - "
	if stackItem.ItemContext() != expected {
		t.Fatalf("expected item context to eq\n%q, but was:\n%q", expected, stackItem.ItemContext())
	}
}

func TestPanicMethodAndPathEmpty(t *testing.T) {
	stackItem := &StackItem{Name: ""}
	expected := ": - "
	if stackItem.ItemContext() != expected {
		t.Fatalf("expected item context to eq\n%q, but was:\n%q", expected, stackItem.ItemContext())
	}
}
