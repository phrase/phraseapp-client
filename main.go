package main

import (
	"os"

	"fmt"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

func main() {
	Run()
}

func Run() {
	var cfg *phraseapp.Config
	defer func() {
		if recovery := recover(); recovery != nil {
			ReportError("PhraseApp Client Error", recovery, cfg)
			printErr(fmt.Errorf("This should not have happened: %s - Contact support: support@phraseapp.com", recovery))
			os.Exit(1)
		}
	}()

	phraseapp.ClientVersion = PHRASEAPP_CLIENT_VERSION
	ValidateVersion()

	cfg, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	r, err := router(cfg)
	if err != nil {
		printErr(err)
		os.Exit(3)
	}

	switch err := r.RunWithArgs(); err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		printErr(err)
		os.Exit(1)
	}
}
