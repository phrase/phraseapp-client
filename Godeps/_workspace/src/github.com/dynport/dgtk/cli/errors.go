package cli

import (
	"fmt"
)

var (
	ErrorNoRoute       = fmt.Errorf("no route matched")
	ErrorHelpRequested = fmt.Errorf("help requested")
)
