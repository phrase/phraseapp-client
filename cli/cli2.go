package cli

import (
	"os"
)

// Run the given action with given arguments. This is useful for programms that don't need the routing features, i.e.
// only have one action to run. The example on the short notation uses this feature.
func RunAction(runner Runner, args ...string) (e error) {
	var a *action
	if a, e = newAction("", runner, ""); e != nil {
		return e
	}
	if e = a.parseArgs(args); e != nil {
		a.showHelp()
		return e
	}
	return a.runner.Run()
}

// Run the given action with the arguments given on the command line.
func RunActionWithArgs(runner Runner) (e error) {
	return RunAction(runner, os.Args[1:]...)
}
