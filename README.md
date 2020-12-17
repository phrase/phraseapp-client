# [Deprecated]
This CLI is deprecated and has been replaced by [CLI v2](https://github.com/phrase/phrase-cli)


# PhraseApp Client (CLI v1)

The PhraseApp Client is available for all major platforms and lets you access all API endpoints as well as easily sync your locale files between your source code and PhraseApp.

Check out our [documentation for more information](https://help.phrase.com/phraseapp-for-developers/phraseapp-client/phraseapp-in-your-terminal).

If you are looking for out newer CLI v2, please go to [https://github.com/phrase/phrase-cli](https://github.com/phrase/phrase-cli)

## Quick Start

This quick start will guide you through the basic steps to get up and running with the PhraseApp Client.

#### 1. Install

[Download and install](https://phrase.com/cli) the client for your platform. See our [detailed installation guide](https://help.phrase.com/phraseapp-for-developers/phraseapp-client/installation) for more information.

##### Homebrew

If you use homebrew, we have provided a tap to make installation easier on Mac OS X:

        brew tap phrase/brewed
        brew install phraseapp

The tap is linked to our Formula collection and will be updated, when you call `brew update` as well.

#### 2. Init

Initialize your project by executing the `init` command. This lets you define your preferred locale file format, source files and more.

    $ cd /path/to/project
    $ phraseapp init

#### 3. Upload your locale files

Use the `push` command to upload your locale files from your defined [sources](https://help.phrase.com/phraseapp-for-developers/phraseapp-client/configuration#push):

    $ phraseapp push

#### 4. Download your locale files

Use the `pull` command to download the most recent locale files back into your project according to your [targets](https://help.phrase.com/phraseapp-for-developers/phraseapp-client/configuration#pull):

    $ phraseapp pull

#### 5. More

To see a list of all available commands, simply execute:

    $ phraseapp

To see all supported options for a command, simple use the `--help` flag:

    $ phraseapp locales list --help

See our [detailed guides](https://help.phrase.com/phraseapp-for-developers/phraseapp-client/phraseapp-in-your-terminal) for in-depth instructions on how to use the PhraseApp Client.

## Contributing

This tool and it's source code are auto-generated from templates that run against a API specification file. Therefore we can not accept any pull requests in this repository. Please use the GitHub Issue Tracker to report bugs.

## Further reading
* [PhraseApp Client Download Page](https://phrase.com/cli)

## Licenses

phraseapp-client is licensed under MIT license. (see [LICENSE](LICENSE))

Parts of phraseapp-client use third party libraries which are vendored and licensed under different licenses:

library | license
---|---
github.com/bgentry/speakeasy | [MIT](https://opensource.org/licenses/mit-license.php) / [Apache 2.0](https://github.com/bgentry/speakeasy/blob/master/LICENSE_WINDOWS)
github.com/daviddengcn/go-colortext | [BSD](https://github.com/daviddengcn/go-colortext/blob/master/LICENSE) / [MIT](https://github.com/daviddengcn/go-colortext/blob/master/LICENSE)
gopkg.in/yaml.v2 | [LGPLv3](https://github.com/go-yaml/yaml/blob/v2/LICENSE)
