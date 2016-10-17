// +build linux darwin

package phraseapp

import (
	"path"
	"os"
)

func defaultConfigDir() string {
	return path.Join(os.Getenv("HomePath"), configName)
}
