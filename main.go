package main

import (
	"fmt"
	"github.com/dynport/dgtk/cli"
	"os"
)

func main() {
	PPRun()
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

func PPRun() {
	p1 := "/Users/sacry1/dev/phrase/**/*"
	p2 := "/Users/sacry1/dev/phrase/phrase/locales/translation_center/*"
	p3 := "/Users/sacry1/dev/phrase/phrase/locales/translation_center/phrase.de.yml"
	ps1, _ := FileStrategy(p1)
	ps2, _ := FileStrategy(p2)
	ps3, _ := FileStrategy(p3)
	fmt.Println("ps1:")
	for _, item := range ps1 {
		fmt.Println("  ", item)
	}
	fmt.Println("ps2:")
	for _, item := range ps2 {
		fmt.Println("  ", item)
	}
	fmt.Println("ps3:")
	for _, item := range ps3 {
		fmt.Println("  ", item)
	}
}
