package main

import (
	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func ApplyNonRestRoutes(r *cli.Router, cfg *phraseapp.Config) {
	r.Register("pull", &PullCommand{Config: cfg}, "Download locales from your PhraseApp project.\n  You can provide parameters supported by the locales#download endpoint http://docs.phraseapp.com/api/v2/locales/#download\n  in your configuration (.phraseapp.yml) for each source.\n  See our configuration guide for more information http://docs.phraseapp.com/developers/cli/configuration/")

	r.Register("push", &PushCommand{Config: cfg}, "Upload locales to your PhraseApp project.\n  You can provide parameters supported by the uploads#create endpoint http://docs.phraseapp.com/api/v2/uploads/#create\n  in your configuration (.phraseapp.yml) for each source.\n  See our configuration guide for more information http://docs.phraseapp.com/developers/cli/configuration/")

	r.Register("init", &WizardCommand{}, "Configure your PhraseApp client.")

	r.Register("upload/cleanup", &UploadCleanupCommand{Config: cfg}, "Delete unmentioned keys for given upload")

	r.RegisterFunc("info", infoCommand, "Info about version and revision of this client")
}
