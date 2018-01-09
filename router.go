package main

import (
	"encoding/json"
	"os"

	"github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const (
	RevisionDocs      = ""
	RevisionGenerator = ""
)

func router(cfg *phraseapp.Config) (*cli.Router, error) {
	r := cli.NewRouter()

	r.Register("account/show", newAccountShow(cfg), "Get details on a single account.")

	r.Register("accounts/list", newAccountsList(cfg), "List all accounts the current user has access to.")

	if cmd, err := newAuthorizationCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("authorization/create", cmd, "Create a new authorization.")
	}

	r.Register("authorization/delete", newAuthorizationDelete(cfg), "Delete an existing authorization. API calls using that token will stop working.")

	r.Register("authorization/show", newAuthorizationShow(cfg), "Get details on a single authorization.")

	if cmd, err := newAuthorizationUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("authorization/update", cmd, "Update an existing authorization.")
	}

	r.Register("authorizations/list", newAuthorizationsList(cfg), "List all your authorizations.")

	if cmd, err := newBlacklistedKeyCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("blacklisted_key/create", cmd, "Create a new rule for blacklisting keys.")
	}

	r.Register("blacklisted_key/delete", newBlacklistedKeyDelete(cfg), "Delete an existing rule for blacklisting keys.")

	r.Register("blacklisted_key/show", newBlacklistedKeyShow(cfg), "Get details on a single rule for blacklisting keys for a given project.")

	if cmd, err := newBlacklistedKeyUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("blacklisted_key/update", cmd, "Update an existing rule for blacklisting keys.")
	}

	r.Register("blacklisted_keys/list", newBlacklistedKeysList(cfg), "List all rules for blacklisting keys for the given project.")

	r.Register("branch/compare", newBranchCompare(cfg), "Compare branch to current state of project")

	if cmd, err := newBranchCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("branch/create", cmd, "Create a new branch.")
	}

	r.Register("branch/merge", newBranchMerge(cfg), "Merge an existing branch.")

	if cmd, err := newBranchUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("branch/update", cmd, "Update an existing branch.")
	}

	r.Register("branches/list", newBranchesList(cfg), "List all branches the of the current project.")

	if cmd, err := newCommentCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/create", cmd, "Create a new comment for a key.")
	}

	r.Register("comment/delete", newCommentDelete(cfg), "Delete an existing comment.")

	r.Register("comment/mark/check", newCommentMarkCheck(cfg), "Check if comment was marked as read. Returns 204 if read, 404 if unread.")

	r.Register("comment/mark/read", newCommentMarkRead(cfg), "Mark a comment as read.")

	r.Register("comment/mark/unread", newCommentMarkUnread(cfg), "Mark a comment as unread.")

	r.Register("comment/show", newCommentShow(cfg), "Get details on a single comment.")

	if cmd, err := newCommentUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/update", cmd, "Update an existing comment.")
	}

	r.Register("comments/list", newCommentsList(cfg), "List all comments for a key.")

	r.Register("formats/list", newFormatsList(cfg), "Get a handy list of all localization file formats supported in PhraseApp.")

	r.Register("glossaries/list", newGlossariesList(cfg), "List all glossaries the current user has access to.")

	if cmd, err := newGlossaryCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary/create", cmd, "Create a new glossary.")
	}

	r.Register("glossary/delete", newGlossaryDelete(cfg), "Delete an existing glossary.")

	r.Register("glossary/show", newGlossaryShow(cfg), "Get details on a single glossary.")

	if cmd, err := newGlossaryUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary/update", cmd, "Update an existing glossary.")
	}

	if cmd, err := newGlossaryTermCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary_term/create", cmd, "Create a new glossary term.")
	}

	r.Register("glossary_term/delete", newGlossaryTermDelete(cfg), "Delete an existing glossary term.")

	r.Register("glossary_term/show", newGlossaryTermShow(cfg), "Get details on a single glossary term.")

	if cmd, err := newGlossaryTermUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary_term/update", cmd, "Update an existing glossary term.")
	}

	if cmd, err := newGlossaryTermTranslationCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary_term_translation/create", cmd, "Create a new glossary term translation.")
	}

	r.Register("glossary_term_translation/delete", newGlossaryTermTranslationDelete(cfg), "Delete an existing glossary term translation.")

	if cmd, err := newGlossaryTermTranslationUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("glossary_term_translation/update", cmd, "Update an existing glossary term translation.")
	}

	r.Register("glossary_terms/list", newGlossaryTermsList(cfg), "List all glossary terms the current user has access to.")

	if cmd, err := newInvitationCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("invitation/create", cmd, "Invite a person to an account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.")
	}

	r.Register("invitation/delete", newInvitationDelete(cfg), "Delete an existing invitation (must not be accepted yet). Access token scope must include <code>team.manage</code>.")

	r.Register("invitation/resend", newInvitationResend(cfg), "Resend the invitation email (must not be accepted yet). Access token scope must include <code>team.manage</code>.")

	r.Register("invitation/show", newInvitationShow(cfg), "Get details on a single invitation. Access token scope must include <code>team.manage</code>.")

	if cmd, err := newInvitationUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("invitation/update", cmd, "Update an existing invitation (must not be accepted yet). The <code>email</code> cannot be updated. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.")
	}

	r.Register("invitations/list", newInvitationsList(cfg), "List invitations for an account. It will also list the accessible resources like projects and locales the invited user has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.")

	r.Register("job/complete", newJobComplete(cfg), "Mark a job as completed.")

	if cmd, err := newJobCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/create", cmd, "Create a new job.")
	}

	r.Register("job/delete", newJobDelete(cfg), "Delete an existing job.")

	if cmd, err := newJobKeysCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/keys/create", cmd, "Add multiple keys to a existing job.")
	}

	if cmd, err := newJobKeysDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/keys/delete", cmd, "Remove multiple keys from existing job.")
	}

	r.Register("job/show", newJobShow(cfg), "Get details on a single job for a given project.")

	r.Register("job/start", newJobStart(cfg), "Starts an existing job in state draft.")

	if cmd, err := newJobUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/update", cmd, "Update an existing job.")
	}

	r.Register("job_locale/complete", newJobLocaleComplete(cfg), "Mark a JobLocale as completed.")

	r.Register("job_locale/delete", newJobLocaleDelete(cfg), "Delete an existing JobLocale.")

	r.Register("job_locale/show", newJobLocaleShow(cfg), "Get a single JobLocale for a given job.")

	if cmd, err := newJobLocaleUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/update", cmd, "Update an existing job.")
	}

	if cmd, err := newJobLocalesCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locales/create", cmd, "Create a new JobLocale.")
	}

	r.Register("job_locales/list", newJobLocalesList(cfg), "List all JobLocales for a given job.")

	if cmd, err := newJobsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("jobs/list", cmd, "List all jobs for the given project.")
	}

	if cmd, err := newKeyCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("key/create", cmd, "Create a new key.")
	}

	r.Register("key/delete", newKeyDelete(cfg), "Delete an existing key.")

	r.Register("key/show", newKeyShow(cfg), "Get details on a single key for a given project.")

	if cmd, err := newKeyUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("key/update", cmd, "Update an existing key.")
	}

	if cmd, err := newKeysDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("keys/delete", cmd, "Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.")
	}

	if cmd, err := newKeysList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("keys/list", cmd, "List all keys for the given project. Alternatively you can POST requests to /search.")
	}

	if cmd, err := newKeysSearch(cfg); err != nil {
		return nil, err
	} else {
		r.Register("keys/search", cmd, "Search keys for the given project matching query.")
	}

	if cmd, err := newKeysTag(cfg); err != nil {
		return nil, err
	} else {
		r.Register("keys/tag", cmd, "Tags all keys matching query. Same constraints as list.")
	}

	if cmd, err := newKeysUntag(cfg); err != nil {
		return nil, err
	} else {
		r.Register("keys/untag", cmd, "Removes specified tags from keys matching query.")
	}

	if cmd, err := newLocaleCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locale/create", cmd, "Create a new locale.")
	}

	r.Register("locale/delete", newLocaleDelete(cfg), "Delete an existing locale.")

	if cmd, err := newLocaleDownload(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locale/download", cmd, "Download a locale in a specific file format.")
	}

	if cmd, err := newLocaleShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locale/show", cmd, "Get details on a single locale for a given project.")
	}

	if cmd, err := newLocaleUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locale/update", cmd, "Update an existing locale.")
	}

	if cmd, err := newLocalesList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locales/list", cmd, "List all locales for the given project.")
	}

	r.Register("member/delete", newMemberDelete(cfg), "Remove a user from the account. The user will be removed from the account but not deleted from PhraseApp. Access token scope must include <code>team.manage</code>.")

	r.Register("member/show", newMemberShow(cfg), "Get details on a single user in the account. Access token scope must include <code>team.manage</code>.")

	if cmd, err := newMemberUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("member/update", cmd, "Update user permissions in the account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.")
	}

	r.Register("members/list", newMembersList(cfg), "Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.")

	r.Register("order/confirm", newOrderConfirm(cfg), "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")

	if cmd, err := newOrderCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("order/create", cmd, "Create a new order. Access token scope must include <code>orders.create</code>.")
	}

	r.Register("order/delete", newOrderDelete(cfg), "Cancel an existing order. Must not yet be confirmed.")

	r.Register("order/show", newOrderShow(cfg), "Get details on a single order.")

	r.Register("orders/list", newOrdersList(cfg), "List all orders for the given project.")

	if cmd, err := newProjectCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("project/create", cmd, "Create a new project.")
	}

	r.Register("project/delete", newProjectDelete(cfg), "Delete an existing project.")

	r.Register("project/show", newProjectShow(cfg), "Get details on a single project.")

	if cmd, err := newProjectUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("project/update", cmd, "Update an existing project.")
	}

	r.Register("projects/list", newProjectsList(cfg), "List all projects the current user has access to.")

	r.Register("show/user", newShowUser(cfg), "Show details for current User.")

	if cmd, err := newStyleguideCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("styleguide/create", cmd, "Create a new style guide.")
	}

	r.Register("styleguide/delete", newStyleguideDelete(cfg), "Delete an existing style guide.")

	r.Register("styleguide/show", newStyleguideShow(cfg), "Get details on a single style guide.")

	if cmd, err := newStyleguideUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("styleguide/update", cmd, "Update an existing style guide.")
	}

	r.Register("styleguides/list", newStyleguidesList(cfg), "List all styleguides for the given project.")

	if cmd, err := newTagCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("tag/create", cmd, "Create a new tag.")
	}

	r.Register("tag/delete", newTagDelete(cfg), "Delete an existing tag.")

	r.Register("tag/show", newTagShow(cfg), "Get details and progress information on a single tag for a given project.")

	r.Register("tags/list", newTagsList(cfg), "List all tags for the given project.")

	if cmd, err := newTranslationCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/create", cmd, "Create a translation.")
	}

	r.Register("translation/show", newTranslationShow(cfg), "Get details on a single translation.")

	if cmd, err := newTranslationUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/update", cmd, "Update an existing translation.")
	}

	if cmd, err := newTranslationsByKey(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/by_key", cmd, "List translations for a specific key.")
	}

	if cmd, err := newTranslationsByLocale(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/by_locale", cmd, "List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")
	}

	if cmd, err := newTranslationsExclude(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/exclude", cmd, "Exclude translations matching query from locale export.")
	}

	if cmd, err := newTranslationsInclude(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/include", cmd, "Include translations matching query in locale export.")
	}

	if cmd, err := newTranslationsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/list", cmd, "List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")
	}

	if cmd, err := newTranslationsSearch(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/search", cmd, "List translations for the given project if you exceed GET request limitations on translations list. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")
	}

	if cmd, err := newTranslationsUnverify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/unverify", cmd, "<div class='alert alert-info'>Only available in the <a href='https://phraseapp.com/pricing' target='_blank'>Control Package</a>.</div>Mark translations matching query as unverified.")
	}

	if cmd, err := newTranslationsVerify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/verify", cmd, "<div class='alert alert-info'>Only available in the <a href='https://phraseapp.com/pricing' target='_blank'>Control Package</a>.</div>Verify translations matching query.")
	}

	if cmd, err := newUploadCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("upload/create", cmd, "Upload a new language file. Creates necessary resources in your project.")
	}

	if cmd, err := newUploadShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("upload/show", cmd, "View details and summary for a single upload.")
	}

	r.Register("uploads/list", newUploadsList(cfg), "List all uploads for the given project.")

	r.Register("version/show", newVersionShow(cfg), "Get details on a single version.")

	r.Register("versions/list", newVersionsList(cfg), "List all versions for the given translation.")

	if cmd, err := newWebhookCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("webhook/create", cmd, "Create a new webhook.")
	}

	r.Register("webhook/delete", newWebhookDelete(cfg), "Delete an existing webhook.")

	r.Register("webhook/show", newWebhookShow(cfg), "Get details on a single webhook.")

	r.Register("webhook/test", newWebhookTest(cfg), "Perform a test request for a webhook.")

	if cmd, err := newWebhookUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("webhook/update", cmd, "Update an existing webhook.")
	}

	r.Register("webhooks/list", newWebhooksList(cfg), "List all webhooks for the given project.")

	ApplyNonRestRoutes(r, cfg)

	return r, nil
}

