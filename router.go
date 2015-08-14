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

	projectId, _ := defaults["ProjectId"]

	r.Register("authorization/create", &AuthorizationCreate{}, "Create a new authorization.")

	r.Register("authorization/delete", &AuthorizationDelete{}, "Delete an existing authorization. API calls using that token will stop working.")

	r.Register("authorization/show", &AuthorizationShow{}, "Get details on a single authorization.")

	r.Register("authorization/update", &AuthorizationUpdate{}, "Update an existing authorization.")

	r.Register("authorizations/list", &AuthorizationsList{}, "List all your authorizations.")

	r.Register("comment/create", &CommentCreate{ProjectId: projectId}, "Create a new comment for a key.")

	r.Register("comment/delete", &CommentDelete{ProjectId: projectId}, "Delete an existing comment.")

	r.Register("comment/mark/check", &CommentMarkCheck{ProjectId: projectId}, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")

	r.Register("comment/mark/read", &CommentMarkRead{ProjectId: projectId}, "Mark a comment as read.")

	r.Register("comment/mark/unread", &CommentMarkUnread{ProjectId: projectId}, "Mark a comment as unread.")

	r.Register("comment/show", &CommentShow{ProjectId: projectId}, "Get details on a single comment.")

	r.Register("comment/update", &CommentUpdate{ProjectId: projectId}, "Update an existing comment.")

	r.Register("comments/list", &CommentsList{ProjectId: projectId}, "List all comments for a key.")

	r.Register("exclude_rule/create", &ExcludeRuleCreate{ProjectId: projectId}, "Create a new blacklisted key.")

	r.Register("exclude_rule/delete", &ExcludeRuleDelete{ProjectId: projectId}, "Delete an existing blacklisted key.")

	r.Register("exclude_rule/show", &ExcludeRuleShow{ProjectId: projectId}, "Get details on a single blacklisted key for a given project.")

	r.Register("exclude_rule/update", &ExcludeRuleUpdate{ProjectId: projectId}, "Update an existing blacklisted key.")

	r.Register("exclude_rules/index", &ExcludeRulesIndex{ProjectId: projectId}, "List all blacklisted keys for the given project.")

	r.Register("formats/list", &FormatsList{}, "Get a handy list of all localization file formats supported in PhraseApp.")

	r.Register("key/create", &KeyCreate{ProjectId: projectId}, "Create a new key.")

	r.Register("key/delete", &KeyDelete{ProjectId: projectId}, "Delete an existing key.")

	r.Register("key/show", &KeyShow{ProjectId: projectId}, "Get details on a single key for a given project.")

	r.Register("key/update", &KeyUpdate{ProjectId: projectId}, "Update an existing key.")

	r.Register("keys/delete", &KeysDelete{ProjectId: projectId}, "Delete all keys matching query. Same constraints as list.")

	r.Register("keys/list", &KeysList{ProjectId: projectId}, "List all keys for the given project. Alternatively you can POST requests to /search.")

	r.Register("keys/search", &KeysSearch{ProjectId: projectId}, "Search keys for the given project matching query.")

	r.Register("keys/tag", &KeysTag{ProjectId: projectId}, "Tags all keys matching query. Same constraints as list.")

	r.Register("keys/untag", &KeysUntag{ProjectId: projectId}, "Removes specified tags from keys matching query.")

	r.Register("locale/create", &LocaleCreate{ProjectId: projectId}, "Create a new locale.")

	r.Register("locale/delete", &LocaleDelete{ProjectId: projectId}, "Delete an existing locale.")

	r.Register("locale/download", &LocaleDownload{ProjectId: projectId}, "Download a locale in a specific file format.")

	r.Register("locale/show", &LocaleShow{ProjectId: projectId}, "Get details on a single locale for a given project.")

	r.Register("locale/update", &LocaleUpdate{ProjectId: projectId}, "Update an existing locale.")

	r.Register("locales/list", &LocalesList{ProjectId: projectId}, "List all locales for the given project.")

	r.Register("order/confirm", &OrderConfirm{ProjectId: projectId}, "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")

	r.Register("order/create", &OrderCreate{ProjectId: projectId}, "Create a new order. Access token scope must include <code>orders.create</code>.")

	r.Register("order/delete", &OrderDelete{ProjectId: projectId}, "Cancel an existing order. Must not yet be confirmed.")

	r.Register("order/show", &OrderShow{ProjectId: projectId}, "Get details on a single order.")

	r.Register("orders/list", &OrdersList{ProjectId: projectId}, "List all orders for the given project.")

	r.Register("project/create", &ProjectCreate{}, "Create a new project.")

	r.Register("project/delete", &ProjectDelete{}, "Delete an existing project.")

	r.Register("project/show", &ProjectShow{}, "Get details on a single project.")

	r.Register("project/update", &ProjectUpdate{}, "Update an existing project.")

	r.Register("projects/list", &ProjectsList{}, "List all projects the current user has access to.")

	r.Register("show/user", &ShowUser{}, "Show details for current User.")

	r.Register("styleguide/create", &StyleguideCreate{ProjectId: projectId}, "Create a new style guide.")

	r.Register("styleguide/delete", &StyleguideDelete{ProjectId: projectId}, "Delete an existing style guide.")

	r.Register("styleguide/show", &StyleguideShow{ProjectId: projectId}, "Get details on a single style guide.")

	r.Register("styleguide/update", &StyleguideUpdate{ProjectId: projectId}, "Update an existing style guide.")

	r.Register("styleguides/list", &StyleguidesList{ProjectId: projectId}, "List all styleguides for the given project.")

	r.Register("tag/create", &TagCreate{ProjectId: projectId}, "Create a new tag.")

	r.Register("tag/delete", &TagDelete{ProjectId: projectId}, "Delete an existing tag.")

	r.Register("tag/show", &TagShow{ProjectId: projectId}, "Get details and progress information on a single tag for a given project.")

	r.Register("tags/list", &TagsList{ProjectId: projectId}, "List all tags for the given project.")

	r.Register("translation/create", &TranslationCreate{ProjectId: projectId}, "Create a translation.")

	r.Register("translation/machine_translate", &TranslationMachineTranslate{ProjectId: projectId}, "Update a translation with machine translation")

	r.Register("translation/show", &TranslationShow{ProjectId: projectId}, "Get details on a single translation.")

	r.Register("translation/update", &TranslationUpdate{ProjectId: projectId}, "Update an existing translation.")

	r.Register("translations/by_key", &TranslationsByKey{ProjectId: projectId}, "List translations for a specific key.")

	r.Register("translations/by_locale", &TranslationsByLocale{ProjectId: projectId}, "List translations for a specific locale.")

	r.Register("translations/exclude", &TranslationsExclude{ProjectId: projectId}, "Exclude translations matching query from locale export.")

	r.Register("translations/include", &TranslationsInclude{ProjectId: projectId}, "Include translations matching query in locale export.")

	r.Register("translations/list", &TranslationsList{ProjectId: projectId}, "List translations for the given project. Alternatively, POST request to /search")

	r.Register("translations/search", &TranslationsSearch{ProjectId: projectId}, "List translations for the given project if you exceed GET request limitations on translations list.")

	r.Register("translations/unverify", &TranslationsUnverify{ProjectId: projectId}, "Mark translations matching query as unverified.")

	r.Register("translations/verify", &TranslationsVerify{ProjectId: projectId}, "Verify translations matching query.")

	r.Register("upload/create", &UploadCreate{ProjectId: projectId}, "Upload a new language file. Creates necessary resources in your project.")

	r.Register("upload/show", &UploadShow{ProjectId: projectId}, "View details and summary for a single upload.")

	r.Register("version/show", &VersionShow{ProjectId: projectId}, "Get details on a single version.")

	r.Register("versions/list", &VersionsList{ProjectId: projectId}, "List all versions for the given translation.")

	r.Register("pull", &PullCommand{}, "Download locales from your PhraseApp project.")

	r.Register("push", &PushCommand{}, "Upload locales to your PhraseApp project.")

	r.Register("init", &WizardCommand{}, "Configure your PhraseApp client.")

	r.RegisterFunc("info", infoCommand, "Info about version and revision of this client")

	return r
}

