package main

import (
	"encoding/json"
	"os"

	"github.com/phrase/phraseapp-client/cli"
	"github.com/phrase/phraseapp-go/phraseapp"
)

const (
	RevisionDocs      = "b33e5c3d0e590c6aed24f39a2c98731c6b2116d9"
	RevisionGenerator = "HEAD/2020-04-02T173135/soenke"
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

	if cmd, err := newBitbucketSyncExport(cfg); err != nil {
		return nil, err
	} else {
		r.Register("bitbucket_sync/export", cmd, "Export translations from Phrase to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.")
	}

	if cmd, err := newBitbucketSyncImport(cfg); err != nil {
		return nil, err
	} else {
		r.Register("bitbucket_sync/import", cmd, "Import translations from Bitbucket to Phrase according to the .phraseapp.yml file within the Bitbucket repository.")
	}

	if cmd, err := newBitbucketSyncsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("bitbucket_syncs/list", cmd, "List all Bitbucket repositories for which synchronisation with Phrase is activated.")
	}

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

	if cmd, err := newBranchCompare(cfg); err != nil {
		return nil, err
	} else {
		r.Register("branch/compare", cmd, "Compare branch with main branch.")
	}

	if cmd, err := newBranchCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("branch/create", cmd, "Create a new branch.")
	}

	r.Register("branch/delete", newBranchDelete(cfg), "Delete an existing branch.")

	if cmd, err := newBranchMerge(cfg); err != nil {
		return nil, err
	} else {
		r.Register("branch/merge", cmd, "Merge an existing branch.")
	}

	r.Register("branch/show", newBranchShow(cfg), "Get details on a single branch for a given project.")

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

	if cmd, err := newCommentDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/delete", cmd, "Delete an existing comment.")
	}

	if cmd, err := newCommentMarkCheck(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/mark/check", cmd, "Check if comment was marked as read. Returns 204 if read, 404 if unread.")
	}

	if cmd, err := newCommentMarkRead(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/mark/read", cmd, "Mark a comment as read.")
	}

	if cmd, err := newCommentMarkUnread(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/mark/unread", cmd, "Mark a comment as unread.")
	}

	if cmd, err := newCommentShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/show", cmd, "Get details on a single comment.")
	}

	if cmd, err := newCommentUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comment/update", cmd, "Update an existing comment.")
	}

	if cmd, err := newCommentsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("comments/list", cmd, "List all comments for a key.")
	}

	if cmd, err := newDistributionCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("distribution/create", cmd, "Create a new distribution.")
	}

	r.Register("distribution/delete", newDistributionDelete(cfg), "Delete an existing distribution.")

	r.Register("distribution/show", newDistributionShow(cfg), "Get details on a single distribution.")

	if cmd, err := newDistributionUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("distribution/update", cmd, "Update an existing distribution.")
	}

	r.Register("distributions/list", newDistributionsList(cfg), "List all distributions for the given account.")

	r.Register("formats/list", newFormatsList(cfg), "Get a handy list of all localization file formats supported in Phrase.")

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

	if cmd, err := newJobComplete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/complete", cmd, "Mark a job as completed.")
	}

	if cmd, err := newJobCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/create", cmd, "Create a new job.")
	}

	if cmd, err := newJobDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/delete", cmd, "Delete an existing job.")
	}

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

	if cmd, err := newJobReopen(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/reopen", cmd, "Mark a job as uncompleted.")
	}

	if cmd, err := newJobShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/show", cmd, "Get details on a single job for a given project.")
	}

	if cmd, err := newJobStart(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/start", cmd, "Starts an existing job in state draft.")
	}

	if cmd, err := newJobUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job/update", cmd, "Update an existing job.")
	}

	if cmd, err := newJobLocaleComplete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/complete", cmd, "Mark a job locale as completed.")
	}

	if cmd, err := newJobLocaleDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/delete", cmd, "Delete an existing job locale.")
	}

	if cmd, err := newJobLocaleReopen(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/reopen", cmd, "Mark a job locale as uncompleted.")
	}

	if cmd, err := newJobLocaleShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/show", cmd, "Get a single job locale for a given job.")
	}

	if cmd, err := newJobLocaleUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locale/update", cmd, "Update an existing job locale.")
	}

	if cmd, err := newJobLocalesCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locales/create", cmd, "Create a new job locale.")
	}

	if cmd, err := newJobLocalesList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("job_locales/list", cmd, "List all job locales for a given job.")
	}

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

	if cmd, err := newKeyDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("key/delete", cmd, "Delete an existing key.")
	}

	if cmd, err := newKeyShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("key/show", cmd, "Get details on a single key for a given project.")
	}

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

	if cmd, err := newLocaleDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("locale/delete", cmd, "Delete an existing locale.")
	}

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

	r.Register("member/delete", newMemberDelete(cfg), "Remove a user from the account. The user will be removed from the account but not deleted from Phrase. Access token scope must include <code>team.manage</code>.")

	r.Register("member/show", newMemberShow(cfg), "Get details on a single user in the account. Access token scope must include <code>team.manage</code>.")

	if cmd, err := newMemberUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("member/update", cmd, "Update user permissions in the account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.")
	}

	r.Register("members/list", newMembersList(cfg), "Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.")

	if cmd, err := newOrderConfirm(cfg); err != nil {
		return nil, err
	} else {
		r.Register("order/confirm", cmd, "Confirm an existing order and send it to the provider for translation. Same constraints as for create.")
	}

	if cmd, err := newOrderCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("order/create", cmd, "Create a new order. Access token scope must include <code>orders.create</code>.")
	}

	if cmd, err := newOrderDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("order/delete", cmd, "Cancel an existing order. Must not yet be confirmed.")
	}

	if cmd, err := newOrderShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("order/show", cmd, "Get details on a single order.")
	}

	if cmd, err := newOrdersList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("orders/list", cmd, "List all orders for the given project.")
	}

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

	if cmd, err := newReleaseCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("release/create", cmd, "Create a new release.")
	}

	r.Register("release/delete", newReleaseDelete(cfg), "Delete an existing release.")

	r.Register("release/publish", newReleasePublish(cfg), "Publish a release for production.")

	r.Register("release/show", newReleaseShow(cfg), "Get details on a single release.")

	if cmd, err := newReleaseUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("release/update", cmd, "Update an existing release.")
	}

	r.Register("releases/list", newReleasesList(cfg), "List all releases for the given distribution.")

	if cmd, err := newScreenshotCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("screenshot/create", cmd, "Create a new screenshot.")
	}

	r.Register("screenshot/delete", newScreenshotDelete(cfg), "Delete an existing screenshot.")

	r.Register("screenshot/show", newScreenshotShow(cfg), "Get details on a single screenshot for a given project.")

	if cmd, err := newScreenshotUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("screenshot/update", cmd, "Update an existing screenshot.")
	}

	if cmd, err := newScreenshotMarkerCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("screenshot_marker/create", cmd, "Create a new screenshot marker.")
	}

	r.Register("screenshot_marker/delete", newScreenshotMarkerDelete(cfg), "Delete an existing screenshot marker.")

	r.Register("screenshot_marker/show", newScreenshotMarkerShow(cfg), "Get details on a single screenshot marker for a given project.")

	if cmd, err := newScreenshotMarkerUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("screenshot_marker/update", cmd, "Update an existing screenshot marker.")
	}

	r.Register("screenshot_markers/list", newScreenshotMarkersList(cfg), "List all screenshot markers for the given project.")

	r.Register("screenshots/list", newScreenshotsList(cfg), "List all screenshots for the given project.")

	r.Register("show/user", newShowUser(cfg), "Show details for current User.")

	if cmd, err := newSpaceCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("space/create", cmd, "Create a new Space.")
	}

	r.Register("space/delete", newSpaceDelete(cfg), "Delete the specified Space.")

	r.Register("space/show", newSpaceShow(cfg), "Show the specified Space.")

	if cmd, err := newSpaceUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("space/update", cmd, "Update the specified Space.")
	}

	r.Register("spaces/list", newSpacesList(cfg), "List all Spaces for the given account.")

	if cmd, err := newSpacesProjectsCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("spaces/projects/create", cmd, "Adds an existing project to the space.")
	}

	r.Register("spaces/projects/delete", newSpacesProjectsDelete(cfg), "Removes a specified project from the specified space.")

	r.Register("spaces/projects/list", newSpacesProjectsList(cfg), "List all projects for the specified Space.")

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

	if cmd, err := newTagDelete(cfg); err != nil {
		return nil, err
	} else {
		r.Register("tag/delete", cmd, "Delete an existing tag.")
	}

	if cmd, err := newTagShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("tag/show", cmd, "Get details and progress information on a single tag for a given project.")
	}

	if cmd, err := newTagsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("tags/list", cmd, "List all tags for the given project.")
	}

	if cmd, err := newTranslationCreate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/create", cmd, "Create a translation.")
	}

	if cmd, err := newTranslationExclude(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/exclude", cmd, "Set exclude from export flag on an existing translation.")
	}

	if cmd, err := newTranslationInclude(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/include", cmd, "Remove exclude from export flag from an existing translation.")
	}

	if cmd, err := newTranslationReview(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/review", cmd, "Mark an existing translation as reviewed.")
	}

	if cmd, err := newTranslationShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/show", cmd, "Get details on a single translation.")
	}

	if cmd, err := newTranslationUnverify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/unverify", cmd, "Mark an existing translation as unverified.")
	}

	if cmd, err := newTranslationUpdate(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/update", cmd, "Update an existing translation.")
	}

	if cmd, err := newTranslationVerify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translation/verify", cmd, "Verify an existing translation.")
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

	if cmd, err := newTranslationsReview(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/review", cmd, "Review translations matching query.")
	}

	if cmd, err := newTranslationsSearch(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/search", cmd, "Search translations for the given project. Provides the same search interface as <code>translations#index</code> but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.")
	}

	if cmd, err := newTranslationsUnverify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/unverify", cmd, "Mark translations matching query as unverified.")
	}

	if cmd, err := newTranslationsVerify(cfg); err != nil {
		return nil, err
	} else {
		r.Register("translations/verify", cmd, "Verify translations matching query.")
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

	if cmd, err := newUploadsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("uploads/list", cmd, "List all uploads for the given project.")
	}

	if cmd, err := newVersionShow(cfg); err != nil {
		return nil, err
	} else {
		r.Register("version/show", cmd, "Get details on a single version.")
	}

	if cmd, err := newVersionsList(cfg); err != nil {
		return nil, err
	} else {
		r.Register("versions/list", cmd, "List all versions for the given translation.")
	}

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

