package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/phrase/phraseapp-client/cli"
	"github.com/phrase/phraseapp-client/internal/print"
	"github.com/phrase/phraseapp-client/internal/updatechecker"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const phraseAppSupport = "support@phrase.com"

var updateChecker = updatechecker.New(
	PHRASEAPP_CLIENT_VERSION,
	filepath.Join(os.TempDir(), ".phraseapp.version"),
	"https://github.com/phrase/phraseapp-client/releases/latest",
	os.Stderr,
)

func main() {
	Run()
}

func Run() {
	var cfg *phraseapp.Config
	defer func() {
		if recovered := recover(); recovered != nil {
			if Debug {
				fmt.Fprintf(os.Stderr, "%v\n%s", recovered, debug.Stack())
			}
			print.Error(fmt.Errorf("This should not have happened: %s - Contact support: %s", recovered, phraseAppSupport))
			os.Exit(1)
		}
	}()

	phraseapp.ClientVersion = PHRASEAPP_CLIENT_VERSION
	updateChecker.Check()

	cfg, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	r, err := router(cfg)
	if err != nil {
		print.Error(err)
		os.Exit(3)
	}

	switch err := r.RunWithArgs(); err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		print.Error(err)
		os.Exit(1)
	}
}
