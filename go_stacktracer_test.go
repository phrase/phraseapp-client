package main

import "testing"

func TestPanicInMain(t *testing.T) {
	stackTrace := NewStackTrace(stackInMainFile)
	if stackTrace.RealStack != stackInMainFile {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", stackInMainFile, stackTrace.RealStack)
	}

	expected := "value.go:228 - reflect.Value.Set()"
	if stackTrace.ErrorContext() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorContext())
	}

	if len(stackTrace.Errors()) != 14 {
		t.Fatalf("expected: %d - was: %d", 14, len(stackTrace.Errors()))
	}
}

func TestPanicInOtherFile(t *testing.T) {
	stackTrace := NewStackTrace(stackInOtherFile)
	if stackTrace.RealStack != stackInOtherFile {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", stackInOtherFile, stackTrace.RealStack)
	}

	expected := "info_command.go:38 - main.infoCommand()"
	if stackTrace.ErrorContext() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorContext())
	}

	if len(stackTrace.Errors()) != 8 {
		t.Fatalf("expected: %d - was: %d", 8, len(stackTrace.Errors()))
	}
}

func TestPanicWithEmptyStack(t *testing.T) {
	stackTrace := NewStackTrace("")
	if stackTrace.RealStack != "" {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "", stackTrace.RealStack)
	}

	expected := ""
	if stackTrace.ErrorContext() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorContext())
	}

	if len(stackTrace.Errors()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.Errors()))
	}
}

func TestPanicWithoutTrace(t *testing.T) {
	stackTrace := NewStackTrace("panic.go")
	if stackTrace.RealStack != "panic.go" {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "panic.go", stackTrace.RealStack)
	}

	expected := ""
	if stackTrace.ErrorContext() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorContext())
	}

	if len(stackTrace.Errors()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.Errors()))
	}
}

var stackInMainFile = `goroutine 1 [running]:
runtime/debug.Stack(0x0, 0x0, 0xc420051480)
  /usr/local/go/src/runtime/debug/stack.go:24 +0x79
main.createBody(0x779a53, 0x16, 0xc420010ac0, 0x34, 0x0, 0x0, 0x0, 0x0, 0xc420051500, 0x8b0a9d, ...)
  /go/src/github.com/phrase/phraseapp-client/error_handler.go:73 +0x26
main.ReportError(0x779a53, 0x16, 0x702820, 0xc4200138d0, 0x0)
  /go/src/github.com/phrase/phraseapp-client/error_handler.go:51 +0x12f
main.Run.func1(0xc420051f18)
  /go/src/github.com/phrase/phraseapp-client/main.go:23 +0x9d
panic(0x702820, 0xc4200138d0)
  /usr/local/go/src/runtime/panic.go:458 +0x243
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.handleErr(0xc420051e88)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/yaml.go:153 +0xe7
panic(0x702820, 0xc4200138d0)
  /usr/local/go/src/runtime/panic.go:458 +0x243
reflect.flag.mustBeAssignable(0x16)
  /usr/local/go/src/reflect/value.go:228 +0x102
reflect.Value.Set(0x71a4c0, 0xc42001b110, 0x16, 0x71a4c0, 0x0, 0x16)
  /usr/local/go/src/reflect/value.go:1327 +0x2f
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).scalar(0xc4200105c0, 0xc42001c4e0, 0x71a4c0, 0xc42001b110, 0x16, 0x71a4c0)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:352 +0x196
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).unmarshal(0xc4200105c0, 0xc42001c4e0, 0x71a4c0, 0xc42001b110, 0x16, 0xc42001b110)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:290 +0x122
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).mappingStruct(0xc4200105c0, 0xc42001c420, 0x71e1a0, 0xc42001b110, 0x19, 0xc420030068)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:635 +0x641
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).mapping(0xc4200105c0, 0xc42001c420, 0x71e1a0, 0xc42001b110, 0x19, 0x71e1a0)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:513 +0xaad
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).unmarshal(0xc4200105c0, 0xc42001c420, 0x71e1a0, 0xc42001b110, 0x19, 0xc42001c420)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:292 +0x216
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).document(0xc4200105c0, 0xc42001c3c0, 0x71e1a0, 0xc42001b110, 0x19, 0x487cb0)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:304 +0x84
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.(*decoder).unmarshal(0xc4200105c0, 0xc42001c3c0, 0x71e1a0, 0xc42001b110, 0x19, 0x5475d5)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/decode.go:280 +0x268
github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml%2ev2.Unmarshal(0xc42007b400, 0x407, 0x607, 0x71e1a0, 0xc42001b110, 0x0, 0x0)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/gopkg.in/yaml.v2/yaml.go:90 +0x2ba
github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp.ReadConfig(0x8, 0x7a0fe8, 0xc420051f18)
  /go/src/github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp/config.go:41 +0xb8
main.Run()
  /go/src/github.com/phrase/phraseapp-client/main.go:33 +0x9e
main.main()
  /go/src/github.com/phrase/phraseapp-client/main.go:13 +0x14`

var stackInOtherFile = `goroutine 1 [running]:
runtime/debug.Stack(0xc4200519c0, 0x31ca80, 0xc4200101c0)
  /usr/local/Cellar/go/1.7.1/libexec/src/runtime/debug/stack.go:24 +0x79
main.Run.func1(0xc420051f18)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/main.go:23 +0x6e
panic(0x31ca80, 0xc4200101c0)
  /usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go:458 +0x243
main.GetInfo(0x0, 0xc42000c4a0)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/info_command.go:38 +0x34
main.infoCommand(0x0, 0xc4201311d0)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/info_command.go:56 +0x26
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*annonymousAction).Run(0xc4201c1260, 0x0, 0x0)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:69 +0x2a
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*Router).Run(0xc420011780, 0xc42000c490, 0x1, 0x1, 0xc4200cf300, 0xc4200ceda0)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:54 +0x12a
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*Router).RunWithArgs(0xc420011780, 0xc420011780, 0x0)
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:60 +0x74
main.Run()
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/main.go:47 +0xee
main.main()
  /Users/sacry1/dev/gows/src/github.com/phrase/phraseapp-client/main.go:14 +0x14`

func TestNewStackItem(t *testing.T) {
	stackItem := NewStackItem("/usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go:458 +0x243", "panic(0x31ca80, 0xc4200101c0)")

	expected := "panic.go:458 - panic()"
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

	expected = "panic.go"
	if stackItem.Name != expected {
		t.Fatalf("expected name to eq\n%q, but was:\n%q", expected, stackItem.Name)
	}

	expected = "panic()"
	if stackItem.Method != expected {
		t.Fatalf("expected method to eq\n%q, but was:\n%q", expected, stackItem.Method)
	}

	expected = "/usr/local/Cellar/go/1.7.1/libexec/src/runtime/panic.go:458 - panic()"
	if stackItem.Line() != expected {
		t.Fatalf("expected line to eq\n%q, but was:\n%q", expected, stackItem.Line())
	}
}