type BitbucketSyncExport struct {
	phraseapp.Config

	phraseapp.BitbucketSyncParams

	ID string `cli:"arg required"`
}

func newBitbucketSyncExport(cfg *phraseapp.Config) (*BitbucketSyncExport, error) {

	actionBitbucketSyncExport := &BitbucketSyncExport{Config: *cfg}

	val, defaultsPresent := actionBitbucketSyncExport.Config.Defaults["bitbucket_sync/export"]
	if defaultsPresent {
		if err := actionBitbucketSyncExport.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBitbucketSyncExport, nil
}

func (cmd *BitbucketSyncExport) Run() error {
	params := &cmd.BitbucketSyncParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BitbucketSyncExport(cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BitbucketSyncImport struct {
	phraseapp.Config

	phraseapp.BitbucketSyncParams

	ID string `cli:"arg required"`
}

func newBitbucketSyncImport(cfg *phraseapp.Config) (*BitbucketSyncImport, error) {

	actionBitbucketSyncImport := &BitbucketSyncImport{Config: *cfg}

	val, defaultsPresent := actionBitbucketSyncImport.Config.Defaults["bitbucket_sync/import"]
	if defaultsPresent {
		if err := actionBitbucketSyncImport.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBitbucketSyncImport, nil
}

func (cmd *BitbucketSyncImport) Run() error {
	params := &cmd.BitbucketSyncParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BitbucketSyncImport(cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type BitbucketSyncsList struct {
	phraseapp.Config

	phraseapp.BitbucketSyncParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`
}

func newBitbucketSyncsList(cfg *phraseapp.Config) (*BitbucketSyncsList, error) {

	actionBitbucketSyncsList := &BitbucketSyncsList{Config: *cfg}
	if cfg.Page != nil {
		actionBitbucketSyncsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionBitbucketSyncsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionBitbucketSyncsList.Config.Defaults["bitbucket_syncs/list"]
	if defaultsPresent {
		if err := actionBitbucketSyncsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBitbucketSyncsList, nil
}

func (cmd *BitbucketSyncsList) Run() error {
	params := &cmd.BitbucketSyncParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BitbucketSyncsList(cmd.Page, cmd.PerPage, params)

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

	phraseapp.BranchParams

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchCompare(cfg *phraseapp.Config) (*BranchCompare, error) {

	actionBranchCompare := &BranchCompare{Config: *cfg}
	actionBranchCompare.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchCompare.Config.Defaults["branch/compare"]
	if defaultsPresent {
		if err := actionBranchCompare.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBranchCompare, nil
}

func (cmd *BranchCompare) Run() error {
	params := &cmd.BranchParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchCompare(cmd.ProjectID, cmd.Name, params)

	if err != nil {
		return err
	}

	return nil
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

type BranchDelete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchDelete(cfg *phraseapp.Config) *BranchDelete {

	actionBranchDelete := &BranchDelete{Config: *cfg}
	actionBranchDelete.ProjectID = cfg.DefaultProjectID

	return actionBranchDelete
}

func (cmd *BranchDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchDelete(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return nil
}

type BranchMerge struct {
	phraseapp.Config

	phraseapp.BranchMergeParams

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchMerge(cfg *phraseapp.Config) (*BranchMerge, error) {

	actionBranchMerge := &BranchMerge{Config: *cfg}
	actionBranchMerge.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchMerge.Config.Defaults["branch/merge"]
	if defaultsPresent {
		if err := actionBranchMerge.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionBranchMerge, nil
}

func (cmd *BranchMerge) Run() error {
	params := &cmd.BranchMergeParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchMerge(cmd.ProjectID, cmd.Name, params)

	if err != nil {
		return err
	}

	return nil
}

type BranchShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newBranchShow(cfg *phraseapp.Config) *BranchShow {

	actionBranchShow := &BranchShow{Config: *cfg}
	actionBranchShow.ProjectID = cfg.DefaultProjectID

	return actionBranchShow
}

func (cmd *BranchShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchShow(cmd.ProjectID, cmd.Name)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
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

	phraseapp.CommentDeleteParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentDelete(cfg *phraseapp.Config) (*CommentDelete, error) {

	actionCommentDelete := &CommentDelete{Config: *cfg}
	actionCommentDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentDelete.Config.Defaults["comment/delete"]
	if defaultsPresent {
		if err := actionCommentDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentDelete, nil
}

func (cmd *CommentDelete) Run() error {
	params := &cmd.CommentDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.CommentDelete(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkCheck struct {
	phraseapp.Config

	phraseapp.CommentMarkCheckParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkCheck(cfg *phraseapp.Config) (*CommentMarkCheck, error) {

	actionCommentMarkCheck := &CommentMarkCheck{Config: *cfg}
	actionCommentMarkCheck.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentMarkCheck.Config.Defaults["comment/mark/check"]
	if defaultsPresent {
		if err := actionCommentMarkCheck.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentMarkCheck, nil
}

func (cmd *CommentMarkCheck) Run() error {
	params := &cmd.CommentMarkCheckParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.CommentMarkCheck(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkRead struct {
	phraseapp.Config

	phraseapp.CommentMarkReadParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkRead(cfg *phraseapp.Config) (*CommentMarkRead, error) {

	actionCommentMarkRead := &CommentMarkRead{Config: *cfg}
	actionCommentMarkRead.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentMarkRead.Config.Defaults["comment/mark/read"]
	if defaultsPresent {
		if err := actionCommentMarkRead.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentMarkRead, nil
}

func (cmd *CommentMarkRead) Run() error {
	params := &cmd.CommentMarkReadParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.CommentMarkRead(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type CommentMarkUnread struct {
	phraseapp.Config

	phraseapp.CommentMarkUnreadParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentMarkUnread(cfg *phraseapp.Config) (*CommentMarkUnread, error) {

	actionCommentMarkUnread := &CommentMarkUnread{Config: *cfg}
	actionCommentMarkUnread.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentMarkUnread.Config.Defaults["comment/mark/unread"]
	if defaultsPresent {
		if err := actionCommentMarkUnread.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentMarkUnread, nil
}

func (cmd *CommentMarkUnread) Run() error {
	params := &cmd.CommentMarkUnreadParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.CommentMarkUnread(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type CommentShow struct {
	phraseapp.Config

	phraseapp.CommentShowParams

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newCommentShow(cfg *phraseapp.Config) (*CommentShow, error) {

	actionCommentShow := &CommentShow{Config: *cfg}
	actionCommentShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentShow.Config.Defaults["comment/show"]
	if defaultsPresent {
		if err := actionCommentShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentShow, nil
}

func (cmd *CommentShow) Run() error {
	params := &cmd.CommentShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.CommentShow(cmd.ProjectID, cmd.KeyID, cmd.ID, params)

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

	phraseapp.CommentsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	KeyID     string `cli:"arg required"`
}

func newCommentsList(cfg *phraseapp.Config) (*CommentsList, error) {

	actionCommentsList := &CommentsList{Config: *cfg}
	actionCommentsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionCommentsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionCommentsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionCommentsList.Config.Defaults["comments/list"]
	if defaultsPresent {
		if err := actionCommentsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionCommentsList, nil
}

func (cmd *CommentsList) Run() error {
	params := &cmd.CommentsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.CommentsList(cmd.ProjectID, cmd.KeyID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type DistributionCreate struct {
	phraseapp.Config

	phraseapp.DistributionsParams

	AccountID string `cli:"arg required"`
}

func newDistributionCreate(cfg *phraseapp.Config) (*DistributionCreate, error) {

	actionDistributionCreate := &DistributionCreate{Config: *cfg}

	val, defaultsPresent := actionDistributionCreate.Config.Defaults["distribution/create"]
	if defaultsPresent {
		if err := actionDistributionCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionDistributionCreate, nil
}

func (cmd *DistributionCreate) Run() error {
	params := &cmd.DistributionsParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.DistributionCreate(cmd.AccountID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type DistributionDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newDistributionDelete(cfg *phraseapp.Config) *DistributionDelete {

	actionDistributionDelete := &DistributionDelete{Config: *cfg}

	return actionDistributionDelete
}

func (cmd *DistributionDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.DistributionDelete(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type DistributionShow struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newDistributionShow(cfg *phraseapp.Config) *DistributionShow {

	actionDistributionShow := &DistributionShow{Config: *cfg}

	return actionDistributionShow
}

func (cmd *DistributionShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.DistributionShow(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type DistributionUpdate struct {
	phraseapp.Config

	phraseapp.DistributionsParams

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newDistributionUpdate(cfg *phraseapp.Config) (*DistributionUpdate, error) {

	actionDistributionUpdate := &DistributionUpdate{Config: *cfg}

	val, defaultsPresent := actionDistributionUpdate.Config.Defaults["distribution/update"]
	if defaultsPresent {
		if err := actionDistributionUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionDistributionUpdate, nil
}

func (cmd *DistributionUpdate) Run() error {
	params := &cmd.DistributionsParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.DistributionUpdate(cmd.AccountID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type DistributionsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
}

func newDistributionsList(cfg *phraseapp.Config) *DistributionsList {

	actionDistributionsList := &DistributionsList{Config: *cfg}
	if cfg.Page != nil {
		actionDistributionsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionDistributionsList.PerPage = *cfg.PerPage
	}

	return actionDistributionsList
}

func (cmd *DistributionsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.DistributionsList(cmd.AccountID, cmd.Page, cmd.PerPage)

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

	phraseapp.JobCompleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobComplete(cfg *phraseapp.Config) (*JobComplete, error) {

	actionJobComplete := &JobComplete{Config: *cfg}
	actionJobComplete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobComplete.Config.Defaults["job/complete"]
	if defaultsPresent {
		if err := actionJobComplete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobComplete, nil
}

func (cmd *JobComplete) Run() error {
	params := &cmd.JobCompleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobComplete(cmd.ProjectID, cmd.ID, params)

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

	phraseapp.JobDeleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobDelete(cfg *phraseapp.Config) (*JobDelete, error) {

	actionJobDelete := &JobDelete{Config: *cfg}
	actionJobDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobDelete.Config.Defaults["job/delete"]
	if defaultsPresent {
		if err := actionJobDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobDelete, nil
}

func (cmd *JobDelete) Run() error {
	params := &cmd.JobDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.JobDelete(cmd.ProjectID, cmd.ID, params)

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

type JobReopen struct {
	phraseapp.Config

	phraseapp.JobReopenParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobReopen(cfg *phraseapp.Config) (*JobReopen, error) {

	actionJobReopen := &JobReopen{Config: *cfg}
	actionJobReopen.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobReopen.Config.Defaults["job/reopen"]
	if defaultsPresent {
		if err := actionJobReopen.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobReopen, nil
}

func (cmd *JobReopen) Run() error {
	params := &cmd.JobReopenParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobReopen(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobShow struct {
	phraseapp.Config

	phraseapp.JobShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobShow(cfg *phraseapp.Config) (*JobShow, error) {

	actionJobShow := &JobShow{Config: *cfg}
	actionJobShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobShow.Config.Defaults["job/show"]
	if defaultsPresent {
		if err := actionJobShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobShow, nil
}

func (cmd *JobShow) Run() error {
	params := &cmd.JobShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobShow(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobStart struct {
	phraseapp.Config

	phraseapp.JobStartParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobStart(cfg *phraseapp.Config) (*JobStart, error) {

	actionJobStart := &JobStart{Config: *cfg}
	actionJobStart.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobStart.Config.Defaults["job/start"]
	if defaultsPresent {
		if err := actionJobStart.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobStart, nil
}

func (cmd *JobStart) Run() error {
	params := &cmd.JobStartParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobStart(cmd.ProjectID, cmd.ID, params)

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

	phraseapp.JobLocaleCompleteParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleComplete(cfg *phraseapp.Config) (*JobLocaleComplete, error) {

	actionJobLocaleComplete := &JobLocaleComplete{Config: *cfg}
	actionJobLocaleComplete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleComplete.Config.Defaults["job_locale/complete"]
	if defaultsPresent {
		if err := actionJobLocaleComplete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocaleComplete, nil
}

func (cmd *JobLocaleComplete) Run() error {
	params := &cmd.JobLocaleCompleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleComplete(cmd.ProjectID, cmd.JobID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleDelete struct {
	phraseapp.Config

	phraseapp.JobLocaleDeleteParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleDelete(cfg *phraseapp.Config) (*JobLocaleDelete, error) {

	actionJobLocaleDelete := &JobLocaleDelete{Config: *cfg}
	actionJobLocaleDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleDelete.Config.Defaults["job_locale/delete"]
	if defaultsPresent {
		if err := actionJobLocaleDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocaleDelete, nil
}

func (cmd *JobLocaleDelete) Run() error {
	params := &cmd.JobLocaleDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.JobLocaleDelete(cmd.ProjectID, cmd.JobID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type JobLocaleReopen struct {
	phraseapp.Config

	phraseapp.JobLocaleReopenParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleReopen(cfg *phraseapp.Config) (*JobLocaleReopen, error) {

	actionJobLocaleReopen := &JobLocaleReopen{Config: *cfg}
	actionJobLocaleReopen.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleReopen.Config.Defaults["job_locale/reopen"]
	if defaultsPresent {
		if err := actionJobLocaleReopen.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocaleReopen, nil
}

func (cmd *JobLocaleReopen) Run() error {
	params := &cmd.JobLocaleReopenParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleReopen(cmd.ProjectID, cmd.JobID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleShow struct {
	phraseapp.Config

	phraseapp.JobLocaleShowParams

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleShow(cfg *phraseapp.Config) (*JobLocaleShow, error) {

	actionJobLocaleShow := &JobLocaleShow{Config: *cfg}
	actionJobLocaleShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleShow.Config.Defaults["job_locale/show"]
	if defaultsPresent {
		if err := actionJobLocaleShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocaleShow, nil
}

func (cmd *JobLocaleShow) Run() error {
	params := &cmd.JobLocaleShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleShow(cmd.ProjectID, cmd.JobID, cmd.ID, params)

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

	phraseapp.JobLocalesListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
}

func newJobLocalesList(cfg *phraseapp.Config) (*JobLocalesList, error) {

	actionJobLocalesList := &JobLocalesList{Config: *cfg}
	actionJobLocalesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionJobLocalesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionJobLocalesList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionJobLocalesList.Config.Defaults["job_locales/list"]
	if defaultsPresent {
		if err := actionJobLocalesList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionJobLocalesList, nil
}

func (cmd *JobLocalesList) Run() error {
	params := &cmd.JobLocalesListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocalesList(cmd.ProjectID, cmd.JobID, cmd.Page, cmd.PerPage, params)

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

	phraseapp.KeyDeleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newKeyDelete(cfg *phraseapp.Config) (*KeyDelete, error) {

	actionKeyDelete := &KeyDelete{Config: *cfg}
	actionKeyDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyDelete.Config.Defaults["key/delete"]
	if defaultsPresent {
		if err := actionKeyDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeyDelete, nil
}

func (cmd *KeyDelete) Run() error {
	params := &cmd.KeyDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.KeyDelete(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type KeyShow struct {
	phraseapp.Config

	phraseapp.KeyShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newKeyShow(cfg *phraseapp.Config) (*KeyShow, error) {

	actionKeyShow := &KeyShow{Config: *cfg}
	actionKeyShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyShow.Config.Defaults["key/show"]
	if defaultsPresent {
		if err := actionKeyShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionKeyShow, nil
}

func (cmd *KeyShow) Run() error {
	params := &cmd.KeyShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.KeyShow(cmd.ProjectID, cmd.ID, params)

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

	phraseapp.LocaleDeleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newLocaleDelete(cfg *phraseapp.Config) (*LocaleDelete, error) {

	actionLocaleDelete := &LocaleDelete{Config: *cfg}
	actionLocaleDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleDelete.Config.Defaults["locale/delete"]
	if defaultsPresent {
		if err := actionLocaleDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionLocaleDelete, nil
}

func (cmd *LocaleDelete) Run() error {
	params := &cmd.LocaleDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.LocaleDelete(cmd.ProjectID, cmd.ID, params)

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

	phraseapp.OrderConfirmParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderConfirm(cfg *phraseapp.Config) (*OrderConfirm, error) {

	actionOrderConfirm := &OrderConfirm{Config: *cfg}
	actionOrderConfirm.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionOrderConfirm.Config.Defaults["order/confirm"]
	if defaultsPresent {
		if err := actionOrderConfirm.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionOrderConfirm, nil
}

func (cmd *OrderConfirm) Run() error {
	params := &cmd.OrderConfirmParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.OrderConfirm(cmd.ProjectID, cmd.ID, params)

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

	phraseapp.OrderDeleteParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderDelete(cfg *phraseapp.Config) (*OrderDelete, error) {

	actionOrderDelete := &OrderDelete{Config: *cfg}
	actionOrderDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionOrderDelete.Config.Defaults["order/delete"]
	if defaultsPresent {
		if err := actionOrderDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionOrderDelete, nil
}

func (cmd *OrderDelete) Run() error {
	params := &cmd.OrderDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.OrderDelete(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return nil
}

type OrderShow struct {
	phraseapp.Config

	phraseapp.OrderShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newOrderShow(cfg *phraseapp.Config) (*OrderShow, error) {

	actionOrderShow := &OrderShow{Config: *cfg}
	actionOrderShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionOrderShow.Config.Defaults["order/show"]
	if defaultsPresent {
		if err := actionOrderShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionOrderShow, nil
}

func (cmd *OrderShow) Run() error {
	params := &cmd.OrderShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.OrderShow(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type OrdersList struct {
	phraseapp.Config

	phraseapp.OrdersListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newOrdersList(cfg *phraseapp.Config) (*OrdersList, error) {

	actionOrdersList := &OrdersList{Config: *cfg}
	actionOrdersList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionOrdersList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionOrdersList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionOrdersList.Config.Defaults["orders/list"]
	if defaultsPresent {
		if err := actionOrdersList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionOrdersList, nil
}

func (cmd *OrdersList) Run() error {
	params := &cmd.OrdersListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.OrdersList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

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

type ReleaseCreate struct {
	phraseapp.Config

	phraseapp.ReleasesParams

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
}

func newReleaseCreate(cfg *phraseapp.Config) (*ReleaseCreate, error) {

	actionReleaseCreate := &ReleaseCreate{Config: *cfg}

	val, defaultsPresent := actionReleaseCreate.Config.Defaults["release/create"]
	if defaultsPresent {
		if err := actionReleaseCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionReleaseCreate, nil
}

func (cmd *ReleaseCreate) Run() error {
	params := &cmd.ReleasesParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ReleaseCreate(cmd.AccountID, cmd.DistributionID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ReleaseDelete struct {
	phraseapp.Config

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
	ID             string `cli:"arg required"`
}

func newReleaseDelete(cfg *phraseapp.Config) *ReleaseDelete {

	actionReleaseDelete := &ReleaseDelete{Config: *cfg}

	return actionReleaseDelete
}

func (cmd *ReleaseDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.ReleaseDelete(cmd.AccountID, cmd.DistributionID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type ReleasePublish struct {
	phraseapp.Config

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
	ID             string `cli:"arg required"`
}

func newReleasePublish(cfg *phraseapp.Config) *ReleasePublish {

	actionReleasePublish := &ReleasePublish{Config: *cfg}

	return actionReleasePublish
}

func (cmd *ReleasePublish) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ReleasePublish(cmd.AccountID, cmd.DistributionID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ReleaseShow struct {
	phraseapp.Config

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
	ID             string `cli:"arg required"`
}

func newReleaseShow(cfg *phraseapp.Config) *ReleaseShow {

	actionReleaseShow := &ReleaseShow{Config: *cfg}

	return actionReleaseShow
}

func (cmd *ReleaseShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ReleaseShow(cmd.AccountID, cmd.DistributionID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ReleaseUpdate struct {
	phraseapp.Config

	phraseapp.ReleasesParams

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
	ID             string `cli:"arg required"`
}

func newReleaseUpdate(cfg *phraseapp.Config) (*ReleaseUpdate, error) {

	actionReleaseUpdate := &ReleaseUpdate{Config: *cfg}

	val, defaultsPresent := actionReleaseUpdate.Config.Defaults["release/update"]
	if defaultsPresent {
		if err := actionReleaseUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionReleaseUpdate, nil
}

func (cmd *ReleaseUpdate) Run() error {
	params := &cmd.ReleasesParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ReleaseUpdate(cmd.AccountID, cmd.DistributionID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ReleasesList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID      string `cli:"arg required"`
	DistributionID string `cli:"arg required"`
}

func newReleasesList(cfg *phraseapp.Config) *ReleasesList {

	actionReleasesList := &ReleasesList{Config: *cfg}
	if cfg.Page != nil {
		actionReleasesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionReleasesList.PerPage = *cfg.PerPage
	}

	return actionReleasesList
}

func (cmd *ReleasesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ReleasesList(cmd.AccountID, cmd.DistributionID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotCreate struct {
	phraseapp.Config

	phraseapp.ScreenshotParams

	ProjectID string `cli:"arg required"`
}

func newScreenshotCreate(cfg *phraseapp.Config) (*ScreenshotCreate, error) {

	actionScreenshotCreate := &ScreenshotCreate{Config: *cfg}
	actionScreenshotCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionScreenshotCreate.Config.Defaults["screenshot/create"]
	if defaultsPresent {
		if err := actionScreenshotCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionScreenshotCreate, nil
}

func (cmd *ScreenshotCreate) Run() error {
	params := &cmd.ScreenshotParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotCreate(cmd.ProjectID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotDelete struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newScreenshotDelete(cfg *phraseapp.Config) *ScreenshotDelete {

	actionScreenshotDelete := &ScreenshotDelete{Config: *cfg}
	actionScreenshotDelete.ProjectID = cfg.DefaultProjectID

	return actionScreenshotDelete
}

func (cmd *ScreenshotDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.ScreenshotDelete(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type ScreenshotShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newScreenshotShow(cfg *phraseapp.Config) *ScreenshotShow {

	actionScreenshotShow := &ScreenshotShow{Config: *cfg}
	actionScreenshotShow.ProjectID = cfg.DefaultProjectID

	return actionScreenshotShow
}

func (cmd *ScreenshotShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotUpdate struct {
	phraseapp.Config

	phraseapp.ScreenshotParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newScreenshotUpdate(cfg *phraseapp.Config) (*ScreenshotUpdate, error) {

	actionScreenshotUpdate := &ScreenshotUpdate{Config: *cfg}
	actionScreenshotUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionScreenshotUpdate.Config.Defaults["screenshot/update"]
	if defaultsPresent {
		if err := actionScreenshotUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionScreenshotUpdate, nil
}

func (cmd *ScreenshotUpdate) Run() error {
	params := &cmd.ScreenshotParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotUpdate(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotMarkerCreate struct {
	phraseapp.Config

	phraseapp.ScreenshotMarkerParams

	ProjectID    string `cli:"arg required"`
	ScreenshotID string `cli:"arg required"`
}

func newScreenshotMarkerCreate(cfg *phraseapp.Config) (*ScreenshotMarkerCreate, error) {

	actionScreenshotMarkerCreate := &ScreenshotMarkerCreate{Config: *cfg}
	actionScreenshotMarkerCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionScreenshotMarkerCreate.Config.Defaults["screenshot_marker/create"]
	if defaultsPresent {
		if err := actionScreenshotMarkerCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionScreenshotMarkerCreate, nil
}

func (cmd *ScreenshotMarkerCreate) Run() error {
	params := &cmd.ScreenshotMarkerParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotMarkerCreate(cmd.ProjectID, cmd.ScreenshotID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotMarkerDelete struct {
	phraseapp.Config

	ProjectID    string `cli:"arg required"`
	ScreenshotID string `cli:"arg required"`
}

func newScreenshotMarkerDelete(cfg *phraseapp.Config) *ScreenshotMarkerDelete {

	actionScreenshotMarkerDelete := &ScreenshotMarkerDelete{Config: *cfg}
	actionScreenshotMarkerDelete.ProjectID = cfg.DefaultProjectID

	return actionScreenshotMarkerDelete
}

func (cmd *ScreenshotMarkerDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.ScreenshotMarkerDelete(cmd.ProjectID, cmd.ScreenshotID)

	if err != nil {
		return err
	}

	return nil
}

type ScreenshotMarkerShow struct {
	phraseapp.Config

	ProjectID    string `cli:"arg required"`
	ScreenshotID string `cli:"arg required"`
	ID           string `cli:"arg required"`
}

func newScreenshotMarkerShow(cfg *phraseapp.Config) *ScreenshotMarkerShow {

	actionScreenshotMarkerShow := &ScreenshotMarkerShow{Config: *cfg}
	actionScreenshotMarkerShow.ProjectID = cfg.DefaultProjectID

	return actionScreenshotMarkerShow
}

func (cmd *ScreenshotMarkerShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotMarkerShow(cmd.ProjectID, cmd.ScreenshotID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotMarkerUpdate struct {
	phraseapp.Config

	phraseapp.ScreenshotMarkerParams

	ProjectID    string `cli:"arg required"`
	ScreenshotID string `cli:"arg required"`
}

func newScreenshotMarkerUpdate(cfg *phraseapp.Config) (*ScreenshotMarkerUpdate, error) {

	actionScreenshotMarkerUpdate := &ScreenshotMarkerUpdate{Config: *cfg}
	actionScreenshotMarkerUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionScreenshotMarkerUpdate.Config.Defaults["screenshot_marker/update"]
	if defaultsPresent {
		if err := actionScreenshotMarkerUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionScreenshotMarkerUpdate, nil
}

func (cmd *ScreenshotMarkerUpdate) Run() error {
	params := &cmd.ScreenshotMarkerParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotMarkerUpdate(cmd.ProjectID, cmd.ScreenshotID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotMarkersList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newScreenshotMarkersList(cfg *phraseapp.Config) *ScreenshotMarkersList {

	actionScreenshotMarkersList := &ScreenshotMarkersList{Config: *cfg}
	actionScreenshotMarkersList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionScreenshotMarkersList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionScreenshotMarkersList.PerPage = *cfg.PerPage
	}

	return actionScreenshotMarkersList
}

func (cmd *ScreenshotMarkersList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotMarkersList(cmd.ProjectID, cmd.ID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type ScreenshotsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newScreenshotsList(cfg *phraseapp.Config) *ScreenshotsList {

	actionScreenshotsList := &ScreenshotsList{Config: *cfg}
	actionScreenshotsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionScreenshotsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionScreenshotsList.PerPage = *cfg.PerPage
	}

	return actionScreenshotsList
}

func (cmd *ScreenshotsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.ScreenshotsList(cmd.ProjectID, cmd.Page, cmd.PerPage)

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

type SpaceCreate struct {
	phraseapp.Config

	phraseapp.SpaceCreateParams

	AccountID string `cli:"arg required"`
}

func newSpaceCreate(cfg *phraseapp.Config) (*SpaceCreate, error) {

	actionSpaceCreate := &SpaceCreate{Config: *cfg}

	val, defaultsPresent := actionSpaceCreate.Config.Defaults["space/create"]
	if defaultsPresent {
		if err := actionSpaceCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionSpaceCreate, nil
}

func (cmd *SpaceCreate) Run() error {
	params := &cmd.SpaceCreateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.SpaceCreate(cmd.AccountID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type SpaceDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newSpaceDelete(cfg *phraseapp.Config) *SpaceDelete {

	actionSpaceDelete := &SpaceDelete{Config: *cfg}

	return actionSpaceDelete
}

func (cmd *SpaceDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.SpaceDelete(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type SpaceShow struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newSpaceShow(cfg *phraseapp.Config) *SpaceShow {

	actionSpaceShow := &SpaceShow{Config: *cfg}

	return actionSpaceShow
}

func (cmd *SpaceShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.SpaceShow(cmd.AccountID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type SpaceUpdate struct {
	phraseapp.Config

	phraseapp.SpaceUpdateParams

	AccountID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newSpaceUpdate(cfg *phraseapp.Config) (*SpaceUpdate, error) {

	actionSpaceUpdate := &SpaceUpdate{Config: *cfg}

	val, defaultsPresent := actionSpaceUpdate.Config.Defaults["space/update"]
	if defaultsPresent {
		if err := actionSpaceUpdate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionSpaceUpdate, nil
}

func (cmd *SpaceUpdate) Run() error {
	params := &cmd.SpaceUpdateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.SpaceUpdate(cmd.AccountID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type SpacesList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
}

func newSpacesList(cfg *phraseapp.Config) *SpacesList {

	actionSpacesList := &SpacesList{Config: *cfg}
	if cfg.Page != nil {
		actionSpacesList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionSpacesList.PerPage = *cfg.PerPage
	}

	return actionSpacesList
}

func (cmd *SpacesList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.SpacesList(cmd.AccountID, cmd.Page, cmd.PerPage)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type SpacesProjectsCreate struct {
	phraseapp.Config

	phraseapp.SpacesProjectsCreateParams

	AccountID string `cli:"arg required"`
	SpaceID   string `cli:"arg required"`
}

func newSpacesProjectsCreate(cfg *phraseapp.Config) (*SpacesProjectsCreate, error) {

	actionSpacesProjectsCreate := &SpacesProjectsCreate{Config: *cfg}

	val, defaultsPresent := actionSpacesProjectsCreate.Config.Defaults["spaces/projects/create"]
	if defaultsPresent {
		if err := actionSpacesProjectsCreate.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionSpacesProjectsCreate, nil
}

func (cmd *SpacesProjectsCreate) Run() error {
	params := &cmd.SpacesProjectsCreateParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.SpacesProjectsCreate(cmd.AccountID, cmd.SpaceID, params)

	if err != nil {
		return err
	}

	return nil
}

type SpacesProjectsDelete struct {
	phraseapp.Config

	AccountID string `cli:"arg required"`
	SpaceID   string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newSpacesProjectsDelete(cfg *phraseapp.Config) *SpacesProjectsDelete {

	actionSpacesProjectsDelete := &SpacesProjectsDelete{Config: *cfg}

	return actionSpacesProjectsDelete
}

func (cmd *SpacesProjectsDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.SpacesProjectsDelete(cmd.AccountID, cmd.SpaceID, cmd.ID)

	if err != nil {
		return err
	}

	return nil
}

type SpacesProjectsList struct {
	phraseapp.Config

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	AccountID string `cli:"arg required"`
	SpaceID   string `cli:"arg required"`
}

func newSpacesProjectsList(cfg *phraseapp.Config) *SpacesProjectsList {

	actionSpacesProjectsList := &SpacesProjectsList{Config: *cfg}
	if cfg.Page != nil {
		actionSpacesProjectsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionSpacesProjectsList.PerPage = *cfg.PerPage
	}

	return actionSpacesProjectsList
}

func (cmd *SpacesProjectsList) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.SpacesProjectsList(cmd.AccountID, cmd.SpaceID, cmd.Page, cmd.PerPage)

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

	phraseapp.TagDeleteParams

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newTagDelete(cfg *phraseapp.Config) (*TagDelete, error) {

	actionTagDelete := &TagDelete{Config: *cfg}
	actionTagDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTagDelete.Config.Defaults["tag/delete"]
	if defaultsPresent {
		if err := actionTagDelete.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTagDelete, nil
}

func (cmd *TagDelete) Run() error {
	params := &cmd.TagDeleteParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.TagDelete(cmd.ProjectID, cmd.Name, params)

	if err != nil {
		return err
	}

	return nil
}

type TagShow struct {
	phraseapp.Config

	phraseapp.TagShowParams

	ProjectID string `cli:"arg required"`
	Name      string `cli:"arg required"`
}

func newTagShow(cfg *phraseapp.Config) (*TagShow, error) {

	actionTagShow := &TagShow{Config: *cfg}
	actionTagShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTagShow.Config.Defaults["tag/show"]
	if defaultsPresent {
		if err := actionTagShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTagShow, nil
}

func (cmd *TagShow) Run() error {
	params := &cmd.TagShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TagShow(cmd.ProjectID, cmd.Name, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TagsList struct {
	phraseapp.Config

	phraseapp.TagsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newTagsList(cfg *phraseapp.Config) (*TagsList, error) {

	actionTagsList := &TagsList{Config: *cfg}
	actionTagsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTagsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionTagsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionTagsList.Config.Defaults["tags/list"]
	if defaultsPresent {
		if err := actionTagsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTagsList, nil
}

func (cmd *TagsList) Run() error {
	params := &cmd.TagsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TagsList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

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

type TranslationExclude struct {
	phraseapp.Config

	phraseapp.TranslationExcludeParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationExclude(cfg *phraseapp.Config) (*TranslationExclude, error) {

	actionTranslationExclude := &TranslationExclude{Config: *cfg}
	actionTranslationExclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationExclude.Config.Defaults["translation/exclude"]
	if defaultsPresent {
		if err := actionTranslationExclude.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationExclude, nil
}

func (cmd *TranslationExclude) Run() error {
	params := &cmd.TranslationExcludeParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationExclude(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationInclude struct {
	phraseapp.Config

	phraseapp.TranslationIncludeParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationInclude(cfg *phraseapp.Config) (*TranslationInclude, error) {

	actionTranslationInclude := &TranslationInclude{Config: *cfg}
	actionTranslationInclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationInclude.Config.Defaults["translation/include"]
	if defaultsPresent {
		if err := actionTranslationInclude.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationInclude, nil
}

func (cmd *TranslationInclude) Run() error {
	params := &cmd.TranslationIncludeParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationInclude(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationReview struct {
	phraseapp.Config

	phraseapp.TranslationReviewParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationReview(cfg *phraseapp.Config) (*TranslationReview, error) {

	actionTranslationReview := &TranslationReview{Config: *cfg}
	actionTranslationReview.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationReview.Config.Defaults["translation/review"]
	if defaultsPresent {
		if err := actionTranslationReview.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationReview, nil
}

func (cmd *TranslationReview) Run() error {
	params := &cmd.TranslationReviewParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationReview(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationShow struct {
	phraseapp.Config

	phraseapp.TranslationShowParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationShow(cfg *phraseapp.Config) (*TranslationShow, error) {

	actionTranslationShow := &TranslationShow{Config: *cfg}
	actionTranslationShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationShow.Config.Defaults["translation/show"]
	if defaultsPresent {
		if err := actionTranslationShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationShow, nil
}

func (cmd *TranslationShow) Run() error {
	params := &cmd.TranslationShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationShow(cmd.ProjectID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type TranslationUnverify struct {
	phraseapp.Config

	phraseapp.TranslationUnverifyParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationUnverify(cfg *phraseapp.Config) (*TranslationUnverify, error) {

	actionTranslationUnverify := &TranslationUnverify{Config: *cfg}
	actionTranslationUnverify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationUnverify.Config.Defaults["translation/unverify"]
	if defaultsPresent {
		if err := actionTranslationUnverify.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationUnverify, nil
}

func (cmd *TranslationUnverify) Run() error {
	params := &cmd.TranslationUnverifyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationUnverify(cmd.ProjectID, cmd.ID, params)

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

type TranslationVerify struct {
	phraseapp.Config

	phraseapp.TranslationVerifyParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newTranslationVerify(cfg *phraseapp.Config) (*TranslationVerify, error) {

	actionTranslationVerify := &TranslationVerify{Config: *cfg}
	actionTranslationVerify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationVerify.Config.Defaults["translation/verify"]
	if defaultsPresent {
		if err := actionTranslationVerify.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationVerify, nil
}

func (cmd *TranslationVerify) Run() error {
	params := &cmd.TranslationVerifyParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationVerify(cmd.ProjectID, cmd.ID, params)

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

type TranslationsReview struct {
	phraseapp.Config

	phraseapp.TranslationsReviewParams

	ProjectID string `cli:"arg required"`
}

func newTranslationsReview(cfg *phraseapp.Config) (*TranslationsReview, error) {

	actionTranslationsReview := &TranslationsReview{Config: *cfg}
	actionTranslationsReview.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsReview.Config.Defaults["translations/review"]
	if defaultsPresent {
		if err := actionTranslationsReview.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionTranslationsReview, nil
}

func (cmd *TranslationsReview) Run() error {
	params := &cmd.TranslationsReviewParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.TranslationsReview(cmd.ProjectID, params)

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

	phraseapp.UploadsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID string `cli:"arg required"`
}

func newUploadsList(cfg *phraseapp.Config) (*UploadsList, error) {

	actionUploadsList := &UploadsList{Config: *cfg}
	actionUploadsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionUploadsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionUploadsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionUploadsList.Config.Defaults["uploads/list"]
	if defaultsPresent {
		if err := actionUploadsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionUploadsList, nil
}

func (cmd *UploadsList) Run() error {
	params := &cmd.UploadsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.UploadsList(cmd.ProjectID, cmd.Page, cmd.PerPage, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionShow struct {
	phraseapp.Config

	phraseapp.VersionShowParams

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
	ID            string `cli:"arg required"`
}

func newVersionShow(cfg *phraseapp.Config) (*VersionShow, error) {

	actionVersionShow := &VersionShow{Config: *cfg}
	actionVersionShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionVersionShow.Config.Defaults["version/show"]
	if defaultsPresent {
		if err := actionVersionShow.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionVersionShow, nil
}

func (cmd *VersionShow) Run() error {
	params := &cmd.VersionShowParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.VersionShow(cmd.ProjectID, cmd.TranslationID, cmd.ID, params)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type VersionsList struct {
	phraseapp.Config

	phraseapp.VersionsListParams

	Page    int `cli:"opt --page default=1"`
	PerPage int `cli:"opt --per-page default=25"`

	ProjectID     string `cli:"arg required"`
	TranslationID string `cli:"arg required"`
}

func newVersionsList(cfg *phraseapp.Config) (*VersionsList, error) {

	actionVersionsList := &VersionsList{Config: *cfg}
	actionVersionsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionVersionsList.Page = *cfg.Page
	}
	if cfg.PerPage != nil {
		actionVersionsList.PerPage = *cfg.PerPage
	}

	val, defaultsPresent := actionVersionsList.Config.Defaults["versions/list"]
	if defaultsPresent {
		if err := actionVersionsList.ApplyValuesFromMap(val); err != nil {
			return nil, err
		}
	}
	return actionVersionsList, nil
}

func (cmd *VersionsList) Run() error {
	params := &cmd.VersionsListParams

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.VersionsList(cmd.ProjectID, cmd.TranslationID, cmd.Page, cmd.PerPage, params)

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
