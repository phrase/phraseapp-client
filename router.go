package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const PHRASEAPP_CLIENT_VERSION = "1.1.10"

func router(cfg *phraseapp.Config) *cli.Router {
	r := cli.NewRouter()

	r.Register("authorization/create", &AuthorizationCreate{Config: cfg}, "Create a new authorization.")

	r.Register("authorization/delete", &AuthorizationDelete{Config: cfg}, "Delete an existing authorization. API calls using that token will stop working.")

	r.Register("authorization/show", &AuthorizationShow{Config: cfg}, "Get details on a single authorization.")

	r.Register("authorization/update", &AuthorizationUpdate{Config: cfg}, "Update an existing authorization.")

	actionAuthorizationsList := &AuthorizationsList{Config: cfg}
	if cfg.Page != nil {
		actionAuthorizationsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionAuthorizationsList.PerPage = *cfg.PerPage
	}
	r.Register("authorizations/list", actionAuthorizationsList, "List all your authorizations.")

	actionBlacklistedKeyCreate := &BlacklistedKeyCreate{Config: cfg}
	actionBlacklistedKeyCreate.ProjectID = cfg.ProjectID
	r.Register("blacklisted_key/create", actionBlacklistedKeyCreate, "Create a new rule for blacklisting keys.")

	actionBlacklistedKeyDelete := &BlacklistedKeyDelete{Config: cfg}
	actionBlacklistedKeyDelete.ProjectID = cfg.ProjectID
	r.Register("blacklisted_key/delete", actionBlacklistedKeyDelete, "Delete an existing rule for blacklisting keys.")

	actionBlacklistedKeyShow := &BlacklistedKeyShow{Config: cfg}
	actionBlacklistedKeyShow.ProjectID = cfg.ProjectID
	r.Register("blacklisted_key/show", actionBlacklistedKeyShow, "Get details on a single rule for blacklisting keys for a given project.")

	actionBlacklistedKeyUpdate := &BlacklistedKeyUpdate{Config: cfg}
	actionBlacklistedKeyUpdate.ProjectID = cfg.ProjectID
	r.Register("blacklisted_key/update", actionBlacklistedKeyUpdate, "Update an existing rule for blacklisting keys.")

	actionBlacklistedKeysList := &BlacklistedKeysList{Config: cfg}
	actionBlacklistedKeysList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionBlacklistedKeysList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionBlacklistedKeysList.PerPage = *cfg.PerPage
	}
	r.Register("blacklisted_keys/list", actionBlacklistedKeysList, "List all rules for blacklisting keys for the given project.")

	actionCommentCreate := &CommentCreate{Config: cfg}
	actionCommentCreate.ProjectID = cfg.ProjectID
	r.Register("comment/create", actionCommentCreate, "Create a new comment for a key.")

	actionCommentDelete := &CommentDelete{Config: cfg}
	actionCommentDelete.ProjectID = cfg.ProjectID
	r.Register("comment/delete", actionCommentDelete, "Delete an existing comment.")

	actionCommentMarkCheck := &CommentMarkCheck{Config: cfg}
	actionCommentMarkCheck.ProjectID = cfg.ProjectID
	r.Register("comment/mark/check", actionCommentMarkCheck, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")

	actionCommentMarkRead := &CommentMarkRead{Config: cfg}
	actionCommentMarkRead.ProjectID = cfg.ProjectID
	r.Register("comment/mark/read", actionCommentMarkRead, "Mark a comment as read.")

	actionCommentMarkUnread := &CommentMarkUnread{Config: cfg}
	actionCommentMarkUnread.ProjectID = cfg.ProjectID
	r.Register("comment/mark/unread", actionCommentMarkUnread, "Mark a comment as unread.")

	actionCommentShow := &CommentShow{Config: cfg}
	actionCommentShow.ProjectID = cfg.ProjectID
	r.Register("comment/show", actionCommentShow, "Get details on a single comment.")

	actionCommentUpdate := &CommentUpdate{Config: cfg}
	actionCommentUpdate.ProjectID = cfg.ProjectID
	r.Register("comment/update", actionCommentUpdate, "Update an existing comment.")

	actionCommentsList := &CommentsList{Config: cfg}
	actionCommentsList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionCommentsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionCommentsList.PerPage = *cfg.PerPage
	}
	r.Register("comments/list", actionCommentsList, "List all comments for a key.")

	actionFormatsList := &FormatsList{Config: cfg}
	if cfg.Page != nil {
		actionFormatsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionFormatsList.PerPage = *cfg.PerPage
	}
	r.Register("formats/list", actionFormatsList, "Get a handy list of all localization file formats supported in PhraseApp.")

	actionKeyCreate := &KeyCreate{Config: cfg}
	actionKeyCreate.ProjectID = cfg.ProjectID
	r.Register("key/create", actionKeyCreate, "Create a new key.")

	actionKeyDelete := &KeyDelete{Config: cfg}
	actionKeyDelete.ProjectID = cfg.ProjectID
	r.Register("key/delete", actionKeyDelete, "Delete an existing key.")

	actionKeyShow := &KeyShow{Config: cfg}
	actionKeyShow.ProjectID = cfg.ProjectID
	r.Register("key/show", actionKeyShow, "Get details on a single key for a given project.")

	actionKeyUpdate := &KeyUpdate{Config: cfg}
	actionKeyUpdate.ProjectID = cfg.ProjectID
	r.Register("key/update", actionKeyUpdate, "Update an existing key.")

	actionKeysDelete := &KeysDelete{Config: cfg}
	actionKeysDelete.ProjectID = cfg.ProjectID
	r.Register("keys/delete", actionKeysDelete, "Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.")

	actionKeysList := &KeysList{Config: cfg}
	actionKeysList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionKeysList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionKeysList.PerPage = *cfg.PerPage
	}
	r.Register("keys/list", actionKeysList, "List all keys for the given project. Alternatively you can POST requests to /search.")

	actionKeysSearch := &KeysSearch{Config: cfg}
	actionKeysSearch.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionKeysSearch.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionKeysSearch.PerPage = *cfg.PerPage
	}
	r.Register("keys/search", actionKeysSearch, "Search keys for the given project matching query.")

	actionKeysTag := &KeysTag{Config: cfg}
	actionKeysTag.ProjectID = cfg.ProjectID
	r.Register("keys/tag", actionKeysTag, "Tags all keys matching query. Same constraints as list.")

	actionKeysUntag := &KeysUntag{Config: cfg}
	actionKeysUntag.ProjectID = cfg.ProjectID
	r.Register("keys/untag", actionKeysUntag, "Removes specified tags from keys matching query.")

	actionLocaleCreate := &LocaleCreate{Config: cfg}
	actionLocaleCreate.ProjectID = cfg.ProjectID
	r.Register("locale/create", actionLocaleCreate, "Create a new locale.")

	actionLocaleDelete := &LocaleDelete{Config: cfg}
	actionLocaleDelete.ProjectID = cfg.ProjectID
	r.Register("locale/delete", actionLocaleDelete, "Delete an existing locale.")

	actionLocaleDownload := &LocaleDownload{Config: cfg}
	actionLocaleDownload.ProjectID = cfg.ProjectID
	r.Register("locale/download", actionLocaleDownload, "Download a locale in a specific file format.")

	actionLocaleShow := &LocaleShow{Config: cfg}
	actionLocaleShow.ProjectID = cfg.ProjectID
	r.Register("locale/show", actionLocaleShow, "Get details on a single locale for a given project.")

	actionLocaleUpdate := &LocaleUpdate{Config: cfg}
	actionLocaleUpdate.ProjectID = cfg.ProjectID
	r.Register("locale/update", actionLocaleUpdate, "Update an existing locale.")

	actionLocalesList := &LocalesList{Config: cfg}
	actionLocalesList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionLocalesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionLocalesList.PerPage = *cfg.PerPage
	}
	r.Register("locales/list", actionLocalesList, "List all locales for the given project.")

	actionOrderConfirm := &OrderConfirm{Config: cfg}
	actionOrderConfirm.ProjectID = cfg.ProjectID
	r.Register("order/confirm", actionOrderConfirm, "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")

	actionOrderCreate := &OrderCreate{Config: cfg}
	actionOrderCreate.ProjectID = cfg.ProjectID
	r.Register("order/create", actionOrderCreate, "Create a new order. Access token scope must include <code>orders.create</code>.")

	actionOrderDelete := &OrderDelete{Config: cfg}
	actionOrderDelete.ProjectID = cfg.ProjectID
	r.Register("order/delete", actionOrderDelete, "Cancel an existing order. Must not yet be confirmed.")

	actionOrderShow := &OrderShow{Config: cfg}
	actionOrderShow.ProjectID = cfg.ProjectID
	r.Register("order/show", actionOrderShow, "Get details on a single order.")

	actionOrdersList := &OrdersList{Config: cfg}
	actionOrdersList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionOrdersList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionOrdersList.PerPage = *cfg.PerPage
	}
	r.Register("orders/list", actionOrdersList, "List all orders for the given project.")

	r.Register("project/create", &ProjectCreate{Config: cfg}, "Create a new project.")

	r.Register("project/delete", &ProjectDelete{Config: cfg}, "Delete an existing project.")

	r.Register("project/show", &ProjectShow{Config: cfg}, "Get details on a single project.")

	r.Register("project/update", &ProjectUpdate{Config: cfg}, "Update an existing project.")

	actionProjectsList := &ProjectsList{Config: cfg}
	if cfg.Page != nil {
		actionProjectsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionProjectsList.PerPage = *cfg.PerPage
	}
	r.Register("projects/list", actionProjectsList, "List all projects the current user has access to.")

	r.Register("show/user", &ShowUser{Config: cfg}, "Show details for current User.")

	actionStyleguideCreate := &StyleguideCreate{Config: cfg}
	actionStyleguideCreate.ProjectID = cfg.ProjectID
	r.Register("styleguide/create", actionStyleguideCreate, "Create a new style guide.")

	actionStyleguideDelete := &StyleguideDelete{Config: cfg}
	actionStyleguideDelete.ProjectID = cfg.ProjectID
	r.Register("styleguide/delete", actionStyleguideDelete, "Delete an existing style guide.")

	actionStyleguideShow := &StyleguideShow{Config: cfg}
	actionStyleguideShow.ProjectID = cfg.ProjectID
	r.Register("styleguide/show", actionStyleguideShow, "Get details on a single style guide.")

	actionStyleguideUpdate := &StyleguideUpdate{Config: cfg}
	actionStyleguideUpdate.ProjectID = cfg.ProjectID
	r.Register("styleguide/update", actionStyleguideUpdate, "Update an existing style guide.")

	actionStyleguidesList := &StyleguidesList{Config: cfg}
	actionStyleguidesList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionStyleguidesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionStyleguidesList.PerPage = *cfg.PerPage
	}
	r.Register("styleguides/list", actionStyleguidesList, "List all styleguides for the given project.")

	actionTagCreate := &TagCreate{Config: cfg}
	actionTagCreate.ProjectID = cfg.ProjectID
	r.Register("tag/create", actionTagCreate, "Create a new tag.")

	actionTagDelete := &TagDelete{Config: cfg}
	actionTagDelete.ProjectID = cfg.ProjectID
	r.Register("tag/delete", actionTagDelete, "Delete an existing tag.")

	actionTagShow := &TagShow{Config: cfg}
	actionTagShow.ProjectID = cfg.ProjectID
	r.Register("tag/show", actionTagShow, "Get details and progress information on a single tag for a given project.")

	actionTagsList := &TagsList{Config: cfg}
	actionTagsList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionTagsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTagsList.PerPage = *cfg.PerPage
	}
	r.Register("tags/list", actionTagsList, "List all tags for the given project.")

	actionTranslationCreate := &TranslationCreate{Config: cfg}
	actionTranslationCreate.ProjectID = cfg.ProjectID
	r.Register("translation/create", actionTranslationCreate, "Create a translation.")

	actionTranslationShow := &TranslationShow{Config: cfg}
	actionTranslationShow.ProjectID = cfg.ProjectID
	r.Register("translation/show", actionTranslationShow, "Get details on a single translation.")

	actionTranslationUpdate := &TranslationUpdate{Config: cfg}
	actionTranslationUpdate.ProjectID = cfg.ProjectID
	r.Register("translation/update", actionTranslationUpdate, "Update an existing translation.")

	actionTranslationsByKey := &TranslationsByKey{Config: cfg}
	actionTranslationsByKey.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionTranslationsByKey.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsByKey.PerPage = *cfg.PerPage
	}
	r.Register("translations/by_key", actionTranslationsByKey, "List translations for a specific key.")

	actionTranslationsByLocale := &TranslationsByLocale{Config: cfg}
	actionTranslationsByLocale.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionTranslationsByLocale.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsByLocale.PerPage = *cfg.PerPage
	}
	r.Register("translations/by_locale", actionTranslationsByLocale, "List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")

	actionTranslationsExclude := &TranslationsExclude{Config: cfg}
	actionTranslationsExclude.ProjectID = cfg.ProjectID
	r.Register("translations/exclude", actionTranslationsExclude, "Exclude translations matching query from locale export.")

	actionTranslationsInclude := &TranslationsInclude{Config: cfg}
	actionTranslationsInclude.ProjectID = cfg.ProjectID
	r.Register("translations/include", actionTranslationsInclude, "Include translations matching query in locale export.")

	actionTranslationsList := &TranslationsList{Config: cfg}
	actionTranslationsList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionTranslationsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsList.PerPage = *cfg.PerPage
	}
	r.Register("translations/list", actionTranslationsList, "List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")

	actionTranslationsSearch := &TranslationsSearch{Config: cfg}
	actionTranslationsSearch.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionTranslationsSearch.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsSearch.PerPage = *cfg.PerPage
	}
	r.Register("translations/search", actionTranslationsSearch, "List translations for the given project if you exceed GET request limitations on translations list. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")

	actionTranslationsUnverify := &TranslationsUnverify{Config: cfg}
	actionTranslationsUnverify.ProjectID = cfg.ProjectID
	r.Register("translations/unverify", actionTranslationsUnverify, "Mark translations matching query as unverified.")

	actionTranslationsVerify := &TranslationsVerify{Config: cfg}
	actionTranslationsVerify.ProjectID = cfg.ProjectID
	r.Register("translations/verify", actionTranslationsVerify, "Verify translations matching query.")

	actionUploadCreate := &UploadCreate{Config: cfg}
	actionUploadCreate.ProjectID = cfg.ProjectID
	r.Register("upload/create", actionUploadCreate, "Upload a new language file. Creates necessary resources in your project.")

	actionUploadShow := &UploadShow{Config: cfg}
	actionUploadShow.ProjectID = cfg.ProjectID
	r.Register("upload/show", actionUploadShow, "View details and summary for a single upload.")

	actionUploadsList := &UploadsList{Config: cfg}
	actionUploadsList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionUploadsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionUploadsList.PerPage = *cfg.PerPage
	}
	r.Register("uploads/list", actionUploadsList, "List all uploads for the given project.")

	actionVersionShow := &VersionShow{Config: cfg}
	actionVersionShow.ProjectID = cfg.ProjectID
	r.Register("version/show", actionVersionShow, "Get details on a single version.")

	actionVersionsList := &VersionsList{Config: cfg}
	actionVersionsList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionVersionsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionVersionsList.PerPage = *cfg.PerPage
	}
	r.Register("versions/list", actionVersionsList, "List all versions for the given translation.")

	actionWebhookCreate := &WebhookCreate{Config: cfg}
	actionWebhookCreate.ProjectID = cfg.ProjectID
	r.Register("webhook/create", actionWebhookCreate, "Create a new webhook.")

	actionWebhookDelete := &WebhookDelete{Config: cfg}
	actionWebhookDelete.ProjectID = cfg.ProjectID
	r.Register("webhook/delete", actionWebhookDelete, "Delete an existing webhook.")

	actionWebhookShow := &WebhookShow{Config: cfg}
	actionWebhookShow.ProjectID = cfg.ProjectID
	r.Register("webhook/show", actionWebhookShow, "Get details on a single webhook.")

	actionWebhookTest := &WebhookTest{Config: cfg}
	actionWebhookTest.ProjectID = cfg.ProjectID
	r.Register("webhook/test", actionWebhookTest, "Perform a test request for a webhook.")

	actionWebhookUpdate := &WebhookUpdate{Config: cfg}
	actionWebhookUpdate.ProjectID = cfg.ProjectID
	r.Register("webhook/update", actionWebhookUpdate, "Update an existing webhook.")

	actionWebhooksList := &WebhooksList{Config: cfg}
	actionWebhooksList.ProjectID = cfg.ProjectID
	if cfg.Page != nil {
		actionWebhooksList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionWebhooksList.PerPage = *cfg.PerPage
	}
	r.Register("webhooks/list", actionWebhooksList, "List all webhooks for the given project.")

	r.Register("pull", &PullCommand{Config: cfg}, "Download locales from your PhraseApp project.\n  You can provide parameters supported by the locales#download endpoint http://docs.phraseapp.com/api/v2/locales/#download\n  in your configuration (.phraseapp.yml) for each source.\n  See our configuration guide for more information http://docs.phraseapp.com/developers/cli/configuration/")

	r.Register("push", &PushCommand{Config: cfg}, "Upload locales to your PhraseApp project.\n  You can provide parameters supported by the uploads#create endpoint http://docs.phraseapp.com/api/v2/uploads/#create\n  in your configuration (.phraseapp.yml) for each source.\n  See our configuration guide for more information http://docs.phraseapp.com/developers/cli/configuration/")

	r.Register("init", &WizardCommand{}, "Configure your PhraseApp client.")

	r.RegisterFunc("info", infoCommand, "Info about version and revision of this client")

	return r
}

