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
	r.Register("authorization/delete", &AuthorizationDelete{}, "Delete an existing authorization. Please note that this will revoke access for that token, so API calls using that token will stop working.")
	r.Register("authorization/list", &AuthorizationList{}, "List all your authorizations.")
	r.Register("authorization/show", &AuthorizationShow{}, "Get details on a single authorization.")
	r.Register("authorization/update", &AuthorizationUpdate{}, "Update an existing authorization.")
	r.Register("blacklist/key/create", &BlacklistKeyCreate{}, "Create a new blacklisted key.")
	r.Register("blacklist/key/delete", &BlacklistKeyDelete{}, "Delete an existing blacklisted key.")
	r.Register("blacklist/key/show", &BlacklistKeyShow{}, "Get details on a single blacklisted key for a given project.")
	r.Register("blacklist/key/update", &BlacklistKeyUpdate{}, "Update an existing blacklisted key.")
	r.Register("blacklist/show", &BlacklistShow{}, "List all blacklisted keys for the given project.")
	r.Register("comment/create", &CommentCreate{}, "Create a new comment for a key.")
	r.Register("comment/delete", &CommentDelete{}, "Delete an existing comment.")
	r.Register("comment/list", &CommentList{}, "List all comments for a key.")
	r.Register("comment/mark/check", &CommentMarkCheck{}, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")
	r.Register("comment/mark/read", &CommentMarkRead{}, "Mark a comment as read")
	r.Register("comment/mark/unread", &CommentMarkUnread{}, "Mark a comment as unread")
	r.Register("comment/show", &CommentShow{}, "Get details on a single comment.")
	r.Register("comment/update", &CommentUpdate{}, "Update an existing comment.")
	r.Register("key/create", &KeyCreate{}, "Create a new key.")
	r.Register("key/delete", &KeyDelete{}, "Delete an existing key.")
	r.Register("key/list", &KeyList{}, "List all keys for the given project.")
	r.Register("key/show", &KeyShow{}, "Get details on a single key for a given project.")
	r.Register("key/update", &KeyUpdate{}, "Update an existing key.")
	r.Register("locale/create", &LocaleCreate{}, "Create a new locale.")
	r.Register("locale/delete", &LocaleDelete{}, "Delete an existing locale.")
	r.Register("locale/download", &LocaleDownload{}, "Download a locale in a specific file format.")
	r.Register("locale/list", &LocaleList{}, "List all locales for the given project.")
	r.Register("locale/show", &LocaleShow{}, "Get details on a single locale for a given project.")
	r.Register("locale/update", &LocaleUpdate{}, "Update an existing locale.")
	r.Register("order/confirm", &OrderConfirm{}, "Confirm an existing order. Sends the order to the language service provider for processing. Please note that your access token must include the <code>orders.create</code> scope to confirm orders.")
	r.Register("order/create", &OrderCreate{}, "Create a new order. Please note that your access token must include the <code>orders.create</code> scope to create orders.")
	r.Register("order/delete", &OrderDelete{}, "Cancel an existing order. Must not yet be confirmed.")
	r.Register("order/list", &OrderList{}, "List all orders for the given project.")
	r.Register("order/show", &OrderShow{}, "Get details on a single order.")
	r.Register("project/create", &ProjectCreate{}, "Create a new project.")
	r.Register("project/delete", &ProjectDelete{}, "Delete an existing project.")
	r.Register("project/list", &ProjectList{}, "List all projects the current user has access to.")
	r.Register("project/show", &ProjectShow{}, "Get details on a single project.")
	r.Register("project/update", &ProjectUpdate{}, "Update an existing project.")
	r.Register("show/user", &ShowUser{}, "Show details for current User.")
	r.Register("styleguide/create", &StyleguideCreate{}, "Create a new style guide.")
	r.Register("styleguide/delete", &StyleguideDelete{}, "Delete an existing style guide.")
	r.Register("styleguide/list", &StyleguideList{}, "List all styleguides for the given project.")
	r.Register("styleguide/show", &StyleguideShow{}, "Get details on a single style guide.")
	r.Register("styleguide/update", &StyleguideUpdate{}, "Update an existing style guide.")
	r.Register("tag/create", &TagCreate{}, "Create a new tag.")
	r.Register("tag/delete", &TagDelete{}, "Delete an existing tag.")
	r.Register("tag/list", &TagList{}, "List all tags for the given project.")
	r.Register("tag/show", &TagShow{}, "Get details and progress information on a single tag for a given project.")
	r.Register("translation/create", &TranslationCreate{}, "Create a translation.")
	r.Register("translation/list/all", &TranslationListAll{}, "List translations for the given project.")
	r.Register("translation/list/key", &TranslationListKey{}, "List translations for a specific key.")
	r.Register("translation/list/locale", &TranslationListLocale{}, "List translations for a specific locale.")
	r.Register("translation/show", &TranslationShow{}, "Get details on a single translation.")
	r.Register("translation/update", &TranslationUpdate{}, "Update an existing translation.")
	r.Register("upload/create", &UploadCreate{}, "Upload a new file to your project. This will extract all new content such as keys, translations, locales, tags etc. and store them in your project.")
	r.Register("upload/show", &UploadShow{}, "View details and summary for a single upload.")
	r.Register("version/list", &VersionList{}, "List all versions for the given translation.")
	r.Register("version/show", &VersionShow{}, "Get details on a single version.")

	r.RegisterFunc("help", helpCommand, "Help for this client")

	return r
}

