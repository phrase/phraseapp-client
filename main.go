package main

import (
	"os"

	"fmt"
	"github.com/dynport/dgtk/cli"
)

func main() {
	Run()
}

func Run() {
	callArgs, _ := ConfigCallArgs()
	err := router(callArgs).RunWithArgs()
	switch err {
	case cli.ErrorHelpRequested, cli.ErrorNoRoute:
		os.Exit(1)
	case nil:
		os.Exit(0)
	default:
		printErr(err)
		fmt.Println("HERE")
		os.Exit(1)
	}
}
