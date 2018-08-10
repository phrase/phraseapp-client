package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime/debug"

	"github.com/phrase/phraseapp-client/internal/print"
	"github.com/phrase/phraseapp-client/internal/updatechecker"
	"github.com/phrase/phraseapp-go/phraseapp"
	"github.com/urfave/cli"
)

const phraseAppSupport = "support@phraseapp.com"

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
	defer func() {
		if recovered := recover(); recovered != nil {
			if Debug {
				fmt.Fprintf(os.Stderr, "%v\n%s", recovered, debug.Stack())
			}
			print.Error(fmt.Errorf("This should not have happened: %s - Contact support: %s", recovered, phraseAppSupport))
			os.Exit(1)
		}
	}()

	updateChecker.Check()
	_, err := phraseapp.ReadConfig()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s\n", err)
		os.Exit(2)
	}

	app := cli.NewApp()
	app.Version = PHRASEAPP_CLIENT_VERSION
	app.Commands = CLICommands
	err = app.Run(os.Args)
	if err != nil {
		print.Error(err)
		os.Exit(3)
	}

	// r, err := router(cfg)
	// if err != nil {
	// 	print.Error(err)
	// 	os.Exit(3)
	// }

	// switch err := r.RunWithArgs(); err {
	// case cli.ErrorHelpRequested, cli.ErrorNoRoute:
	// 	os.Exit(1)
	// case nil:
	// 	os.Exit(0)
	// default:
	// 	print.Error(err)
	// 	os.Exit(1)
	// }
}
