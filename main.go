package main

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"os"
)

func main() {
	err := PullRun()
	if err != nil {
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
	}

	// Run(os.Args[1:]...)
}

func Run(args ...string) {
	callArgs, _ := ConfigCallArgs()
	err := router(callArgs).RunWithArgs()
	switch err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
		os.Exit(1)
	}
}

func PullRun() error {
	config, err := ConfigPushPull()
	if err != nil {
		return err
	}
	targets := config.Phraseapp.Pull.Targets

	for _, target := range targets {
		p := PathComponents(target.File)
		paths, err := PullStrategy(p, target)
		fmt.Println("Error", err)
		fmt.Println("Paths", paths)
	}

	return nil
}

func PushRun() error {
	config, err := ConfigPushPull()
	if err != nil {
		return err
	}
	sources := config.Phraseapp.Push.Sources
	for _, source := range sources {
		p := PathComponents(source.File)
		fmt.Println(p)
	}

	return nil
}
