package main

import (
	"fmt"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

func infoCommand() error {
	fmt.Printf("Built at:                            %s\n", BUILT_AT)
	fmt.Printf("PhraseApp Client version:            %s\n", PHRASEAPP_CLIENT_VERSION)
	fmt.Printf("PhraseApp Client revision:           %s\n", REVISION)
	fmt.Printf("PhraseApp Library revision:          %s\n", LIBRARY_REVISION)
	fmt.Printf("PhraseApp Docs revision client:      %s\n", RevisionDocs)
	fmt.Printf("PhraseApp Docs revision lib:         %s\n", phraseapp.RevisionDocs)
	fmt.Printf("PhraseApp Generator revision client: %s\n", RevisionGenerator)
	fmt.Printf("PhraseApp Generator revision lib:    %s\n", phraseapp.RevisionGenerator)
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
