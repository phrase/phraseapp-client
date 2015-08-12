package cli

// Struct used to configure an action.
type ExampleRunner struct {
	Verbose bool   `cli:"type=opt short=v long=verbose"`               // Flag (boolean option) example. This is either set or not.
	Command string `cli:"type=opt short=c long=command required=true"` // Option that has a default value.
	Hosts   string `cli:"type=arg required=true"`                      // Argument with at least one required.
}

// Run the action. Called when the cli.Run function is called with a route matching the one of this action.
func (er *ExampleRunner) Run() error {
	// Called when action matches route.
	if er.Verbose {
		logger.Printf("Going to execute %q at the following hosts: %v", er.Command, er.Hosts)
	}
	// [..] Executing the SSH command is left to the reader.
	return nil
}

// Basic example that shows how to register an action to a route and execute it.
func Example_basic() {
	router := NewRouter()
	router.Register("run/on/hosts", &ExampleRunner{}, "This is an example that pretends to run a given command on a set of hosts.")
	router.Run("run", "on", "host", "-v", "-c", "uname -a", "192.168.1.1")
	router.RunWithArgs() // Run with args given on the command line.
}
