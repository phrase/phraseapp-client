package errors

import (
	"encoding/json"

	bugsnag "github.com/bugsnag/bugsnag-go"
)

// Version contains version information of a client and its dependencies.
type Version struct {
	BuiltAt                  string `json:"built_at"`
	Revision                 string `json:"revision"`
	GeneratorRevision        string `json:"generator_revision"`
	DocRevision              string `json:"doc_revision"`
	LibraryRevision          string `json:"library_revision"`
	LibraryGeneratorRevision string `json:"library_generator_revision"`
	LibraryDocRevision       string `json:"library_doc_revision"`
	GoVersion                string `json:"go_version"`
}

// Environment describes the runtime environment the error occured in.
type Environment struct {
	OS   string `json:"os"`
	Arch string `json:"architecture"`
}

// Project contains the Project ID. (It's organized like this so it shows
// up in a tab on bugsnag.com.)
type Project struct {
	ID string `json:"id,omitempty"`
}

// User describes who's using the CLI tool.
type User struct {
	AccessToken string `json:"access_token,omitempty"`
	Username    string `json:"username,omitempty"`
}

// MetaData holds contextual information about an error besides the stack.
type MetaData struct {
	User        User        `json:"user"`
	Project     Project     `json:"project"`
	Version     Version     `json:"version"`
	Environment Environment `json:"environment"`
}

// ToBugsnagMetaData converts md to the bugsnag-compatible meta data type.
func (md *MetaData) ToBugsnagMetaData() bugsnag.MetaData {
	// simply encode to JSON ...
	metaDataJSON, err := json.Marshal(md)
	if err != nil {
		panic(err) // should never happen
	}

	// ... and decode into map
	metaDataMap := map[string]map[string]interface{}{}
	err = json.Unmarshal(metaDataJSON, &metaDataMap)
	if err != nil {
		panic(err) // should never happen
	}

	return bugsnag.MetaData(metaDataMap)
}
