package main

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"os"
)

func main() {
	Run(os.Args[1:]...)
}

func Run(args ...string) {
	err := router().Run(args...)
	switch err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		if err.Error() == "required argument not set" {
			fmt.Fprintf(os.Stderr, "ProjectId not set retrying with config...\n")
			callArgs, err := ConfigCallArgs()
			if err != nil {
				fmt.Fprintf(os.Stderr, "ERROR: could not read from config\n")
			}
			args = append(args, callArgs.Phraseapp.ProjectId)
			Run(args...)
		} else {
			fmt.Fprintf(os.Stderr, "ERROR: %s\n", err)
			os.Exit(1)
		}
	}
}