func helpCommand() error {
	fmt.Printf("Built at 2015-04-21 09:25:48.973559884 +0200 CEST\n")
	return cli.ErrorHelpRequested
}

type AuthorizationCreate struct {
	phraseapp.AuthHandler

	Note   string   `cli:"opt --note"`
	Scopes []string `cli:"opt --scopes"`
}

func (cmd *AuthorizationCreate) Run() error {
	params := new(phraseapp.AuthorizationParams)
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

type AuthorizationList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *AuthorizationList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationList(cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

	Note   string   `cli:"opt --note"`
	Scopes []string `cli:"opt --scopes"`

	Id string `cli:"arg required"`
}

func (cmd *AuthorizationUpdate) Run() error {
	params := new(phraseapp.AuthorizationParams)
	params.Note = cmd.Note
	params.Scopes = cmd.Scopes

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.AuthorizationUpdate(cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistKeyCreate struct {
	phraseapp.AuthHandler

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
}

func (cmd *BlacklistKeyCreate) Run() error {
	params := new(phraseapp.BlacklistedKeyParams)
	params.Name = cmd.Name

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.BlacklistKeyCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistKeyDelete struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *BlacklistKeyDelete) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.BlacklistKeyDelete(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return nil
}

type BlacklistKeyShow struct {
	phraseapp.AuthHandler

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *BlacklistKeyShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.BlacklistKeyShow(cmd.ProjectId, cmd.Id)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistKeyUpdate struct {
	phraseapp.AuthHandler

	Name string `cli:"opt --name"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *BlacklistKeyUpdate) Run() error {
	params := new(phraseapp.BlacklistedKeyParams)
	params.Name = cmd.Name

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.BlacklistKeyUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BlacklistShow struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *BlacklistShow) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.BlacklistShow(cmd.ProjectId, cmd.Page, cmd.PerPage)
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

type CommentList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *CommentList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentList(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type KeyCreate struct {
	phraseapp.AuthHandler

	DataType             *string  `cli:"opt --data-type"`
	Description          *string  `cli:"opt --description"`
	FormatValueType      *string  `cli:"opt --format-value-type"`
	MaxCharactersAllowed *int64   `cli:"opt --max-characters-allowed"`
	Name                 string   `cli:"opt --name"`
	NamePlural           *string  `cli:"opt --name-plural"`
	OriginalFile         *string  `cli:"opt --original-file"`
	Plural               *bool    `cli:"opt --plural"`
	RemoveScreenshot     *bool    `cli:"opt --remove-screenshot"`
	Screenshot           *string  `cli:"opt --screenshot"`
	Tags                 []string `cli:"opt --tags"`
	Unformatted          *bool    `cli:"opt --unformatted"`
	XmlSpacePreserve     *bool    `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeyCreate) Run() error {
	params := new(phraseapp.TranslationKeyParams)
	params.DataType = cmd.DataType
	params.Description = cmd.Description
	params.FormatValueType = cmd.FormatValueType
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

type KeyList struct {
	phraseapp.AuthHandler

	LocaleId   *string `cli:"opt --locale-id"`
	Order      *string `cli:"opt --order"`
	Sort       *string `cli:"opt --sort"`
	Translated *bool   `cli:"opt --translated"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeyList) Run() error {
	params := new(phraseapp.KeyListParams)
	params.LocaleId = cmd.LocaleId
	params.Order = cmd.Order
	params.Sort = cmd.Sort
	params.Translated = cmd.Translated

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeyList(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

	DataType             *string  `cli:"opt --data-type"`
	Description          *string  `cli:"opt --description"`
	FormatValueType      *string  `cli:"opt --format-value-type"`
	MaxCharactersAllowed *int64   `cli:"opt --max-characters-allowed"`
	Name                 string   `cli:"opt --name"`
	NamePlural           *string  `cli:"opt --name-plural"`
	OriginalFile         *string  `cli:"opt --original-file"`
	Plural               *bool    `cli:"opt --plural"`
	RemoveScreenshot     *bool    `cli:"opt --remove-screenshot"`
	Screenshot           *string  `cli:"opt --screenshot"`
	Tags                 []string `cli:"opt --tags"`
	Unformatted          *bool    `cli:"opt --unformatted"`
	XmlSpacePreserve     *bool    `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *KeyUpdate) Run() error {
	params := new(phraseapp.TranslationKeyParams)
	params.DataType = cmd.DataType
	params.Description = cmd.Description
	params.FormatValueType = cmd.FormatValueType
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
	Format                   string                  `cli:"opt --format"`
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
	params.Format = cmd.Format
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

type LocaleList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *LocaleList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.LocaleList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type OrderList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *OrderList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.OrderList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type ProjectList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *ProjectList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.ProjectList(cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type StyleguideList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *StyleguideList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.StyleguideList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type TagList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TagList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TagList(cmd.ProjectId, cmd.Page, cmd.PerPage)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

type TranslationListAll struct {
	phraseapp.AuthHandler

	Order      *string    `cli:"opt --order"`
	Since      *time.Time `cli:"opt --since"`
	Sort       *string    `cli:"opt --sort"`
	Unverified *bool      `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationListAll) Run() error {
	params := new(phraseapp.TranslationListAllParams)
	params.Order = cmd.Order
	params.Since = cmd.Since
	params.Sort = cmd.Sort
	params.Unverified = cmd.Unverified

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationListAll(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationListKey struct {
	phraseapp.AuthHandler

	Order      *string    `cli:"opt --order"`
	Since      *time.Time `cli:"opt --since"`
	Sort       *string    `cli:"opt --sort"`
	Unverified *bool      `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *TranslationListKey) Run() error {
	params := new(phraseapp.TranslationListKeyParams)
	params.Order = cmd.Order
	params.Since = cmd.Since
	params.Sort = cmd.Sort
	params.Unverified = cmd.Unverified

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationListKey(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationListLocale struct {
	phraseapp.AuthHandler

	Order      *string    `cli:"opt --order"`
	Since      *time.Time `cli:"opt --since"`
	Sort       *string    `cli:"opt --sort"`
	Unverified *bool      `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	LocaleId  string `cli:"arg required"`
}

func (cmd *TranslationListLocale) Run() error {
	params := new(phraseapp.TranslationListLocaleParams)
	params.Order = cmd.Order
	params.Since = cmd.Since
	params.Sort = cmd.Sort
	params.Unverified = cmd.Unverified

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationListLocale(cmd.ProjectId, cmd.LocaleId, cmd.Page, cmd.PerPage, params)
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

type UploadCreate struct {
	phraseapp.AuthHandler

	ConvertEmoji       *bool                   `cli:"opt --convert-emoji"`
	File               string                  `cli:"opt --file"`
	Format             *string                 `cli:"opt --format"`
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
	params.Format = cmd.Format
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

type VersionList struct {
	phraseapp.AuthHandler

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId     string `cli:"arg required"`
	TranslationId string `cli:"arg required"`
}

func (cmd *VersionList) Run() error {

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.VersionList(cmd.ProjectId, cmd.TranslationId, cmd.Page, cmd.PerPage)
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
