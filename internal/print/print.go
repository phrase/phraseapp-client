package print

import (
	"fmt"
	"io"
	"os"

	ct "github.com/daviddengcn/go-colortext"
)

func Success(msg string, args ...interface{}) {
	WithColor(ct.Green, msg, args...)
}

func Failure(msg string, args ...interface{}) {
	WithColor(ct.Red, msg, args...)
}

func WithColor(color ct.Color, msg string, args ...interface{}) {
	fprintWithColor(os.Stdout, color, msg, args...)
}

func Error(err error) {
	fprintWithColor(os.Stderr, ct.Red, "ERROR: %s", err)
}

func fprintWithColor(w io.Writer, color ct.Color, msg string, args ...interface{}) {
	ct.Foreground(color, true)
	fmt.Fprintf(w, msg, args...)
	fmt.Fprintln(w)
	ct.ResetColor()
}
