// auto generated from the api spec
package main

import (
	"encoding/json"
	"os"

	"github.com/phrase/phraseapp-go/phraseapp"
	"github.com/urfave/cli"
)

const (
	RevisionDocs      = ""
	RevisionGenerator = ""
)

var CLICommands = []cli.Command{
	{
		Name:     "account",
		Category: "User Management",
		Subcommands: []cli.Command{
			{
				Name:        "show",
				Usage:       "Get a single account",
				Description: "Get details on a single account.",
				Action:      newAccountShow,
			},
			{
				Name:        "list",
				Usage:       "List accounts",
				Description: "List all accounts the current user has access to.",
				Action:      newAccountsList,
			},
		},
	},
	{
		Name:     "authorization",
		Category: "User Management",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create an authorization",
				Description: "Create a new authorization.",
				Action:      newAuthorizationCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete an authorization",
				Description: "Delete an existing authorization. API calls using that token will stop working.",
				Action:      newAuthorizationDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single authorization",
				Description: "Get details on a single authorization.",
				Action:      newAuthorizationShow,
			},
			{
				Name:        "update",
				Usage:       "Update an authorization",
				Description: "Update an existing authorization.",
				Action:      newAuthorizationUpdate,
			},
			{
				Name:        "list",
				Usage:       "List authorizations",
				Description: "List all your authorizations.",
				Action:      newAuthorizationsList,
			},
		},
	},
	{
		Name:     "bitbucket_sync",
		Category: "Integrations",
		Subcommands: []cli.Command{
			{
				Name:        "export",
				Usage:       "Export from PhraseApp to Bitbucket",
				Description: "Export translations from PhraseApp to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.",
				Action:      newBitbucketSyncExport,
			},
			{
				Name:        "import",
				Usage:       "Import to PhraseApp from Bitbucket",
				Description: "Import translations from Bitbucket to PhraseApp according to the .phraseapp.yml file within the Bitbucket repository.",
				Action:      newBitbucketSyncImport,
			},
			{
				Name:        "list",
				Usage:       "List Bitbucket syncs",
				Description: "List all Bitbucket repositories for which synchronisation with PhraseApp is activated.",
				Action:      newBitbucketSyncsList,
			},
		},
	},
	{
		Name:     "blacklisted_key",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a blacklisted key",
				Description: "Create a new rule for blacklisting keys.",
				Action:      newBlacklistedKeyCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a blacklisted key",
				Description: "Delete an existing rule for blacklisting keys.",
				Action:      newBlacklistedKeyDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single blacklisted key",
				Description: "Get details on a single rule for blacklisting keys for a given project.",
				Action:      newBlacklistedKeyShow,
			},
			{
				Name:        "update",
				Usage:       "Update a blacklisted key",
				Description: "Update an existing rule for blacklisting keys.",
				Action:      newBlacklistedKeyUpdate,
			},
			{
				Name:        "list",
				Usage:       "List blacklisted keys",
				Description: "List all rules for blacklisting keys for the given project.",
				Action:      newBlacklistedKeysList,
			},
		},
	},
	{
		Name:     "branch",
		Category: "Workflows",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a branch",
				Description: "Create a new branch.",
				Action:      newBranchCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a branch",
				Description: "Delete an existing branch.",
				Action:      newBranchDelete,
			},
			{
				Name:        "merge",
				Usage:       "Merge a branch",
				Description: "Merge an existing branch.",
				Action:      newBranchMerge,
			},
			{
				Name:        "show",
				Usage:       "Get a single branch",
				Description: "Get details on a single branch for a given project.",
				Action:      newBranchShow,
			},
			{
				Name:        "update",
				Usage:       "Update a branch",
				Description: "Update an existing branch.",
				Action:      newBranchUpdate,
			},
			{
				Name:        "list",
				Usage:       "List branches",
				Description: "List all branches the of the current project.",
				Action:      newBranchesList,
			},
		},
	},
	{
		Name:     "comment",
		Category: "Workflows",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a comment",
				Description: "Create a new comment for a key.",
				Action:      newCommentCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a comment",
				Description: "Delete an existing comment.",
				Action:      newCommentDelete,
			},
			{
				Name:        "mark",
				Usage:       "Check if comment is read",
				Description: "Check if comment was marked as read. Returns 204 if read, 404 if unread.",
				Action:      newCommentMarkCheck,
			},
			{
				Name:        "mark",
				Usage:       "Mark a comment as read",
				Description: "Mark a comment as read.",
				Action:      newCommentMarkRead,
			},
			{
				Name:        "mark",
				Usage:       "Mark a comment as unread",
				Description: "Mark a comment as unread.",
				Action:      newCommentMarkUnread,
			},
			{
				Name:        "show",
				Usage:       "Get a single comment",
				Description: "Get details on a single comment.",
				Action:      newCommentShow,
			},
			{
				Name:        "update",
				Usage:       "Update a comment",
				Description: "Update an existing comment.",
				Action:      newCommentUpdate,
			},
			{
				Name:        "list",
				Usage:       "List comments",
				Description: "List all comments for a key.",
				Action:      newCommentsList,
			},
		},
	},
	{
		Name: "format",

		Subcommands: []cli.Command{
			{
				Name:        "list",
				Usage:       "List formats",
				Description: "Get a handy list of all localization file formats supported in PhraseApp.",
				Action:      newFormatsList,
			},
		},
	},
	{
		Name:     "glossary",
		Category: "Quality",
		Subcommands: []cli.Command{
			{
				Name:        "list",
				Usage:       "List glossaries",
				Description: "List all glossaries the current user has access to.",
				Action:      newGlossariesList,
			},
			{
				Name:        "create",
				Usage:       "Create a glossary",
				Description: "Create a new glossary.",
				Action:      newGlossaryCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a glossary",
				Description: "Delete an existing glossary.",
				Action:      newGlossaryDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single glossary",
				Description: "Get details on a single glossary.",
				Action:      newGlossaryShow,
			},
			{
				Name:        "update",
				Usage:       "Update a glossary",
				Description: "Update an existing glossary.",
				Action:      newGlossaryUpdate,
			},
		},
	},
	{
		Name:     "glossary_term",
		Category: "Quality",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a glossary term",
				Description: "Create a new glossary term.",
				Action:      newGlossaryTermCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a glossary term",
				Description: "Delete an existing glossary term.",
				Action:      newGlossaryTermDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single glossary term",
				Description: "Get details on a single glossary term.",
				Action:      newGlossaryTermShow,
			},
			{
				Name:        "update",
				Usage:       "Update a glossary term",
				Description: "Update an existing glossary term.",
				Action:      newGlossaryTermUpdate,
			},
			{
				Name:        "list",
				Usage:       "List glossary terms",
				Description: "List all glossary terms the current user has access to.",
				Action:      newGlossaryTermsList,
			},
		},
	},
	{
		Name:     "glossary_term_translation",
		Category: "Quality",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a glossary term translation",
				Description: "Create a new glossary term translation.",
				Action:      newGlossaryTermTranslationCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a glossary term translation",
				Description: "Delete an existing glossary term translation.",
				Action:      newGlossaryTermTranslationDelete,
			},
			{
				Name:        "update",
				Usage:       "Update a glossary term translation",
				Description: "Update an existing glossary term translation.",
				Action:      newGlossaryTermTranslationUpdate,
			},
		},
	},
	{
		Name:     "invitation",
		Category: "User Management",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a new invitation",
				Description: "Invite a person to an account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete an invitation",
				Description: "Delete an existing invitation (must not be accepted yet). Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationDelete,
			},
			{
				Name:        "resend",
				Usage:       "Resend an invitation",
				Description: "Resend the invitation email (must not be accepted yet). Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationResend,
			},
			{
				Name:        "show",
				Usage:       "Get a single invitation",
				Description: "Get details on a single invitation. Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationShow,
			},
			{
				Name:        "update",
				Usage:       "Update an invitation",
				Description: "Update an existing invitation (must not be accepted yet). The <code>email</code> cannot be updated. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationUpdate,
			},
			{
				Name:        "list",
				Usage:       "List invitations",
				Description: "List invitations for an account. It will also list the accessible resources like projects and locales the invited user has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.",
				Action:      newInvitationsList,
			},
		},
	},
	{
		Name:     "job",
		Category: "Workflows",
		Subcommands: []cli.Command{
			{
				Name:        "complete",
				Usage:       "Complete a job",
				Description: "Mark a job as completed.",
				Action:      newJobComplete,
			},
			{
				Name:        "create",
				Usage:       "Create a job",
				Description: "Create a new job.",
				Action:      newJobCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a job",
				Description: "Delete an existing job.",
				Action:      newJobDelete,
			},
			{
				Name:        "keys",
				Usage:       "Add keys to job",
				Description: "Add multiple keys to a existing job.",
				Action:      newJobKeysCreate,
			},
			{
				Name:        "keys",
				Usage:       "Remove keys from job",
				Description: "Remove multiple keys from existing job.",
				Action:      newJobKeysDelete,
			},
			{
				Name:        "reopen",
				Usage:       "Reopen a job",
				Description: "Mark a job as uncompleted.",
				Action:      newJobReopen,
			},
			{
				Name:        "show",
				Usage:       "Get a single job",
				Description: "Get details on a single job for a given project.",
				Action:      newJobShow,
			},
			{
				Name:        "start",
				Usage:       "Start a job",
				Description: "Starts an existing job in state draft.",
				Action:      newJobStart,
			},
			{
				Name:        "update",
				Usage:       "Update a job",
				Description: "Update an existing job.",
				Action:      newJobUpdate,
			},
			{
				Name:        "list",
				Usage:       "List jobs",
				Description: "List all jobs for the given project.",
				Action:      newJobsList,
			},
		},
	},
	{
		Name:     "job_locale",
		Category: "Workflows",
		Subcommands: []cli.Command{
			{
				Name:        "complete",
				Usage:       "Complete a job locale",
				Description: "Mark a job locale as completed.",
				Action:      newJobLocaleComplete,
			},
			{
				Name:        "delete",
				Usage:       "Delete a job locale",
				Description: "Delete an existing job locale.",
				Action:      newJobLocaleDelete,
			},
			{
				Name:        "reopen",
				Usage:       "Reopen a job locale",
				Description: "Mark a job locale as uncompleted.",
				Action:      newJobLocaleReopen,
			},
			{
				Name:        "show",
				Usage:       "Get a single job locale",
				Description: "Get a single job locale for a given job.",
				Action:      newJobLocaleShow,
			},
			{
				Name:        "update",
				Usage:       "Update a job locale",
				Description: "Update an existing job locale.",
				Action:      newJobLocaleUpdate,
			},
			{
				Name:        "create",
				Usage:       "Create a job locale",
				Description: "Create a new job locale.",
				Action:      newJobLocalesCreate,
			},
			{
				Name:        "list",
				Usage:       "List job locales",
				Description: "List all job locales for a given job.",
				Action:      newJobLocalesList,
			},
		},
	},
	{
		Name:     "locale",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a locale",
				Description: "Create a new locale.",
				Action:      newLocaleCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a locale",
				Description: "Delete an existing locale.",
				Action:      newLocaleDelete,
			},
			{
				Name:        "download",
				Usage:       "Download a locale",
				Description: "Download a locale in a specific file format.",
				Action:      newLocaleDownload,
			},
			{
				Name:        "show",
				Usage:       "Get a single locale",
				Description: "Get details on a single locale for a given project.",
				Action:      newLocaleShow,
			},
			{
				Name:        "update",
				Usage:       "Update a locale",
				Description: "Update an existing locale.",
				Action:      newLocaleUpdate,
			},
			{
				Name:        "list",
				Usage:       "List locales",
				Description: "List all locales for the given project.",
				Action:      newLocalesList,
			},
		},
	},
	{
		Name:     "member",
		Category: "User Management",
		Subcommands: []cli.Command{
			{
				Name:        "delete",
				Usage:       "Remove a user from the account",
				Description: "Remove a user from the account. The user will be removed from the account but not deleted from PhraseApp. Access token scope must include <code>team.manage</code>.",
				Action:      newMemberDelete,
			},
			{
				Name:        "show",
				Usage:       "Get single member",
				Description: "Get details on a single user in the account. Access token scope must include <code>team.manage</code>.",
				Action:      newMemberShow,
			},
			{
				Name:        "update",
				Usage:       "Update a member",
				Description: "Update user permissions in the account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.",
				Action:      newMemberUpdate,
			},
			{
				Name:        "list",
				Usage:       "List members",
				Description: "Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.",
				Action:      newMembersList,
			},
		},
	},
	{
		Name:     "project",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a project",
				Description: "Create a new project.",
				Action:      newProjectCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a project",
				Description: "Delete an existing project.",
				Action:      newProjectDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single project",
				Description: "Get details on a single project.",
				Action:      newProjectShow,
			},
			{
				Name:        "update",
				Usage:       "Update a project",
				Description: "Update an existing project.",
				Action:      newProjectUpdate,
			},
			{
				Name:        "list",
				Usage:       "List projects",
				Description: "List all projects the current user has access to.",
				Action:      newProjectsList,
			},
		},
	},
	{
		Name:     "styleguide",
		Category: "Ordering",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a style guide",
				Description: "Create a new style guide.",
				Action:      newStyleguideCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a style guide",
				Description: "Delete an existing style guide.",
				Action:      newStyleguideDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single style guide",
				Description: "Get details on a single style guide.",
				Action:      newStyleguideShow,
			},
			{
				Name:        "update",
				Usage:       "Update a style guide",
				Description: "Update an existing style guide.",
				Action:      newStyleguideUpdate,
			},
			{
				Name:        "list",
				Usage:       "List style guides",
				Description: "List all styleguides for the given project.",
				Action:      newStyleguidesList,
			},
		},
	},
	{
		Name:     "tag",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a tag",
				Description: "Create a new tag.",
				Action:      newTagCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a tag",
				Description: "Delete an existing tag.",
				Action:      newTagDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single tag",
				Description: "Get details and progress information on a single tag for a given project.",
				Action:      newTagShow,
			},
			{
				Name:        "list",
				Usage:       "List tags",
				Description: "List all tags for the given project.",
				Action:      newTagsList,
			},
		},
	},
	{
		Name:     "translation",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a translation",
				Description: "Create a translation.",
				Action:      newTranslationCreate,
			},
			{
				Name:        "show",
				Usage:       "Get a single translation",
				Description: "Get details on a single translation.",
				Action:      newTranslationShow,
			},
			{
				Name:        "update",
				Usage:       "Update a translation",
				Description: "Update an existing translation.",
				Action:      newTranslationUpdate,
			},
			{
				Name:        "by_key",
				Usage:       "List translations by key",
				Description: "List translations for a specific key.",
				Action:      newTranslationsByKey,
			},
			{
				Name:        "by_locale",
				Usage:       "List translations by locale",
				Description: "List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.",
				Action:      newTranslationsByLocale,
			},
			{
				Name:        "exclude",
				Usage:       "Set exclude from export flag on translations selected by query",
				Description: "Exclude translations matching query from locale export.",
				Action:      newTranslationsExclude,
			},
			{
				Name:        "include",
				Usage:       "Remove exlude from import flag from translations selected by query",
				Description: "Include translations matching query in locale export.",
				Action:      newTranslationsInclude,
			},
			{
				Name:        "list",
				Usage:       "List all translations",
				Description: "List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.",
				Action:      newTranslationsList,
			},
			{
				Name:        "search",
				Usage:       "List all translations",
				Description: "List translations for the given project if you exceed GET request limitations on translations list. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.",
				Action:      newTranslationsSearch,
			},
			{
				Name:        "unverify",
				Usage:       "Mark translations selected by query as unverified",
				Description: "<div class='alert alert-info'>Only available in the <a href='https://phraseapp.com/pricing' target='_blank'>Control Package</a>.</div>Mark translations matching query as unverified.",
				Action:      newTranslationsUnverify,
			},
			{
				Name:        "verify",
				Usage:       "Verify translations selected by query",
				Description: "<div class='alert alert-info'>Only available in the <a href='https://phraseapp.com/pricing' target='_blank'>Control Package</a>.</div>Verify translations matching query.",
				Action:      newTranslationsVerify,
			},
		},
	},
	{
		Name:     "translation_key",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a key",
				Description: "Create a new key.",
				Action:      newKeyCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a key",
				Description: "Delete an existing key.",
				Action:      newKeyDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single key",
				Description: "Get details on a single key for a given project.",
				Action:      newKeyShow,
			},
			{
				Name:        "update",
				Usage:       "Update a key",
				Description: "Update an existing key.",
				Action:      newKeyUpdate,
			},
			{
				Name:        "delete",
				Usage:       "Delete collection of keys",
				Description: "Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.",
				Action:      newKeysDelete,
			},
			{
				Name:        "list",
				Usage:       "List keys",
				Description: "List all keys for the given project. Alternatively you can POST requests to /search.",
				Action:      newKeysList,
			},
			{
				Name:        "search",
				Usage:       "Search keys",
				Description: "Search keys for the given project matching query.",
				Action:      newKeysSearch,
			},
			{
				Name:        "tag",
				Usage:       "Add tags to collection of keys",
				Description: "Tags all keys matching query. Same constraints as list.",
				Action:      newKeysTag,
			},
			{
				Name:        "untag",
				Usage:       "Remove tags from collection of keys",
				Description: "Removes specified tags from keys matching query.",
				Action:      newKeysUntag,
			},
		},
	},
	{
		Name:     "translation_order",
		Category: "Ordering",
		Subcommands: []cli.Command{
			{
				Name:        "confirm",
				Usage:       "Confirm an order",
				Description: "Confirm an existing order and send it to the provider for translation. Same constraints as for create.",
				Action:      newOrderConfirm,
			},
			{
				Name:        "create",
				Usage:       "Create a new order",
				Description: "Create a new order. Access token scope must include <code>orders.create</code>.",
				Action:      newOrderCreate,
			},
			{
				Name:        "delete",
				Usage:       "Cancel an order",
				Description: "Cancel an existing order. Must not yet be confirmed.",
				Action:      newOrderDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single order",
				Description: "Get details on a single order.",
				Action:      newOrderShow,
			},
			{
				Name:        "list",
				Usage:       "List orders",
				Description: "List all orders for the given project.",
				Action:      newOrdersList,
			},
		},
	},
	{
		Name:     "translation_version",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "show",
				Usage:       "Get a single version",
				Description: "Get details on a single version.",
				Action:      newVersionShow,
			},
			{
				Name:        "list",
				Usage:       "List all versions",
				Description: "List all versions for the given translation.",
				Action:      newVersionsList,
			},
		},
	},
	{
		Name:     "upload",
		Category: "Core Resources",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Upload a new file",
				Description: "Upload a new language file. Creates necessary resources in your project.",
				Action:      newUploadCreate,
			},
			{
				Name:        "show",
				Usage:       "View upload details",
				Description: "View details and summary for a single upload.",
				Action:      newUploadShow,
			},
			{
				Name:        "list",
				Usage:       "List uploads",
				Description: "List all uploads for the given project.",
				Action:      newUploadsList,
			},
		},
	},
	{
		Name:     "user",
		Category: "User Management",
		Subcommands: []cli.Command{
			{
				Name:        "user",
				Usage:       "Show current User",
				Description: "Show details for current User.",
				Action:      newShowUser,
			},
		},
	},
	{
		Name:     "webhook",
		Category: "Integrations",
		Subcommands: []cli.Command{
			{
				Name:        "create",
				Usage:       "Create a webhook",
				Description: "Create a new webhook.",
				Action:      newWebhookCreate,
			},
			{
				Name:        "delete",
				Usage:       "Delete a webhook",
				Description: "Delete an existing webhook.",
				Action:      newWebhookDelete,
			},
			{
				Name:        "show",
				Usage:       "Get a single webhook",
				Description: "Get details on a single webhook.",
				Action:      newWebhookShow,
			},
			{
				Name:        "test",
				Usage:       "Test a webhook",
				Description: "Perform a test request for a webhook.",
				Action:      newWebhookTest,
			},
			{
				Name:        "update",
				Usage:       "Update a webhook",
				Description: "Update an existing webhook.",
				Action:      newWebhookUpdate,
			},
			{
				Name:        "list",
				Usage:       "List webhooks",
				Description: "List all webhooks for the given project.",
				Action:      newWebhooksList,
			},
		},
	},
}

