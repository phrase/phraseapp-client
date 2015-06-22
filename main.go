package main

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"os"
)

func main() {
	err := PushPullRun()
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

func PushPullRun() error {
	config, err := ConfigPushPull()
	if err != nil {
		return err
	}
	projectId := config.Phraseapp.ProjectId
	accessToken := config.Phraseapp.AccessToken
	sources := config.Phraseapp.Push.Sources
	targets := config.Phraseapp.Pull.Targets

	PrettyPrint(projectId, accessToken, sources, targets)

	for _, source := range sources {
		_, err := FileStrategy(source.File, source.Format)
		if err != nil {
			return err
		}
	}

	for _, target := range targets {
		_, err := FileStrategy(target.File, target.Format)
		if err != nil {
			return err
		}
	}

	return nil
}