func infoCommand() error {
	fmt.Printf("Built at 2016-02-08 12:22:23.17965581 +0100 CET\n")
	fmt.Println("PhraseApp Client version:", "1.1.10")
	fmt.Println("PhraseApp API Client revision:", "bc3bb842f1ef88d846a66cd3f2e64d3e3ff3edc4")
	fmt.Println("PhraseApp Client revision:", "3a26373ae3efe050d85d7d0833398fde87c85f89")
	fmt.Println("PhraseApp Docs revision:", "509eaae478f03a2110146345db57da95cd0eccd9")
	return nil
}

type AuthorizationCreate struct {
	*phraseapp.Config

	ExpiresAt **time.Time `cli:"opt --expires-at"`
	Note      *string     `cli:"opt --note"`
	Scopes    []string    `cli:"opt --scopes"`
}

func (cmd *AuthorizationCreate) Run() error {

	params := new(phraseapp.AuthorizationParams)

	val, defaultsPresent := cmd.Config.Defaults["authorization/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ExpiresAt **time.Time `cli:"opt --expires-at"`
	Note      *string     `cli:"opt --note"`
	Scopes    []string    `cli:"opt --scopes"`

	ID string `cli:"arg required"`
}

func (cmd *AuthorizationUpdate) Run() error {

	params := new(phraseapp.AuthorizationParams)

	val, defaultsPresent := cmd.Config.Defaults["authorization/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *AuthorizationsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
}

func (cmd *BlacklistedKeyCreate) Run() error {

	params := new(phraseapp.BlacklistedKeyParams)

	val, defaultsPresent := cmd.Config.Defaults["blacklisted_key/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *BlacklistedKeyUpdate) Run() error {

	params := new(phraseapp.BlacklistedKeyParams)

	val, defaultsPresent := cmd.Config.Defaults["blacklisted_key/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *BlacklistedKeysList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Message *string `cli:"opt --message"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *CommentCreate) Run() error {

	params := new(phraseapp.CommentParams)

	val, defaultsPresent := cmd.Config.Defaults["comment/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Message != nil {
		params.Message = cmd.Message
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkCheck) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkRead) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentMarkUnread) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Message *string `cli:"opt --message"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *CommentUpdate) Run() error {

	params := new(phraseapp.CommentParams)

	val, defaultsPresent := cmd.Config.Defaults["comment/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Message != nil {
		params.Message = cmd.Message
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *CommentsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *FormatsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

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

	params := new(phraseapp.TranslationKeyParams)

	val, defaultsPresent := cmd.Config.Defaults["key/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *KeyDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *KeyShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

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

	params := new(phraseapp.TranslationKeyParams)

	val, defaultsPresent := cmd.Config.Defaults["key/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysDelete) Run() error {

	params := new(phraseapp.KeysDeleteParams)

	val, defaultsPresent := cmd.Config.Defaults["keys/delete"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.LocaleID != nil {
		params.LocaleID = cmd.LocaleID
	}

	if cmd.Q != nil {
		params.Q = cmd.Q
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	LocaleID *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query -q"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysList) Run() error {

	params := new(phraseapp.KeysListParams)

	val, defaultsPresent := cmd.Config.Defaults["keys/list"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	LocaleID *string `cli:"opt --locale-id"`
	Order    *string `cli:"opt --order"`
	Q        *string `cli:"opt --query -q"`
	Sort     *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysSearch) Run() error {

	params := new(phraseapp.KeysSearchParams)

	val, defaultsPresent := cmd.Config.Defaults["keys/search"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`
	Tags     *string `cli:"opt --tags"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysTag) Run() error {

	params := new(phraseapp.KeysTagParams)

	val, defaultsPresent := cmd.Config.Defaults["keys/tag"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	LocaleID *string `cli:"opt --locale-id"`
	Q        *string `cli:"opt --query -q"`
	Tags     *string `cli:"opt --tags"`

	ProjectID string `cli:"arg required"`
}

func (cmd *KeysUntag) Run() error {

	params := new(phraseapp.KeysUntagParams)

	val, defaultsPresent := cmd.Config.Defaults["keys/untag"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Code           *string `cli:"opt --code"`
	Default        *bool   `cli:"opt --default"`
	Main           *bool   `cli:"opt --main"`
	Name           *string `cli:"opt --name"`
	Rtl            *bool   `cli:"opt --rtl"`
	SourceLocaleID *string `cli:"opt --source-locale-id"`

	ProjectID string `cli:"arg required"`
}

func (cmd *LocaleCreate) Run() error {

	params := new(phraseapp.LocaleParams)

	val, defaultsPresent := cmd.Config.Defaults["locale/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ConvertEmoji               bool                    `cli:"opt --convert-emoji"`
	Encoding                   *string                 `cli:"opt --encoding"`
	FallbackLocaleID           *string                 `cli:"opt --fallback-locale-id"`
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

	params := new(phraseapp.LocaleDownloadParams)

	val, defaultsPresent := cmd.Config.Defaults["locale/download"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	params.ConvertEmoji = cmd.ConvertEmoji

	if cmd.Encoding != nil {
		params.Encoding = cmd.Encoding
	}

	if cmd.FallbackLocaleID != nil {
		params.FallbackLocaleID = cmd.FallbackLocaleID
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *LocaleShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

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

	params := new(phraseapp.LocaleParams)

	val, defaultsPresent := cmd.Config.Defaults["locale/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *LocalesList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderConfirm) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Category                         *string  `cli:"opt --category"`
	IncludeUntranslatedKeys          *bool    `cli:"opt --include-untranslated-keys"`
	IncludeUnverifiedTranslations    *bool    `cli:"opt --include-unverified-translations"`
	Lsp                              *string  `cli:"opt --lsp"`
	Message                          *string  `cli:"opt --message"`
	Priority                         *bool    `cli:"opt --priority"`
	Quality                          *bool    `cli:"opt --quality"`
	SourceLocaleID                   *string  `cli:"opt --source-locale-id"`
	StyleguideID                     *string  `cli:"opt --styleguide-id"`
	Tag                              *string  `cli:"opt --tag"`
	TargetLocaleIDs                  []string `cli:"opt --target-locale-ids"`
	TranslationType                  *string  `cli:"opt --translation-type"`
	UnverifyTranslationsUponDelivery *bool    `cli:"opt --unverify-translations-upon-delivery"`

	ProjectID string `cli:"arg required"`
}

func (cmd *OrderCreate) Run() error {

	params := new(phraseapp.TranslationOrderParams)

	val, defaultsPresent := cmd.Config.Defaults["order/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *OrderShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *OrdersList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	MainFormat              *string `cli:"opt --main-format"`
	Name                    *string `cli:"opt --name"`
	SharesTranslationMemory *bool   `cli:"opt --shares-translation-memory"`
}

func (cmd *ProjectCreate) Run() error {

	params := new(phraseapp.ProjectParams)

	val, defaultsPresent := cmd.Config.Defaults["project/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ID string `cli:"arg required"`
}

func (cmd *ProjectDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ID string `cli:"arg required"`
}

func (cmd *ProjectShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	MainFormat              *string `cli:"opt --main-format"`
	Name                    *string `cli:"opt --name"`
	SharesTranslationMemory *bool   `cli:"opt --shares-translation-memory"`

	ID string `cli:"arg required"`
}

func (cmd *ProjectUpdate) Run() error {

	params := new(phraseapp.ProjectParams)

	val, defaultsPresent := cmd.Config.Defaults["project/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func (cmd *ProjectsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config
}

func (cmd *ShowUser) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

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

	params := new(phraseapp.StyleguideParams)

	val, defaultsPresent := cmd.Config.Defaults["styleguide/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *StyleguideDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *StyleguideShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

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

	params := new(phraseapp.StyleguideParams)

	val, defaultsPresent := cmd.Config.Defaults["styleguide/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *StyleguidesList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Name *string `cli:"opt --name"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TagCreate) Run() error {

	params := new(phraseapp.TagParams)

	val, defaultsPresent := cmd.Config.Defaults["tag/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Name != nil {
		params.Name = cmd.Name
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func (cmd *TagShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TagsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Content      *string `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	KeyID        *string `cli:"opt --key-id"`
	LocaleID     *string `cli:"opt --locale-id"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationCreate) Run() error {

	params := new(phraseapp.TranslationParams)

	val, defaultsPresent := cmd.Config.Defaults["translation/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *TranslationShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Content      *string `cli:"opt --content"`
	Excluded     *bool   `cli:"opt --excluded"`
	PluralSuffix *string `cli:"opt --plural-suffix"`
	Unverified   *bool   `cli:"opt --unverified"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *TranslationUpdate) Run() error {

	params := new(phraseapp.TranslationUpdateParams)

	val, defaultsPresent := cmd.Config.Defaults["translation/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func (cmd *TranslationsByKey) Run() error {

	params := new(phraseapp.TranslationsByKeyParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/by_key"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	LocaleID  string `cli:"arg required"`
}

func (cmd *TranslationsByLocale) Run() error {

	params := new(phraseapp.TranslationsByLocaleParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/by_locale"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsExclude) Run() error {

	params := new(phraseapp.TranslationsExcludeParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/exclude"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsInclude) Run() error {

	params := new(phraseapp.TranslationsIncludeParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/include"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsList) Run() error {

	params := new(phraseapp.TranslationsListParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/list"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsSearch) Run() error {

	params := new(phraseapp.TranslationsSearchParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/search"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsUnverify) Run() error {

	params := new(phraseapp.TranslationsUnverifyParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/unverify"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Order *string `cli:"opt --order"`
	Q     *string `cli:"opt --query -q"`
	Sort  *string `cli:"opt --sort"`

	ProjectID string `cli:"arg required"`
}

func (cmd *TranslationsVerify) Run() error {

	params := new(phraseapp.TranslationsVerifyParams)

	val, defaultsPresent := cmd.Config.Defaults["translations/verify"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ConvertEmoji       *bool   `cli:"opt --convert-emoji"`
	File               *string `cli:"opt --file"`
	FileEncoding       *string `cli:"opt --file-encoding"`
	FileFormat         *string `cli:"opt --file-format"`
	LocaleID           *string `cli:"opt --locale-id"`
	SkipUnverification *bool   `cli:"opt --skip-unverification"`
	SkipUploadTags     *bool   `cli:"opt --skip-upload-tags"`
	Tags               *string `cli:"opt --tags"`
	UpdateTranslations *bool   `cli:"opt --update-translations"`

	ProjectID string `cli:"arg required"`
}

func (cmd *UploadCreate) Run() error {

	params := new(phraseapp.UploadParams)

	val, defaultsPresent := cmd.Config.Defaults["upload/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.ConvertEmoji != nil {
		params.ConvertEmoji = cmd.ConvertEmoji
	}

	if cmd.File != nil {
		params.File = cmd.File
	}

	if cmd.FileEncoding != nil {
		params.FileEncoding = cmd.FileEncoding
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

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *UploadShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *UploadsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
	ID            string `cli:"arg required"`
}

func (cmd *VersionShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
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
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
}

func (cmd *VersionsList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	res, err := client.VersionsList(cmd.ProjectID, cmd.TranslationID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type WebhookCreate struct {
	*phraseapp.Config

	Active      *bool   `cli:"opt --active"`
	CallbackUrl *string `cli:"opt --callback-url"`
	Description *string `cli:"opt --description"`
	Events      *string `cli:"opt --events"`

	ProjectID string `cli:"arg required"`
}

func (cmd *WebhookCreate) Run() error {

	params := new(phraseapp.WebhookParams)

	val, defaultsPresent := cmd.Config.Defaults["webhook/create"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Active != nil {
		params.Active = cmd.Active
	}

	if cmd.CallbackUrl != nil {
		params.CallbackUrl = cmd.CallbackUrl
	}

	if cmd.Description != nil {
		params.Description = cmd.Description
	}

	if cmd.Events != nil {
		params.Events = cmd.Events
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	res, err := client.WebhookCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type WebhookDelete struct {
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *WebhookDelete) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	err = client.WebhookDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type WebhookShow struct {
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *WebhookShow) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	res, err := client.WebhookShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type WebhookTest struct {
	*phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *WebhookTest) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	err = client.WebhookTest(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type WebhookUpdate struct {
	*phraseapp.Config

	Active      *bool   `cli:"opt --active"`
	CallbackUrl *string `cli:"opt --callback-url"`
	Description *string `cli:"opt --description"`
	Events      *string `cli:"opt --events"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func (cmd *WebhookUpdate) Run() error {

	params := new(phraseapp.WebhookParams)

	val, defaultsPresent := cmd.Config.Defaults["webhook/update"]
	if defaultsPresent {
		if err := params.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}

	if cmd.Active != nil {
		params.Active = cmd.Active
	}

	if cmd.CallbackUrl != nil {
		params.CallbackUrl = cmd.CallbackUrl
	}

	if cmd.Description != nil {
		params.Description = cmd.Description
	}

	if cmd.Events != nil {
		params.Events = cmd.Events
	}

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	res, err := client.WebhookUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type WebhooksList struct {
	*phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func (cmd *WebhooksList) Run() error {

	client, err := phraseapp.NewClient(cmd.Config.Credentials)
	if err != nil {
		return err
	}

	res, err := client.WebhooksList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}
