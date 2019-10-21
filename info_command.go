package main

import (
	"fmt"
	"runtime"
	"strings"
)

func GetInfo() string {
	info := []string{
		fmt.Sprintf("Phrase client version:            %s", PHRASE_CLIENT_VERSION),
		fmt.Sprintf("Phrase client revision:           %s", REVISION),
		fmt.Sprintf("Phrase library revision:          %s", LIBRARY_REVISION),
		fmt.Sprintf("Last change at:                   %s", LAST_CHANGE),
		fmt.Sprintf("Go version:                       %s", runtime.Version()),
	}
	return fmt.Sprintf("%s\n", strings.Join(info, "\n"))
}

func infoCommand() error {
	fmt.Print(GetInfo())
	return nil
}

var (
	LAST_CHANGE           = "LIVE"
	REVISION              = "DEV"
	LIBRARY_REVISION      = "DEV"
	PHRASE_CLIENT_VERSION = "DEV"
)