type AccountShow struct {
	phraseapp.Config

	ID string `cli:"arg required"`
}

func newAccountShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAccountShow := &AccountShow{Config: *cfg}

	return actionAccountShow.Run()
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

func newAccountsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAccountsList := &AccountsList{Config: *cfg}
	if cfg.Page != nil {
		actionAccountsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionAccountsList.PerPage = *cfg.PerPage
	}

	return actionAccountsList.Run()
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

func newAuthorizationCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAuthorizationCreate := &AuthorizationCreate{Config: *cfg}

	val, defaultsPresent := actionAuthorizationCreate.Config.Defaults["authorization/create"]
	if defaultsPresent {
		if err := actionAuthorizationCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionAuthorizationCreate.Run()
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

func newAuthorizationDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAuthorizationDelete := &AuthorizationDelete{Config: *cfg}

	return actionAuthorizationDelete.Run()
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

func newAuthorizationShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAuthorizationShow := &AuthorizationShow{Config: *cfg}

	return actionAuthorizationShow.Run()
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

func newAuthorizationUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAuthorizationUpdate := &AuthorizationUpdate{Config: *cfg}

	val, defaultsPresent := actionAuthorizationUpdate.Config.Defaults["authorization/update"]
	if defaultsPresent {
		if err := actionAuthorizationUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionAuthorizationUpdate.Run()
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

func newAuthorizationsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionAuthorizationsList := &AuthorizationsList{Config: *cfg}
	if cfg.Page != nil {
		actionAuthorizationsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionAuthorizationsList.PerPage = *cfg.PerPage
	}

	return actionAuthorizationsList.Run()
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

func newBitbucketSyncExport(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBitbucketSyncExport := &BitbucketSyncExport{Config: *cfg}

	val, defaultsPresent := actionBitbucketSyncExport.Config.Defaults["bitbucket_sync/export"]
	if defaultsPresent {
		if err := actionBitbucketSyncExport.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBitbucketSyncExport.Run()
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

func newBitbucketSyncImport(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBitbucketSyncImport := &BitbucketSyncImport{Config: *cfg}

	val, defaultsPresent := actionBitbucketSyncImport.Config.Defaults["bitbucket_sync/import"]
	if defaultsPresent {
		if err := actionBitbucketSyncImport.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBitbucketSyncImport.Run()
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

func newBitbucketSyncsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionBitbucketSyncsList.Run()
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

func newBlacklistedKeyCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBlacklistedKeyCreate := &BlacklistedKeyCreate{Config: *cfg}
	actionBlacklistedKeyCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBlacklistedKeyCreate.Config.Defaults["blacklisted_key/create"]
	if defaultsPresent {
		if err := actionBlacklistedKeyCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBlacklistedKeyCreate.Run()
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

func newBlacklistedKeyDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBlacklistedKeyDelete := &BlacklistedKeyDelete{Config: *cfg}
	actionBlacklistedKeyDelete.ProjectID = cfg.DefaultProjectID

	return actionBlacklistedKeyDelete.Run()
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

func newBlacklistedKeyShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBlacklistedKeyShow := &BlacklistedKeyShow{Config: *cfg}
	actionBlacklistedKeyShow.ProjectID = cfg.DefaultProjectID

	return actionBlacklistedKeyShow.Run()
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

func newBlacklistedKeyUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBlacklistedKeyUpdate := &BlacklistedKeyUpdate{Config: *cfg}
	actionBlacklistedKeyUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBlacklistedKeyUpdate.Config.Defaults["blacklisted_key/update"]
	if defaultsPresent {
		if err := actionBlacklistedKeyUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBlacklistedKeyUpdate.Run()
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

func newBlacklistedKeysList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBlacklistedKeysList := &BlacklistedKeysList{Config: *cfg}
	actionBlacklistedKeysList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionBlacklistedKeysList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionBlacklistedKeysList.PerPage = *cfg.PerPage
	}

	return actionBlacklistedKeysList.Run()
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

type BranchCreate struct {
	phraseapp.Config
	phraseapp.BranchParams

	ProjectID string `cli:"arg required"`
}

func newBranchCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchCreate := &BranchCreate{Config: *cfg}
	actionBranchCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchCreate.Config.Defaults["branch/create"]
	if defaultsPresent {
		if err := actionBranchCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBranchCreate.Run()
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
	ID        string `cli:"arg required"`
}

func newBranchDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchDelete := &BranchDelete{Config: *cfg}
	actionBranchDelete.ProjectID = cfg.DefaultProjectID

	return actionBranchDelete.Run()
}

func (cmd *BranchDelete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchDelete(cmd.ProjectID, cmd.ID)
	if err != nil {
		return err
	}

	return nil
}

type BranchMerge struct {
	phraseapp.Config
	phraseapp.BranchMergeParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBranchMerge(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchMerge := &BranchMerge{Config: *cfg}
	actionBranchMerge.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchMerge.Config.Defaults["branch/merge"]
	if defaultsPresent {
		if err := actionBranchMerge.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBranchMerge.Run()
}

func (cmd *BranchMerge) Run() error {
	params := &cmd.BranchMergeParams
	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	err = client.BranchMerge(cmd.ProjectID, cmd.ID, params)
	if err != nil {
		return err
	}

	return nil
}

type BranchShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBranchShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchShow := &BranchShow{Config: *cfg}
	actionBranchShow.ProjectID = cfg.DefaultProjectID

	return actionBranchShow.Run()
}

func (cmd *BranchShow) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchShow(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type BranchUpdate struct {
	phraseapp.Config
	phraseapp.BranchParams

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newBranchUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchUpdate := &BranchUpdate{Config: *cfg}
	actionBranchUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionBranchUpdate.Config.Defaults["branch/update"]
	if defaultsPresent {
		if err := actionBranchUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionBranchUpdate.Run()
}

func (cmd *BranchUpdate) Run() error {
	params := &cmd.BranchParams
	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.BranchUpdate(cmd.ProjectID, cmd.ID, params)

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

func newBranchesList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionBranchesList := &BranchesList{Config: *cfg}
	actionBranchesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionBranchesList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionBranchesList.PerPage = *cfg.PerPage
	}

	return actionBranchesList.Run()
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

func newCommentCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentCreate := &CommentCreate{Config: *cfg}
	actionCommentCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentCreate.Config.Defaults["comment/create"]
	if defaultsPresent {
		if err := actionCommentCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionCommentCreate.Run()
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

func newCommentDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentDelete := &CommentDelete{Config: *cfg}
	actionCommentDelete.ProjectID = cfg.DefaultProjectID

	return actionCommentDelete.Run()
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

func newCommentMarkCheck(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentMarkCheck := &CommentMarkCheck{Config: *cfg}
	actionCommentMarkCheck.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkCheck.Run()
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

func newCommentMarkRead(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentMarkRead := &CommentMarkRead{Config: *cfg}
	actionCommentMarkRead.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkRead.Run()
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

func newCommentMarkUnread(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentMarkUnread := &CommentMarkUnread{Config: *cfg}
	actionCommentMarkUnread.ProjectID = cfg.DefaultProjectID

	return actionCommentMarkUnread.Run()
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

func newCommentShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentShow := &CommentShow{Config: *cfg}
	actionCommentShow.ProjectID = cfg.DefaultProjectID

	return actionCommentShow.Run()
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

func newCommentUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentUpdate := &CommentUpdate{Config: *cfg}
	actionCommentUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionCommentUpdate.Config.Defaults["comment/update"]
	if defaultsPresent {
		if err := actionCommentUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionCommentUpdate.Run()
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

func newCommentsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionCommentsList := &CommentsList{Config: *cfg}
	actionCommentsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionCommentsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionCommentsList.PerPage = *cfg.PerPage
	}

	return actionCommentsList.Run()
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

func newFormatsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionFormatsList := &FormatsList{Config: *cfg}
	if cfg.Page != nil {
		actionFormatsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionFormatsList.PerPage = *cfg.PerPage
	}

	return actionFormatsList.Run()
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

func newGlossariesList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossariesList := &GlossariesList{Config: *cfg}
	if cfg.Page != nil {
		actionGlossariesList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionGlossariesList.PerPage = *cfg.PerPage
	}

	return actionGlossariesList.Run()
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

func newGlossaryCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryCreate := &GlossaryCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryCreate.Config.Defaults["glossary/create"]
	if defaultsPresent {
		if err := actionGlossaryCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryCreate.Run()
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

func newGlossaryDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryDelete := &GlossaryDelete{Config: *cfg}

	return actionGlossaryDelete.Run()
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

func newGlossaryShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryShow := &GlossaryShow{Config: *cfg}

	return actionGlossaryShow.Run()
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

func newGlossaryUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryUpdate := &GlossaryUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryUpdate.Config.Defaults["glossary/update"]
	if defaultsPresent {
		if err := actionGlossaryUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryUpdate.Run()
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

func newGlossaryTermCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermCreate := &GlossaryTermCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermCreate.Config.Defaults["glossary_term/create"]
	if defaultsPresent {
		if err := actionGlossaryTermCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryTermCreate.Run()
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

func newGlossaryTermDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermDelete := &GlossaryTermDelete{Config: *cfg}

	return actionGlossaryTermDelete.Run()
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

func newGlossaryTermShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermShow := &GlossaryTermShow{Config: *cfg}

	return actionGlossaryTermShow.Run()
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

func newGlossaryTermUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermUpdate := &GlossaryTermUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermUpdate.Config.Defaults["glossary_term/update"]
	if defaultsPresent {
		if err := actionGlossaryTermUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryTermUpdate.Run()
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

func newGlossaryTermTranslationCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermTranslationCreate := &GlossaryTermTranslationCreate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermTranslationCreate.Config.Defaults["glossary_term_translation/create"]
	if defaultsPresent {
		if err := actionGlossaryTermTranslationCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryTermTranslationCreate.Run()
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

func newGlossaryTermTranslationDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermTranslationDelete := &GlossaryTermTranslationDelete{Config: *cfg}

	return actionGlossaryTermTranslationDelete.Run()
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

func newGlossaryTermTranslationUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermTranslationUpdate := &GlossaryTermTranslationUpdate{Config: *cfg}

	val, defaultsPresent := actionGlossaryTermTranslationUpdate.Config.Defaults["glossary_term_translation/update"]
	if defaultsPresent {
		if err := actionGlossaryTermTranslationUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionGlossaryTermTranslationUpdate.Run()
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

func newGlossaryTermsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionGlossaryTermsList := &GlossaryTermsList{Config: *cfg}
	if cfg.Page != nil {
		actionGlossaryTermsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionGlossaryTermsList.PerPage = *cfg.PerPage
	}

	return actionGlossaryTermsList.Run()
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

func newInvitationCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationCreate := &InvitationCreate{Config: *cfg}

	val, defaultsPresent := actionInvitationCreate.Config.Defaults["invitation/create"]
	if defaultsPresent {
		if err := actionInvitationCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionInvitationCreate.Run()
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

func newInvitationDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationDelete := &InvitationDelete{Config: *cfg}

	return actionInvitationDelete.Run()
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

func newInvitationResend(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationResend := &InvitationResend{Config: *cfg}

	return actionInvitationResend.Run()
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

func newInvitationShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationShow := &InvitationShow{Config: *cfg}

	return actionInvitationShow.Run()
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

func newInvitationUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationUpdate := &InvitationUpdate{Config: *cfg}

	val, defaultsPresent := actionInvitationUpdate.Config.Defaults["invitation/update"]
	if defaultsPresent {
		if err := actionInvitationUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionInvitationUpdate.Run()
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

func newInvitationsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionInvitationsList := &InvitationsList{Config: *cfg}
	if cfg.Page != nil {
		actionInvitationsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionInvitationsList.PerPage = *cfg.PerPage
	}

	return actionInvitationsList.Run()
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

func newJobComplete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobComplete := &JobComplete{Config: *cfg}
	actionJobComplete.ProjectID = cfg.DefaultProjectID

	return actionJobComplete.Run()
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

func newJobCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobCreate := &JobCreate{Config: *cfg}
	actionJobCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobCreate.Config.Defaults["job/create"]
	if defaultsPresent {
		if err := actionJobCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobCreate.Run()
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

func newJobDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobDelete := &JobDelete{Config: *cfg}
	actionJobDelete.ProjectID = cfg.DefaultProjectID

	return actionJobDelete.Run()
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

func newJobKeysCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobKeysCreate := &JobKeysCreate{Config: *cfg}
	actionJobKeysCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobKeysCreate.Config.Defaults["job/keys/create"]
	if defaultsPresent {
		if err := actionJobKeysCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobKeysCreate.Run()
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

func newJobKeysDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobKeysDelete := &JobKeysDelete{Config: *cfg}
	actionJobKeysDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobKeysDelete.Config.Defaults["job/keys/delete"]
	if defaultsPresent {
		if err := actionJobKeysDelete.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobKeysDelete.Run()
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

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobReopen(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobReopen := &JobReopen{Config: *cfg}
	actionJobReopen.ProjectID = cfg.DefaultProjectID

	return actionJobReopen.Run()
}

func (cmd *JobReopen) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobReopen(cmd.ProjectID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobShow := &JobShow{Config: *cfg}
	actionJobShow.ProjectID = cfg.DefaultProjectID

	return actionJobShow.Run()
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

func newJobStart(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobStart := &JobStart{Config: *cfg}
	actionJobStart.ProjectID = cfg.DefaultProjectID

	return actionJobStart.Run()
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

func newJobUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobUpdate := &JobUpdate{Config: *cfg}
	actionJobUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobUpdate.Config.Defaults["job/update"]
	if defaultsPresent {
		if err := actionJobUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobUpdate.Run()
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
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleComplete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocaleComplete := &JobLocaleComplete{Config: *cfg}
	actionJobLocaleComplete.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleComplete.Run()
}

func (cmd *JobLocaleComplete) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleComplete(cmd.ProjectID, cmd.JobID, cmd.ID)

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

func newJobLocaleDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocaleDelete := &JobLocaleDelete{Config: *cfg}
	actionJobLocaleDelete.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleDelete.Run()
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

type JobLocaleReopen struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleReopen(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocaleReopen := &JobLocaleReopen{Config: *cfg}
	actionJobLocaleReopen.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleReopen.Run()
}

func (cmd *JobLocaleReopen) Run() error {

	client, err := newClient(cmd.Config.Credentials, cmd.Config.Debug)
	if err != nil {
		return err
	}

	res, err := client.JobLocaleReopen(cmd.ProjectID, cmd.JobID, cmd.ID)

	if err != nil {
		return err
	}

	return json.NewEncoder(os.Stdout).Encode(&res)
}

type JobLocaleShow struct {
	phraseapp.Config

	ProjectID string `cli:"arg required"`
	JobID     string `cli:"arg required"`
	ID        string `cli:"arg required"`
}

func newJobLocaleShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocaleShow := &JobLocaleShow{Config: *cfg}
	actionJobLocaleShow.ProjectID = cfg.DefaultProjectID

	return actionJobLocaleShow.Run()
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

func newJobLocaleUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocaleUpdate := &JobLocaleUpdate{Config: *cfg}
	actionJobLocaleUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocaleUpdate.Config.Defaults["job_locale/update"]
	if defaultsPresent {
		if err := actionJobLocaleUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobLocaleUpdate.Run()
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

func newJobLocalesCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocalesCreate := &JobLocalesCreate{Config: *cfg}
	actionJobLocalesCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionJobLocalesCreate.Config.Defaults["job_locales/create"]
	if defaultsPresent {
		if err := actionJobLocalesCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionJobLocalesCreate.Run()
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

func newJobLocalesList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionJobLocalesList := &JobLocalesList{Config: *cfg}
	actionJobLocalesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionJobLocalesList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionJobLocalesList.PerPage = *cfg.PerPage
	}

	return actionJobLocalesList.Run()
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

func newJobsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionJobsList.Run()
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

func newKeyCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeyCreate := &KeyCreate{Config: *cfg}
	actionKeyCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyCreate.Config.Defaults["key/create"]
	if defaultsPresent {
		if err := actionKeyCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionKeyCreate.Run()
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

func newKeyDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeyDelete := &KeyDelete{Config: *cfg}
	actionKeyDelete.ProjectID = cfg.DefaultProjectID

	return actionKeyDelete.Run()
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

func newKeyShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeyShow := &KeyShow{Config: *cfg}
	actionKeyShow.ProjectID = cfg.DefaultProjectID

	return actionKeyShow.Run()
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

func newKeyUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeyUpdate := &KeyUpdate{Config: *cfg}
	actionKeyUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeyUpdate.Config.Defaults["key/update"]
	if defaultsPresent {
		if err := actionKeyUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionKeyUpdate.Run()
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

func newKeysDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeysDelete := &KeysDelete{Config: *cfg}
	actionKeysDelete.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysDelete.Config.Defaults["keys/delete"]
	if defaultsPresent {
		if err := actionKeysDelete.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionKeysDelete.Run()
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

func newKeysList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionKeysList.Run()
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

func newKeysSearch(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionKeysSearch.Run()
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

func newKeysTag(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeysTag := &KeysTag{Config: *cfg}
	actionKeysTag.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysTag.Config.Defaults["keys/tag"]
	if defaultsPresent {
		if err := actionKeysTag.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionKeysTag.Run()
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

func newKeysUntag(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionKeysUntag := &KeysUntag{Config: *cfg}
	actionKeysUntag.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionKeysUntag.Config.Defaults["keys/untag"]
	if defaultsPresent {
		if err := actionKeysUntag.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionKeysUntag.Run()
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

func newLocaleCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionLocaleCreate := &LocaleCreate{Config: *cfg}
	actionLocaleCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleCreate.Config.Defaults["locale/create"]
	if defaultsPresent {
		if err := actionLocaleCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionLocaleCreate.Run()
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

func newLocaleDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionLocaleDelete := &LocaleDelete{Config: *cfg}
	actionLocaleDelete.ProjectID = cfg.DefaultProjectID

	return actionLocaleDelete.Run()
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

func newLocaleDownload(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionLocaleDownload := &LocaleDownload{Config: *cfg}
	actionLocaleDownload.ProjectID = cfg.DefaultProjectID
	if cfg.DefaultFileFormat != "" {
		actionLocaleDownload.FileFormat = &cfg.DefaultFileFormat
	}

	val, defaultsPresent := actionLocaleDownload.Config.Defaults["locale/download"]
	if defaultsPresent {
		if err := actionLocaleDownload.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionLocaleDownload.Run()
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

func newLocaleShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionLocaleShow := &LocaleShow{Config: *cfg}
	actionLocaleShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleShow.Config.Defaults["locale/show"]
	if defaultsPresent {
		if err := actionLocaleShow.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionLocaleShow.Run()
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

func newLocaleUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionLocaleUpdate := &LocaleUpdate{Config: *cfg}
	actionLocaleUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionLocaleUpdate.Config.Defaults["locale/update"]
	if defaultsPresent {
		if err := actionLocaleUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionLocaleUpdate.Run()
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

func newLocalesList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionLocalesList.Run()
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

func newMemberDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionMemberDelete := &MemberDelete{Config: *cfg}

	return actionMemberDelete.Run()
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

func newMemberShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionMemberShow := &MemberShow{Config: *cfg}

	return actionMemberShow.Run()
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

func newMemberUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionMemberUpdate := &MemberUpdate{Config: *cfg}

	val, defaultsPresent := actionMemberUpdate.Config.Defaults["member/update"]
	if defaultsPresent {
		if err := actionMemberUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionMemberUpdate.Run()
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

func newMembersList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionMembersList := &MembersList{Config: *cfg}
	if cfg.Page != nil {
		actionMembersList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionMembersList.PerPage = *cfg.PerPage
	}

	return actionMembersList.Run()
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

func newOrderConfirm(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionOrderConfirm := &OrderConfirm{Config: *cfg}
	actionOrderConfirm.ProjectID = cfg.DefaultProjectID

	return actionOrderConfirm.Run()
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

func newOrderCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionOrderCreate := &OrderCreate{Config: *cfg}
	actionOrderCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionOrderCreate.Config.Defaults["order/create"]
	if defaultsPresent {
		if err := actionOrderCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionOrderCreate.Run()
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

func newOrderDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionOrderDelete := &OrderDelete{Config: *cfg}
	actionOrderDelete.ProjectID = cfg.DefaultProjectID

	return actionOrderDelete.Run()
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

func newOrderShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionOrderShow := &OrderShow{Config: *cfg}
	actionOrderShow.ProjectID = cfg.DefaultProjectID

	return actionOrderShow.Run()
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

func newOrdersList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionOrdersList := &OrdersList{Config: *cfg}
	actionOrdersList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionOrdersList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionOrdersList.PerPage = *cfg.PerPage
	}

	return actionOrdersList.Run()
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

func newProjectCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionProjectCreate := &ProjectCreate{Config: *cfg}

	val, defaultsPresent := actionProjectCreate.Config.Defaults["project/create"]
	if defaultsPresent {
		if err := actionProjectCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionProjectCreate.Run()
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

func newProjectDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionProjectDelete := &ProjectDelete{Config: *cfg}

	return actionProjectDelete.Run()
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

func newProjectShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionProjectShow := &ProjectShow{Config: *cfg}

	return actionProjectShow.Run()
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

func newProjectUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionProjectUpdate := &ProjectUpdate{Config: *cfg}

	val, defaultsPresent := actionProjectUpdate.Config.Defaults["project/update"]
	if defaultsPresent {
		if err := actionProjectUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionProjectUpdate.Run()
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

func newProjectsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionProjectsList := &ProjectsList{Config: *cfg}
	if cfg.Page != nil {
		actionProjectsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionProjectsList.PerPage = *cfg.PerPage
	}

	return actionProjectsList.Run()
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

func newShowUser(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionShowUser := &ShowUser{Config: *cfg}

	return actionShowUser.Run()
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

func newStyleguideCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionStyleguideCreate := &StyleguideCreate{Config: *cfg}
	actionStyleguideCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionStyleguideCreate.Config.Defaults["styleguide/create"]
	if defaultsPresent {
		if err := actionStyleguideCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionStyleguideCreate.Run()
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

func newStyleguideDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionStyleguideDelete := &StyleguideDelete{Config: *cfg}
	actionStyleguideDelete.ProjectID = cfg.DefaultProjectID

	return actionStyleguideDelete.Run()
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

func newStyleguideShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionStyleguideShow := &StyleguideShow{Config: *cfg}
	actionStyleguideShow.ProjectID = cfg.DefaultProjectID

	return actionStyleguideShow.Run()
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

func newStyleguideUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionStyleguideUpdate := &StyleguideUpdate{Config: *cfg}
	actionStyleguideUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionStyleguideUpdate.Config.Defaults["styleguide/update"]
	if defaultsPresent {
		if err := actionStyleguideUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionStyleguideUpdate.Run()
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

func newStyleguidesList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionStyleguidesList := &StyleguidesList{Config: *cfg}
	actionStyleguidesList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionStyleguidesList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionStyleguidesList.PerPage = *cfg.PerPage
	}

	return actionStyleguidesList.Run()
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

func newTagCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTagCreate := &TagCreate{Config: *cfg}
	actionTagCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTagCreate.Config.Defaults["tag/create"]
	if defaultsPresent {
		if err := actionTagCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTagCreate.Run()
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

func newTagDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTagDelete := &TagDelete{Config: *cfg}
	actionTagDelete.ProjectID = cfg.DefaultProjectID

	return actionTagDelete.Run()
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

func newTagShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTagShow := &TagShow{Config: *cfg}
	actionTagShow.ProjectID = cfg.DefaultProjectID

	return actionTagShow.Run()
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

func newTagsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTagsList := &TagsList{Config: *cfg}
	actionTagsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionTagsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionTagsList.PerPage = *cfg.PerPage
	}

	return actionTagsList.Run()
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

func newTranslationCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationCreate := &TranslationCreate{Config: *cfg}
	actionTranslationCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationCreate.Config.Defaults["translation/create"]
	if defaultsPresent {
		if err := actionTranslationCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationCreate.Run()
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

func newTranslationShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationShow := &TranslationShow{Config: *cfg}
	actionTranslationShow.ProjectID = cfg.DefaultProjectID

	return actionTranslationShow.Run()
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

func newTranslationUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationUpdate := &TranslationUpdate{Config: *cfg}
	actionTranslationUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationUpdate.Config.Defaults["translation/update"]
	if defaultsPresent {
		if err := actionTranslationUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationUpdate.Run()
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

func newTranslationsByKey(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionTranslationsByKey.Run()
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

func newTranslationsByLocale(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionTranslationsByLocale.Run()
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

func newTranslationsExclude(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationsExclude := &TranslationsExclude{Config: *cfg}
	actionTranslationsExclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsExclude.Config.Defaults["translations/exclude"]
	if defaultsPresent {
		if err := actionTranslationsExclude.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationsExclude.Run()
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

func newTranslationsInclude(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationsInclude := &TranslationsInclude{Config: *cfg}
	actionTranslationsInclude.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsInclude.Config.Defaults["translations/include"]
	if defaultsPresent {
		if err := actionTranslationsInclude.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationsInclude.Run()
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

func newTranslationsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionTranslationsList.Run()
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

func newTranslationsSearch(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
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
			return err
		}
	}
	return actionTranslationsSearch.Run()
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

func newTranslationsUnverify(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationsUnverify := &TranslationsUnverify{Config: *cfg}
	actionTranslationsUnverify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsUnverify.Config.Defaults["translations/unverify"]
	if defaultsPresent {
		if err := actionTranslationsUnverify.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationsUnverify.Run()
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

func newTranslationsVerify(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionTranslationsVerify := &TranslationsVerify{Config: *cfg}
	actionTranslationsVerify.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionTranslationsVerify.Config.Defaults["translations/verify"]
	if defaultsPresent {
		if err := actionTranslationsVerify.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionTranslationsVerify.Run()
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

func newUploadCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionUploadCreate := &UploadCreate{Config: *cfg}
	actionUploadCreate.ProjectID = cfg.DefaultProjectID
	if cfg.DefaultFileFormat != "" {
		actionUploadCreate.FileFormat = &cfg.DefaultFileFormat
	}

	val, defaultsPresent := actionUploadCreate.Config.Defaults["upload/create"]
	if defaultsPresent {
		if err := actionUploadCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionUploadCreate.Run()
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

func newUploadShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionUploadShow := &UploadShow{Config: *cfg}
	actionUploadShow.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionUploadShow.Config.Defaults["upload/show"]
	if defaultsPresent {
		if err := actionUploadShow.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionUploadShow.Run()
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

func newUploadsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionUploadsList := &UploadsList{Config: *cfg}
	actionUploadsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionUploadsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionUploadsList.PerPage = *cfg.PerPage
	}

	return actionUploadsList.Run()
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

func newVersionShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionVersionShow := &VersionShow{Config: *cfg}
	actionVersionShow.ProjectID = cfg.DefaultProjectID

	return actionVersionShow.Run()
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

func newVersionsList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionVersionsList := &VersionsList{Config: *cfg}
	actionVersionsList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionVersionsList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionVersionsList.PerPage = *cfg.PerPage
	}

	return actionVersionsList.Run()
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

func newWebhookCreate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhookCreate := &WebhookCreate{Config: *cfg}
	actionWebhookCreate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionWebhookCreate.Config.Defaults["webhook/create"]
	if defaultsPresent {
		if err := actionWebhookCreate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionWebhookCreate.Run()
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

func newWebhookDelete(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhookDelete := &WebhookDelete{Config: *cfg}
	actionWebhookDelete.ProjectID = cfg.DefaultProjectID

	return actionWebhookDelete.Run()
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

func newWebhookShow(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhookShow := &WebhookShow{Config: *cfg}
	actionWebhookShow.ProjectID = cfg.DefaultProjectID

	return actionWebhookShow.Run()
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

func newWebhookTest(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhookTest := &WebhookTest{Config: *cfg}
	actionWebhookTest.ProjectID = cfg.DefaultProjectID

	return actionWebhookTest.Run()
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

func newWebhookUpdate(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhookUpdate := &WebhookUpdate{Config: *cfg}
	actionWebhookUpdate.ProjectID = cfg.DefaultProjectID

	val, defaultsPresent := actionWebhookUpdate.Config.Defaults["webhook/update"]
	if defaultsPresent {
		if err := actionWebhookUpdate.ApplyValuesFromMap(val); err != nil {
			return err
		}
	}
	return actionWebhookUpdate.Run()
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

func newWebhooksList(cli *cli.Context) error {
	cfg, _ := phraseapp.ReadConfig()
	actionWebhooksList := &WebhooksList{Config: *cfg}
	actionWebhooksList.ProjectID = cfg.DefaultProjectID
	if cfg.Page != nil {
		actionWebhooksList.Page = *cfg.Page
	}

	if cfg.PerPage != nil {
		actionWebhooksList.PerPage = *cfg.PerPage
	}

	return actionWebhooksList.Run()
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