type AccountShow struct {
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newAccountShow(cfg *phraseapp.Config) *AccountShow {

	actionAccountShow := &AccountShow{Config: *cfg}

	return actionAccountShow
}

func (cmd *AccountShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.AccountShow(cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AccountsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func newAccountsList(cfg *phraseapp.Config) *AccountsList {

	actionAccountsList := &AccountsList{Config: *cfg}
	if cfg.Page != nil {
		actionAccountsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionAccountsList.PerPage = *cfg.PerPage
	}

	return actionAccountsList
}

func (cmd *AccountsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.AccountsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type AuthorizationCreate struct {
	phraseapp.Config

	phraseapp.AuthorizationParams
}

func newAuthorizationCreate(cfg *phraseapp.Config) (*AuthorizationCreate, error) {

	actionAuthorizationCreate := &AuthorizationCreate{Config: *cfg}

	val, defaultsPresent := actionAuthorizationCreate.Config.Defaults["authorization/create"]
	if defaultsPresent {
		if err := actionAuthorizationCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionAuthorizationCreate, nil
}

func (cmd *AuthorizationCreate) Run() error {
	params := &cmd.AuthorizationParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newAuthorizationDelete(cfg *phraseapp.Config) *AuthorizationDelete {

	actionAuthorizationDelete := &AuthorizationDelete{Config: *cfg}

	return actionAuthorizationDelete
}

func (cmd *AuthorizationDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newAuthorizationShow(cfg *phraseapp.Config) *AuthorizationShow {

	actionAuthorizationShow := &AuthorizationShow{Config: *cfg}

	return actionAuthorizationShow
}

func (cmd *AuthorizationShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.AuthorizationParams

	ID string `cli:"arg required"`
}

func newAuthorizationUpdate(cfg *phraseapp.Config) (*AuthorizationUpdate, error) {

	actionAuthorizationUpdate := &AuthorizationUpdate{Config: *cfg}

	val, defaultsPresent := actionAuthorizationUpdate.Config.Defaults["authorization/update"]
	if defaultsPresent {
		if err := actionAuthorizationUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionAuthorizationUpdate, nil
}

func (cmd *AuthorizationUpdate) Run() error {
	params := &cmd.AuthorizationParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func newAuthorizationsList(cfg *phraseapp.Config) *AuthorizationsList {

	actionAuthorizationsList := &AuthorizationsList{Config: *cfg}
	if cfg.Page != nil {
		actionAuthorizationsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionAuthorizationsList.PerPage = *cfg.PerPage
	}

	return actionAuthorizationsList
}

func (cmd *AuthorizationsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.BlacklistedKeyParams

	ProjectID string `cli:"arg required"`
}

func newBlacklistedKeyCreate(cfg *phraseapp.Config) (*BlacklistedKeyCreate, error) {

	actionBlacklistedKeyCreate := &BlacklistedKeyCreate{Config: *cfg}
	actionBlacklistedKeyCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBlacklistedKeyCreate.Config.Defaults["blacklisted_key/create"]
	if defaultsPresent {
		if err := actionBlacklistedKeyCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBlacklistedKeyCreate, nil
}

func (cmd *BlacklistedKeyCreate) Run() error {
	params := &cmd.BlacklistedKeyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBlacklistedKeyDelete(cfg *phraseapp.Config) *BlacklistedKeyDelete {

	actionBlacklistedKeyDelete := &BlacklistedKeyDelete{Config: *cfg}
	actionBlacklistedKeyDelete.ProjectID = cfg.DefaultProjectID

	return actionBlacklistedKeyDelete
}

func (cmd *BlacklistedKeyDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBlacklistedKeyShow(cfg *phraseapp.Config) *BlacklistedKeyShow {

	actionBlacklistedKeyShow := &BlacklistedKeyShow{Config: *cfg}
	actionBlacklistedKeyShow.ProjectID = cfg.DefaultProjectID

	return actionBlacklistedKeyShow
}

func (cmd *BlacklistedKeyShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.BlacklistedKeyParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBlacklistedKeyUpdate(cfg *phraseapp.Config) (*BlacklistedKeyUpdate, error) {

	actionBlacklistedKeyUpdate := &BlacklistedKeyUpdate{Config: *cfg}
	actionBlacklistedKeyUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBlacklistedKeyUpdate.Config.Defaults["blacklisted_key/update"]
	if defaultsPresent {
		if err := actionBlacklistedKeyUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBlacklistedKeyUpdate, nil
}

func (cmd *BlacklistedKeyUpdate) Run() error {
	params := &cmd.BlacklistedKeyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newBlacklistedKeysList(cfg *phraseapp.Config) *BlacklistedKeysList {

	actionBlacklistedKeysList := &BlacklistedKeysList{Config: *cfg}
	actionBlacklistedKeysList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionBlacklistedKeysList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionBlacklistedKeysList.PerPage = *cfg.PerPage
	}

	return actionBlacklistedKeysList
}

func (cmd *BlacklistedKeysList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BlacklistedKeysList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BranchCompare struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchCompare(cfg *phraseapp.Config) *BranchCompare {

	actionBranchCompare := &BranchCompare{Config: *cfg}
	actionBranchCompare.ProjectID = cfg.DefaultProjectID

	return actionBranchCompare
}

func (cmd *BranchCompare) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchCompare(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BranchCreate struct {
	phraseapp.Config

	phraseapp.BranchParams

	ProjectID string `cli:"arg required"`
}

func newBranchCreate(cfg *phraseapp.Config) (*BranchCreate, error) {

	actionBranchCreate := &BranchCreate{Config: *cfg}
	actionBranchCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchCreate.Config.Defaults["branch/create"]
	if defaultsPresent {
		if err := actionBranchCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBranchCreate, nil
}

func (cmd *BranchCreate) Run() error {
	params := &cmd.BranchParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BranchMerge struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchMerge(cfg *phraseapp.Config) *BranchMerge {

	actionBranchMerge := &BranchMerge{Config: *cfg}
	actionBranchMerge.ProjectID = cfg.DefaultProjectID

	return actionBranchMerge
}

func (cmd *BranchMerge) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchMerge(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return nil
}

type BranchUpdate struct {
	phraseapp.Config

	phraseapp.BranchParams

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchUpdate(cfg *phraseapp.Config) (*BranchUpdate, error) {

	actionBranchUpdate := &BranchUpdate{Config: *cfg}
	actionBranchUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchUpdate.Config.Defaults["branch/update"]
	if defaultsPresent {
		if err := actionBranchUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBranchUpdate, nil
}

func (cmd *BranchUpdate) Run() error {
	params := &cmd.BranchParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchUpdate(cmd.ProjectID, cmd.Name, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BranchesList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newBranchesList(cfg *phraseapp.Config) *BranchesList {

	actionBranchesList := &BranchesList{Config: *cfg}
	actionBranchesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionBranchesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionBranchesList.PerPage = *cfg.PerPage
	}

	return actionBranchesList
}

func (cmd *BranchesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchesList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type CommentCreate struct {
	phraseapp.Config

	phraseapp.CommentParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func newCommentCreate(cfg *phraseapp.Config) (*CommentCreate, error) {

	actionCommentCreate := &CommentCreate{Config: *cfg}
	actionCommentCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentCreate.Config.Defaults["comment/create"]
	if defaultsPresent {
		if err := actionCommentCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentCreate, nil
}

func (cmd *CommentCreate) Run() error {
	params := &cmd.CommentParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentDelete(cfg *phraseapp.Config) *CommentDelete {

	actionCommentDelete := &CommentDelete{Config: *cfg}
	actionCommentDelete.ProjectID = cfg.DefaultProjectID

	return actionCommentDelete
}

func (cmd *CommentDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkCheck(cfg *phraseapp.Config) *CommentMarkCheck {

	actionCommentMarkCheck := &CommentMarkCheck{Config: *cfg}
	actionCommentMarkCheck.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkCheck
}

func (cmd *CommentMarkCheck) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkRead(cfg *phraseapp.Config) *CommentMarkRead {

	actionCommentMarkRead := &CommentMarkRead{Config: *cfg}
	actionCommentMarkRead.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkRead
}

func (cmd *CommentMarkRead) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkUnread(cfg *phraseapp.Config) *CommentMarkUnread {

	actionCommentMarkUnread := &CommentMarkUnread{Config: *cfg}
	actionCommentMarkUnread.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkUnread
}

func (cmd *CommentMarkUnread) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentShow(cfg *phraseapp.Config) *CommentShow {

	actionCommentShow := &CommentShow{Config: *cfg}
	actionCommentShow.ProjectID = cfg.DefaultProjectID

	return actionCommentShow
}

func (cmd *CommentShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.CommentParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentUpdate(cfg *phraseapp.Config) (*CommentUpdate, error) {

	actionCommentUpdate := &CommentUpdate{Config: *cfg}
	actionCommentUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentUpdate.Config.Defaults["comment/update"]
	if defaultsPresent {
		if err := actionCommentUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentUpdate, nil
}

func (cmd *CommentUpdate) Run() error {
	params := &cmd.CommentParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func newCommentsList(cfg *phraseapp.Config) *CommentsList {

	actionCommentsList := &CommentsList{Config: *cfg}
	actionCommentsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionCommentsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionCommentsList.PerPage = *cfg.PerPage
	}

	return actionCommentsList
}

func (cmd *CommentsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func newFormatsList(cfg *phraseapp.Config) *FormatsList {

	actionFormatsList := &FormatsList{Config: *cfg}
	if cfg.Page != nil {
		actionFormatsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionFormatsList.PerPage = *cfg.PerPage
	}

	return actionFormatsList
}

func (cmd *FormatsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.FormatsList(cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossariesList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
}

func newGlossariesList(cfg *phraseapp.Config) *GlossariesList {

	actionGlossariesList := &GlossariesList{Config: *cfg}
	if cfg.Page != nil {
		actionGlossariesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionGlossariesList.PerPage = *cfg.PerPage
	}

	return actionGlossariesList
}

func (cmd *GlossariesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossariesList(cmd.AccountID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryCreate struct {
	phraseapp.Config

	phraseapp.GlossaryParams

	AccountID string `cli:"arg required"`
}

func newGlossaryCreate(cfg *phraseapp.Config) (*GlossaryCreate, error) {

	actionGlossaryCreate := &GlossaryCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryCreate.Config.Defaults["glossary/create"]
	if defaultsPresent {
		if err := actionGlossaryCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryCreate, nil
}

func (cmd *GlossaryCreate) Run() error {
	params := &cmd.GlossaryParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryCreate(cmd.AccountID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newGlossaryDelete(cfg *phraseapp.Config) *GlossaryDelete {

	actionGlossaryDelete := &GlossaryDelete{Config: *cfg}

	return actionGlossaryDelete
}

func (cmd *GlossaryDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.GlossaryDelete(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type GlossaryShow struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newGlossaryShow(cfg *phraseapp.Config) *GlossaryShow {

	actionGlossaryShow := &GlossaryShow{Config: *cfg}

	return actionGlossaryShow
}

func (cmd *GlossaryShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryShow(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryUpdate struct {
	phraseapp.Config

	phraseapp.GlossaryParams

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newGlossaryUpdate(cfg *phraseapp.Config) (*GlossaryUpdate, error) {

	actionGlossaryUpdate := &GlossaryUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryUpdate.Config.Defaults["glossary/update"]
	if defaultsPresent {
		if err := actionGlossaryUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryUpdate, nil
}

func (cmd *GlossaryUpdate) Run() error {
	params := &cmd.GlossaryParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryUpdate(cmd.AccountID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermCreate struct {
	phraseapp.Config

	phraseapp.GlossaryTermParams

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
}

func newGlossaryTermCreate(cfg *phraseapp.Config) (*GlossaryTermCreate, error) {

	actionGlossaryTermCreate := &GlossaryTermCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermCreate.Config.Defaults["glossary_term/create"]
	if defaultsPresent {
		if err := actionGlossaryTermCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryTermCreate, nil
}

func (cmd *GlossaryTermCreate) Run() error {
	params := &cmd.GlossaryTermParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermCreate(cmd.AccountID, cmd.GlossaryID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermDelete struct {
	phraseapp.Config

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	ID         string `cli:"arg required"`
}

func newGlossaryTermDelete(cfg *phraseapp.Config) *GlossaryTermDelete {

	actionGlossaryTermDelete := &GlossaryTermDelete{Config: *cfg}

	return actionGlossaryTermDelete
}

func (cmd *GlossaryTermDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.GlossaryTermDelete(cmd.AccountID, cmd.GlossaryID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type GlossaryTermShow struct {
	phraseapp.Config

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	ID         string `cli:"arg required"`
}

func newGlossaryTermShow(cfg *phraseapp.Config) *GlossaryTermShow {

	actionGlossaryTermShow := &GlossaryTermShow{Config: *cfg}

	return actionGlossaryTermShow
}

func (cmd *GlossaryTermShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermShow(cmd.AccountID, cmd.GlossaryID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermUpdate struct {
	phraseapp.Config

	phraseapp.GlossaryTermParams

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	ID         string `cli:"arg required"`
}

func newGlossaryTermUpdate(cfg *phraseapp.Config) (*GlossaryTermUpdate, error) {

	actionGlossaryTermUpdate := &GlossaryTermUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermUpdate.Config.Defaults["glossary_term/update"]
	if defaultsPresent {
		if err := actionGlossaryTermUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryTermUpdate, nil
}

func (cmd *GlossaryTermUpdate) Run() error {
	params := &cmd.GlossaryTermParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermUpdate(cmd.AccountID, cmd.GlossaryID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermTranslationCreate struct {
	phraseapp.Config

	phraseapp.GlossaryTermTranslationParams

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	TermID     string `cli:"arg required"`
}

func newGlossaryTermTranslationCreate(cfg *phraseapp.Config) (*GlossaryTermTranslationCreate, error) {

	actionGlossaryTermTranslationCreate := &GlossaryTermTranslationCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermTranslationCreate.Config.Defaults["glossary_term_translation/create"]
	if defaultsPresent {
		if err := actionGlossaryTermTranslationCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryTermTranslationCreate, nil
}

func (cmd *GlossaryTermTranslationCreate) Run() error {
	params := &cmd.GlossaryTermTranslationParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermTranslationCreate(cmd.AccountID, cmd.GlossaryID, cmd.TermID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermTranslationDelete struct {
	phraseapp.Config

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	TermID     string `cli:"arg required"`
	ID         string `cli:"arg required"`
}

func newGlossaryTermTranslationDelete(cfg *phraseapp.Config) *GlossaryTermTranslationDelete {

	actionGlossaryTermTranslationDelete := &GlossaryTermTranslationDelete{Config: *cfg}

	return actionGlossaryTermTranslationDelete
}

func (cmd *GlossaryTermTranslationDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.GlossaryTermTranslationDelete(cmd.AccountID, cmd.GlossaryID, cmd.TermID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type GlossaryTermTranslationUpdate struct {
	phraseapp.Config

	phraseapp.GlossaryTermTranslationParams

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
	TermID     string `cli:"arg required"`
	ID         string `cli:"arg required"`
}

func newGlossaryTermTranslationUpdate(cfg *phraseapp.Config) (*GlossaryTermTranslationUpdate, error) {

	actionGlossaryTermTranslationUpdate := &GlossaryTermTranslationUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermTranslationUpdate.Config.Defaults["glossary_term_translation/update"]
	if defaultsPresent {
		if err := actionGlossaryTermTranslationUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionGlossaryTermTranslationUpdate, nil
}

func (cmd *GlossaryTermTranslationUpdate) Run() error {
	params := &cmd.GlossaryTermTranslationParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermTranslationUpdate(cmd.AccountID, cmd.GlossaryID, cmd.TermID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type GlossaryTermsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID  string `cli:"arg required"`
	GlossaryID string `cli:"arg required"`
}

func newGlossaryTermsList(cfg *phraseapp.Config) *GlossaryTermsList {

	actionGlossaryTermsList := &GlossaryTermsList{Config: *cfg}
	if cfg.Page != nil {
		actionGlossaryTermsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionGlossaryTermsList.PerPage = *cfg.PerPage
	}

	return actionGlossaryTermsList
}

func (cmd *GlossaryTermsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.GlossaryTermsList(cmd.AccountID, cmd.GlossaryID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type InvitationCreate struct {
	phraseapp.Config

	phraseapp.InvitationCreateParams

	AccountID string `cli:"arg required"`
}

func newInvitationCreate(cfg *phraseapp.Config) (*InvitationCreate, error) {

	actionInvitationCreate := &InvitationCreate{Config: *cfg}

	val, defaultsPresent := actionInvitationCreate.Config.Defaults["invitation/create"]
	if defaultsPresent {
		if err := actionInvitationCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionInvitationCreate, nil
}

func (cmd *InvitationCreate) Run() error {
	params := &cmd.InvitationCreateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.InvitationCreate(cmd.AccountID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type InvitationDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newInvitationDelete(cfg *phraseapp.Config) *InvitationDelete {

	actionInvitationDelete := &InvitationDelete{Config: *cfg}

	return actionInvitationDelete
}

func (cmd *InvitationDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.InvitationDelete(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type InvitationResend struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newInvitationResend(cfg *phraseapp.Config) *InvitationResend {

	actionInvitationResend := &InvitationResend{Config: *cfg}

	return actionInvitationResend
}

func (cmd *InvitationResend) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.InvitationResend(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type InvitationShow struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newInvitationShow(cfg *phraseapp.Config) *InvitationShow {

	actionInvitationShow := &InvitationShow{Config: *cfg}

	return actionInvitationShow
}

func (cmd *InvitationShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.InvitationShow(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type InvitationUpdate struct {
	phraseapp.Config

	phraseapp.InvitationUpdateParams

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newInvitationUpdate(cfg *phraseapp.Config) (*InvitationUpdate, error) {

	actionInvitationUpdate := &InvitationUpdate{Config: *cfg}

	val, defaultsPresent := actionInvitationUpdate.Config.Defaults["invitation/update"]
	if defaultsPresent {
		if err := actionInvitationUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionInvitationUpdate, nil
}

func (cmd *InvitationUpdate) Run() error {
	params := &cmd.InvitationUpdateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.InvitationUpdate(cmd.AccountID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type InvitationsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
}

func newInvitationsList(cfg *phraseapp.Config) *InvitationsList {

	actionInvitationsList := &InvitationsList{Config: *cfg}
	if cfg.Page != nil {
		actionInvitationsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionInvitationsList.PerPage = *cfg.PerPage
	}

	return actionInvitationsList
}

func (cmd *InvitationsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.InvitationsList(cmd.AccountID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobComplete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobComplete(cfg *phraseapp.Config) *JobComplete {

	actionJobComplete := &JobComplete{Config: *cfg}
	actionJobComplete.ProjectID = cfg.DefaultProjectID

	return actionJobComplete
}

func (cmd *JobComplete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobComplete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobCreate struct {
	phraseapp.Config

	phraseapp.JobParams

	ProjectID string `cli:"arg required"`
}

func newJobCreate(cfg *phraseapp.Config) (*JobCreate, error) {

	actionJobCreate := &JobCreate{Config: *cfg}
	actionJobCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobCreate.Config.Defaults["job/create"]
	if defaultsPresent {
		if err := actionJobCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobCreate, nil
}

func (cmd *JobCreate) Run() error {
	params := &cmd.JobParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobDelete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobDelete(cfg *phraseapp.Config) *JobDelete {

	actionJobDelete := &JobDelete{Config: *cfg}
	actionJobDelete.ProjectID = cfg.DefaultProjectID

	return actionJobDelete
}

func (cmd *JobDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.JobDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type JobKeysCreate struct {
	phraseapp.Config

	phraseapp.JobKeysCreateParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobKeysCreate(cfg *phraseapp.Config) (*JobKeysCreate, error) {

	actionJobKeysCreate := &JobKeysCreate{Config: *cfg}
	actionJobKeysCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobKeysCreate.Config.Defaults["job/keys/create"]
	if defaultsPresent {
		if err := actionJobKeysCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobKeysCreate, nil
}

func (cmd *JobKeysCreate) Run() error {
	params := &cmd.JobKeysCreateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobKeysCreate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobKeysDelete struct {
	phraseapp.Config

	phraseapp.JobKeysDeleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobKeysDelete(cfg *phraseapp.Config) (*JobKeysDelete, error) {

	actionJobKeysDelete := &JobKeysDelete{Config: *cfg}
	actionJobKeysDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobKeysDelete.Config.Defaults["job/keys/delete"]
	if defaultsPresent {
		if err := actionJobKeysDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobKeysDelete, nil
}

func (cmd *JobKeysDelete) Run() error {
	params := &cmd.JobKeysDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.JobKeysDelete(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type JobShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobShow(cfg *phraseapp.Config) *JobShow {

	actionJobShow := &JobShow{Config: *cfg}
	actionJobShow.ProjectID = cfg.DefaultProjectID

	return actionJobShow
}

func (cmd *JobShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobStart struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobStart(cfg *phraseapp.Config) *JobStart {

	actionJobStart := &JobStart{Config: *cfg}
	actionJobStart.ProjectID = cfg.DefaultProjectID

	return actionJobStart
}

func (cmd *JobStart) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobStart(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobUpdate struct {
	phraseapp.Config

	phraseapp.JobUpdateParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobUpdate(cfg *phraseapp.Config) (*JobUpdate, error) {

	actionJobUpdate := &JobUpdate{Config: *cfg}
	actionJobUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobUpdate.Config.Defaults["job/update"]
	if defaultsPresent {
		if err := actionJobUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobUpdate, nil
}

func (cmd *JobUpdate) Run() error {
	params := &cmd.JobUpdateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleComplete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleComplete(cfg *phraseapp.Config) *JobLocaleComplete {

	actionJobLocaleComplete := &JobLocaleComplete{Config: *cfg}
	actionJobLocaleComplete.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleComplete
}

func (cmd *JobLocaleComplete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleComplete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleDelete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleDelete(cfg *phraseapp.Config) *JobLocaleDelete {

	actionJobLocaleDelete := &JobLocaleDelete{Config: *cfg}
	actionJobLocaleDelete.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleDelete
}

func (cmd *JobLocaleDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.JobLocaleDelete(cmd.ProjectID, cmd.JobID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type JobLocaleShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleShow(cfg *phraseapp.Config) *JobLocaleShow {

	actionJobLocaleShow := &JobLocaleShow{Config: *cfg}
	actionJobLocaleShow.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleShow
}

func (cmd *JobLocaleShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleShow(cmd.ProjectID, cmd.JobID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleUpdate struct {
	phraseapp.Config

	phraseapp.JobLocaleParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleUpdate(cfg *phraseapp.Config) (*JobLocaleUpdate, error) {

	actionJobLocaleUpdate := &JobLocaleUpdate{Config: *cfg}
	actionJobLocaleUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleUpdate.Config.Defaults["job_locale/update"]
	if defaultsPresent {
		if err := actionJobLocaleUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocaleUpdate, nil
}

func (cmd *JobLocaleUpdate) Run() error {
	params := &cmd.JobLocaleParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleUpdate(cmd.ProjectID, cmd.JobID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocalesCreate struct {
	phraseapp.Config

	phraseapp.JobLocaleParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
}

func newJobLocalesCreate(cfg *phraseapp.Config) (*JobLocalesCreate, error) {

	actionJobLocalesCreate := &JobLocalesCreate{Config: *cfg}
	actionJobLocalesCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocalesCreate.Config.Defaults["job_locales/create"]
	if defaultsPresent {
		if err := actionJobLocalesCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocalesCreate, nil
}

func (cmd *JobLocalesCreate) Run() error {
	params := &cmd.JobLocaleParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocalesCreate(cmd.ProjectID, cmd.JobID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocalesList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
}

func newJobLocalesList(cfg *phraseapp.Config) *JobLocalesList {

	actionJobLocalesList := &JobLocalesList{Config: *cfg}
	actionJobLocalesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionJobLocalesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionJobLocalesList.PerPage = *cfg.PerPage
	}

	return actionJobLocalesList
}

func (cmd *JobLocalesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocalesList(cmd.ProjectID, cmd.JobID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobsList struct {
	phraseapp.Config

	phraseapp.JobsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newJobsList(cfg *phraseapp.Config) (*JobsList, error) {

	actionJobsList := &JobsList{Config: *cfg}
	actionJobsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionJobsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionJobsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionJobsList.Config.Defaults["jobs/list"]
	if defaultsPresent {
		if err := actionJobsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobsList, nil
}

func (cmd *JobsList) Run() error {
	params := &cmd.JobsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobsList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type KeyCreate struct {
	phraseapp.Config

	phraseapp.TranslationKeyParams

	ProjectID string `cli:"arg required"`
}

func newKeyCreate(cfg *phraseapp.Config) (*KeyCreate, error) {

	actionKeyCreate := &KeyCreate{Config: *cfg}
	actionKeyCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyCreate.Config.Defaults["key/create"]
	if defaultsPresent {
		if err := actionKeyCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeyCreate, nil
}

func (cmd *KeyCreate) Run() error {
	params := &cmd.TranslationKeyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newKeyDelete(cfg *phraseapp.Config) *KeyDelete {

	actionKeyDelete := &KeyDelete{Config: *cfg}
	actionKeyDelete.ProjectID = cfg.DefaultProjectID

	return actionKeyDelete
}

func (cmd *KeyDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newKeyShow(cfg *phraseapp.Config) *KeyShow {

	actionKeyShow := &KeyShow{Config: *cfg}
	actionKeyShow.ProjectID = cfg.DefaultProjectID

	return actionKeyShow
}

func (cmd *KeyShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationKeyParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newKeyUpdate(cfg *phraseapp.Config) (*KeyUpdate, error) {

	actionKeyUpdate := &KeyUpdate{Config: *cfg}
	actionKeyUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyUpdate.Config.Defaults["key/update"]
	if defaultsPresent {
		if err := actionKeyUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeyUpdate, nil
}

func (cmd *KeyUpdate) Run() error {
	params := &cmd.TranslationKeyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.KeysDeleteParams

	ProjectID string `cli:"arg required"`
}

func newKeysDelete(cfg *phraseapp.Config) (*KeysDelete, error) {

	actionKeysDelete := &KeysDelete{Config: *cfg}
	actionKeysDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysDelete.Config.Defaults["keys/delete"]
	if defaultsPresent {
		if err := actionKeysDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeysDelete, nil
}

func (cmd *KeysDelete) Run() error {
	params := &cmd.KeysDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.KeysListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newKeysList(cfg *phraseapp.Config) (*KeysList, error) {

	actionKeysList := &KeysList{Config: *cfg}
	actionKeysList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionKeysList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionKeysList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionKeysList.Config.Defaults["keys/list"]
	if defaultsPresent {
		if err := actionKeysList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeysList, nil
}

func (cmd *KeysList) Run() error {
	params := &cmd.KeysListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.KeysSearchParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newKeysSearch(cfg *phraseapp.Config) (*KeysSearch, error) {

	actionKeysSearch := &KeysSearch{Config: *cfg}
	actionKeysSearch.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionKeysSearch.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionKeysSearch.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionKeysSearch.Config.Defaults["keys/search"]
	if defaultsPresent {
		if err := actionKeysSearch.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeysSearch, nil
}

func (cmd *KeysSearch) Run() error {
	params := &cmd.KeysSearchParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.KeysTagParams

	ProjectID string `cli:"arg required"`
}

func newKeysTag(cfg *phraseapp.Config) (*KeysTag, error) {

	actionKeysTag := &KeysTag{Config: *cfg}
	actionKeysTag.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysTag.Config.Defaults["keys/tag"]
	if defaultsPresent {
		if err := actionKeysTag.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeysTag, nil
}

func (cmd *KeysTag) Run() error {
	params := &cmd.KeysTagParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.KeysUntagParams

	ProjectID string `cli:"arg required"`
}

func newKeysUntag(cfg *phraseapp.Config) (*KeysUntag, error) {

	actionKeysUntag := &KeysUntag{Config: *cfg}
	actionKeysUntag.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysUntag.Config.Defaults["keys/untag"]
	if defaultsPresent {
		if err := actionKeysUntag.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeysUntag, nil
}

func (cmd *KeysUntag) Run() error {
	params := &cmd.KeysUntagParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.LocaleParams

	ProjectID string `cli:"arg required"`
}

func newLocaleCreate(cfg *phraseapp.Config) (*LocaleCreate, error) {

	actionLocaleCreate := &LocaleCreate{Config: *cfg}
	actionLocaleCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleCreate.Config.Defaults["locale/create"]
	if defaultsPresent {
		if err := actionLocaleCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocaleCreate, nil
}

func (cmd *LocaleCreate) Run() error {
	params := &cmd.LocaleParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newLocaleDelete(cfg *phraseapp.Config) *LocaleDelete {

	actionLocaleDelete := &LocaleDelete{Config: *cfg}
	actionLocaleDelete.ProjectID = cfg.DefaultProjectID

	return actionLocaleDelete
}

func (cmd *LocaleDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.LocaleDownloadParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newLocaleDownload(cfg *phraseapp.Config) (*LocaleDownload, error) {

	actionLocaleDownload := &LocaleDownload{Config: *cfg}
	actionLocaleDownload.ProjectID = cfg.DefaultProjectID
	if cfg.DefaultFileFormat != "" {
		actionLocaleDownload.FileFormat = &cfg.DefaultFileFormat
	}

	val, defaultsPresent := actionLocaleDownload.Config.Defaults["locale/download"]
	if defaultsPresent {
		if err := actionLocaleDownload.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocaleDownload, nil
}

func (cmd *LocaleDownload) Run() error {
	params := &cmd.LocaleDownloadParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.LocaleDownload(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	os.Stdout.Write(res)
	return nil
}

type LocaleShow struct {
	phraseapp.Config

	phraseapp.LocaleShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newLocaleShow(cfg *phraseapp.Config) (*LocaleShow, error) {

	actionLocaleShow := &LocaleShow{Config: *cfg}
	actionLocaleShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleShow.Config.Defaults["locale/show"]
	if defaultsPresent {
		if err := actionLocaleShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocaleShow, nil
}

func (cmd *LocaleShow) Run() error {
	params := &cmd.LocaleShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.LocaleShow(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type LocaleUpdate struct {
	phraseapp.Config

	phraseapp.LocaleParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newLocaleUpdate(cfg *phraseapp.Config) (*LocaleUpdate, error) {

	actionLocaleUpdate := &LocaleUpdate{Config: *cfg}
	actionLocaleUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleUpdate.Config.Defaults["locale/update"]
	if defaultsPresent {
		if err := actionLocaleUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocaleUpdate, nil
}

func (cmd *LocaleUpdate) Run() error {
	params := &cmd.LocaleParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.LocalesListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newLocalesList(cfg *phraseapp.Config) (*LocalesList, error) {

	actionLocalesList := &LocalesList{Config: *cfg}
	actionLocalesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionLocalesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionLocalesList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionLocalesList.Config.Defaults["locales/list"]
	if defaultsPresent {
		if err := actionLocalesList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocalesList, nil
}

func (cmd *LocalesList) Run() error {
	params := &cmd.LocalesListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.LocalesList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type MemberDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newMemberDelete(cfg *phraseapp.Config) *MemberDelete {

	actionMemberDelete := &MemberDelete{Config: *cfg}

	return actionMemberDelete
}

func (cmd *MemberDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.MemberDelete(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type MemberShow struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newMemberShow(cfg *phraseapp.Config) *MemberShow {

	actionMemberShow := &MemberShow{Config: *cfg}

	return actionMemberShow
}

func (cmd *MemberShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.MemberShow(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type MemberUpdate struct {
	phraseapp.Config

	phraseapp.MemberUpdateParams

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newMemberUpdate(cfg *phraseapp.Config) (*MemberUpdate, error) {

	actionMemberUpdate := &MemberUpdate{Config: *cfg}

	val, defaultsPresent := actionMemberUpdate.Config.Defaults["member/update"]
	if defaultsPresent {
		if err := actionMemberUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionMemberUpdate, nil
}

func (cmd *MemberUpdate) Run() error {
	params := &cmd.MemberUpdateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.MemberUpdate(cmd.AccountID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type MembersList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
}

func newMembersList(cfg *phraseapp.Config) *MembersList {

	actionMembersList := &MembersList{Config: *cfg}
	if cfg.Page != nil {
		actionMembersList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionMembersList.PerPage = *cfg.PerPage
	}

	return actionMembersList
}

func (cmd *MembersList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.MembersList(cmd.AccountID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrderConfirm struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderConfirm(cfg *phraseapp.Config) *OrderConfirm {

	actionOrderConfirm := &OrderConfirm{Config: *cfg}
	actionOrderConfirm.ProjectID = cfg.DefaultProjectID

	return actionOrderConfirm
}

func (cmd *OrderConfirm) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationOrderParams

	ProjectID string `cli:"arg required"`
}

func newOrderCreate(cfg *phraseapp.Config) (*OrderCreate, error) {

	actionOrderCreate := &OrderCreate{Config: *cfg}
	actionOrderCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionOrderCreate.Config.Defaults["order/create"]
	if defaultsPresent {
		if err := actionOrderCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionOrderCreate, nil
}

func (cmd *OrderCreate) Run() error {
	params := &cmd.TranslationOrderParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderDelete(cfg *phraseapp.Config) *OrderDelete {

	actionOrderDelete := &OrderDelete{Config: *cfg}
	actionOrderDelete.ProjectID = cfg.DefaultProjectID

	return actionOrderDelete
}

func (cmd *OrderDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderShow(cfg *phraseapp.Config) *OrderShow {

	actionOrderShow := &OrderShow{Config: *cfg}
	actionOrderShow.ProjectID = cfg.DefaultProjectID

	return actionOrderShow
}

func (cmd *OrderShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newOrdersList(cfg *phraseapp.Config) *OrdersList {

	actionOrdersList := &OrdersList{Config: *cfg}
	actionOrdersList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionOrdersList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionOrdersList.PerPage = *cfg.PerPage
	}

	return actionOrdersList
}

func (cmd *OrdersList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.ProjectParams
}

func newProjectCreate(cfg *phraseapp.Config) (*ProjectCreate, error) {

	actionProjectCreate := &ProjectCreate{Config: *cfg}

	val, defaultsPresent := actionProjectCreate.Config.Defaults["project/create"]
	if defaultsPresent {
		if err := actionProjectCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionProjectCreate, nil
}

func (cmd *ProjectCreate) Run() error {
	params := &cmd.ProjectParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newProjectDelete(cfg *phraseapp.Config) *ProjectDelete {

	actionProjectDelete := &ProjectDelete{Config: *cfg}

	return actionProjectDelete
}

func (cmd *ProjectDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newProjectShow(cfg *phraseapp.Config) *ProjectShow {

	actionProjectShow := &ProjectShow{Config: *cfg}

	return actionProjectShow
}

func (cmd *ProjectShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.ProjectParams

	ID string `cli:"arg required"`
}

func newProjectUpdate(cfg *phraseapp.Config) (*ProjectUpdate, error) {

	actionProjectUpdate := &ProjectUpdate{Config: *cfg}

	val, defaultsPresent := actionProjectUpdate.Config.Defaults["project/update"]
	if defaultsPresent {
		if err := actionProjectUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionProjectUpdate, nil
}

func (cmd *ProjectUpdate) Run() error {
	params := &cmd.ProjectParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func newProjectsList(cfg *phraseapp.Config) *ProjectsList {

	actionProjectsList := &ProjectsList{Config: *cfg}
	if cfg.Page != nil {
		actionProjectsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionProjectsList.PerPage = *cfg.PerPage
	}

	return actionProjectsList
}

func (cmd *ProjectsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config
}

func newShowUser(cfg *phraseapp.Config) *ShowUser {

	actionShowUser := &ShowUser{Config: *cfg}

	return actionShowUser
}

func (cmd *ShowUser) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.StyleguideParams

	ProjectID string `cli:"arg required"`
}

func newStyleguideCreate(cfg *phraseapp.Config) (*StyleguideCreate, error) {

	actionStyleguideCreate := &StyleguideCreate{Config: *cfg}
	actionStyleguideCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionStyleguideCreate.Config.Defaults["styleguide/create"]
	if defaultsPresent {
		if err := actionStyleguideCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionStyleguideCreate, nil
}

func (cmd *StyleguideCreate) Run() error {
	params := &cmd.StyleguideParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newStyleguideDelete(cfg *phraseapp.Config) *StyleguideDelete {

	actionStyleguideDelete := &StyleguideDelete{Config: *cfg}
	actionStyleguideDelete.ProjectID = cfg.DefaultProjectID

	return actionStyleguideDelete
}

func (cmd *StyleguideDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newStyleguideShow(cfg *phraseapp.Config) *StyleguideShow {

	actionStyleguideShow := &StyleguideShow{Config: *cfg}
	actionStyleguideShow.ProjectID = cfg.DefaultProjectID

	return actionStyleguideShow
}

func (cmd *StyleguideShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.StyleguideParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newStyleguideUpdate(cfg *phraseapp.Config) (*StyleguideUpdate, error) {

	actionStyleguideUpdate := &StyleguideUpdate{Config: *cfg}
	actionStyleguideUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionStyleguideUpdate.Config.Defaults["styleguide/update"]
	if defaultsPresent {
		if err := actionStyleguideUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionStyleguideUpdate, nil
}

func (cmd *StyleguideUpdate) Run() error {
	params := &cmd.StyleguideParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newStyleguidesList(cfg *phraseapp.Config) *StyleguidesList {

	actionStyleguidesList := &StyleguidesList{Config: *cfg}
	actionStyleguidesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionStyleguidesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionStyleguidesList.PerPage = *cfg.PerPage
	}

	return actionStyleguidesList
}

func (cmd *StyleguidesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TagParams

	ProjectID string `cli:"arg required"`
}

func newTagCreate(cfg *phraseapp.Config) (*TagCreate, error) {

	actionTagCreate := &TagCreate{Config: *cfg}
	actionTagCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTagCreate.Config.Defaults["tag/create"]
	if defaultsPresent {
		if err := actionTagCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTagCreate, nil
}

func (cmd *TagCreate) Run() error {
	params := &cmd.TagParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newTagDelete(cfg *phraseapp.Config) *TagDelete {

	actionTagDelete := &TagDelete{Config: *cfg}
	actionTagDelete.ProjectID = cfg.DefaultProjectID

	return actionTagDelete
}

func (cmd *TagDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newTagShow(cfg *phraseapp.Config) *TagShow {

	actionTagShow := &TagShow{Config: *cfg}
	actionTagShow.ProjectID = cfg.DefaultProjectID

	return actionTagShow
}

func (cmd *TagShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newTagsList(cfg *phraseapp.Config) *TagsList {

	actionTagsList := &TagsList{Config: *cfg}
	actionTagsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTagsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTagsList.PerPage = *cfg.PerPage
	}

	return actionTagsList
}

func (cmd *TagsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationParams

	ProjectID string `cli:"arg required"`
}

func newTranslationCreate(cfg *phraseapp.Config) (*TranslationCreate, error) {

	actionTranslationCreate := &TranslationCreate{Config: *cfg}
	actionTranslationCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationCreate.Config.Defaults["translation/create"]
	if defaultsPresent {
		if err := actionTranslationCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationCreate, nil
}

func (cmd *TranslationCreate) Run() error {
	params := &cmd.TranslationParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationShow(cfg *phraseapp.Config) *TranslationShow {

	actionTranslationShow := &TranslationShow{Config: *cfg}
	actionTranslationShow.ProjectID = cfg.DefaultProjectID

	return actionTranslationShow
}

func (cmd *TranslationShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationUpdateParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationUpdate(cfg *phraseapp.Config) (*TranslationUpdate, error) {

	actionTranslationUpdate := &TranslationUpdate{Config: *cfg}
	actionTranslationUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationUpdate.Config.Defaults["translation/update"]
	if defaultsPresent {
		if err := actionTranslationUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationUpdate, nil
}

func (cmd *TranslationUpdate) Run() error {
	params := &cmd.TranslationUpdateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsByKeyParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func newTranslationsByKey(cfg *phraseapp.Config) (*TranslationsByKey, error) {

	actionTranslationsByKey := &TranslationsByKey{Config: *cfg}
	actionTranslationsByKey.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTranslationsByKey.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsByKey.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionTranslationsByKey.Config.Defaults["translations/by_key"]
	if defaultsPresent {
		if err := actionTranslationsByKey.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsByKey, nil
}

func (cmd *TranslationsByKey) Run() error {
	params := &cmd.TranslationsByKeyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsByLocaleParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	LocaleID  string `cli:"arg required"`
}

func newTranslationsByLocale(cfg *phraseapp.Config) (*TranslationsByLocale, error) {

	actionTranslationsByLocale := &TranslationsByLocale{Config: *cfg}
	actionTranslationsByLocale.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTranslationsByLocale.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsByLocale.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionTranslationsByLocale.Config.Defaults["translations/by_locale"]
	if defaultsPresent {
		if err := actionTranslationsByLocale.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsByLocale, nil
}

func (cmd *TranslationsByLocale) Run() error {
	params := &cmd.TranslationsByLocaleParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsExcludeParams

	ProjectID string `cli:"arg required"`
}

func newTranslationsExclude(cfg *phraseapp.Config) (*TranslationsExclude, error) {

	actionTranslationsExclude := &TranslationsExclude{Config: *cfg}
	actionTranslationsExclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsExclude.Config.Defaults["translations/exclude"]
	if defaultsPresent {
		if err := actionTranslationsExclude.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsExclude, nil
}

func (cmd *TranslationsExclude) Run() error {
	params := &cmd.TranslationsExcludeParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsIncludeParams

	ProjectID string `cli:"arg required"`
}

func newTranslationsInclude(cfg *phraseapp.Config) (*TranslationsInclude, error) {

	actionTranslationsInclude := &TranslationsInclude{Config: *cfg}
	actionTranslationsInclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsInclude.Config.Defaults["translations/include"]
	if defaultsPresent {
		if err := actionTranslationsInclude.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsInclude, nil
}

func (cmd *TranslationsInclude) Run() error {
	params := &cmd.TranslationsIncludeParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newTranslationsList(cfg *phraseapp.Config) (*TranslationsList, error) {

	actionTranslationsList := &TranslationsList{Config: *cfg}
	actionTranslationsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTranslationsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionTranslationsList.Config.Defaults["translations/list"]
	if defaultsPresent {
		if err := actionTranslationsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsList, nil
}

func (cmd *TranslationsList) Run() error {
	params := &cmd.TranslationsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsSearchParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newTranslationsSearch(cfg *phraseapp.Config) (*TranslationsSearch, error) {

	actionTranslationsSearch := &TranslationsSearch{Config: *cfg}
	actionTranslationsSearch.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTranslationsSearch.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTranslationsSearch.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionTranslationsSearch.Config.Defaults["translations/search"]
	if defaultsPresent {
		if err := actionTranslationsSearch.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsSearch, nil
}

func (cmd *TranslationsSearch) Run() error {
	params := &cmd.TranslationsSearchParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsUnverifyParams

	ProjectID string `cli:"arg required"`
}

func newTranslationsUnverify(cfg *phraseapp.Config) (*TranslationsUnverify, error) {

	actionTranslationsUnverify := &TranslationsUnverify{Config: *cfg}
	actionTranslationsUnverify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsUnverify.Config.Defaults["translations/unverify"]
	if defaultsPresent {
		if err := actionTranslationsUnverify.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsUnverify, nil
}

func (cmd *TranslationsUnverify) Run() error {
	params := &cmd.TranslationsUnverifyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.TranslationsVerifyParams

	ProjectID string `cli:"arg required"`
}

func newTranslationsVerify(cfg *phraseapp.Config) (*TranslationsVerify, error) {

	actionTranslationsVerify := &TranslationsVerify{Config: *cfg}
	actionTranslationsVerify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsVerify.Config.Defaults["translations/verify"]
	if defaultsPresent {
		if err := actionTranslationsVerify.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsVerify, nil
}

func (cmd *TranslationsVerify) Run() error {
	params := &cmd.TranslationsVerifyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.UploadParams

	ProjectID string `cli:"arg required"`
}

func newUploadCreate(cfg *phraseapp.Config) (*UploadCreate, error) {

	actionUploadCreate := &UploadCreate{Config: *cfg}
	actionUploadCreate.ProjectID = cfg.DefaultProjectID
	if cfg.DefaultFileFormat != "" {
		actionUploadCreate.FileFormat = &cfg.DefaultFileFormat
	}

	val, defaultsPresent := actionUploadCreate.Config.Defaults["upload/create"]
	if defaultsPresent {
		if err := actionUploadCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionUploadCreate, nil
}

func (cmd *UploadCreate) Run() error {
	params := &cmd.UploadParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.UploadShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newUploadShow(cfg *phraseapp.Config) (*UploadShow, error) {

	actionUploadShow := &UploadShow{Config: *cfg}
	actionUploadShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionUploadShow.Config.Defaults["upload/show"]
	if defaultsPresent {
		if err := actionUploadShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionUploadShow, nil
}

func (cmd *UploadShow) Run() error {
	params := &cmd.UploadShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.UploadShow(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type UploadsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newUploadsList(cfg *phraseapp.Config) *UploadsList {

	actionUploadsList := &UploadsList{Config: *cfg}
	actionUploadsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionUploadsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionUploadsList.PerPage = *cfg.PerPage
	}

	return actionUploadsList
}

func (cmd *UploadsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
	ID            string `cli:"arg required"`
}

func newVersionShow(cfg *phraseapp.Config) *VersionShow {

	actionVersionShow := &VersionShow{Config: *cfg}
	actionVersionShow.ProjectID = cfg.DefaultProjectID

	return actionVersionShow
}

func (cmd *VersionShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
}

func newVersionsList(cfg *phraseapp.Config) *VersionsList {

	actionVersionsList := &VersionsList{Config: *cfg}
	actionVersionsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionVersionsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionVersionsList.PerPage = *cfg.PerPage
	}

	return actionVersionsList
}

func (cmd *VersionsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.WebhookParams

	ProjectID string `cli:"arg required"`
}

func newWebhookCreate(cfg *phraseapp.Config) (*WebhookCreate, error) {

	actionWebhookCreate := &WebhookCreate{Config: *cfg}
	actionWebhookCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionWebhookCreate.Config.Defaults["webhook/create"]
	if defaultsPresent {
		if err := actionWebhookCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionWebhookCreate, nil
}

func (cmd *WebhookCreate) Run() error {
	params := &cmd.WebhookParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newWebhookDelete(cfg *phraseapp.Config) *WebhookDelete {

	actionWebhookDelete := &WebhookDelete{Config: *cfg}
	actionWebhookDelete.ProjectID = cfg.DefaultProjectID

	return actionWebhookDelete
}

func (cmd *WebhookDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newWebhookShow(cfg *phraseapp.Config) *WebhookShow {

	actionWebhookShow := &WebhookShow{Config: *cfg}
	actionWebhookShow.ProjectID = cfg.DefaultProjectID

	return actionWebhookShow
}

func (cmd *WebhookShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newWebhookTest(cfg *phraseapp.Config) *WebhookTest {

	actionWebhookTest := &WebhookTest{Config: *cfg}
	actionWebhookTest.ProjectID = cfg.DefaultProjectID

	return actionWebhookTest
}

func (cmd *WebhookTest) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	phraseapp.WebhookParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newWebhookUpdate(cfg *phraseapp.Config) (*WebhookUpdate, error) {

	actionWebhookUpdate := &WebhookUpdate{Config: *cfg}
	actionWebhookUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionWebhookUpdate.Config.Defaults["webhook/update"]
	if defaultsPresent {
		if err := actionWebhookUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionWebhookUpdate, nil
}

func (cmd *WebhookUpdate) Run() error {
	params := &cmd.WebhookParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
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
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newWebhooksList(cfg *phraseapp.Config) *WebhooksList {

	actionWebhooksList := &WebhooksList{Config: *cfg}
	actionWebhooksList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionWebhooksList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionWebhooksList.PerPage = *cfg.PerPage
	}

	return actionWebhooksList
}

func (cmd *WebhooksList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.WebhooksList(cmd.ProjectID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}
