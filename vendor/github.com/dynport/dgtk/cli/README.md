# cli

A library to easily create command line interfaces. This is driven by the desire to have a simple but powerful way to
generate those. The existing frameworks failed to deliver features for extended parsing of arguments.

The core ideas of cli are:

* Have routes to actions (like in a web API). There is a fuzzy matching for routes, that allows for giving short cuts
  (like `d s` for `do something`), as long as the given short cuts are unambiguous.
* Actions have parameters like options, flags and arguments. Each action is associated with a struct that has annotated
  fields. The annotations are used to fill the associated fields with the value given on the command line (or by the
  defaults). Action struct must implement the `Runner` interface.
* Options are given via a handle in short or long form (`-c` vs. `--config-file`) and have a value.
* Flags are options with a boolean value, that is set to `true` if the flag is given.
* Arguments are given additionally with out a special handle. This is why order and existence are essential!
* Actions can have hierarchies to facilitate reuse of code.


## Examples

The following example creates a simple CLI action for running commands at a remote host (like `example run on host -c
"uname -a" 192.168.1.1` given the binary has the `example` name and the `uname` command should be executed on
`192.168.1.1`).

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
			log.Printf("Going to execute %q at the following hosts: %v", er.Command, er.Hosts)
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

This example used the long notation of the annotation parser. The following would have the very same effect, but be much
more concise:

	// Struct used to configure an action.
	type ExampleRunner struct {
		Verbose bool   `cli:"opt -v --verbose"`          // Flag (boolean option) example. This is either set or not.
		Command string `cli:"opt -c --command required"` // Option that has a default value.
		Hosts   string `cli:"arg required"`              // Argument with at least one required.
	}

If an action doesn't need parameters it's also possible to directly register a function:

	router.RegisterFunc("do/something", func() error { return fmt.Errorf("I should do something") }, "do something")

If there is only a single action for a program the router is not required and this action can be registered directly:

	cli.RunActionWithArgs(&ExampleRunner{})

If you must know whether an option was set, or not you can use pointer values. Note that this has implications, as values must first be tested for nil (aka the option wasn't given on the command line) and must be dereferenced, to get the actual value. For booleans the flag mechanisms is deactivated, i.e. a value (true or false) must be given.

	// Struct used to configure an action.
	type ExampleRunner struct {
		Opt1 *string `cli:"opt -o"`
		Opt2 *int    `cli:"opt -i"`
		Opt3 *bool   `cli:"opt -t"`
	}

