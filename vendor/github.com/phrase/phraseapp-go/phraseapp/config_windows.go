// +build linux darwin

package phraseapp

import (
	"os"
	"path/filepath"
)

func defaultConfigDir() string {
	return os.Getenv("HomePath")
}
