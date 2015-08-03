# PhraseApp API v2 CLI Client

This is the <a href="http://docs.phraseapp.com/guides/working-with-phraseapp/command-line-client/" target="_blank">command-line-client</a> for the [PhraseApp API v2](http://docs.phraseapp.com/api/v2). It provides a clean command line interface to [PhraseApp](https://phraseapp.com).


## Installing

**[Download the latest release of the PhraseApp API command-line client for your platform.](https://github.com/phrase/phraseapp-client/releases/latest)** 

Here is the official guide for **<a href="http://docs.phraseapp.com/guides/working-with-phraseapp/command-line-client/" target="_blank">working with the command-line-client</a>**

Executables are available for OS X, Linux and Windows. Other platforms that have a [Go](http://golang.org/) port, can build our client from source. If you're using Xcode, Android Studio or Visual Studio, check out our IDE integration plugins for syncing locale files between your repository and PhraseApp.

Store downloaded executable at location, where it can be found by the system (see [PATH variable](http://en.wikipedia.org/wiki/PATH_%28variable%29)). On OS X/Linux it might be necessary to set the executable flag:

	cd path/to/phraseapp/executable && chmod +x phraseapp

## Configuration

Start the setup dialog with:

		$ phraseapp init

You will need an API access token which you can create inside the Translation Center on your User Profile.

The init dialog will guide you through the setup of the CLI client, allow you to select a project from your account, the localization format of your project and allows you to specify the location of your locale files inside your project's codebase.

At the end of the configuration a .phraseapp.yml config file is written which you can later adapt to your needs. You can find more on the configuration file format at the end of this page.

## Configuration file

 The configuration file is located at *.phraseapp.yml* and is used to specify

 * API authentication credentials
 * PhraseApp Project ID
 * Localization format
 * File naming rules for push and pull of locales between your source code and PhraseApp
 * Default settings for API actions

The following is an example for a non-trivial setup with two different projects and localization formats:
<pre>
phraseapp:
	access_token: "3d7e6598d955c45c040459df5692ac4c32a99cbfcab1049f237af1b928a17793"
	project_id: "5c05692ac45c0c32a995c0cbfcab1"
	push:
		sources:
			-
			  file: ./config/locales/<locale_name>.yml
				params:
					file_format: yml
	pull:
		targets:
			-
			  file: ./config/locales/<locale_name>.yml
				params:
					file_format: yml
			-
			  file: ./config/locales/en.stringsdict
				project_id: 2a99cbfcab1049f23
				params:
					locale_id: en
					file_format: stringsdict

	defaults:
		"locales/download":
			file_format: xlf
</pre>

Note: Defaults will only be used, if nothing es is specified for a certain action. Global defaults directly under the *phraseapp* namespace will be overridden by defaults set in push, pull or specific subcommand defaults.

## Usage
Generally the PhraseApp client allows to access all API endpoints specifified for API v2 and responds with JSON or a Status Code on empty responses.

List aall available commands by calling **phraseapp**. This will list all available commands with a short description. You can also call help on any command with **phraseapp locales list --help** to see what parameters may be required.

## Authentication

Every request to the PhraseApp API must be authenticated. There are two methods available for authentication: using your PhraseApp credentials (username and password) or using a OAuth token.

**Authentication using PhraseApp credentials**

Example:

<pre>
$ phraseapp projects list --username user@example.com
Password: ********
</pre>

If two-factor authentication is activated for the account, a extra token must be provided. This is requested via the --tfa flag:

<pre>
$ phraseapp projects list --username user@example.com --tfa
Password: ********
TFA: ********
</pre>

**Authentication via OAuth Token**

Example:

<pre>
$ phraseapp projects list --token 3d7e6598d955c45c040459df5692ac4c32a99cbfcab1049f237af1b928a17793
</pre>


## Further reading
* [Usage Example](http://docs.phraseapp.com/api/v2/examples/)
* [API v2 Specification](http://docs.phraseapp.com/api/v2/)
* [PhraseApp API command-line client on GitHub (Bug reports, Feature requests))](https://github.com/phrase/phraseapp-client)
