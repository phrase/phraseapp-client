package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func router() *cli.Router {
	r := cli.NewRouter()

	r.Register("authorization/create", &AuthorizationCreate{}, "Create a new authorization.")
	r.Register("authorization/delete", &AuthorizationDelete{}, "Delete an existing authorization. API calls using that token will stop working.")
	r.Register("authorization/show", &AuthorizationShow{}, "Get details on a single authorization.")
	r.Register("authorization/update", &AuthorizationUpdate{}, "Update an existing authorization.")
	r.Register("authorizations/list", &AuthorizationsList{}, "List all your authorizations.")
	r.Register("comment/create", &CommentCreate{}, "Create a new comment for a key.")
	r.Register("comment/delete", &CommentDelete{}, "Delete an existing comment.")
	r.Register("comment/mark/check", &CommentMarkCheck{}, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")
	r.Register("comment/mark/read", &CommentMarkRead{}, "Mark a comment as read")
	r.Register("comment/mark/unread", &CommentMarkUnread{}, "Mark a comment as unread")
	r.Register("comment/show", &CommentShow{}, "Get details on a single comment.")
	r.Register("comment/update", &CommentUpdate{}, "Update an existing comment.")
	r.Register("comments/list", &CommentsList{}, "List all comments for a key.")
	r.Register("exclude_rule/create", &ExcludeRuleCreate{}, "Create a new blacklisted key.")
	r.Register("exclude_rule/delete", &ExcludeRuleDelete{}, "Delete an existing blacklisted key.")
	r.Register("exclude_rule/show", &ExcludeRuleShow{}, "Get details on a single blacklisted key for a given project.")
	r.Register("exclude_rule/update", &ExcludeRuleUpdate{}, "Update an existing blacklisted key.")
	r.Register("exclude_rules/index", &ExcludeRulesIndex{}, "List all blacklisted keys for the given project.")
	r.Register("formats/list", &FormatsList{}, "Get a handy list of all localization file formats supported in PhraseApp.")
	r.Register("key/create", &KeyCreate{}, "Create a new key.")
	r.Register("key/delete", &KeyDelete{}, "Delete an existing key.")
	r.Register("key/show", &KeyShow{}, "Get details on a single key for a given project.")
	r.Register("key/update", &KeyUpdate{}, "Update an existing key.")
	r.Register("keys/delete", &KeysDelete{}, "Delete all keys matching query. Same constraints as list.")
	r.Register("keys/list", &KeysList{}, "List all keys for the given project. Alternatively you can POST requests to /search.")
	r.Register("keys/search", &KeysSearch{}, "List all keys for the given project matching query.")
	r.Register("keys/tag", &KeysTag{}, "Tags all keys matching query. Same constraints as list.")
	r.Register("keys/untag", &KeysUntag{}, "Removes specified tags from keys matching query.")
	r.Register("locale/create", &LocaleCreate{}, "Create a new locale.")
	r.Register("locale/delete", &LocaleDelete{}, "Delete an existing locale.")
	r.Register("locale/download", &LocaleDownload{}, "Download a locale in a specific file format.")
	r.Register("locale/show", &LocaleShow{}, "Get details on a single locale for a given project.")
	r.Register("locale/update", &LocaleUpdate{}, "Update an existing locale.")
	r.Register("locales/list", &LocalesList{}, "List all locales for the given project.")
	r.Register("order/confirm", &OrderConfirm{}, "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")
	r.Register("order/create", &OrderCreate{}, "Create a new order. Access token scope must include <code>orders.create</code>.")
	r.Register("order/delete", &OrderDelete{}, "Cancel an existing order. Must not yet be confirmed.")
	r.Register("order/show", &OrderShow{}, "Get details on a single order.")
	r.Register("orders/list", &OrdersList{}, "List all orders for the given project.")
	r.Register("project/create", &ProjectCreate{}, "Create a new project.")
	r.Register("project/delete", &ProjectDelete{}, "Delete an existing project.")
	r.Register("project/show", &ProjectShow{}, "Get details on a single project.")
	r.Register("project/update", &ProjectUpdate{}, "Update an existing project.")
	r.Register("projects/list", &ProjectsList{}, "List all projects the current user has access to.")
	r.Register("show/user", &ShowUser{}, "Show details for current User.")
	r.Register("styleguide/create", &StyleguideCreate{}, "Create a new style guide.")
	r.Register("styleguide/delete", &StyleguideDelete{}, "Delete an existing style guide.")
	r.Register("styleguide/show", &StyleguideShow{}, "Get details on a single style guide.")
	r.Register("styleguide/update", &StyleguideUpdate{}, "Update an existing style guide.")
	r.Register("styleguides/list", &StyleguidesList{}, "List all styleguides for the given project.")
	r.Register("tag/create", &TagCreate{}, "Create a new tag.")
	r.Register("tag/delete", &TagDelete{}, "Delete an existing tag.")
	r.Register("tag/show", &TagShow{}, "Get details and progress information on a single tag for a given project.")
	r.Register("tags/list", &TagsList{}, "List all tags for the given project.")
	r.Register("translation/create", &TranslationCreate{}, "Create a translation.")
	r.Register("translation/show", &TranslationShow{}, "Get details on a single translation.")
	r.Register("translation/update", &TranslationUpdate{}, "Update an existing translation.")
	r.Register("translations/by_key", &TranslationsByKey{}, "List translations for a specific key.")
	r.Register("translations/by_locale", &TranslationsByLocale{}, "List translations for a specific locale.")
	r.Register("translations/exclude", &TranslationsExclude{}, "Exclude translations matching query from locale export.")
	r.Register("translations/include", &TranslationsInclude{}, "Include translations matching query in locale export")
	r.Register("translations/list", &TranslationsList{}, "List translations for the given project. Alternatively, POST request to /search")
	r.Register("translations/search", &TranslationsSearch{}, "List translations for the given project if you exceed GET request limitations on translations list.")
	r.Register("translations/unverify", &TranslationsUnverify{}, "Mark translations matching query as unverified")
	r.Register("translations/verify", &TranslationsVerify{}, "Verify translations matching query.")
	r.Register("upload/create", &UploadCreate{}, "Upload a new language file. Creates necessary resources in your project.")
	r.Register("upload/show", &UploadShow{}, "View details and summary for a single upload.")
	r.Register("version/show", &VersionShow{}, "Get details on a single version.")
	r.Register("versions/list", &VersionsList{}, "List all versions for the given translation.")

	r.RegisterFunc("help", helpCommand, "Help for this client")

	return r
}

func helpCommand() error {
	fmt.Printf("Built at 2015-06-11 14:50:01.969678142 +0200 CEST\n")
	return cli.ErrorHelpRequested
}

type AuthorizationCreate struct {
	phraseapp.AuthHandler

	ExpiresAt *time.Time `cli:"opt --expires-at"`
	Note      string     `cli:"opt --note"`
	Scopes    []string   `cli:"opt --scopes"`
}

func (cmd *AuthorizationCreate) Run() error {
	params := new(phraseapp.AuthorizationParams)
	params.ExpiresAt = cmd.ExpiresAt
	params.Note = cmd.Note
	params.Scopes = cmd.Scopes

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationCreate(params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationDelete struct {
	phraseapp.AuthHandler

	Id string `cli:"arg required"`
}

func (cmd *AuthorizationDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.AuthorizationDelete(cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type AuthorizationShow struct {
	phraseapp.AuthHandler

	Id string `cli:"arg required"`
}

func (cmd *AuthorizationShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationShow(cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationUpdate struct {
	phraseapp.AuthHandler

	ExpiresAt *time.Time `cli:"opt --expires-at"`
	Note      string     `cli:"opt --note"`
	Scopes    []string   `cli:"opt --scopes"`

	Id string `cli:"arg required"`
}

func (cmd *AuthorizationUpdate) Run() error {
	params := new(phraseapp.AuthorizationParams)
	params.ExpiresAt = cmd.ExpiresAt
	params.Note = cmd.Note
	params.Scopes = cmd.Scopes

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationUpdate(cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *AuthorizationsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationsList(cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentCreate struct {
	phraseapp.AuthHandler

	Message string `cli:"opt --message"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *CommentCreate) Run() error {
	params := new(phraseapp.CommentParams)
	params.Message = cmd.Message

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentCreate(cmd.ProjectId, cmd.KeyId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.CommentDelete(cmd.ProjectId, cmd.KeyId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type CommentMarkCheck struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentMarkCheck) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.CommentMarkCheck(cmd.ProjectId, cmd.KeyId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type CommentMarkRead struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentMarkRead) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.CommentMarkRead(cmd.ProjectId, cmd.KeyId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type CommentMarkUnread struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentMarkUnread) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.CommentMarkUnread(cmd.ProjectId, cmd.KeyId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type CommentShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentShow(cmd.ProjectId, cmd.KeyId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentUpdate struct {
	phraseapp.AuthHandler

	Message string `cli:"opt --message"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *CommentUpdate) Run() error {
	params := new(phraseapp.CommentParams)
	params.Message = cmd.Message

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentUpdate(cmd.ProjectId, cmd.KeyId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *CommentsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentsList(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleCreate struct {
	phraseapp.AuthHandler

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
}

func (cmd *ExcludeRuleCreate) Run() error {
	params := new(phraseapp.ExcludeRuleParams)
	params.Name = cmd.Name

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ExcludeRuleCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.ExcludeRuleDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type ExcludeRuleShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ExcludeRuleShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRuleUpdate struct {
	phraseapp.AuthHandler

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *ExcludeRuleUpdate) Run() error {
	params := new(phraseapp.ExcludeRuleParams)
	params.Name = cmd.Name

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ExcludeRuleUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ExcludeRulesIndex struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *ExcludeRulesIndex) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ExcludeRulesIndex(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type FormatsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *FormatsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.FormatsList(cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyCreate struct {
	phraseapp.AuthHandler

	DataType              *string  `cli:"opt --data-type"`
	Description           *string  `cli:"opt --description"`
	LocalizedFormatKey    *string  `cli:"opt --localized-format-key"`
	LocalizedFormatString *string  `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64   `cli:"opt --max-characters-allowed"`
	Name                  string   `cli:"opt --name"`
	NamePlural            *string  `cli:"opt --name-plural"`
	OriginalFile          *string  `cli:"opt --original-file"`
	Plural                *bool    `cli:"opt --plural"`
	RemoveScreenshot      *bool    `cli:"opt --remove-screenshot"`
	Screenshot            *string  `cli:"opt --screenshot"`
	Tags                  []string `cli:"opt --tags"`
	Unformatted           *bool    `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool    `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeyCreate) Run() error {
	params := new(phraseapp.TranslationKeyParams)
	params.DataType = cmd.DataType
	params.Description = cmd.Description
	params.LocalizedFormatKey = cmd.LocalizedFormatKey
	params.LocalizedFormatString = cmd.LocalizedFormatString
	params.MaxCharactersAllowed = cmd.MaxCharactersAllowed
	params.Name = cmd.Name
	params.NamePlural = cmd.NamePlural
	params.OriginalFile = cmd.OriginalFile
	params.Plural = cmd.Plural
	params.RemoveScreenshot = cmd.RemoveScreenshot
	params.Screenshot = cmd.Screenshot
	params.Tags = cmd.Tags
	params.Unformatted = cmd.Unformatted
	params.XmlSpacePreserve = cmd.XmlSpacePreserve

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeyCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *KeyDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.KeyDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type KeyShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *KeyShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeyShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyUpdate struct {
	phraseapp.AuthHandler

	DataType              *string  `cli:"opt --data-type"`
	Description           *string  `cli:"opt --description"`
	LocalizedFormatKey    *string  `cli:"opt --localized-format-key"`
	LocalizedFormatString *string  `cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64   `cli:"opt --max-characters-allowed"`
	Name                  string   `cli:"opt --name"`
	NamePlural            *string  `cli:"opt --name-plural"`
	OriginalFile          *string  `cli:"opt --original-file"`
	Plural                *bool    `cli:"opt --plural"`
	RemoveScreenshot      *bool    `cli:"opt --remove-screenshot"`
	Screenshot            *string  `cli:"opt --screenshot"`
	Tags                  []string `cli:"opt --tags"`
	Unformatted           *bool    `cli:"opt --unformatted"`
	XmlSpacePreserve      *bool    `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *KeyUpdate) Run() error {
	params := new(phraseapp.TranslationKeyParams)
	params.DataType = cmd.DataType
	params.Description = cmd.Description
	params.LocalizedFormatKey = cmd.LocalizedFormatKey
	params.LocalizedFormatString = cmd.LocalizedFormatString
	params.MaxCharactersAllowed = cmd.MaxCharactersAllowed
	params.Name = cmd.Name
	params.NamePlural = cmd.NamePlural
	params.OriginalFile = cmd.OriginalFile
	params.Plural = cmd.Plural
	params.RemoveScreenshot = cmd.RemoveScreenshot
	params.Screenshot = cmd.Screenshot
	params.Tags = cmd.Tags
	params.Unformatted = cmd.Unformatted
	params.XmlSpacePreserve = cmd.XmlSpacePreserve

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeyUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysDelete struct {
	phraseapp.AuthHandler

	LocaleId *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeysDelete) Run() error {
	params := new(phraseapp.KeysDeleteParams)
	params.LocaleId = cmd.LocaleId
	params.Q = cmd.Q

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.KeysDelete(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return nil
}

type KeysList struct {
	phraseapp.AuthHandler

	LocaleId *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeysList) Run() error {
	params := new(phraseapp.KeysListParams)
	params.LocaleId = cmd.LocaleId
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeysList(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysSearch struct {
	phraseapp.AuthHandler

	LocaleId *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeysSearch) Run() error {
	params := new(phraseapp.KeysSearchParams)
	params.LocaleId = cmd.LocaleId
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeysSearch(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeysTag struct {
	phraseapp.AuthHandler

	LocaleId *string  `cli:"opt --locale-id"`
	Q        *string  `cli:"opt --query"`
	Tags     []string `cli:"opt --tags"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeysTag) Run() error {
	params := new(phraseapp.KeysTagParams)
	params.LocaleId = cmd.LocaleId
	params.Q = cmd.Q
	params.Tags = cmd.Tags

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.KeysTag(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return nil
}

type KeysUntag struct {
	phraseapp.AuthHandler

	LocaleId *string  `cli:"opt --locale-id"`
	Q        *string  `cli:"opt --query"`
	Tags     []string `cli:"opt --tags"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeysUntag) Run() error {
	params := new(phraseapp.KeysUntagParams)
	params.LocaleId = cmd.LocaleId
	params.Q = cmd.Q
	params.Tags = cmd.Tags

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.KeysUntag(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return nil
}

type LocaleCreate struct {
	phraseapp.AuthHandler

	Code           string  `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           string  `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleId *string `cli:"opt --source-locale-id"`

	ProjectId string `cli:"arg required"`
}

func (cmd *LocaleCreate) Run() error {
	params := new(phraseapp.LocaleParams)
	params.Code = cmd.Code
	params.Default = cmd.Default
	params.Main = cmd.Main
	params.Name = cmd.Name
	params.Rtl = cmd.Rtl
	params.SourceLocaleId = cmd.SourceLocaleId

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocaleCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *LocaleDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.LocaleDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type LocaleDownload struct {
	phraseapp.AuthHandler

	ConvertEmoji             *bool                   `cli:"opt --convert-emoji"`
	FileFormat               string                  `cli:"opt --file-format"`
	FormatOptions            *map[string]interface{} `cli:"opt --format-options"`
	IncludeEmptyTranslations *bool                   `cli:"opt --include-empty-translations"`
	KeepNotranslateTags      *bool                   `cli:"opt --keep-notranslate-tags"`
	TagId                    *string                 `cli:"opt --tag-id"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *LocaleDownload) Run() error {
	params := new(phraseapp.LocaleDownloadParams)
	params.ConvertEmoji = cmd.ConvertEmoji
	params.FileFormat = cmd.FileFormat
	params.FormatOptions = cmd.FormatOptions
	params.IncludeEmptyTranslations = cmd.IncludeEmptyTranslations
	params.KeepNotranslateTags = cmd.KeepNotranslateTags
	params.TagId = cmd.TagId

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocaleDownload(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	fmt.Println(string(res))
	return nil
}

type LocaleShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *LocaleShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocaleShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleUpdate struct {
	phraseapp.AuthHandler

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
	params := new(phraseapp.LocaleParams)
	params.Code = cmd.Code
	params.Default = cmd.Default
	params.Main = cmd.Main
	params.Name = cmd.Name
	params.Rtl = cmd.Rtl
	params.SourceLocaleId = cmd.SourceLocaleId

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocaleUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocalesList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *LocalesList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocalesList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderConfirm struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *OrderConfirm) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.OrderConfirm(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderCreate struct {
	phraseapp.AuthHandler

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
	params := new(phraseapp.TranslationOrderParams)
	params.Category = cmd.Category
	params.IncludeUntranslatedKeys = cmd.IncludeUntranslatedKeys
	params.IncludeUnverifiedTranslations = cmd.IncludeUnverifiedTranslations
	params.Lsp = cmd.Lsp
	params.Message = cmd.Message
	params.Priority = cmd.Priority
	params.Quality = cmd.Quality
	params.SourceLocaleId = cmd.SourceLocaleId
	params.StyleguideId = cmd.StyleguideId
	params.Tag = cmd.Tag
	params.TargetLocaleIds = cmd.TargetLocaleIds
	params.TranslationType = cmd.TranslationType
	params.UnverifyTranslationsUponDelivery = cmd.UnverifyTranslationsUponDelivery

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.OrderCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *OrderDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.OrderDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type OrderShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *OrderShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.OrderShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrdersList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *OrdersList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.OrdersList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectCreate struct {
	phraseapp.AuthHandler

	Name                    string `cli:"opt --name"`
	SharesTranslationMemory *bool  `cli:"opt --shares-translation-memory"`
}

func (cmd *ProjectCreate) Run() error {
	params := new(phraseapp.ProjectParams)
	params.Name = cmd.Name
	params.SharesTranslationMemory = cmd.SharesTranslationMemory

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ProjectCreate(params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectDelete struct {
	phraseapp.AuthHandler

	Id string `cli:"arg required"`
}

func (cmd *ProjectDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.ProjectDelete(cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type ProjectShow struct {
	phraseapp.AuthHandler

	Id string `cli:"arg required"`
}

func (cmd *ProjectShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ProjectShow(cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectUpdate struct {
	phraseapp.AuthHandler

	Name                    string `cli:"opt --name"`
	SharesTranslationMemory *bool  `cli:"opt --shares-translation-memory"`

	Id string `cli:"arg required"`
}

func (cmd *ProjectUpdate) Run() error {
	params := new(phraseapp.ProjectParams)
	params.Name = cmd.Name
	params.SharesTranslationMemory = cmd.SharesTranslationMemory

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ProjectUpdate(cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ProjectsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *ProjectsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ProjectsList(cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ShowUser struct {
	phraseapp.AuthHandler
}

func (cmd *ShowUser) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ShowUser()
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideCreate struct {
	phraseapp.AuthHandler

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
	params := new(phraseapp.StyleguideParams)
	params.Audience = cmd.Audience
	params.Business = cmd.Business
	params.CompanyBranding = cmd.CompanyBranding
	params.Formatting = cmd.Formatting
	params.GlossaryTerms = cmd.GlossaryTerms
	params.GrammarConsistency = cmd.GrammarConsistency
	params.GrammaticalPerson = cmd.GrammaticalPerson
	params.LiteralTranslation = cmd.LiteralTranslation
	params.OverallTone = cmd.OverallTone
	params.Samples = cmd.Samples
	params.TargetAudience = cmd.TargetAudience
	params.Title = cmd.Title
	params.VocabularyType = cmd.VocabularyType

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.StyleguideCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *StyleguideDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.StyleguideDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type StyleguideShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *StyleguideShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.StyleguideShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguideUpdate struct {
	phraseapp.AuthHandler

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
	params := new(phraseapp.StyleguideParams)
	params.Audience = cmd.Audience
	params.Business = cmd.Business
	params.CompanyBranding = cmd.CompanyBranding
	params.Formatting = cmd.Formatting
	params.GlossaryTerms = cmd.GlossaryTerms
	params.GrammarConsistency = cmd.GrammarConsistency
	params.GrammaticalPerson = cmd.GrammaticalPerson
	params.LiteralTranslation = cmd.LiteralTranslation
	params.OverallTone = cmd.OverallTone
	params.Samples = cmd.Samples
	params.TargetAudience = cmd.TargetAudience
	params.Title = cmd.Title
	params.VocabularyType = cmd.VocabularyType

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.StyleguideUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type StyleguidesList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *StyleguidesList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.StyleguidesList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagCreate struct {
	phraseapp.AuthHandler

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TagCreate) Run() error {
	params := new(phraseapp.TagParams)
	params.Name = cmd.Name

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TagCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.TagDelete(cmd.ProjectId, cmd.Name)
	if err != nil {
		return err
	}

	return nil
}

type TagShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TagShow(cmd.ProjectId, cmd.Name)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TagsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TagsList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationCreate struct {
	phraseapp.AuthHandler

	Content      string  `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	KeyId        string  `cli:"opt --key-id"`
	LocaleId     string  `cli:"opt --locale-id"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationCreate) Run() error {
	params := new(phraseapp.TranslationParams)
	params.Content = cmd.Content
	params.Excluded = cmd.Excluded
	params.KeyId = cmd.KeyId
	params.LocaleId = cmd.LocaleId
	params.PluralSuffix = cmd.PluralSuffix
	params.Unverified = cmd.Unverified

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *TranslationShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationUpdate struct {
	phraseapp.AuthHandler

	Content      string  `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *TranslationUpdate) Run() error {
	params := new(phraseapp.TranslationUpdateParams)
	params.Content = cmd.Content
	params.Excluded = cmd.Excluded
	params.PluralSuffix = cmd.PluralSuffix
	params.Unverified = cmd.Unverified

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByKey struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *TranslationsByKey) Run() error {
	params := new(phraseapp.TranslationsByKeyParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsByKey(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsByLocale struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	LocaleId  string `cli:"arg required"`
}

func (cmd *TranslationsByLocale) Run() error {
	params := new(phraseapp.TranslationsByLocaleParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsByLocale(cmd.ProjectId, cmd.LocaleId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsExclude struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsExclude) Run() error {
	params := new(phraseapp.TranslationsExcludeParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsExclude(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsInclude struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsInclude) Run() error {
	params := new(phraseapp.TranslationsIncludeParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsInclude(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsList struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsList) Run() error {
	params := new(phraseapp.TranslationsListParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsList(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsSearch struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsSearch) Run() error {
	params := new(phraseapp.TranslationsSearchParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsSearch(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsUnverify struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsUnverify) Run() error {
	params := new(phraseapp.TranslationsUnverifyParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsUnverify(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationsVerify struct {
	phraseapp.AuthHandler

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query"`
	Sort  *string `cli:"opt --sort"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationsVerify) Run() error {
	params := new(phraseapp.TranslationsVerifyParams)
	params.Order = cmd.Order
	params.Q = cmd.Q
	params.Sort = cmd.Sort

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationsVerify(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadCreate struct {
	phraseapp.AuthHandler

	ConvertEmoji       *bool                   `cli:"opt --convert-emoji"`
	File               string                  `cli:"opt --file"`
	FileFormat         *string                 `cli:"opt --file-format"`
	FormatOptions      *map[string]interface{} `cli:"opt --format-options"`
	LocaleId           *string                 `cli:"opt --locale-id"`
	SkipUnverification *bool                   `cli:"opt --skip-unverification"`
	SkipUploadTags     *bool                   `cli:"opt --skip-upload-tags"`
	Tags               []string                `cli:"opt --tags"`
	UpdateTranslations *bool                   `cli:"opt --update-translations"`

	ProjectId string `cli:"arg required"`
}

func (cmd *UploadCreate) Run() error {
	params := new(phraseapp.LocaleFileImportParams)
	params.ConvertEmoji = cmd.ConvertEmoji
	params.File = cmd.File
	params.FileFormat = cmd.FileFormat
	params.FormatOptions = cmd.FormatOptions
	params.LocaleId = cmd.LocaleId
	params.SkipUnverification = cmd.SkipUnverification
	params.SkipUploadTags = cmd.SkipUploadTags
	params.Tags = cmd.Tags
	params.UpdateTranslations = cmd.UpdateTranslations

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.UploadCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *UploadShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.UploadShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionShow struct {
	phraseapp.AuthHandler

	ProjectId     string `cli:"arg required"`
	TranslationId string `cli:"arg required"`
	Id            string `cli:"arg required"`
}

func (cmd *VersionShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.VersionShow(cmd.ProjectId, cmd.TranslationId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionsList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId     string `cli:"arg required"`
	TranslationId string `cli:"arg required"`
}

func (cmd *VersionsList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.VersionsList(cmd.ProjectId, cmd.TranslationId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}
