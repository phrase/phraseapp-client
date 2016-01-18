// +build linux darwin

package main

import (
	"path"
	"os"
)

func defaultConfigDir() string {
	return path.Join(os.Getenv("HomePath"), configName)
}