func infoCommand() error {
	fmt.Printf("Built at 2015-08-14 15:27:11.825876922 +0200 CEST\n")
	fmt.Println("PhraseApp Client version:", "test")
	fmt.Println("PhraseApp API Client revision:", "3d1142bef0cc318d5cd913e2f944959e8eaddfe7")
	fmt.Println("PhraseApp Client revision:", "948e7fd1576ee6818ec033130bd080564300a067")
	fmt.Println("PhraseApp Docs revision:", "8f1d9ef516480148c220f54bd26ed2dd8d947857")
	return nil
}

type AuthorizationCreate struct {
	phraseapp.Credentials

	ExpiresAt *time.Time `cli:"opt --expires-at"`
	Note      string     `cli:"opt --note"`
	Scopes    []string   `cli:"opt --scopes"`
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

	if cmd.Note != "" {
		params.Note = cmd.Note
	}

	if cmd.Scopes != nil {
		params.Scopes = cmd.Scopes
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
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
	phraseapp.Credentials

	Id string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.AuthorizationDelete(cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type AuthorizationShow struct {
	phraseapp.Credentials

	Id string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationShow(cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationUpdate struct {
	phraseapp.Credentials

	ExpiresAt *time.Time `cli:"opt --expires-at"`
	Note      string     `cli:"opt --note"`
	Scopes    []string   `cli:"opt --scopes"`

	Id string `cli:"arg required"`
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

	if cmd.Note != "" {
		params.Note = cmd.Note
	}

	if cmd.Scopes != nil {
		params.Scopes = cmd.Scopes
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationUpdate(cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationsList struct {
	phraseapp.Credentials

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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.AuthorizationsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentCreate struct {
	phraseapp.Credentials

	Message string `cli:"opt --message"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
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

	if cmd.Message != "" {
		params.Message = cmd.Message
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentCreate(cmd.ProjectId, cmd.KeyId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentDelete(cmd.ProjectId, cmd.KeyId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkCheck struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkCheck(cmd.ProjectId, cmd.KeyId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkRead struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkRead(cmd.ProjectId, cmd.KeyId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkUnread struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.CommentMarkUnread(cmd.ProjectId, cmd.KeyId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type CommentShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentShow(cmd.ProjectId, cmd.KeyId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentUpdate struct {
	phraseapp.Credentials

	Message string `cli:"opt --message"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.Message != "" {
		params.Message = cmd.Message
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentUpdate(cmd.ProjectId, cmd.KeyId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentsList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.CommentsList(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleCreate struct {
	phraseapp.Credentials

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
}

func (cmd *ExcludeRuleCreate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.ExcludeRuleParams)

	val, defaultsPresent := defaults["exclude_rule/create"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ExcludeRuleCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleDelete) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.ExcludeRuleDelete(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type ExcludeRuleShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleShow) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ExcludeRuleShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleUpdate struct {
	phraseapp.Credentials

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleUpdate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	params := new(phraseapp.ExcludeRuleParams)

	val, defaultsPresent := defaults["exclude_rule/update"]

	if defaultsPresent {
		params, e = params.ApplyDefaults(val)
		if e != nil {
			return e
		}
	}

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ExcludeRuleUpdate(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRulesIndex struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *ExcludeRulesIndex) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ExcludeRulesIndex(cmd.ProjectId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type FormatsList struct {
	phraseapp.Credentials

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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
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
	phraseapp.Credentials

	DataType              *string `cli:"opt --data-type"`
	Description           *string `cli:"opt --description"`
	LocalizedFormatKey    *string `cli:"opt --localized-format-key"`
	LocalizedFormatString *string `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64  `cli:"opt --max-characters-allowed"`
	Name                  string  `cli:"opt --name"`
	NamePlural            *string `cli:"opt --name-plural"`
	OriginalFile          *string `cli:"opt --original-file"`
	Plural                *bool   `cli:"opt --plural"`
	RemoveScreenshot      *bool   `cli:"opt --remove-screenshot"`
	Screenshot            *string `cli:"opt --screenshot"`
	Tags                  *string `cli:"opt --tags"`
	Unformatted           *bool   `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool   `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Name != "" {
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.KeyDelete(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type KeyShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyUpdate struct {
	phraseapp.Credentials

	DataType              *string `cli:"opt --data-type"`
	Description           *string `cli:"opt --description"`
	LocalizedFormatKey    *string `cli:"opt --localized-format-key"`
	LocalizedFormatString *string `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64  `cli:"opt --max-characters-allowed"`
	Name                  string  `cli:"opt --name"`
	NamePlural            *string `cli:"opt --name-plural"`
	OriginalFile          *string `cli:"opt --original-file"`
	Plural                *bool   `cli:"opt --plural"`
	RemoveScreenshot      *bool   `cli:"opt --remove-screenshot"`
	Screenshot            *string `cli:"opt --screenshot"`
	Tags                  *string `cli:"opt --tags"`
	Unformatted           *bool   `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool   `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.Name != "" {
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeyUpdate(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysDelete struct {
	phraseapp.Credentials

	LocaleId *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query"`

	ProjectId string `cli:"arg required"`
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

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysDelete(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysList struct {
	phraseapp.Credentials

	LocaleId *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysList(cmd.ProjectId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysSearch struct {
	phraseapp.Credentials

	LocaleId *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysSearch(cmd.ProjectId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysTag struct {
	phraseapp.Credentials

	LocaleId *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query"`
	Tags     string  `cli:"opt --tags"`

	ProjectId string `cli:"arg required"`
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

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Tags != "" {
		params.Tags = cmd.Tags
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysTag(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysUntag struct {
	phraseapp.Credentials

	LocaleId *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query"`
	Tags     string  `cli:"opt --tags"`

	ProjectId string `cli:"arg required"`
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

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	if cmd.Tags != "" {
		params.Tags = cmd.Tags
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.KeysUntag(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleCreate struct {
	phraseapp.Credentials

	Code           string  `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           string  `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleId *string `cli:"opt --source-locale-id"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Code != "" {
		params.Code = cmd.Code
	}

	if cmd.Default != nil {
		params.Default = cmd.Default
	}

	if cmd.Main != nil {
		params.Main = cmd.Main
	}

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	if cmd.Rtl != nil {
		params.Rtl = cmd.Rtl
	}

	if cmd.SourceLocaleId != nil {
		params.SourceLocaleId = cmd.SourceLocaleId
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.LocaleDelete(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type LocaleDownload struct {
	phraseapp.Credentials

	ConvertEmoji             *bool                   `cli:"opt --convert-emoji"`
	FileFormat               string                  `cli:"opt --file-format"`
	FormatOptions            *map[string]interface{} `cli:"opt --format-options"`
	IncludeEmptyTranslations *bool                   `cli:"opt --include-empty-translations"`
	KeepNotranslateTags      *bool                   `cli:"opt --keep-notranslate-tags"`
	Tag                      *string                 `cli:"opt --tag"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.ConvertEmoji != nil {
		params.ConvertEmoji = cmd.ConvertEmoji
	}

	if cmd.FileFormat != "" {
		params.FileFormat = cmd.FileFormat
	}

	if cmd.FormatOptions != nil {
		params.FormatOptions = cmd.FormatOptions
	}

	if cmd.IncludeEmptyTranslations != nil {
		params.IncludeEmptyTranslations = cmd.IncludeEmptyTranslations
	}

	if cmd.KeepNotranslateTags != nil {
		params.KeepNotranslateTags = cmd.KeepNotranslateTags
	}

	if cmd.Tag != nil {
		params.Tag = cmd.Tag
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleDownload(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}

type LocaleShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleUpdate struct {
	phraseapp.Credentials

	Code           string  `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           string  `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleId *string `cli:"opt --source-locale-id"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.Code != "" {
		params.Code = cmd.Code
	}

	if cmd.Default != nil {
		params.Default = cmd.Default
	}

	if cmd.Main != nil {
		params.Main = cmd.Main
	}

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	if cmd.Rtl != nil {
		params.Rtl = cmd.Rtl
	}

	if cmd.SourceLocaleId != nil {
		params.SourceLocaleId = cmd.SourceLocaleId
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocaleUpdate(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocalesList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.LocalesList(cmd.ProjectId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderConfirm struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderConfirm(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderCreate struct {
	phraseapp.Credentials

	Category                         string   `cli:"opt --category"`
	IncludeUntranslatedKeys          *bool    `cli:"opt --include-untranslated-keys"`
	IncludeUnverifiedTranslations    *bool    `cli:"opt --include-unverified-translations"`
	Lsp                              string   `cli:"opt --lsp"`
	Message                          *string  `cli:"opt --message"`
	Priority                         *bool    `cli:"opt --priority"`
	Quality                          *bool    `cli:"opt --quality"`
	SourceLocaleId                   string   `cli:"opt --source-locale-id"`
	StyleguideId                     *string  `cli:"opt --styleguide-id"`
	Tag                              *string  `cli:"opt --tag"`
	TargetLocaleIds                  []string `cli:"opt --target-locale-ids"`
	TranslationType                  string   `cli:"opt --translation-type"`
	UnverifyTranslationsUponDelivery *bool    `cli:"opt --unverify-translations-upon-delivery"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Category != "" {
		params.Category = cmd.Category
	}

	if cmd.IncludeUntranslatedKeys != nil {
		params.IncludeUntranslatedKeys = cmd.IncludeUntranslatedKeys
	}

	if cmd.IncludeUnverifiedTranslations != nil {
		params.IncludeUnverifiedTranslations = cmd.IncludeUnverifiedTranslations
	}

	if cmd.Lsp != "" {
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

	if cmd.SourceLocaleId != "" {
		params.SourceLocaleId = cmd.SourceLocaleId
	}

	if cmd.StyleguideId != nil {
		params.StyleguideId = cmd.StyleguideId
	}

	if cmd.Tag != nil {
		params.Tag = cmd.Tag
	}

	if cmd.TargetLocaleIds != nil {
		params.TargetLocaleIds = cmd.TargetLocaleIds
	}

	if cmd.TranslationType != "" {
		params.TranslationType = cmd.TranslationType
	}

	if cmd.UnverifyTranslationsUponDelivery != nil {
		params.UnverifyTranslationsUponDelivery = cmd.UnverifyTranslationsUponDelivery
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.OrderDelete(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type OrderShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrderShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrdersList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.OrdersList(cmd.ProjectId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectCreate struct {
	phraseapp.Credentials

	Name                    string `cli:"opt --name"`
	SharesTranslationMemory *bool  `cli:"opt --shares-translation-memory"`
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

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	if cmd.SharesTranslationMemory != nil {
		params.SharesTranslationMemory = cmd.SharesTranslationMemory
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
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
	phraseapp.Credentials

	Id string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.ProjectDelete(cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type ProjectShow struct {
	phraseapp.Credentials

	Id string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectShow(cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectUpdate struct {
	phraseapp.Credentials

	Name                    string `cli:"opt --name"`
	SharesTranslationMemory *bool  `cli:"opt --shares-translation-memory"`

	Id string `cli:"arg required"`
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

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	if cmd.SharesTranslationMemory != nil {
		params.SharesTranslationMemory = cmd.SharesTranslationMemory
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.ProjectUpdate(cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectsList struct {
	phraseapp.Credentials

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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
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
	phraseapp.Credentials
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
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
	phraseapp.Credentials

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
	Title              string  `cli:"opt --title"`
	VocabularyType     *string `cli:"opt --vocabulary-type"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Title != "" {
		params.Title = cmd.Title
	}

	if cmd.VocabularyType != nil {
		params.VocabularyType = cmd.VocabularyType
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.StyleguideDelete(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type StyleguideShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideUpdate struct {
	phraseapp.Credentials

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
	Title              string  `cli:"opt --title"`
	VocabularyType     *string `cli:"opt --vocabulary-type"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.Title != "" {
		params.Title = cmd.Title
	}

	if cmd.VocabularyType != nil {
		params.VocabularyType = cmd.VocabularyType
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguideUpdate(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguidesList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.StyleguidesList(cmd.ProjectId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagCreate struct {
	phraseapp.Credentials

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Name != "" {
		params.Name = cmd.Name
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagDelete struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.TagDelete(cmd.ProjectId, cmd.Name)

	if err != nil {
		return err
	}

	return nil
}

type TagShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagShow(cmd.ProjectId, cmd.Name)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagsList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TagsList(cmd.ProjectId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationCreate struct {
	phraseapp.Credentials

	Content      string  `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	KeyId        string  `cli:"opt --key-id"`
	LocaleId     string  `cli:"opt --locale-id"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
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

	if cmd.Content != "" {
		params.Content = cmd.Content
	}

	if cmd.Excluded != nil {
		params.Excluded = cmd.Excluded
	}

	if cmd.KeyId != "" {
		params.KeyId = cmd.KeyId
	}

	if cmd.LocaleId != "" {
		params.LocaleId = cmd.LocaleId
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationMachineTranslate struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *TranslationMachineTranslate) Run() error {

	defaults, e := ConfigDefaultParams()
	if e != nil {
		_ = defaults
		return e
	}

	defaultCredentials, e := ConfigDefaultCredentials()
	if e != nil {
		return e
	}

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	err = client.TranslationMachineTranslate(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return nil
}

type TranslationShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationUpdate struct {
	phraseapp.Credentials

	Content      string  `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	if cmd.Content != "" {
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationUpdate(cmd.ProjectId, cmd.Id, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByKey struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsByKey(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByLocale struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	LocaleId  string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsByLocale(cmd.ProjectId, cmd.LocaleId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsExclude struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsExclude(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsInclude struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsInclude(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsList struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsList(cmd.ProjectId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsSearch struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsSearch(cmd.ProjectId, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsUnverify struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsUnverify(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsVerify struct {
	phraseapp.Credentials

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.TranslationsVerify(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadCreate struct {
	phraseapp.Credentials

	ConvertEmoji       *bool   `cli:"opt --convert-emoji"`
	File               string  `cli:"opt --file"`
	FileFormat         *string `cli:"opt --file-format"`
	LocaleId           *string `cli:"opt --locale-id"`
	SkipUnverification *bool   `cli:"opt --skip-unverification"`
	SkipUploadTags     *bool   `cli:"opt --skip-upload-tags"`
	Tags               *string `cli:"opt --tags"`
	UpdateTranslations *bool   `cli:"opt --update-translations"`

	ProjectId string `cli:"arg required"`
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

	if cmd.File != "" {
		params.File = cmd.File
	}

	if cmd.FileFormat != nil {
		params.FileFormat = cmd.FileFormat
	}

	if cmd.LocaleId != nil {
		params.LocaleId = cmd.LocaleId
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.UploadCreate(cmd.ProjectId, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadShow struct {
	phraseapp.Credentials

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.UploadShow(cmd.ProjectId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionShow struct {
	phraseapp.Credentials

	ProjectId     string `cli:"arg required"`
	TranslationId string `cli:"arg required"`
	Id            string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.VersionShow(cmd.ProjectId, cmd.TranslationId, cmd.Id)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionsList struct {
	phraseapp.Credentials

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId     string `cli:"arg required"`
	TranslationId string `cli:"arg required"`
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

	client, err := phraseapp.NewClient(cmd.Credentials, defaultCredentials)
	if err != nil {
		return err
	}

	res, err := client.VersionsList(cmd.ProjectId, cmd.TranslationId, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}
