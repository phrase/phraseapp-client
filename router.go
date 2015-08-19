package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func router(defaults map[string]string) *cli.Router {
	r := cli.NewRouter()

	projectID, _ := defaults["ProjectID"]

	r.Register("authorization/create", &AuthorizationCreate{}, "Create a new authorization.")

	r.Register("authorization/delete", &AuthorizationDelete{}, "Delete an existing authorization. API calls using that token will stop working.")

	r.Register("authorization/show", &AuthorizationShow{}, "Get details on a single authorization.")

	r.Register("authorization/update", &AuthorizationUpdate{}, "Update an existing authorization.")

	r.Register("authorizations/list", &AuthorizationsList{}, "List all your authorizations.")

	r.Register("blacklisted_key/create", &BlacklistedKeyCreate{ProjectID: projectID}, "Create a new rule for blacklisting keys.")

	r.Register("blacklisted_key/delete", &BlacklistedKeyDelete{ProjectID: projectID}, "Delete an existing rule for blacklisting keys.")

	r.Register("blacklisted_key/show", &BlacklistedKeyShow{ProjectID: projectID}, "Get details on a single rule for blacklisting keys for a given project.")

	r.Register("blacklisted_key/update", &BlacklistedKeyUpdate{ProjectID: projectID}, "Update an existing rule for blacklisting keys.")

	r.Register("blacklisted_keys/list", &BlacklistedKeysList{ProjectID: projectID}, "List all rules for blacklisting keys for the given project.")

	r.Register("comment/create", &CommentCreate{ProjectID: projectID}, "Create a new comment for a key.")

	r.Register("comment/delete", &CommentDelete{ProjectID: projectID}, "Delete an existing comment.")

	r.Register("comment/mark/check", &CommentMarkCheck{ProjectID: projectID}, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")

	r.Register("comment/mark/read", &CommentMarkRead{ProjectID: projectID}, "Mark a comment as read.")

	r.Register("comment/mark/unread", &CommentMarkUnread{ProjectID: projectID}, "Mark a comment as unread.")

	r.Register("comment/show", &CommentShow{ProjectID: projectID}, "Get details on a single comment.")

	r.Register("comment/update", &CommentUpdate{ProjectID: projectID}, "Update an existing comment.")

	r.Register("comments/list", &CommentsList{ProjectID: projectID}, "List all comments for a key.")

	r.Register("formats/list", &FormatsList{}, "Get a handy list of all localization file formats supported in PhraseApp.")

	r.Register("key/create", &KeyCreate{ProjectID: projectID}, "Create a new key.")

	r.Register("key/delete", &KeyDelete{ProjectID: projectID}, "Delete an existing key.")

	r.Register("key/show", &KeyShow{ProjectID: projectID}, "Get details on a single key for a given project.")

	r.Register("key/update", &KeyUpdate{ProjectID: projectID}, "Update an existing key.")

	r.Register("keys/delete", &KeysDelete{ProjectID: projectID}, "Delete all keys matching query. Same constraints as list.")

	r.Register("keys/list", &KeysList{ProjectID: projectID}, "List all keys for the given project. Alternatively you can POST requests to /search.")

	r.Register("keys/search", &KeysSearch{ProjectID: projectID}, "Search keys for the given project matching query.")

	r.Register("keys/tag", &KeysTag{ProjectID: projectID}, "Tags all keys matching query. Same constraints as list.")

	r.Register("keys/untag", &KeysUntag{ProjectID: projectID}, "Removes specified tags from keys matching query.")

	r.Register("locale/create", &LocaleCreate{ProjectID: projectID}, "Create a new locale.")

	r.Register("locale/delete", &LocaleDelete{ProjectID: projectID}, "Delete an existing locale.")

	r.Register("locale/download", &LocaleDownload{ProjectID: projectID}, "Download a locale in a specific file format.")

	r.Register("locale/show", &LocaleShow{ProjectID: projectID}, "Get details on a single locale for a given project.")

	r.Register("locale/update", &LocaleUpdate{ProjectID: projectID}, "Update an existing locale.")

	r.Register("locales/list", &LocalesList{ProjectID: projectID}, "List all locales for the given project.")

	r.Register("order/confirm", &OrderConfirm{ProjectID: projectID}, "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")

	r.Register("order/create", &OrderCreate{ProjectID: projectID}, "Create a new order. Access token scope must include <code>orders.create</code>.")

	r.Register("order/delete", &OrderDelete{ProjectID: projectID}, "Cancel an existing order. Must not yet be confirmed.")

	r.Register("order/show", &OrderShow{ProjectID: projectID}, "Get details on a single order.")

	r.Register("orders/list", &OrdersList{ProjectID: projectID}, "List all orders for the given project.")

	r.Register("project/create", &ProjectCreate{}, "Create a new project.")

	r.Register("project/delete", &ProjectDelete{}, "Delete an existing project.")

	r.Register("project/show", &ProjectShow{}, "Get details on a single project.")

	r.Register("project/update", &ProjectUpdate{}, "Update an existing project.")

	r.Register("projects/list", &ProjectsList{}, "List all projects the current user has access to.")

	r.Register("show/user", &ShowUser{}, "Show details for current User.")

	r.Register("styleguide/create", &StyleguideCreate{ProjectID: projectID}, "Create a new style guide.")

	r.Register("styleguide/delete", &StyleguideDelete{ProjectID: projectID}, "Delete an existing style guide.")

	r.Register("styleguide/show", &StyleguideShow{ProjectID: projectID}, "Get details on a single style guide.")

	r.Register("styleguide/update", &StyleguideUpdate{ProjectID: projectID}, "Update an existing style guide.")

	r.Register("styleguides/list", &StyleguidesList{ProjectID: projectID}, "List all styleguides for the given project.")

	r.Register("tag/create", &TagCreate{ProjectID: projectID}, "Create a new tag.")

	r.Register("tag/delete", &TagDelete{ProjectID: projectID}, "Delete an existing tag.")

	r.Register("tag/show", &TagShow{ProjectID: projectID}, "Get details and progress information on a single tag for a given project.")

	r.Register("tags/list", &TagsList{ProjectID: projectID}, "List all tags for the given project.")

	r.Register("translation/create", &TranslationCreate{ProjectID: projectID}, "Create a translation.")

	r.Register("translation/show", &TranslationShow{ProjectID: projectID}, "Get details on a single translation.")

	r.Register("translation/update", &TranslationUpdate{ProjectID: projectID}, "Update an existing translation.")

	r.Register("translations/by_key", &TranslationsByKey{ProjectID: projectID}, "List translations for a specific key.")

	r.Register("translations/by_locale", &TranslationsByLocale{ProjectID: projectID}, "List translations for a specific locale.")

	r.Register("translations/exclude", &TranslationsExclude{ProjectID: projectID}, "Exclude translations matching query from locale export.")

	r.Register("translations/include", &TranslationsInclude{ProjectID: projectID}, "Include translations matching query in locale export.")

	r.Register("translations/list", &TranslationsList{ProjectID: projectID}, "List translations for the given project. Alternatively, POST request to /search")

	r.Register("translations/search", &TranslationsSearch{ProjectID: projectID}, "List translations for the given project if you exceed GET request limitations on translations list.")

	r.Register("translations/unverify", &TranslationsUnverify{ProjectID: projectID}, "Mark translations matching query as unverified.")

	r.Register("translations/verify", &TranslationsVerify{ProjectID: projectID}, "Verify translations matching query.")

	r.Register("upload/create", &UploadCreate{ProjectID: projectID}, "Upload a new language file. Creates necessary resources in your project.")

	r.Register("upload/show", &UploadShow{ProjectID: projectID}, "View details and summary for a single upload.")

	r.Register("uploads/list", &UploadsList{ProjectID: projectID}, "List all uploads for the given project.")

	r.Register("version/show", &VersionShow{ProjectID: projectID}, "Get details on a single version.")

	r.Register("versions/list", &VersionsList{ProjectID: projectID}, "List all versions for the given translation.")

	r.Register("pull", &PullCommand{}, "Download locales from your PhraseApp project.")

	r.Register("push", &PushCommand{}, "Upload locales to your PhraseApp project.")

	r.Register("init", &WizardCommand{}, "Configure your PhraseApp client.")

	r.RegisterFunc("info", infoCommand, "Info about version and revision of this client")

	return r
}

func infoCommand() error {
	fmt.Printf("Built at 2015-08-18 19:28:28.860032443 +0200 CEST\n")
	fmt.Println("PhraseApp Client version:", "test")
	fmt.Println("PhraseApp API Client revision:", "8c3f2127836724c1d428896ab7f6eb113707f862")
	fmt.Println("PhraseApp Client revision:", "6491b83a75fbcabc5b9b174d51e240dd853ae9ae")
	fmt.Println("PhraseApp Docs revision:", "26152cec812022f63b4d442871dcdf880ef61891")
	return nil
}

type AuthorizationCreate struct {
	Credentials

	ExpiresAt **time.Time `cli:"opt --expires-at"`
	Note      *string     `cli:"opt --note"`
	Scopes    []string    `cli:"opt --scopes"`
}

func (cmd *AuthorizationCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.AuthorizationParams)

	val, defaultsPresent := defaults["authorization/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.ExpiresAt != nil {
		params.ExpiresAt = cmd.ExpiresAt
	}

	if cmd.Note != nil {
		params.Note = cmd.Note
	}

	if cmd.Scopes != nil {
		params.Scopes = cmd.Scopes
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationCreate(params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationDelete struct {
	Credentials

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.AuthorizationDelete(cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type AuthorizationShow struct {
	Credentials

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationShow(cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationUpdate struct {
	Credentials

	ExpiresAt **time.Time `cli:"opt --expires-at"`
	Note      *string     `cli:"opt --note"`
	Scopes    []string    `cli:"opt --scopes"`

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.AuthorizationParams)

	val, defaultsPresent := defaults["authorization/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.ExpiresAt != nil {
		params.ExpiresAt = cmd.ExpiresAt
	}

	if cmd.Note != nil {
		params.Note = cmd.Note
	}

	if cmd.Scopes != nil {
		params.Scopes = cmd.Scopes
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationUpdate(cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *AuthorizationsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistedKeyCreate struct {
	Credentials

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
}

func (cmd *BlacklistedKeyCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.BlacklistedKeyParams)

	val, defaultsPresent := defaults["blacklisted_key/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.BlacklistedKeyCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistedKeyDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.BlacklistedKeyDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type BlacklistedKeyShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.BlacklistedKeyShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistedKeyUpdate struct {
	Credentials

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.BlacklistedKeyParams)

	val, defaultsPresent := defaults["blacklisted_key/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.BlacklistedKeyUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistedKeysList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *BlacklistedKeysList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.BlacklistedKeysList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentCreate struct {
	Credentials

	Message *string `cli:"opt --message"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *CommentCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.CommentParams)

	val, defaultsPresent := defaults["comment/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Message != nil {
		params.Message = cmd.Message
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentCreate(cmd.ProjectID, cmd.KeyID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentDelete(cmd.ProjectID, cmd.KeyID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkCheck struct {
	Credentials

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkCheck) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkCheck(cmd.ProjectID, cmd.KeyID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkRead struct {
	Credentials

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkRead) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkRead(cmd.ProjectID, cmd.KeyID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkUnread struct {
	Credentials

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkUnread) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkUnread(cmd.ProjectID, cmd.KeyID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type CommentShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentShow(cmd.ProjectID, cmd.KeyID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentUpdate struct {
	Credentials

	Message *string `cli:"opt --message"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.CommentParams)

	val, defaultsPresent := defaults["comment/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Message != nil {
		params.Message = cmd.Message
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentUpdate(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *CommentsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentsList(cmd.ProjectID, cmd.KeyID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type FormatsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *FormatsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.FormatsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyCreate struct {
	Credentials

	DataType              *string `cli:"opt --data-type"`
	Description           *string `cli:"opt --description"`
	LocalizedFormatKey    *string `cli:"opt --localized-format-key"`
	LocalizedFormatString *string `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64  `cli:"opt --max-characters-allowed"`
	Name                  *string `cli:"opt --name"`
	NamePlural            *string `cli:"opt --name-plural"`
	OriginalFile          *string `cli:"opt --original-file"`
	Plural                *bool   `cli:"opt --plural"`
	RemoveScreenshot      *bool   `cli:"opt --remove-screenshot"`
	Screenshot            *string `cli:"opt --screenshot"`
	Tags                  *string `cli:"opt --tags"`
	Unformatted           *bool   `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool   `cli:"opt --xml-space-preserve"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeyCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationKeyParams)

	val, defaultsPresent := defaults["key/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.DataType != nil {
		params.DataType = cmd.DataType
	}

	if cmd.Description != nil {
		params.Description = cmd.Description
	}

	if cmd.LocalizedFormatKey != nil {
		params.LocalizedFormatKey = cmd.LocalizedFormatKey
	}

	if cmd.LocalizedFormatString != nil {
		params.LocalizedFormatString = cmd.LocalizedFormatString
	}

	if cmd.MaxCharactersAllowed != nil {
		params.MaxCharactersAllowed = cmd.MaxCharactersAllowed
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.NamePlural != nil {
		params.NamePlural = cmd.NamePlural
	}

	if cmd.OriginalFile != nil {
		params.OriginalFile = cmd.OriginalFile
	}

	if cmd.Plural != nil {
		params.Plural = cmd.Plural
	}

	if cmd.RemoveScreenshot != nil {
		params.RemoveScreenshot = cmd.RemoveScreenshot
	}

	if cmd.Screenshot != nil {
		params.Screenshot = cmd.Screenshot
	}

	if cmd.Tags != nil {
		params.Tags = cmd.Tags
	}

	if cmd.Unformatted != nil {
		params.Unformatted = cmd.Unformatted
	}

	if cmd.XmlSpacePreserve != nil {
		params.XmlSpacePreserve = cmd.XmlSpacePreserve
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *KeyDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.KeyDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type KeyShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *KeyShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyUpdate struct {
	Credentials

	DataType              *string `cli:"opt --data-type"`
	Description           *string `cli:"opt --description"`
	LocalizedFormatKey    *string `cli:"opt --localized-format-key"`
	LocalizedFormatString *string `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64  `cli:"opt --max-characters-allowed"`
	Name                  *string `cli:"opt --name"`
	NamePlural            *string `cli:"opt --name-plural"`
	OriginalFile          *string `cli:"opt --original-file"`
	Plural                *bool   `cli:"opt --plural"`
	RemoveScreenshot      *bool   `cli:"opt --remove-screenshot"`
	Screenshot            *string `cli:"opt --screenshot"`
	Tags                  *string `cli:"opt --tags"`
	Unformatted           *bool   `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool   `cli:"opt --xml-space-preserve"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *KeyUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationKeyParams)

	val, defaultsPresent := defaults["key/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.DataType != nil {
		params.DataType = cmd.DataType
	}

	if cmd.Description != nil {
		params.Description = cmd.Description
	}

	if cmd.LocalizedFormatKey != nil {
		params.LocalizedFormatKey = cmd.LocalizedFormatKey
	}

	if cmd.LocalizedFormatString != nil {
		params.LocalizedFormatString = cmd.LocalizedFormatString
	}

	if cmd.MaxCharactersAllowed != nil {
		params.MaxCharactersAllowed = cmd.MaxCharactersAllowed
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.NamePlural != nil {
		params.NamePlural = cmd.NamePlural
	}

	if cmd.OriginalFile != nil {
		params.OriginalFile = cmd.OriginalFile
	}

	if cmd.Plural != nil {
		params.Plural = cmd.Plural
	}

	if cmd.RemoveScreenshot != nil {
		params.RemoveScreenshot = cmd.RemoveScreenshot
	}

	if cmd.Screenshot != nil {
		params.Screenshot = cmd.Screenshot
	}

	if cmd.Tags != nil {
		params.Tags = cmd.Tags
	}

	if cmd.Unformatted != nil {
		params.Unformatted = cmd.Unformatted
	}

	if cmd.XmlSpacePreserve != nil {
		params.XmlSpacePreserve = cmd.XmlSpacePreserve
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysDelete struct {
	Credentials

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.KeysDeleteParams)

	val, defaultsPresent := defaults["keys/delete"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysDelete(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysList struct {
	Credentials

	LocaleID *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query -q"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.KeysListParams)

	val, defaultsPresent := defaults["keys/list"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysSearch struct {
	Credentials

	LocaleID *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query -q"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysSearch) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.KeysSearchParams)

	val, defaultsPresent := defaults["keys/search"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysSearch(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysTag struct {
	Credentials

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`
	Tags     *string `cli:"opt --tags"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysTag) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.KeysTagParams)

	val, defaultsPresent := defaults["keys/tag"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Tags != nil {
		params.Tags = cmd.Tags
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysTag(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysUntag struct {
	Credentials

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`
	Tags     *string `cli:"opt --tags"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysUntag) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.KeysUntagParams)

	val, defaultsPresent := defaults["keys/untag"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Tags != nil {
		params.Tags = cmd.Tags
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysUntag(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleCreate struct {
	Credentials

	Code           *string `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           *string `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleID *string `cli:"opt --source-locale-id"`

	ProjectID string `cli:"arg required"`
}

func (cmd *LocaleCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.LocaleParams)

	val, defaultsPresent := defaults["locale/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Code != nil {
		params.Code = cmd.Code
	}

	if cmd.Default != nil {
		params.Default = cmd.Default
	}

	if cmd.Main != nil {
		params.Main = cmd.Main
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.Rtl != nil {
		params.Rtl = cmd.Rtl
	}

	if cmd.SourceLocaleID != nil {
		params.SourceLocaleID = cmd.SourceLocaleID
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.LocaleDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type LocaleDownload struct {
	Credentials

	ConvertEmoji               bool                    `cli:"opt --convert-emoji"`
	Encoding                   *string                 `cli:"opt --encoding"`
	FileFormat                 *string                 `cli:"opt --file-format"`
	FormatOptions              *map[string]interface{} `cli:"opt --format-options"`
	IncludeEmptyTranslations   bool                    `cli:"opt --include-empty-translations"`
	KeepNotranslateTags        bool                    `cli:"opt --keep-notranslate-tags"`
	SkipUnverifiedTranslations bool                    `cli:"opt --skip-unverified-translations"`
	Tag                        *string                 `cli:"opt --tag"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleDownload) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.LocaleDownloadParams)

	val, defaultsPresent := defaults["locale/download"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	params.ConvertEmoji = cmd.ConvertEmoji

	if cmd.Encoding != nil {
		params.Encoding = cmd.Encoding
	}

	if cmd.FileFormat != nil {
		params.FileFormat = cmd.FileFormat
	}

	if cmd.FormatOptions != nil {
		params.FormatOptions = cmd.FormatOptions
	}

	params.IncludeEmptyTranslations = cmd.IncludeEmptyTranslations

	params.KeepNotranslateTags = cmd.KeepNotranslateTags

	params.SkipUnverifiedTranslations = cmd.SkipUnverifiedTranslations

	if cmd.Tag != nil {
		params.Tag = cmd.Tag
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleDownload(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}

type LocaleShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleUpdate struct {
	Credentials

	Code           *string `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           *string `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleID *string `cli:"opt --source-locale-id"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.LocaleParams)

	val, defaultsPresent := defaults["locale/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Code != nil {
		params.Code = cmd.Code
	}

	if cmd.Default != nil {
		params.Default = cmd.Default
	}

	if cmd.Main != nil {
		params.Main = cmd.Main
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.Rtl != nil {
		params.Rtl = cmd.Rtl
	}

	if cmd.SourceLocaleID != nil {
		params.SourceLocaleID = cmd.SourceLocaleID
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocalesList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *LocalesList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocalesList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderConfirm struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderConfirm) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderConfirm(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderCreate struct {
	Credentials

	Category                         *string `cli:"opt --category"`
	IncludeUntranslatedKeys          *bool   `cli:"opt --include-untranslated-keys"`
	IncludeUnverifiedTranslations    *bool   `cli:"opt --include-unverified-translations"`
	Lsp                              *string `cli:"opt --lsp"`
	Message                          *string `cli:"opt --message"`
	Priority                         *bool   `cli:"opt --priority"`
	Quality                          *bool   `cli:"opt --quality"`
	SourceLocaleID                   *string `cli:"opt --source-locale-id"`
	StyleguideID                     *string `cli:"opt --styleguide-id"`
	Tag                              *string `cli:"opt --tag"`
	TargetLocaleIDs                  *string `cli:"opt --target-locale-ids"`
	TranslationType                  *string `cli:"opt --translation-type"`
	UnverifyTranslationsUponDelivery *bool   `cli:"opt --unverify-translations-upon-delivery"`

	ProjectID string `cli:"arg required"`
}

func (cmd *OrderCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationOrderParams)

	val, defaultsPresent := defaults["order/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Category != nil {
		params.Category = cmd.Category
	}

	if cmd.IncludeUntranslatedKeys != nil {
		params.IncludeUntranslatedKeys = cmd.IncludeUntranslatedKeys
	}

	if cmd.IncludeUnverifiedTranslations != nil {
		params.IncludeUnverifiedTranslations = cmd.IncludeUnverifiedTranslations
	}

	if cmd.Lsp != nil {
		params.Lsp = cmd.Lsp
	}

	if cmd.Message != nil {
		params.Message = cmd.Message
	}

	if cmd.Priority != nil {
		params.Priority = cmd.Priority
	}

	if cmd.Quality != nil {
		params.Quality = cmd.Quality
	}

	if cmd.SourceLocaleID != nil {
		params.SourceLocaleID = cmd.SourceLocaleID
	}

	if cmd.StyleguideID != nil {
		params.StyleguideID = cmd.StyleguideID
	}

	if cmd.Tag != nil {
		params.Tag = cmd.Tag
	}

	if cmd.TargetLocaleIDs != nil {
		params.TargetLocaleIDs = cmd.TargetLocaleIDs
	}

	if cmd.TranslationType != nil {
		params.TranslationType = cmd.TranslationType
	}

	if cmd.UnverifyTranslationsUponDelivery != nil {
		params.UnverifyTranslationsUponDelivery = cmd.UnverifyTranslationsUponDelivery
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.OrderDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type OrderShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrdersList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *OrdersList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrdersList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectCreate struct {
	Credentials

	MainFormat              *string `cli:"opt --main-format"`
	Name                    *string `cli:"opt --name"`
	SharesTranslationMemory *bool   `cli:"opt --shares-translation-memory"`
}

func (cmd *ProjectCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.ProjectParams)

	val, defaultsPresent := defaults["project/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.MainFormat != nil {
		params.MainFormat = cmd.MainFormat
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.SharesTranslationMemory != nil {
		params.SharesTranslationMemory = cmd.SharesTranslationMemory
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectCreate(params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectDelete struct {
	Credentials

	ID string `cli:"arg required"`
}

func (cmd *ProjectDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.ProjectDelete(cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type ProjectShow struct {
	Credentials

	ID string `cli:"arg required"`
}

func (cmd *ProjectShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectShow(cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectUpdate struct {
	Credentials

	MainFormat              *string `cli:"opt --main-format"`
	Name                    *string `cli:"opt --name"`
	SharesTranslationMemory *bool   `cli:"opt --shares-translation-memory"`

	ID string `cli:"arg required"`
}

func (cmd *ProjectUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.ProjectParams)

	val, defaultsPresent := defaults["project/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.MainFormat != nil {
		params.MainFormat = cmd.MainFormat
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	if cmd.SharesTranslationMemory != nil {
		params.SharesTranslationMemory = cmd.SharesTranslationMemory
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectUpdate(cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *ProjectsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ShowUser struct {
	Credentials
}

func (cmd *ShowUser) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ShowUser()

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideCreate struct {
	Credentials

	Audience           *string `cli:"opt --audience"`
	Business           *string `cli:"opt --business"`
	CompanyBranding    *string `cli:"opt --company-branding"`
	Formatting         *string `cli:"opt --formatting"`
	GlossaryTerms      *string `cli:"opt --glossary-terms"`
	GrammarConsistency *string `cli:"opt --grammar-consistency"`
	GrammaticalPerson  *string `cli:"opt --grammatical-person"`
	LiteralTranslation *string `cli:"opt --literal-translation"`
	OverallTone        *string `cli:"opt --overall-tone"`
	Samples            *string `cli:"opt --samples"`
	TargetAudience     *string `cli:"opt --target-audience"`
	Title              *string `cli:"opt --title"`
	VocabularyType     *string `cli:"opt --vocabulary-type"`

	ProjectID string `cli:"arg required"`
}

func (cmd *StyleguideCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.StyleguideParams)

	val, defaultsPresent := defaults["styleguide/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Audience != nil {
		params.Audience = cmd.Audience
	}

	if cmd.Business != nil {
		params.Business = cmd.Business
	}

	if cmd.CompanyBranding != nil {
		params.CompanyBranding = cmd.CompanyBranding
	}

	if cmd.Formatting != nil {
		params.Formatting = cmd.Formatting
	}

	if cmd.GlossaryTerms != nil {
		params.GlossaryTerms = cmd.GlossaryTerms
	}

	if cmd.GrammarConsistency != nil {
		params.GrammarConsistency = cmd.GrammarConsistency
	}

	if cmd.GrammaticalPerson != nil {
		params.GrammaticalPerson = cmd.GrammaticalPerson
	}

	if cmd.LiteralTranslation != nil {
		params.LiteralTranslation = cmd.LiteralTranslation
	}

	if cmd.OverallTone != nil {
		params.OverallTone = cmd.OverallTone
	}

	if cmd.Samples != nil {
		params.Samples = cmd.Samples
	}

	if cmd.TargetAudience != nil {
		params.TargetAudience = cmd.TargetAudience
	}

	if cmd.Title != nil {
		params.Title = cmd.Title
	}

	if cmd.VocabularyType != nil {
		params.VocabularyType = cmd.VocabularyType
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *StyleguideDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.StyleguideDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type StyleguideShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *StyleguideShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideUpdate struct {
	Credentials

	Audience           *string `cli:"opt --audience"`
	Business           *string `cli:"opt --business"`
	CompanyBranding    *string `cli:"opt --company-branding"`
	Formatting         *string `cli:"opt --formatting"`
	GlossaryTerms      *string `cli:"opt --glossary-terms"`
	GrammarConsistency *string `cli:"opt --grammar-consistency"`
	GrammaticalPerson  *string `cli:"opt --grammatical-person"`
	LiteralTranslation *string `cli:"opt --literal-translation"`
	OverallTone        *string `cli:"opt --overall-tone"`
	Samples            *string `cli:"opt --samples"`
	TargetAudience     *string `cli:"opt --target-audience"`
	Title              *string `cli:"opt --title"`
	VocabularyType     *string `cli:"opt --vocabulary-type"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *StyleguideUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.StyleguideParams)

	val, defaultsPresent := defaults["styleguide/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Audience != nil {
		params.Audience = cmd.Audience
	}

	if cmd.Business != nil {
		params.Business = cmd.Business
	}

	if cmd.CompanyBranding != nil {
		params.CompanyBranding = cmd.CompanyBranding
	}

	if cmd.Formatting != nil {
		params.Formatting = cmd.Formatting
	}

	if cmd.GlossaryTerms != nil {
		params.GlossaryTerms = cmd.GlossaryTerms
	}

	if cmd.GrammarConsistency != nil {
		params.GrammarConsistency = cmd.GrammarConsistency
	}

	if cmd.GrammaticalPerson != nil {
		params.GrammaticalPerson = cmd.GrammaticalPerson
	}

	if cmd.LiteralTranslation != nil {
		params.LiteralTranslation = cmd.LiteralTranslation
	}

	if cmd.OverallTone != nil {
		params.OverallTone = cmd.OverallTone
	}

	if cmd.Samples != nil {
		params.Samples = cmd.Samples
	}

	if cmd.TargetAudience != nil {
		params.TargetAudience = cmd.TargetAudience
	}

	if cmd.Title != nil {
		params.Title = cmd.Title
	}

	if cmd.VocabularyType != nil {
		params.VocabularyType = cmd.VocabularyType
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguidesList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *StyleguidesList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguidesList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagCreate struct {
	Credentials

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TagCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TagParams)

	val, defaultsPresent := defaults["tag/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagDelete struct {
	Credentials

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.TagDelete(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return nil
}

type TagShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagShow(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TagsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagsList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationCreate struct {
	Credentials

	Content      *string `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	KeyID        *string `cli:"opt --key-id"`
	LocaleID     *string `cli:"opt --locale-id"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationParams)

	val, defaultsPresent := defaults["translation/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Content != nil {
		params.Content = cmd.Content
	}

	if cmd.Excluded != nil {
		params.Excluded = cmd.Excluded
	}

	if cmd.KeyID != nil {
		params.KeyID = cmd.KeyID
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.PluralSuffix != nil {
		params.PluralSuffix = cmd.PluralSuffix
	}

	if cmd.Unverified != nil {
		params.Unverified = cmd.Unverified
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *TranslationShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationUpdate struct {
	Credentials

	Content      *string `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *TranslationUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationUpdateParams)

	val, defaultsPresent := defaults["translation/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Content != nil {
		params.Content = cmd.Content
	}

	if cmd.Excluded != nil {
		params.Excluded = cmd.Excluded
	}

	if cmd.PluralSuffix != nil {
		params.PluralSuffix = cmd.PluralSuffix
	}

	if cmd.Unverified != nil {
		params.Unverified = cmd.Unverified
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByKey struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *TranslationsByKey) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsByKeyParams)

	val, defaultsPresent := defaults["translations/by_key"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsByKey(cmd.ProjectID, cmd.KeyID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByLocale struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	LocaleID  string `cli:"arg required"`
}

func (cmd *TranslationsByLocale) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsByLocaleParams)

	val, defaultsPresent := defaults["translations/by_locale"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsByLocale(cmd.ProjectID, cmd.LocaleID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsExclude struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsExclude) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsExcludeParams)

	val, defaultsPresent := defaults["translations/exclude"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsExclude(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsInclude struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsInclude) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsIncludeParams)

	val, defaultsPresent := defaults["translations/include"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsInclude(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsList struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsListParams)

	val, defaultsPresent := defaults["translations/list"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsSearch struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsSearch) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsSearchParams)

	val, defaultsPresent := defaults["translations/search"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsSearch(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsUnverify struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsUnverify) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsUnverifyParams)

	val, defaultsPresent := defaults["translations/unverify"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsUnverify(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsVerify struct {
	Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsVerify) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.TranslationsVerifyParams)

	val, defaultsPresent := defaults["translations/verify"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Order != nil {
		params.Order = cmd.Order
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Sort != nil {
		params.Sort = cmd.Sort
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsVerify(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadCreate struct {
	Credentials

	ConvertEmoji       *bool   `cli:"opt --convert-emoji"`
	File               *string `cli:"opt --file"`
	FileFormat         *string `cli:"opt --file-format"`
	LocaleID           *string `cli:"opt --locale-id"`
	SkipUnverification *bool   `cli:"opt --skip-unverification"`
	SkipUploadTags     *bool   `cli:"opt --skip-upload-tags"`
	Tags               *string `cli:"opt --tags"`
	UpdateTranslations *bool   `cli:"opt --update-translations"`

	ProjectID string `cli:"arg required"`
}

func (cmd *UploadCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.LocaleFileImportParams)

	val, defaultsPresent := defaults["upload/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.ConvertEmoji != nil {
		params.ConvertEmoji = cmd.ConvertEmoji
	}

	if cmd.File != nil {
		params.File = cmd.File
	}

	if cmd.FileFormat != nil {
		params.FileFormat = cmd.FileFormat
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.SkipUnverification != nil {
		params.SkipUnverification = cmd.SkipUnverification
	}

	if cmd.SkipUploadTags != nil {
		params.SkipUploadTags = cmd.SkipUploadTags
	}

	if cmd.Tags != nil {
		params.Tags = cmd.Tags
	}

	if cmd.UpdateTranslations != nil {
		params.UpdateTranslations = cmd.UpdateTranslations
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.UploadCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadShow struct {
	Credentials

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *UploadShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.UploadShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *UploadsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.UploadsList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionShow struct {
	Credentials

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
	ID            string `cli:"arg required"`
}

func (cmd *VersionShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.VersionShow(cmd.ProjectID, cmd.TranslationID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionsList struct {
	Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
}

func (cmd *VersionsList) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	credentials := PhraseAppCredentials(cmd.Credentials)
	client, err := phraseapp.NewClient(credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.VersionsList(cmd.ProjectID, cmd.TranslationID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}
