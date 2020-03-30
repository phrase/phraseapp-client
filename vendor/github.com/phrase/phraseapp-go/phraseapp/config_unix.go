package phraseapp

import (
	"os"
)

func defaultConfigDir() string {
	return os.Getenv("HOME")
}
