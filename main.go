package main

import (
	"fmt"
	"os"

	bserrors "github.com/bugsnag/bugsnag-go/errors"
	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func main() {
	Run()
}

const phraseAppSupport = "support@phraseapp.com"

func Run() {
	var cfg *phraseapp.Config
	defer func() {
		if recovered := recover(); recovered != nil {
			if recovered, ok := recovered.(error); ok && PHRASEAPP_CLIENT_VERSION != "DEV" {
				reportError(bserrors.New(recovered, 3), cfg)
			}
			printError(fmt.Errorf("This should not have happened: %s - Contact support: %s", recovered, phraseAppSupport))
			os.Exit(1)
		}
	}()

	phraseapp.ClientVersion = PHRASEAPP_CLIENT_VERSION
	CheckForUpdate()

	cfg, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	r, err := router(cfg)
	if err != nil {
		printError(err)
		os.Exit(3)
	}

	switch err := r.RunWithArgs(); err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		printError(err)
		os.Exit(1)
	}
}
