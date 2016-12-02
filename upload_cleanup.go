package main

import (
	"fmt"
	"strings"

	"github.com/phrase/phraseapp-go/phraseapp"
)

type UploadCleanupCommand struct {
	*phraseapp.Config
	ID             string `cli:"arg required"`
	SuppressPrompt bool   `cli:"opt --yes desc='Don’t ask for confirmation'"`
}

func (cmd *UploadCleanupCommand) Run() error {

	client, err := newClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	return UploadCleanup(client, cmd)
}

func UploadCleanup(client *phraseapp.Client, cmd *UploadCleanupCommand) error {
	q := "unmentioned_in_upload:" + cmd.ID
	params := &phraseapp.KeysListParams{Q: &q}

	var err error
	page := 1

	keys, err := client.KeysList(cmd.Config.DefaultProjectID, page, 25, params)
	if err != nil {
		return err
	}

	if len(keys) == 0 {
		fmt.Println("There were no keys unmentioned in that upload.")
		return nil
	}

	for len(keys) != 0 {
		ids := make([]string, len(keys), len(keys))
		names := make([]string, len(keys), len(keys))
		for i, key := range keys {
			ids[i] = key.ID
			names[i] = key.Name
		}

		if !cmd.SuppressPrompt {
			fmt.Println("You are about to delete the following key(s) from your project:")
			fmt.Println(strings.Join(names, " "))
			fmt.Print("Are you sure you want to continue? (y/n) [n] ")

			confirmation := prompt()
			if err != nil {
				return err
			}

			if strings.ToLower(confirmation) != "y" {
				fmt.Println("Clean up aborted")
				return nil
			}
		}

		q := "ids:" + strings.Join(ids, ",")
		affected, err := client.KeysDelete(cmd.Config.DefaultProjectID, &phraseapp.KeysDeleteParams{
			Q: &q,
		})

		if err != nil {
			return err
		}

		fmt.Printf("%d key(s) successfully deleted.\n", affected.RecordsAffected)

		page++
		keys, err = client.KeysList(cmd.Config.DefaultProjectID, page, 25, params)
	}

	return nil
}
