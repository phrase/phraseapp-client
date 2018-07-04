package main

import (
	"fmt"
	"runtime"
	"strings"
)

func GetInfo() string {
	info := []string{
		fmt.Sprintf("PhraseApp client version:            %s", PHRASEAPP_CLIENT_VERSION),
		fmt.Sprintf("PhraseApp client revision:           %s", REVISION),
		fmt.Sprintf("PhraseApp library revision:          %s", LIBRARY_REVISION),
		fmt.Sprintf("Go version:                          %s", runtime.Version()),
	}
	return fmt.Sprintf("%s\n", strings.Join(info, "\n"))
}

func infoCommand() error {
	fmt.Print(GetInfo())
	return nil
}

var (
	REVISION                 = "DEV"
	LIBRARY_REVISION         = "DEV"
	PHRASEAPP_CLIENT_VERSION = "DEV"
)
