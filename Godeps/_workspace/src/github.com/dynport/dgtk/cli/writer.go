package cli

import (
	"io"
	"os"
)

var DefaultWriter io.Writer = os.Stderr
