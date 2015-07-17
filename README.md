# PhraseApp API v2 CLI Client

This is the [PhraseApp API v2](http://docs.phraseapp.com/api/v2) command line
client. It provides a clean command line interface to [PhraseApp](https://phraseapp.com).


## Installation

The PhraseApp CLI tool is provided as a self contained binary for
MacOS, Linux and Windows. It can be downloaded from the [Github release
page](https://github.com/phrase/phraseapp-client/releases). Other platforms
that have a [Go](http://golang.org) port, can build it using `go get
github.com/phrase/phraseapp-client` if the go toolchain is properly installed.

The downloaded binary must be moved to a location, where
it can be found by the system (see [Wikipedia on PATH
variable](http://en.wikipedia.org/wiki/PATH_%28variable%29)). Additionally it is
recommended you rename it to `phraseapp` or `phraseapp.exe` respectively, i.e.
to strip the platform identifier from the name. On unix it might be necessary to
set the executable flag (using `chmod +x phraseapp`).

To test your installation run the following command (note that the output might
differ depending on your version of the binary):

	$ phraseapp help
	Built at 2015-03-31 14:10:54.247597256 +0200 CEST


## Authentication

Every request to the PhraseApp API must be authenticated.
There are two basic mechanisms described in the [API v2
documentation](http://docs.phraseapp.com/api/v2), that both are supported by the
client:

* Using your PhraseApp credentials: Specify your username via the `--username`
	option:

		$ phraseapp show user --username gfrey
		Password:
		
 	As shown in the example the user's password is requested on the command
 	line (there is no facility to supply it via command line intentionally). If
 	multi-factor authentication is activated for the account a token must be
 	provided, too. This is requested via the `--tfa` flag:

		$ phraseapp show user --username gfrey --tfa
		Password:
		TFA-Token:
		
	Now the token must be given, too.

* Using an OAuth token: Those tokens can be created in the frontend and are
	provided via the `--token` option.

		$ phraseapp show user --token <OAuth token>
		
	This is the preferred mode, as it is more secure, as a leaked token can be
	revoked easily and the token can be restricted with respect to access rights.

	
## Configuration

As specifying the authentication information with each and every request is
tedious and error prone the client supports a configuration file. This file is
located at `$HOME/.config/phraseapp/config.json` and may contain the following
keys in a JSON hash:

* `username`: The username used for authentication.
* `tfa`: A flag to indicate whether TFA is required for the account (the default
	is `false` if not given).
* `token`: The user's OAuth token.

An example for username and password based authentication with TFA activated
would be:

	{
		"username": "gfrey",
		"tfa": true
	}
	
With the following configuration file OAuth token based authentication would be
used:

	{
		"token": "mysecretoauthtoken"
	}
	
If both token and username are given in the configuration file, the token will
take precedence and be used. If the configuration file exists, but command line
flags are given, these flags will be used (this helps to test other mechanisms,
accounts or tokens).

All example given in the rest of the documentation assume that there is a
configuration file with valid authentication credentials given, i.e. no options
will be given on the command line.


## Usage

The `phraseapp` tool provides a router to determine all different commands
available, i.e. there is a single binary that does different things depending on
the path given as first arguments (similar to what git does). For example to get
the list of all projects, there is a route `projects list` that will receive this
information:

	$ phraseapp projects list

If no route is given, a list of all available routes will be printed. Each route
can have options that are shown in more detail with the `-h/--help` flag:

	$ phraseapp project update -h
	project update [-h|--help] [--username <Username>] [--token <Token>] [--tfa] [--path <Config>] [--name <Name>] [--shares-translation-memory <SharesTranslationMemory>] <Id>
	  Update an existing project.
	  OPTIONS
	    -h --help                 show help for action
	    --username <Username>     username used for authentication
	    --token <Token>           token used for authentication
	    --tfa                     use Two-Factor Authentication
	    --path <Config>           path to the config file (default: $HOME/.config/phraseapp/config.json)
	    --name <Name>
	    --shares-translation-memory <SharesTranslationMemory>
	
	  ARGUMENTS
	    <Id>

The help shows that the `project update` route requires a single argument `Id`
that identifies the project to update. Additionally there are a set of options.
While the options for authentication and config file are there for every route,
the `--name` and `--shares-translation-memory` are specific to the project
update (used to specify the new name or to specify whether the translation
memory should be shared).


