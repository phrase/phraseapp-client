package phraseapp

import (
	"path"
	"os"
)

func defaultConfigDir() string {
	return path.Join(os.Getenv("HOME"), configName)
}
