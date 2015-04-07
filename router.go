package main

import (
	"encoding/json"
	"fmt"
	"os"

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
	fmt.Printf("Built at 2015-04-07 09:45:13.352215775 +0200 CEST\n")
	return cli.ErrorHelpRequested
}

type AuthorizationCreate struct {
	phraseapp.AuthHandler

	Note   string `cli:"opt --note"`
	Scopes string `cli:"opt --scopes"`
}

func (cmd *AuthorizationCreate) Run() error {

	params := new(phraseapp.AuthorizationParams)

	if cmd.Note != "" {
		err := params.SetNote(cmd.Note)
		if err != nil {
			return err
		}
	}
	if cmd.Scopes != "" {
		err := params.SetScopes(cmd.Scopes)
		if err != nil {
			return err
		}
	}

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

	Note   string `cli:"opt --note"`
	Scopes string `cli:"opt --scopes"`

	Id string `cli:"arg required"`
}

func (cmd *AuthorizationUpdate) Run() error {

	params := new(phraseapp.AuthorizationParams)

	if cmd.Note != "" {
		err := params.SetNote(cmd.Note)
		if err != nil {
			return err
		}
	}
	if cmd.Scopes != "" {
		err := params.SetScopes(cmd.Scopes)
		if err != nil {
			return err
		}
	}

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

	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}

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

	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}

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

	if cmd.Message != "" {
		err := params.SetMessage(cmd.Message)
		if err != nil {
			return err
		}
	}

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

	if cmd.Message != "" {
		err := params.SetMessage(cmd.Message)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.CommentUpdate(cmd.ProjectId, cmd.KeyId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyCreate struct {
	phraseapp.AuthHandler

	DataType             string `cli:"opt --data-type"`
	Description          string `cli:"opt --description"`
	FormatValueType      string `cli:"opt --format-value-type"`
	MaxCharactersAllowed string `cli:"opt --max-characters-allowed"`
	Name                 string `cli:"opt --name"`
	NamePlural           string `cli:"opt --name-plural"`
	OriginalFile         string `cli:"opt --original-file"`
	Plural               string `cli:"opt --plural"`
	RemoveScreenshot     string `cli:"opt --remove-screenshot"`
	Screenshot           string `cli:"opt --screenshot"`
	Tags                 string `cli:"opt --tags"`
	Unformatted          string `cli:"opt --unformatted"`
	XmlSpacePreserve     string `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeyCreate) Run() error {

	params := new(phraseapp.TranslationKeyParams)

	if cmd.DataType != "" {
		err := params.SetDataType(cmd.DataType)
		if err != nil {
			return err
		}
	}
	if cmd.Description != "" {
		err := params.SetDescription(cmd.Description)
		if err != nil {
			return err
		}
	}
	if cmd.FormatValueType != "" {
		err := params.SetFormatValueType(cmd.FormatValueType)
		if err != nil {
			return err
		}
	}
	if cmd.MaxCharactersAllowed != "" {
		err := params.SetMaxCharactersAllowed(cmd.MaxCharactersAllowed)
		if err != nil {
			return err
		}
	}
	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.NamePlural != "" {
		err := params.SetNamePlural(cmd.NamePlural)
		if err != nil {
			return err
		}
	}
	if cmd.OriginalFile != "" {
		err := params.SetOriginalFile(cmd.OriginalFile)
		if err != nil {
			return err
		}
	}
	if cmd.Plural != "" {
		err := params.SetPlural(cmd.Plural)
		if err != nil {
			return err
		}
	}
	if cmd.RemoveScreenshot != "" {
		err := params.SetRemoveScreenshot(cmd.RemoveScreenshot)
		if err != nil {
			return err
		}
	}
	if cmd.Screenshot != "" {
		err := params.SetScreenshot(cmd.Screenshot)
		if err != nil {
			return err
		}
	}
	if cmd.Tags != "" {
		err := params.SetTags(cmd.Tags)
		if err != nil {
			return err
		}
	}
	if cmd.Unformatted != "" {
		err := params.SetUnformatted(cmd.Unformatted)
		if err != nil {
			return err
		}
	}
	if cmd.XmlSpacePreserve != "" {
		err := params.SetXmlSpacePreserve(cmd.XmlSpacePreserve)
		if err != nil {
			return err
		}
	}

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

	LocaleId   string `cli:"opt --locale-id"`
	Order      string `cli:"opt --order"`
	Sort       string `cli:"opt --sort"`
	Translated string `cli:"opt --translated"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *KeyList) Run() error {

	params := new(phraseapp.KeyListParams)

	if cmd.LocaleId != "" {
		err := params.SetLocaleId(cmd.LocaleId)
		if err != nil {
			return err
		}
	}
	if cmd.Order != "" {
		err := params.SetOrder(cmd.Order)
		if err != nil {
			return err
		}
	}
	if cmd.Sort != "" {
		err := params.SetSort(cmd.Sort)
		if err != nil {
			return err
		}
	}
	if cmd.Translated != "" {
		err := params.SetTranslated(cmd.Translated)
		if err != nil {
			return err
		}
	}

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

	DataType             string `cli:"opt --data-type"`
	Description          string `cli:"opt --description"`
	FormatValueType      string `cli:"opt --format-value-type"`
	MaxCharactersAllowed string `cli:"opt --max-characters-allowed"`
	Name                 string `cli:"opt --name"`
	NamePlural           string `cli:"opt --name-plural"`
	OriginalFile         string `cli:"opt --original-file"`
	Plural               string `cli:"opt --plural"`
	RemoveScreenshot     string `cli:"opt --remove-screenshot"`
	Screenshot           string `cli:"opt --screenshot"`
	Tags                 string `cli:"opt --tags"`
	Unformatted          string `cli:"opt --unformatted"`
	XmlSpacePreserve     string `cli:"opt --xml-space-preserve"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *KeyUpdate) Run() error {

	params := new(phraseapp.TranslationKeyParams)

	if cmd.DataType != "" {
		err := params.SetDataType(cmd.DataType)
		if err != nil {
			return err
		}
	}
	if cmd.Description != "" {
		err := params.SetDescription(cmd.Description)
		if err != nil {
			return err
		}
	}
	if cmd.FormatValueType != "" {
		err := params.SetFormatValueType(cmd.FormatValueType)
		if err != nil {
			return err
		}
	}
	if cmd.MaxCharactersAllowed != "" {
		err := params.SetMaxCharactersAllowed(cmd.MaxCharactersAllowed)
		if err != nil {
			return err
		}
	}
	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.NamePlural != "" {
		err := params.SetNamePlural(cmd.NamePlural)
		if err != nil {
			return err
		}
	}
	if cmd.OriginalFile != "" {
		err := params.SetOriginalFile(cmd.OriginalFile)
		if err != nil {
			return err
		}
	}
	if cmd.Plural != "" {
		err := params.SetPlural(cmd.Plural)
		if err != nil {
			return err
		}
	}
	if cmd.RemoveScreenshot != "" {
		err := params.SetRemoveScreenshot(cmd.RemoveScreenshot)
		if err != nil {
			return err
		}
	}
	if cmd.Screenshot != "" {
		err := params.SetScreenshot(cmd.Screenshot)
		if err != nil {
			return err
		}
	}
	if cmd.Tags != "" {
		err := params.SetTags(cmd.Tags)
		if err != nil {
			return err
		}
	}
	if cmd.Unformatted != "" {
		err := params.SetUnformatted(cmd.Unformatted)
		if err != nil {
			return err
		}
	}
	if cmd.XmlSpacePreserve != "" {
		err := params.SetXmlSpacePreserve(cmd.XmlSpacePreserve)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.KeyUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleCreate struct {
	phraseapp.AuthHandler

	Code           string `cli:"opt --code"`
	Default        string `cli:"opt --default"`
	Main           string `cli:"opt --main"`
	Name           string `cli:"opt --name"`
	Rtl            string `cli:"opt --rtl"`
	SourceLocaleId string `cli:"opt --source-locale-id"`

	ProjectId string `cli:"arg required"`
}

func (cmd *LocaleCreate) Run() error {

	params := new(phraseapp.LocaleParams)

	if cmd.Code != "" {
		err := params.SetCode(cmd.Code)
		if err != nil {
			return err
		}
	}
	if cmd.Default != "" {
		err := params.SetDefault(cmd.Default)
		if err != nil {
			return err
		}
	}
	if cmd.Main != "" {
		err := params.SetMain(cmd.Main)
		if err != nil {
			return err
		}
	}
	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.Rtl != "" {
		err := params.SetRtl(cmd.Rtl)
		if err != nil {
			return err
		}
	}
	if cmd.SourceLocaleId != "" {
		err := params.SetSourceLocaleId(cmd.SourceLocaleId)
		if err != nil {
			return err
		}
	}

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

	ConvertEmoji             string `cli:"opt --convert-emoji"`
	Format                   string `cli:"opt --format"`
	IncludeEmptyTranslations string `cli:"opt --include-empty-translations"`
	KeepNotranslateTags      string `cli:"opt --keep-notranslate-tags"`
	TagId                    string `cli:"opt --tag-id"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *LocaleDownload) Run() error {

	params := new(phraseapp.LocaleDownloadParams)

	if cmd.ConvertEmoji != "" {
		err := params.SetConvertEmoji(cmd.ConvertEmoji)
		if err != nil {
			return err
		}
	}
	if cmd.Format != "" {
		err := params.SetFormat(cmd.Format)
		if err != nil {
			return err
		}
	}
	if cmd.IncludeEmptyTranslations != "" {
		err := params.SetIncludeEmptyTranslations(cmd.IncludeEmptyTranslations)
		if err != nil {
			return err
		}
	}
	if cmd.KeepNotranslateTags != "" {
		err := params.SetKeepNotranslateTags(cmd.KeepNotranslateTags)
		if err != nil {
			return err
		}
	}
	if cmd.TagId != "" {
		err := params.SetTagId(cmd.TagId)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	err := phraseapp.LocaleDownload(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

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

	Code           string `cli:"opt --code"`
	Default        string `cli:"opt --default"`
	Main           string `cli:"opt --main"`
	Name           string `cli:"opt --name"`
	Rtl            string `cli:"opt --rtl"`
	SourceLocaleId string `cli:"opt --source-locale-id"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *LocaleUpdate) Run() error {

	params := new(phraseapp.LocaleParams)

	if cmd.Code != "" {
		err := params.SetCode(cmd.Code)
		if err != nil {
			return err
		}
	}
	if cmd.Default != "" {
		err := params.SetDefault(cmd.Default)
		if err != nil {
			return err
		}
	}
	if cmd.Main != "" {
		err := params.SetMain(cmd.Main)
		if err != nil {
			return err
		}
	}
	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.Rtl != "" {
		err := params.SetRtl(cmd.Rtl)
		if err != nil {
			return err
		}
	}
	if cmd.SourceLocaleId != "" {
		err := params.SetSourceLocaleId(cmd.SourceLocaleId)
		if err != nil {
			return err
		}
	}

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

	Category                         string `cli:"opt --category"`
	IncludeUntranslatedKeys          string `cli:"opt --include-untranslated-keys"`
	IncludeUnverifiedTranslations    string `cli:"opt --include-unverified-translations"`
	Lsp                              string `cli:"opt --lsp"`
	Message                          string `cli:"opt --message"`
	Priority                         string `cli:"opt --priority"`
	Quality                          string `cli:"opt --quality"`
	SourceLocaleId                   string `cli:"opt --source-locale-id"`
	StyleguideId                     string `cli:"opt --styleguide-id"`
	Tag                              string `cli:"opt --tag"`
	TargetLocaleIds                  string `cli:"opt --target-locale-ids"`
	TranslationType                  string `cli:"opt --translation-type"`
	UnverifyTranslationsUponDelivery string `cli:"opt --unverify-translations-upon-delivery"`

	ProjectId string `cli:"arg required"`
}

func (cmd *OrderCreate) Run() error {

	params := new(phraseapp.TranslationOrderParams)

	if cmd.Category != "" {
		err := params.SetCategory(cmd.Category)
		if err != nil {
			return err
		}
	}
	if cmd.IncludeUntranslatedKeys != "" {
		err := params.SetIncludeUntranslatedKeys(cmd.IncludeUntranslatedKeys)
		if err != nil {
			return err
		}
	}
	if cmd.IncludeUnverifiedTranslations != "" {
		err := params.SetIncludeUnverifiedTranslations(cmd.IncludeUnverifiedTranslations)
		if err != nil {
			return err
		}
	}
	if cmd.Lsp != "" {
		err := params.SetLsp(cmd.Lsp)
		if err != nil {
			return err
		}
	}
	if cmd.Message != "" {
		err := params.SetMessage(cmd.Message)
		if err != nil {
			return err
		}
	}
	if cmd.Priority != "" {
		err := params.SetPriority(cmd.Priority)
		if err != nil {
			return err
		}
	}
	if cmd.Quality != "" {
		err := params.SetQuality(cmd.Quality)
		if err != nil {
			return err
		}
	}
	if cmd.SourceLocaleId != "" {
		err := params.SetSourceLocaleId(cmd.SourceLocaleId)
		if err != nil {
			return err
		}
	}
	if cmd.StyleguideId != "" {
		err := params.SetStyleguideId(cmd.StyleguideId)
		if err != nil {
			return err
		}
	}
	if cmd.Tag != "" {
		err := params.SetTag(cmd.Tag)
		if err != nil {
			return err
		}
	}
	if cmd.TargetLocaleIds != "" {
		err := params.SetTargetLocaleIds(cmd.TargetLocaleIds)
		if err != nil {
			return err
		}
	}
	if cmd.TranslationType != "" {
		err := params.SetTranslationType(cmd.TranslationType)
		if err != nil {
			return err
		}
	}
	if cmd.UnverifyTranslationsUponDelivery != "" {
		err := params.SetUnverifyTranslationsUponDelivery(cmd.UnverifyTranslationsUponDelivery)
		if err != nil {
			return err
		}
	}

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
	SharesTranslationMemory string `cli:"opt --shares-translation-memory"`
}

func (cmd *ProjectCreate) Run() error {

	params := new(phraseapp.ProjectParams)

	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.SharesTranslationMemory != "" {
		err := params.SetSharesTranslationMemory(cmd.SharesTranslationMemory)
		if err != nil {
			return err
		}
	}

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
	SharesTranslationMemory string `cli:"opt --shares-translation-memory"`

	Id string `cli:"arg required"`
}

func (cmd *ProjectUpdate) Run() error {

	params := new(phraseapp.ProjectParams)

	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}
	if cmd.SharesTranslationMemory != "" {
		err := params.SetSharesTranslationMemory(cmd.SharesTranslationMemory)
		if err != nil {
			return err
		}
	}

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

	Audience           string `cli:"opt --audience"`
	Business           string `cli:"opt --business"`
	CompanyBranding    string `cli:"opt --company-branding"`
	Formatting         string `cli:"opt --formatting"`
	GlossaryTerms      string `cli:"opt --glossary-terms"`
	GrammarConsistency string `cli:"opt --grammar-consistency"`
	GrammaticalPerson  string `cli:"opt --grammatical-person"`
	LiteralTranslation string `cli:"opt --literal-translation"`
	OverallTone        string `cli:"opt --overall-tone"`
	Samples            string `cli:"opt --samples"`
	TargetAudience     string `cli:"opt --target-audience"`
	Title              string `cli:"opt --title"`
	VocabularyType     string `cli:"opt --vocabulary-type"`

	ProjectId string `cli:"arg required"`
}

func (cmd *StyleguideCreate) Run() error {

	params := new(phraseapp.StyleguideParams)

	if cmd.Audience != "" {
		err := params.SetAudience(cmd.Audience)
		if err != nil {
			return err
		}
	}
	if cmd.Business != "" {
		err := params.SetBusiness(cmd.Business)
		if err != nil {
			return err
		}
	}
	if cmd.CompanyBranding != "" {
		err := params.SetCompanyBranding(cmd.CompanyBranding)
		if err != nil {
			return err
		}
	}
	if cmd.Formatting != "" {
		err := params.SetFormatting(cmd.Formatting)
		if err != nil {
			return err
		}
	}
	if cmd.GlossaryTerms != "" {
		err := params.SetGlossaryTerms(cmd.GlossaryTerms)
		if err != nil {
			return err
		}
	}
	if cmd.GrammarConsistency != "" {
		err := params.SetGrammarConsistency(cmd.GrammarConsistency)
		if err != nil {
			return err
		}
	}
	if cmd.GrammaticalPerson != "" {
		err := params.SetGrammaticalPerson(cmd.GrammaticalPerson)
		if err != nil {
			return err
		}
	}
	if cmd.LiteralTranslation != "" {
		err := params.SetLiteralTranslation(cmd.LiteralTranslation)
		if err != nil {
			return err
		}
	}
	if cmd.OverallTone != "" {
		err := params.SetOverallTone(cmd.OverallTone)
		if err != nil {
			return err
		}
	}
	if cmd.Samples != "" {
		err := params.SetSamples(cmd.Samples)
		if err != nil {
			return err
		}
	}
	if cmd.TargetAudience != "" {
		err := params.SetTargetAudience(cmd.TargetAudience)
		if err != nil {
			return err
		}
	}
	if cmd.Title != "" {
		err := params.SetTitle(cmd.Title)
		if err != nil {
			return err
		}
	}
	if cmd.VocabularyType != "" {
		err := params.SetVocabularyType(cmd.VocabularyType)
		if err != nil {
			return err
		}
	}

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

	Audience           string `cli:"opt --audience"`
	Business           string `cli:"opt --business"`
	CompanyBranding    string `cli:"opt --company-branding"`
	Formatting         string `cli:"opt --formatting"`
	GlossaryTerms      string `cli:"opt --glossary-terms"`
	GrammarConsistency string `cli:"opt --grammar-consistency"`
	GrammaticalPerson  string `cli:"opt --grammatical-person"`
	LiteralTranslation string `cli:"opt --literal-translation"`
	OverallTone        string `cli:"opt --overall-tone"`
	Samples            string `cli:"opt --samples"`
	TargetAudience     string `cli:"opt --target-audience"`
	Title              string `cli:"opt --title"`
	VocabularyType     string `cli:"opt --vocabulary-type"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *StyleguideUpdate) Run() error {

	params := new(phraseapp.StyleguideParams)

	if cmd.Audience != "" {
		err := params.SetAudience(cmd.Audience)
		if err != nil {
			return err
		}
	}
	if cmd.Business != "" {
		err := params.SetBusiness(cmd.Business)
		if err != nil {
			return err
		}
	}
	if cmd.CompanyBranding != "" {
		err := params.SetCompanyBranding(cmd.CompanyBranding)
		if err != nil {
			return err
		}
	}
	if cmd.Formatting != "" {
		err := params.SetFormatting(cmd.Formatting)
		if err != nil {
			return err
		}
	}
	if cmd.GlossaryTerms != "" {
		err := params.SetGlossaryTerms(cmd.GlossaryTerms)
		if err != nil {
			return err
		}
	}
	if cmd.GrammarConsistency != "" {
		err := params.SetGrammarConsistency(cmd.GrammarConsistency)
		if err != nil {
			return err
		}
	}
	if cmd.GrammaticalPerson != "" {
		err := params.SetGrammaticalPerson(cmd.GrammaticalPerson)
		if err != nil {
			return err
		}
	}
	if cmd.LiteralTranslation != "" {
		err := params.SetLiteralTranslation(cmd.LiteralTranslation)
		if err != nil {
			return err
		}
	}
	if cmd.OverallTone != "" {
		err := params.SetOverallTone(cmd.OverallTone)
		if err != nil {
			return err
		}
	}
	if cmd.Samples != "" {
		err := params.SetSamples(cmd.Samples)
		if err != nil {
			return err
		}
	}
	if cmd.TargetAudience != "" {
		err := params.SetTargetAudience(cmd.TargetAudience)
		if err != nil {
			return err
		}
	}
	if cmd.Title != "" {
		err := params.SetTitle(cmd.Title)
		if err != nil {
			return err
		}
	}
	if cmd.VocabularyType != "" {
		err := params.SetVocabularyType(cmd.VocabularyType)
		if err != nil {
			return err
		}
	}

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

	if cmd.Name != "" {
		err := params.SetName(cmd.Name)
		if err != nil {
			return err
		}
	}

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

	Content      string `cli:"opt --content"`
	Excluded     string `cli:"opt --excluded"`
	KeyId        string `cli:"opt --key-id"`
	LocaleId     string `cli:"opt --locale-id"`
	PluralSuffix string `cli:"opt --plural-suffix"`
	Unverified   string `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationCreate) Run() error {

	params := new(phraseapp.TranslationParams)

	if cmd.Content != "" {
		err := params.SetContent(cmd.Content)
		if err != nil {
			return err
		}
	}
	if cmd.Excluded != "" {
		err := params.SetExcluded(cmd.Excluded)
		if err != nil {
			return err
		}
	}
	if cmd.KeyId != "" {
		err := params.SetKeyId(cmd.KeyId)
		if err != nil {
			return err
		}
	}
	if cmd.LocaleId != "" {
		err := params.SetLocaleId(cmd.LocaleId)
		if err != nil {
			return err
		}
	}
	if cmd.PluralSuffix != "" {
		err := params.SetPluralSuffix(cmd.PluralSuffix)
		if err != nil {
			return err
		}
	}
	if cmd.Unverified != "" {
		err := params.SetUnverified(cmd.Unverified)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationCreate(cmd.ProjectId, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationListAll struct {
	phraseapp.AuthHandler

	Order      string `cli:"opt --order"`
	Since      string `cli:"opt --since"`
	Sort       string `cli:"opt --sort"`
	Unverified string `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
}

func (cmd *TranslationListAll) Run() error {

	params := new(phraseapp.TranslationListAllParams)

	if cmd.Order != "" {
		err := params.SetOrder(cmd.Order)
		if err != nil {
			return err
		}
	}
	if cmd.Since != "" {
		err := params.SetSince(cmd.Since)
		if err != nil {
			return err
		}
	}
	if cmd.Sort != "" {
		err := params.SetSort(cmd.Sort)
		if err != nil {
			return err
		}
	}
	if cmd.Unverified != "" {
		err := params.SetUnverified(cmd.Unverified)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationListAll(cmd.ProjectId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationListKey struct {
	phraseapp.AuthHandler

	Order      string `cli:"opt --order"`
	Since      string `cli:"opt --since"`
	Sort       string `cli:"opt --sort"`
	Unverified string `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	KeyId     string `cli:"arg required"`
}

func (cmd *TranslationListKey) Run() error {

	params := new(phraseapp.TranslationListKeyParams)

	if cmd.Order != "" {
		err := params.SetOrder(cmd.Order)
		if err != nil {
			return err
		}
	}
	if cmd.Since != "" {
		err := params.SetSince(cmd.Since)
		if err != nil {
			return err
		}
	}
	if cmd.Sort != "" {
		err := params.SetSort(cmd.Sort)
		if err != nil {
			return err
		}
	}
	if cmd.Unverified != "" {
		err := params.SetUnverified(cmd.Unverified)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationListKey(cmd.ProjectId, cmd.KeyId, cmd.Page, cmd.PerPage, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationListLocale struct {
	phraseapp.AuthHandler

	Order      string `cli:"opt --order"`
	Since      string `cli:"opt --since"`
	Sort       string `cli:"opt --sort"`
	Unverified string `cli:"opt --unverified"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectId string `cli:"arg required"`
	LocaleId  string `cli:"arg required"`
}

func (cmd *TranslationListLocale) Run() error {

	params := new(phraseapp.TranslationListLocaleParams)

	if cmd.Order != "" {
		err := params.SetOrder(cmd.Order)
		if err != nil {
			return err
		}
	}
	if cmd.Since != "" {
		err := params.SetSince(cmd.Since)
		if err != nil {
			return err
		}
	}
	if cmd.Sort != "" {
		err := params.SetSort(cmd.Sort)
		if err != nil {
			return err
		}
	}
	if cmd.Unverified != "" {
		err := params.SetUnverified(cmd.Unverified)
		if err != nil {
			return err
		}
	}

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

	Content      string `cli:"opt --content"`
	Excluded     string `cli:"opt --excluded"`
	PluralSuffix string `cli:"opt --plural-suffix"`
	Unverified   string `cli:"opt --unverified"`

	ProjectId string `cli:"arg required"`
	Id        string `cli:"arg required"`
}

func (cmd *TranslationUpdate) Run() error {

	params := new(phraseapp.TranslationUpdateParams)

	if cmd.Content != "" {
		err := params.SetContent(cmd.Content)
		if err != nil {
			return err
		}
	}
	if cmd.Excluded != "" {
		err := params.SetExcluded(cmd.Excluded)
		if err != nil {
			return err
		}
	}
	if cmd.PluralSuffix != "" {
		err := params.SetPluralSuffix(cmd.PluralSuffix)
		if err != nil {
			return err
		}
	}
	if cmd.Unverified != "" {
		err := params.SetUnverified(cmd.Unverified)
		if err != nil {
			return err
		}
	}

	phraseapp.RegisterAuthHandler(&cmd.AuthHandler)

	res, err := phraseapp.TranslationUpdate(cmd.ProjectId, cmd.Id, params)
	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadCreate struct {
	phraseapp.AuthHandler

	ConvertEmoji       string `cli:"opt --convert-emoji"`
	File               string `cli:"opt --file"`
	Format             string `cli:"opt --format"`
	FormatOptions      string `cli:"opt --format-options"`
	LocaleId           string `cli:"opt --locale-id"`
	SkipUnverification string `cli:"opt --skip-unverification"`
	SkipUploadTags     string `cli:"opt --skip-upload-tags"`
	Tags               string `cli:"opt --tags"`
	UpdateTranslations string `cli:"opt --update-translations"`

	ProjectId string `cli:"arg required"`
}

func (cmd *UploadCreate) Run() error {

	params := new(phraseapp.LocaleFileImportParams)

	if cmd.ConvertEmoji != "" {
		err := params.SetConvertEmoji(cmd.ConvertEmoji)
		if err != nil {
			return err
		}
	}
	if cmd.File != "" {
		err := params.SetFile(cmd.File)
		if err != nil {
			return err
		}
	}
	if cmd.Format != "" {
		err := params.SetFormat(cmd.Format)
		if err != nil {
			return err
		}
	}
	if cmd.FormatOptions != "" {
		err := params.SetFormatOptions(cmd.FormatOptions)
		if err != nil {
			return err
		}
	}
	if cmd.LocaleId != "" {
		err := params.SetLocaleId(cmd.LocaleId)
		if err != nil {
			return err
		}
	}
	if cmd.SkipUnverification != "" {
		err := params.SetSkipUnverification(cmd.SkipUnverification)
		if err != nil {
			return err
		}
	}
	if cmd.SkipUploadTags != "" {
		err := params.SetSkipUploadTags(cmd.SkipUploadTags)
		if err != nil {
			return err
		}
	}
	if cmd.Tags != "" {
		err := params.SetTags(cmd.Tags)
		if err != nil {
			return err
		}
	}
	if cmd.UpdateTranslations != "" {
		err := params.SetUpdateTranslations(cmd.UpdateTranslations)
		if err != nil {
			return err
		}
	}

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
