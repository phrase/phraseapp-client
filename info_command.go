package main

import (
	"fmt"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
	"strings"
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
	}
	return fmt.Sprintf("%s\n", strings.Join(info, "\n"))
}

func infoCommand() error {
	fmt.Print(GetInfo())
	return nil
}

var (
	REVISION                 string
	LIBRARY_REVISION         string
	BUILT_AT                 string
	PHRASEAPP_CLIENT_VERSION string
)

func init() {
	if PHRASEAPP_CLIENT_VERSION == "" {
		PHRASEAPP_CLIENT_VERSION = "DEV"
	}
	if REVISION == "" {
		REVISION = "DEV"
	}
	if LIBRARY_REVISION == "" {
		LIBRARY_REVISION = "DEV"
	}
	if BUILT_AT == "" {
		BUILT_AT = "LIVE"
	}
}
