package main

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/phrase/phraseapp-go/phraseapp"
)

func GetInfo() string {
	info := []string{
		fmt.Sprintf("Built at:                            %s", BUILT_AT),
		fmt.Sprintf("PhraseApp Client version:            %s", PHRASEAPP_CLIENT_VERSION),
		fmt.Sprintf("PhraseApp Client revision:           %s", REVISION),
		fmt.Sprintf("PhraseApp Library revision:          %s", LIBRARY_REVISION),
		fmt.Sprintf("PhraseApp Docs revision client:      %s", RevisionDocs),
		fmt.Sprintf("PhraseApp Docs revision lib:         %s", phraseapp.RevisionDocs),
		fmt.Sprintf("PhraseApp Generator revision client: %s", RevisionGenerator),
		fmt.Sprintf("PhraseApp Generator revision lib:    %s", phraseapp.RevisionGenerator),
		fmt.Sprintf("GoVersion:                           %s", runtime.Version()),
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
	BUILT_AT                 = "LIVE"
	PHRASEAPP_CLIENT_VERSION = "DEV"
)
