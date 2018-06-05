package main

import (
	"github.com/phrase/phraseapp-go/phraseapp"
)

func projectID(cfg *phraseapp.Config) string {
	if cfg != nil {
		return cfg.DefaultProjectID
	}

	return ""
}
