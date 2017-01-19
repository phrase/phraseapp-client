package main

import (
	"bytes"
	"testing"
)

func TestParseStackTraceWithEmptyStack(t *testing.T) {
	stackTrace := ParseStackTrace([]byte{})
	if len(stackTrace.RealStack) != 0 {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "", stackTrace.RealStack)
	}

	expected := "no error location found"
	if stackTrace.ErrorLocation() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorLocation())
	}

	if len(stackTrace.List()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.List()))
	}
}

func TestParseStackTraceWithoutTrace(t *testing.T) {
	stack := []byte("panic.go")
	stackTrace := ParseStackTrace(stack)
	if !bytes.Equal(stackTrace.RealStack, stack) {
		t.Fatalf("expected real stack to eq\n%q, but was:\n%q", "panic.go", stackTrace.RealStack)
	}

	expected := "no error location found"
	if stackTrace.ErrorLocation() != expected {
		t.Fatalf("expected: %q - was: %q", expected, stackTrace.ErrorLocation())
	}

	if len(stackTrace.List()) != 0 {
		t.Fatalf("expected: %d - was: %d", 0, len(stackTrace.List()))
	}
}

func TestParseStackTraceMain(t *testing.T) {
	s := ParseStackTrace(stackInMainFile)
	tests := []struct{ Has, Want interface{} }{
		{len(s.Items), 20},
		{s.ErrorLocation(), "config.go:41 - github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp.ReadConfig"},
		{s.Items[0].Name, "stack.go"},
		{s.Items[0].Method, "runtime/debug.Stack"},
		{s.Items[0].AbsolutePath, "/usr/local/go/src/runtime/debug/stack.go"},
		{s.Items[0].LineNo, "24"},
		{s.Items[1].Name, "error_handler.go"},
		{s.Items[1].Method, "main.createBody"},
		{s.Items[1].AbsolutePath, "/go/src/github.com/phrase/phraseapp-client/error_handler.go"},
		{s.Items[1].LineNo, "73"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i+1, tc.Want, tc.Has)
		}
	}
}

func TestParseStackTraceOther(t *testing.T) {
	s := ParseStackTrace(stackInOtherFile)
	tests := []struct{ Has, Want interface{} }{
		{len(s.Items), 10},
		{s.ErrorLocation(), "info_command.go:38 - main.GetInfo"},
		{s.Items[0].Name, "stack.go"},
		{s.Items[0].Method, "runtime/debug.Stack"},
		{s.Items[0].AbsolutePath, "/path_to_go/1.7.1/libexec/src/runtime/debug/stack.go"},
		{s.Items[0].LineNo, "24"},
		{s.Items[1].Name, "main.go"},
		{s.Items[1].Method, "main.Run.func1"},
		{s.Items[1].AbsolutePath, "/homepath/src/github.com/phrase/phraseapp-client/main.go"},
		{s.Items[1].LineNo, "23"},
	}
	for i, tc := range tests {
		if tc.Has != tc.Want {
			t.Errorf("%d: want=%#v has=%#v", i+1, tc.Want, tc.Has)
		}
	}
}

var stackInMainFile = []byte(`goroutine 1 [running]:
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
  /go/src/github.com/phrase/phraseapp-client/main.go:13 +0x14`)

var stackInOtherFile = []byte(`goroutine 1 [running]:
runtime/debug.Stack(0xc4200519c0, 0x31ca80, 0xc4200101c0)
  /path_to_go/1.7.1/libexec/src/runtime/debug/stack.go:24 +0x79
main.Run.func1(0xc420051f18)
  /homepath/src/github.com/phrase/phraseapp-client/main.go:23 +0x6e
panic(0x31ca80, 0xc4200101c0)
  /path_to_go/1.7.1/libexec/src/runtime/panic.go:458 +0x243
main.GetInfo(0x0, 0xc42000c4a0)
  /homepath/src/github.com/phrase/phraseapp-client/info_command.go:38 +0x34
main.infoCommand(0x0, 0xc4201311d0)
  /homepath/src/github.com/phrase/phraseapp-client/info_command.go:56 +0x26
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*annonymousAction).Run(0xc4201c1260, 0x0, 0x0)
  /homepath/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:69 +0x2a
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*Router).Run(0xc420011780, 0xc42000c490, 0x1, 0x1, 0xc4200cf300, 0xc4200ceda0)
  /homepath/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:54 +0x12a
github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli.(*Router).RunWithArgs(0xc420011780, 0xc420011780, 0x0)
  /homepath/src/github.com/phrase/phraseapp-client/vendor/github.com/dynport/dgtk/cli/router.go:60 +0x74
main.Run()
  /homepath/src/github.com/phrase/phraseapp-client/main.go:47 +0xee
main.main()
  /homepath/src/github.com/phrase/phraseapp-client/main.go:14 +0x14`)
