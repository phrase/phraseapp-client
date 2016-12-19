package phraseapp

import (
	"os"
	"path/filepath"
)

func defaultConfigDir() string {
	return filepath.Join(os.Getenv("HOME"), configName)
}
