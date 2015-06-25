package main

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"os"
)

func main() {
	err := PullRun()
	if err != nil {
		fmt.Println(err)
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
	targets, err := PullTargetsFromConfig()
	if err != nil {
		return err
	}

	for _, target := range targets {
		p := PathComponents(target.File)
		err := Pull(p, target)
		if err != nil {
			return err
		}
	}

	return nil
}

/*
func PushRun() error {
	sources, err := PushSourcesFromConfig()
	if err != nil {
		return err
	}

	for _, source := range sources {
		p := PathComponents(source.File)
		_, err := Push(p, source)
		if err != nil {
			fmt.Println(err)
		}
		// fmt.Println(paths)
	}

	return nil
}
*/
