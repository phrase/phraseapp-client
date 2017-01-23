package main

import (
	"fmt"
	"log"
	"os"
	"runtime"

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
			if rc, ok := recovered.(string); ok {
				l := log.New(os.Stderr, "", 0)
				stack := make([]uintptr, bserrors.MaxStackDepth)
				length := runtime.Callers(0, stack[:])
				l.Printf("PANIC: %s", rc)
				for i := 0; i < length; i++ {
					f := runtime.FuncForPC(stack[i])
					file, line := f.FileLine(stack[i])
					l.Printf("%s:%d: %s", file, line, f.Name())
				}
			}
			printError(fmt.Errorf("This should not have happened: %s - Contact support: %s", recovered, phraseAppSupport))
			os.Exit(1)
		}
	}()

	if os.Getenv("DO_PANIC") == "true" {
		panic("testing a panic")
	}

	phraseapp.ClientVersion = PHRASEAPP_CLIENT_VERSION
	CheckForUpdate(os.Stderr)

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
