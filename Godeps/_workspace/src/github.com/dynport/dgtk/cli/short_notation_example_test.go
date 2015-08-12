package cli

import (
	"log"
)

// Struct used to configure an action.
type ExampleShortNotationRunner struct {
	Verbose bool   `cli:"opt -v --verbose"`          // Flag (boolean option) example. This is either set or not.
	Command string `cli:"opt -c --command required"` // Option that has a default value.
	Hosts   string `cli:"arg required"`              // Argument with at least one required.
}

// Run the action. Called when the cli.Run function is called with a route matching the one of this action.
func (er *ExampleShortNotationRunner) Run() error {
	// Called when action matches route.
	if er.Verbose {
		log.Printf("Going to execute %q at the following hosts: %v", er.Command, er.Hosts)
	}
	// [..] Executing the SSH command is left to the reader.
	return nil
}

// Example that shows how to run a single action without using a router.
func Example_shortNotation() {
	RunAction(&ExampleShortNotationRunner{}, "-v", "-c", "uname -a", "192.168.1.1")
	RunActionWithArgs(&ExampleShortNotationRunner{})
}
