# PhraseApp Client

The PhraseApp Client is available for all major platforms and lets you access all API endpoints as well as easily sync your locale files between your source code and PhraseApp.

Check out our [documentation for more information](http://docs.phraseapp.com/developers/cli/).

## Quick Start

This quick start will guide you through the basic steps to get up and running with the PhraseApp Client.

#### 1. Install

[Download and install](https://phraseapp.com/cli){:target="_blank"} the client for your platform. See our [detailed installation guide](http://docs.phraseapp.com/developers/cli/installation#download) for more information.

#### 2. Init

[Initialize your project](http://docs.phraseapp.com/developers/cli/installation#initialization) by executing the `init` command. This lets you define your preferred locale file format, source files and more.

    $ cd /path/to/project
    $ phraseapp init

#### 3. Upload your locale files

Use the `push` command to upload your locale files from your defined [sources](http://docs.phraseapp.com/developers/cli/configuration#sources):

    $ phraseapp push

#### 4. Download your locale files

Use the `pull` command to download the most recent locale files back into your project according to your [targets](http://docs.phraseapp.com/developers/cli/configuration#targets):

    $ phraseapp pull

#### 5. More

To see a list of all available commands, simply execute:

    $ phraseapp

To see all supported options for a command, simple use the `--help` flag:

    $ phraseapp locales list --help

See our [detailed guides](http://docs.phraseapp.com/developers/cli/) for in-depth instructions on how to use the PhraseApp Client.

## Further reading
* [PhraseApp Client Download Page](https://phraseapp.com/cli)
