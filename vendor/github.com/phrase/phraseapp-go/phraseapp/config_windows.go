// +build linux darwin

package phraseapp

import (
	"os"
	"path/filepath"
)

func defaultConfigDir() string {
	return filepath.Join(os.Getenv("HomePath"), configName)
}
