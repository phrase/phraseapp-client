package phraseapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	RevisionDocs      = "b33e5c3d0e590c6aed24f39a2c98731c6b2116d9"
	RevisionGenerator = "HEAD/2020-04-02T172629/soenke"
)

type Account struct {
	Company   string     `json:"company"`
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type AccountDetails struct {
	Account

	Slug string `json:"slug"`
}

type AffectedCount struct {
	RecordsAffected int64 `json:"records_affected"`
}

type AffectedResources struct {
	RecordsAffected int64 `json:"records_affected"`
}

type Authorization struct {
	CreatedAt      *time.Time `json:"created_at"`
	ExpiresAt      *time.Time `json:"expires_at"`
	HashedToken    string     `json:"hashed_token"`
	ID             string     `json:"id"`
	Note           string     `json:"note"`
	Scopes         []string   `json:"scopes"`
	TokenLastEight string     `json:"token_last_eight"`
	UpdatedAt      *time.Time `json:"updated_at"`
}

type AuthorizationWithToken struct {
	Authorization

	Token string `json:"token"`
}

type BitbucketSync struct {
	ID                        string          `json:"id"`
	LastExportToBitbucketAt   *time.Time      `json:"last_export_to_bitbucket_at"`
	LastImportFromBitbucketAt *time.Time      `json:"last_import_from_bitbucket_at"`
	PhraseappProjects         []*ProjectShort `json:"phraseapp_projects"`
	RepositoryName            string          `json:"repository_name"`
	ValidPhraseappYaml        bool            `json:"valid_phraseapp_yaml"`
}

type BitbucketSyncExportResponse struct {
	StatusPath string `json:"status_path"`
}

type BlacklistedKey struct {
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Branch struct {
	CreatedAt *time.Time   `json:"created_at"`
	CreatedBy *UserPreview `json:"created_by"`
	MergedAt  *time.Time   `json:"merged_at"`
	MergedBy  *UserPreview `json:"merged_by"`
	Name      string       `json:"name"`
	State     string       `json:"state"`
	UpdatedAt *time.Time   `json:"updated_at"`
}

type Comment struct {
	CreatedAt *time.Time   `json:"created_at"`
	ID        string       `json:"id"`
	Message   string       `json:"message"`
	UpdatedAt *time.Time   `json:"updated_at"`
	User      *UserPreview `json:"user"`
}

type Distribution struct {
	CreatedAt *time.Time        `json:"created_at"`
	DeletedAt *time.Time        `json:"deleted_at"`
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Platforms []string          `json:"platforms"`
	Project   *ProjectShort     `json:"project"`
	Releases  []*ReleasePreview `json:"releases"`
}

type DistributionPreview struct {
	CreatedAt    *time.Time    `json:"created_at"`
	DeletedAt    *time.Time    `json:"deleted_at"`
	ID           string        `json:"id"`
	Name         string        `json:"name"`
	Platforms    []string      `json:"platforms"`
	Project      *ProjectShort `json:"project"`
	ReleaseCount int64         `json:"release_count"`
}

type Format struct {
	ApiName                   string `json:"api_name"`
	DefaultEncoding           string `json:"default_encoding"`
	DefaultFile               string `json:"default_file"`
	Description               string `json:"description"`
	Exportable                bool   `json:"exportable"`
	Extension                 string `json:"extension"`
	Importable                bool   `json:"importable"`
	IncludesLocaleInformation bool   `json:"includes_locale_information"`
	Name                      string `json:"name"`
	RendersDefaultLocale      bool   `json:"renders_default_locale"`
}

type Glossary struct {
	CreatedAt *time.Time      `json:"created_at"`
	ID        string          `json:"id"`
	Name      string          `json:"name"`
	Projects  []*ProjectShort `json:"projects"`
	UpdatedAt *time.Time      `json:"updated_at"`
}

type GlossaryTerm struct {
	CaseSensitive bool                       `json:"case_sensitive"`
	CreatedAt     *time.Time                 `json:"created_at"`
	Description   string                     `json:"description"`
	ID            string                     `json:"id"`
	Term          string                     `json:"term"`
	Translatable  bool                       `json:"translatable"`
	Translations  []*GlossaryTermTranslation `json:"translations"`
	UpdatedAt     *time.Time                 `json:"updated_at"`
}

type GlossaryTermTranslation struct {
	Content    string     `json:"content"`
	CreatedAt  *time.Time `json:"created_at"`
	ID         string     `json:"id"`
	LocaleCode string     `json:"locale_code"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type Invitation struct {
	AcceptedAt  *time.Time        `json:"accepted_at"`
	CreatedAt   *time.Time        `json:"created_at"`
	Email       string            `json:"email"`
	ID          string            `json:"id"`
	Locales     []*LocalePreview  `json:"locales"`
	Permissions map[string]string `json:"permissions"`
	Projects    []*ProjectShort   `json:"projects"`
	Role        string            `json:"role"`
	State       string            `json:"state"`
	UpdatedAt   *time.Time        `json:"updated_at"`
}

type Job struct {
	Briefing  string     `json:"briefing"`
	CreatedAt *time.Time `json:"created_at"`
	DueDate   *time.Time `json:"due_date"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	State     string     `json:"state"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type JobDetails struct {
	Job

	JobTagName string           `json:"job_tag_name"`
	Keys       []*KeyPreview    `json:"keys"`
	Locales    []*LocalePreview `json:"locales"`
	Owner      *UserPreview     `json:"owner"`
}

type JobLocale struct {
	ID     string         `json:"id"`
	Job    *JobPreview    `json:"job"`
	Locale *LocalePreview `json:"locale"`
	Users  []*UserPreview `json:"users"`
}

type JobPreview struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	State string `json:"state"`
}

type KeyPreview struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Plural bool   `json:"plural"`
}

type Locale struct {
	Code         string         `json:"code"`
	CreatedAt    *time.Time     `json:"created_at"`
	Default      bool           `json:"default"`
	ID           string         `json:"id"`
	Main         bool           `json:"main"`
	Name         string         `json:"name"`
	PluralForms  []string       `json:"plural_forms"`
	Rtl          bool           `json:"rtl"`
	SourceLocale *LocalePreview `json:"source_locale"`
	UpdatedAt    *time.Time     `json:"updated_at"`
}

type LocaleDetails struct {
	Locale

	Statistics *LocaleStatistics `json:"statistics"`
}

type LocalePreview struct {
	Code string `json:"code"`
	ID   string `json:"id"`
	Name string `json:"name"`
}

type LocaleStatistics struct {
	KeysTotalCount              int64 `json:"keys_total_count"`
	KeysUntranslatedCount       int64 `json:"keys_untranslated_count"`
	MissingWordsCount           int64 `json:"missing_words_count"`
	TranslationsCompletedCount  int64 `json:"translations_completed_count"`
	TranslationsUnverifiedCount int64 `json:"translations_unverified_count"`
	UnverifiedWordsCount        int64 `json:"unverified_words_count"`
	WordsTotalCount             int64 `json:"words_total_count"`
}

type Member struct {
	Email    string            `json:"email"`
	ID       string            `json:"id"`
	Projects []*ProjectLocales `json:"projects"`
	Role     string            `json:"role"`
	Username string            `json:"username"`
}

type Project struct {
	Account         *Account   `json:"account"`
	CreatedAt       *time.Time `json:"created_at"`
	ID              string     `json:"id"`
	MainFormat      string     `json:"main_format"`
	Name            string     `json:"name"`
	ProjectImageUrl string     `json:"project_image_url"`
	UpdatedAt       *time.Time `json:"updated_at"`
}

type ProjectDetails struct {
	Project

	SharesTranslationMemory bool   `json:"shares_translation_memory"`
	Slug                    string `json:"slug"`
}

type ProjectLocales struct {
	ProjectShort

	Locales []*LocalePreview `json:"locales"`
}

type ProjectShort struct {
	CreatedAt  *time.Time `json:"created_at"`
	ID         string     `json:"id"`
	MainFormat string     `json:"main_format"`
	Name       string     `json:"name"`
	UpdatedAt  *time.Time `json:"updated_at"`
}

type Release struct {
	AppMaxVersion string           `json:"app_max_version"`
	AppMinVersion string           `json:"app_min_version"`
	CreatedAt     *time.Time       `json:"created_at"`
	Description   string           `json:"description"`
	Environments  []string         `json:"environments"`
	ID            string           `json:"id"`
	Locales       []*LocalePreview `json:"locales"`
	Platforms     []string         `json:"platforms"`
	Project       *ProjectShort    `json:"project"`
	UpdatedAt     *time.Time       `json:"updated_at"`
	Version       int64            `json:"version"`
}

type ReleasePreview struct {
	AppMaxVersion string        `json:"app_max_version"`
	AppMinVersion string        `json:"app_min_version"`
	CreatedAt     *time.Time    `json:"created_at"`
	Description   string        `json:"description"`
	Environments  []string      `json:"environments"`
	ID            string        `json:"id"`
	LocaleCodes   []string      `json:"locale_codes"`
	Platforms     []string      `json:"platforms"`
	Project       *ProjectShort `json:"project"`
	UpdatedAt     *time.Time    `json:"updated_at"`
	Version       int64         `json:"version"`
}

type Screenshot struct {
	CreatedAt     *time.Time `json:"created_at"`
	Description   string     `json:"description"`
	ID            string     `json:"id"`
	MarkersCount  int64      `json:"markers_count"`
	Name          string     `json:"name"`
	ScreenshotUrl string     `json:"screenshot_url"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type ScreenshotMarker struct {
	CreatedAt        *time.Time      `json:"created_at"`
	ID               string          `json:"id"`
	Presentation     string          `json:"presentation"`
	PresentationType string          `json:"presentation_type"`
	TranslationKey   *TranslationKey `json:"translation_key"`
	UpdatedAt        *time.Time      `json:"updated_at"`
}

type Space struct {
	CreatedAt     *time.Time `json:"created_at"`
	ID            string     `json:"id"`
	Name          string     `json:"name"`
	Projects      []*Project `json:"projects"`
	ProjectsCount int64      `json:"projects_count"`
	UpdatedAt     *time.Time `json:"updated_at"`
}

type StatisticsListItem struct {
	Locale     *LocalePreview `json:"locale"`
	Statistics StatisticsType `json:"statistics"`
}

type StatisticsType struct {
	KeysTotalCount              int64 `json:"keys_total_count"`
	KeysUntranslatedCount       int64 `json:"keys_untranslated_count"`
	TranslationsCompletedCount  int64 `json:"translations_completed_count"`
	TranslationsUnverifiedCount int64 `json:"translations_unverified_count"`
}

type Styleguide struct {
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Title     string     `json:"title"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type StyleguideDetails struct {
	Styleguide

	Audience           string `json:"audience"`
	Business           string `json:"business"`
	CompanyBranding    string `json:"company_branding"`
	Formatting         string `json:"formatting"`
	GlossaryTerms      string `json:"glossary_terms"`
	GrammarConsistency string `json:"grammar_consistency"`
	GrammaticalPerson  string `json:"grammatical_person"`
	LiteralTranslation string `json:"literal_translation"`
	OverallTone        string `json:"overall_tone"`
	PublicUrl          string `json:"public_url"`
	Samples            string `json:"samples"`
	TargetAudience     string `json:"target_audience"`
	VocabularyType     string `json:"vocabulary_type"`
}

type StyleguidePreview struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

type SummaryType struct {
	LocalesCreated             int64 `json:"locales_created"`
	TagsCreated                int64 `json:"tags_created"`
	TranslationKeysCreated     int64 `json:"translation_keys_created"`
	TranslationKeysIgnored     int64 `json:"translation_keys_ignored"`
	TranslationKeysUnmentioned int64 `json:"translation_keys_unmentioned"`
	TranslationKeysUpdated     int64 `json:"translation_keys_updated"`
	TranslationsCreated        int64 `json:"translations_created"`
	TranslationsUpdated        int64 `json:"translations_updated"`
}

type Tag struct {
	CreatedAt *time.Time `json:"created_at"`
	KeysCount int64      `json:"keys_count"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type TagWithStats struct {
	Tag

	Statistics []*StatisticsListItem `json:"statistics"`
}

type Translation struct {
	Content      string         `json:"content"`
	CreatedAt    *time.Time     `json:"created_at"`
	Excluded     bool           `json:"excluded"`
	ID           string         `json:"id"`
	Key          *KeyPreview    `json:"key"`
	Locale       *LocalePreview `json:"locale"`
	Placeholders []string       `json:"placeholders"`
	PluralSuffix string         `json:"plural_suffix"`
	State        string         `json:"state"`
	Unverified   bool           `json:"unverified"`
	UpdatedAt    *time.Time     `json:"updated_at"`
}

type TranslationDetails struct {
	Translation

	User      *UserPreview `json:"user"`
	WordCount int64        `json:"word_count"`
}

type TranslationKey struct {
	CreatedAt   *time.Time `json:"created_at"`
	DataType    string     `json:"data_type"`
	Description string     `json:"description"`
	ID          string     `json:"id"`
	Name        string     `json:"name"`
	NameHash    string     `json:"name_hash"`
	Plural      bool       `json:"plural"`
	Tags        []string   `json:"tags"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TranslationKeyDetails struct {
	TranslationKey

	CommentsCount        int64        `json:"comments_count"`
	Creator              *UserPreview `json:"creator"`
	FormatValueType      string       `json:"format_value_type"`
	MaxCharactersAllowed int64        `json:"max_characters_allowed"`
	NamePlural           string       `json:"name_plural"`
	OriginalFile         string       `json:"original_file"`
	ScreenshotUrl        string       `json:"screenshot_url"`
	Unformatted          bool         `json:"unformatted"`
	XmlSpacePreserve     bool         `json:"xml_space_preserve"`
}

type TranslationOrder struct {
	AmountInCents                    int64              `json:"amount_in_cents"`
	CreatedAt                        *time.Time         `json:"created_at"`
	Currency                         string             `json:"currency"`
	ID                               string             `json:"id"`
	Lsp                              string             `json:"lsp"`
	Message                          string             `json:"message"`
	Priority                         bool               `json:"priority"`
	ProgressPercent                  int64              `json:"progress_percent"`
	Quality                          bool               `json:"quality"`
	SourceLocale                     *LocalePreview     `json:"source_locale"`
	State                            string             `json:"state"`
	Styleguide                       *StyleguidePreview `json:"styleguide"`
	Tag                              string             `json:"tag"`
	TargetLocales                    []*LocalePreview   `json:"target_locales"`
	TranslationType                  string             `json:"translation_type"`
	UnverifyTranslationsUponDelivery bool               `json:"unverify_translations_upon_delivery"`
	UpdatedAt                        *time.Time         `json:"updated_at"`
}

type TranslationVersion struct {
	ChangedAt    *time.Time     `json:"changed_at"`
	Content      string         `json:"content"`
	CreatedAt    *time.Time     `json:"created_at"`
	ID           string         `json:"id"`
	Key          *KeyPreview    `json:"key"`
	Locale       *LocalePreview `json:"locale"`
	PluralSuffix string         `json:"plural_suffix"`
	UpdatedAt    *time.Time     `json:"updated_at"`
}

type TranslationVersionWithUser struct {
	TranslationVersion

	User *UserPreview `json:"user"`
}

type Upload struct {
	CreatedAt *time.Time  `json:"created_at"`
	Filename  string      `json:"filename"`
	Format    string      `json:"format"`
	ID        string      `json:"id"`
	State     string      `json:"state"`
	Summary   SummaryType `json:"summary"`
	Tag       string      `json:"tag"`
	UpdatedAt *time.Time  `json:"updated_at"`
}

type User struct {
	CreatedAt *time.Time `json:"created_at"`
	Email     string     `json:"email"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	Position  string     `json:"position"`
	UpdatedAt *time.Time `json:"updated_at"`
	Username  string     `json:"username"`
}

type UserPreview struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Username string `json:"username"`
}

type Webhook struct {
	Active      bool       `json:"active"`
	CallbackUrl string     `json:"callback_url"`
	CreatedAt   *time.Time `json:"created_at"`
	Description string     `json:"description"`
	Events      []string   `json:"events"`
	ID          string     `json:"id"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type AuthorizationParams struct {
	ExpiresAt **time.Time `json:"expires_at,omitempty"  cli:"opt --expires-at"`
	Note      *string     `json:"note,omitempty"  cli:"opt --note"`
	Scopes    []string    `json:"scopes,omitempty"  cli:"opt --scopes"`
}

func (params *AuthorizationParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "expires_at":
			val, ok := v.(*time.Time)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ExpiresAt = &val

		case "note":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Note = &val

		case "scopes":
			ok := false
			params.Scopes, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *AuthorizationParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Note != nil && *params.Note != "" {
		queryParams["note"] = *params.Note
	}

	return queryParams
}

type BitbucketSyncParams struct {
	AccountID *string `json:"account_id,omitempty"  cli:"opt --account-id"`
}

func (params *BitbucketSyncParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "account_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.AccountID = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *BitbucketSyncParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.AccountID != nil && *params.AccountID != "" {
		queryParams["account_id"] = *params.AccountID
	}

	return queryParams
}

type BlacklistedKeyParams struct {
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *BlacklistedKeyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *BlacklistedKeyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type BranchParams struct {
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *BranchParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *BranchParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type CommentParams struct {
	Branch  *string `json:"branch,omitempty"  cli:"opt --branch"`
	Message *string `json:"message,omitempty"  cli:"opt --message"`
}

func (params *CommentParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "message":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Message = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Message != nil && *params.Message != "" {
		queryParams["message"] = *params.Message
	}

	return queryParams
}

type DistributionsParams struct {
	FallbackToDefaultLocale     *bool             `json:"fallback_to_default_locale,omitempty"  cli:"opt --fallback-to-default-locale"`
	FallbackToNonRegionalLocale *bool             `json:"fallback_to_non_regional_locale,omitempty"  cli:"opt --fallback-to-non-regional-locale"`
	FormatOptions               map[string]string `json:"format_options,omitempty"  cli:"opt --format-options"`
	Name                        *string           `json:"name,omitempty"  cli:"opt --name"`
	Platforms                   []string          `json:"platforms,omitempty"  cli:"opt --platforms"`
	ProjectID                   *string           `json:"project_id,omitempty"  cli:"opt --project-id"`
	UseLastReviewedVersion      *bool             `json:"use_last_reviewed_version,omitempty"  cli:"opt --use-last-reviewed-version"`
}

func (params *DistributionsParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "fallback_to_default_locale":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FallbackToDefaultLocale = &val

		case "fallback_to_non_regional_locale":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FallbackToNonRegionalLocale = &val

		case "format_options":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.FormatOptions = val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "platforms":
			ok := false
			params.Platforms, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "project_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectID = &val

		case "use_last_reviewed_version":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UseLastReviewedVersion = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *DistributionsParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.FallbackToDefaultLocale != nil {
		queryParams["fallback_to_default_locale"] = strconv.FormatBool(*params.FallbackToDefaultLocale)
	}

	if params.FallbackToNonRegionalLocale != nil {
		queryParams["fallback_to_non_regional_locale"] = strconv.FormatBool(*params.FallbackToNonRegionalLocale)
	}

	for key, value := range convertMapToQueryParams("format_options", params.FormatOptions) {
		queryParams[key] = value
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	if params.ProjectID != nil && *params.ProjectID != "" {
		queryParams["project_id"] = *params.ProjectID
	}

	if params.UseLastReviewedVersion != nil {
		queryParams["use_last_reviewed_version"] = strconv.FormatBool(*params.UseLastReviewedVersion)
	}

	return queryParams
}

type GlossaryParams struct {
	Name       *string `json:"name,omitempty"  cli:"opt --name"`
	ProjectIDs *string `json:"project_ids,omitempty"  cli:"opt --project-ids"`
}

func (params *GlossaryParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "project_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectIDs = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *GlossaryParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	if params.ProjectIDs != nil && *params.ProjectIDs != "" {
		queryParams["project_ids"] = *params.ProjectIDs
	}

	return queryParams
}

type GlossaryTermTranslationParams struct {
	Content    *string `json:"content,omitempty"  cli:"opt --content"`
	LocaleCode *string `json:"locale_code,omitempty"  cli:"opt --locale-code"`
}

func (params *GlossaryTermTranslationParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "content":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Content = &val

		case "locale_code":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleCode = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *GlossaryTermTranslationParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Content != nil && *params.Content != "" {
		queryParams["content"] = *params.Content
	}

	if params.LocaleCode != nil && *params.LocaleCode != "" {
		queryParams["locale_code"] = *params.LocaleCode
	}

	return queryParams
}

type GlossaryTermParams struct {
	CaseSensitive *bool   `json:"case_sensitive,omitempty"  cli:"opt --case-sensitive"`
	Description   *string `json:"description,omitempty"  cli:"opt --description"`
	Term          *string `json:"term,omitempty"  cli:"opt --term"`
	Translatable  *bool   `json:"translatable,omitempty"  cli:"opt --translatable"`
}

func (params *GlossaryTermParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "case_sensitive":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.CaseSensitive = &val

		case "description":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Description = &val

		case "term":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Term = &val

		case "translatable":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Translatable = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *GlossaryTermParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.CaseSensitive != nil {
		queryParams["case_sensitive"] = strconv.FormatBool(*params.CaseSensitive)
	}

	if params.Description != nil && *params.Description != "" {
		queryParams["description"] = *params.Description
	}

	if params.Term != nil && *params.Term != "" {
		queryParams["term"] = *params.Term
	}

	if params.Translatable != nil {
		queryParams["translatable"] = strconv.FormatBool(*params.Translatable)
	}

	return queryParams
}

type JobLocaleParams struct {
	Branch   *string  `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string  `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	UserIDs  []string `json:"user_ids,omitempty"  cli:"opt --user-ids"`
}

func (params *JobLocaleParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "user_ids":
			ok := false
			params.UserIDs, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocaleParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.LocaleID != nil && *params.LocaleID != "" {
		queryParams["locale_id"] = *params.LocaleID
	}

	return queryParams
}

type JobParams struct {
	Branch            *string     `json:"branch,omitempty"  cli:"opt --branch"`
	Briefing          *string     `json:"briefing,omitempty"  cli:"opt --briefing"`
	DueDate           **time.Time `json:"due_date,omitempty"  cli:"opt --due-date"`
	Name              *string     `json:"name,omitempty"  cli:"opt --name"`
	Tags              []string    `json:"tags,omitempty"  cli:"opt --tags"`
	TranslationKeyIDs []string    `json:"translation_key_ids,omitempty"  cli:"opt --translation-key-ids"`
}

func (params *JobParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "briefing":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Briefing = &val

		case "due_date":
			val, ok := v.(*time.Time)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.DueDate = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "tags":
			ok := false
			params.Tags, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "translation_key_ids":
			ok := false
			params.TranslationKeyIDs, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Briefing != nil && *params.Briefing != "" {
		queryParams["briefing"] = *params.Briefing
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type TranslationKeyParams struct {
	Branch                *string `json:"branch,omitempty"  cli:"opt --branch"`
	DataType              *string `json:"data_type,omitempty"  cli:"opt --data-type"`
	Description           *string `json:"description,omitempty"  cli:"opt --description"`
	LocalizedFormatKey    *string `json:"localized_format_key,omitempty"  cli:"opt --localized-format-key"`
	LocalizedFormatString *string `json:"localized_format_string,omitempty"  cli:"opt --localized-format-string"`
	MaxCharactersAllowed  *int64  `json:"max_characters_allowed,omitempty"  cli:"opt --max-characters-allowed"`
	Name                  *string `json:"name,omitempty"  cli:"opt --name"`
	NamePlural            *string `json:"name_plural,omitempty"  cli:"opt --name-plural"`
	OriginalFile          *string `json:"original_file,omitempty"  cli:"opt --original-file"`
	Plural                *bool   `json:"plural,omitempty"  cli:"opt --plural"`
	RemoveScreenshot      *bool   `json:"remove_screenshot,omitempty"  cli:"opt --remove-screenshot"`
	Screenshot            *string `json:"screenshot,omitempty"  cli:"opt --screenshot"`
	Tags                  *string `json:"tags,omitempty"  cli:"opt --tags"`
	Unformatted           *bool   `json:"unformatted,omitempty"  cli:"opt --unformatted"`
	XmlSpacePreserve      *bool   `json:"xml_space_preserve,omitempty"  cli:"opt --xml-space-preserve"`
}

func (params *TranslationKeyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "data_type":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.DataType = &val

		case "description":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Description = &val

		case "localized_format_key":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocalizedFormatKey = &val

		case "localized_format_string":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocalizedFormatString = &val

		case "max_characters_allowed":
			val, ok := v.(int64)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.MaxCharactersAllowed = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "name_plural":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.NamePlural = &val

		case "original_file":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.OriginalFile = &val

		case "plural":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Plural = &val

		case "remove_screenshot":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.RemoveScreenshot = &val

		case "screenshot":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Screenshot = &val

		case "tags":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tags = &val

		case "unformatted":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Unformatted = &val

		case "xml_space_preserve":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.XmlSpacePreserve = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationKeyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.DataType != nil && *params.DataType != "" {
		queryParams["data_type"] = *params.DataType
	}

	if params.Description != nil && *params.Description != "" {
		queryParams["description"] = *params.Description
	}

	if params.LocalizedFormatKey != nil && *params.LocalizedFormatKey != "" {
		queryParams["localized_format_key"] = *params.LocalizedFormatKey
	}

	if params.LocalizedFormatString != nil && *params.LocalizedFormatString != "" {
		queryParams["localized_format_string"] = *params.LocalizedFormatString
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	if params.NamePlural != nil && *params.NamePlural != "" {
		queryParams["name_plural"] = *params.NamePlural
	}

	if params.OriginalFile != nil && *params.OriginalFile != "" {
		queryParams["original_file"] = *params.OriginalFile
	}

	if params.Plural != nil {
		queryParams["plural"] = strconv.FormatBool(*params.Plural)
	}

	if params.RemoveScreenshot != nil {
		queryParams["remove_screenshot"] = strconv.FormatBool(*params.RemoveScreenshot)
	}

	if params.Tags != nil && *params.Tags != "" {
		queryParams["tags"] = *params.Tags
	}

	if params.Unformatted != nil {
		queryParams["unformatted"] = strconv.FormatBool(*params.Unformatted)
	}

	if params.XmlSpacePreserve != nil {
		queryParams["xml_space_preserve"] = strconv.FormatBool(*params.XmlSpacePreserve)
	}

	return queryParams
}

type LocaleParams struct {
	Autotranslate               *bool   `json:"autotranslate,omitempty"  cli:"opt --autotranslate"`
	Branch                      *string `json:"branch,omitempty"  cli:"opt --branch"`
	Code                        *string `json:"code,omitempty"  cli:"opt --code"`
	Default                     *bool   `json:"default,omitempty"  cli:"opt --default"`
	Main                        *bool   `json:"main,omitempty"  cli:"opt --main"`
	Name                        *string `json:"name,omitempty"  cli:"opt --name"`
	Rtl                         *bool   `json:"rtl,omitempty"  cli:"opt --rtl"`
	SourceLocaleID              *string `json:"source_locale_id,omitempty"  cli:"opt --source-locale-id"`
	UnverifyNewTranslations     *bool   `json:"unverify_new_translations,omitempty"  cli:"opt --unverify-new-translations"`
	UnverifyUpdatedTranslations *bool   `json:"unverify_updated_translations,omitempty"  cli:"opt --unverify-updated-translations"`
}

func (params *LocaleParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "autotranslate":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Autotranslate = &val

		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "code":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Code = &val

		case "default":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Default = &val

		case "main":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Main = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "rtl":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Rtl = &val

		case "source_locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SourceLocaleID = &val

		case "unverify_new_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UnverifyNewTranslations = &val

		case "unverify_updated_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UnverifyUpdatedTranslations = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *LocaleParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Autotranslate != nil {
		queryParams["autotranslate"] = strconv.FormatBool(*params.Autotranslate)
	}

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Code != nil && *params.Code != "" {
		queryParams["code"] = *params.Code
	}

	if params.Default != nil {
		queryParams["default"] = strconv.FormatBool(*params.Default)
	}

	if params.Main != nil {
		queryParams["main"] = strconv.FormatBool(*params.Main)
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	if params.Rtl != nil {
		queryParams["rtl"] = strconv.FormatBool(*params.Rtl)
	}

	if params.UnverifyNewTranslations != nil {
		queryParams["unverify_new_translations"] = strconv.FormatBool(*params.UnverifyNewTranslations)
	}

	if params.UnverifyUpdatedTranslations != nil {
		queryParams["unverify_updated_translations"] = strconv.FormatBool(*params.UnverifyUpdatedTranslations)
	}

	return queryParams
}

type TranslationOrderParams struct {
	Branch                           *string  `json:"branch,omitempty"  cli:"opt --branch"`
	Category                         *string  `json:"category,omitempty"  cli:"opt --category"`
	IncludeUntranslatedKeys          *bool    `json:"include_untranslated_keys,omitempty"  cli:"opt --include-untranslated-keys"`
	IncludeUnverifiedTranslations    *bool    `json:"include_unverified_translations,omitempty"  cli:"opt --include-unverified-translations"`
	Lsp                              *string  `json:"lsp,omitempty"  cli:"opt --lsp"`
	Message                          *string  `json:"message,omitempty"  cli:"opt --message"`
	Priority                         *bool    `json:"priority,omitempty"  cli:"opt --priority"`
	Quality                          *bool    `json:"quality,omitempty"  cli:"opt --quality"`
	SourceLocaleID                   *string  `json:"source_locale_id,omitempty"  cli:"opt --source-locale-id"`
	StyleguideID                     *string  `json:"styleguide_id,omitempty"  cli:"opt --styleguide-id"`
	Tag                              *string  `json:"tag,omitempty"  cli:"opt --tag"`
	TargetLocaleIDs                  []string `json:"target_locale_ids,omitempty"  cli:"opt --target-locale-ids"`
	TranslationType                  *string  `json:"translation_type,omitempty"  cli:"opt --translation-type"`
	UnverifyTranslationsUponDelivery *bool    `json:"unverify_translations_upon_delivery,omitempty"  cli:"opt --unverify-translations-upon-delivery"`
}

func (params *TranslationOrderParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "category":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Category = &val

		case "include_untranslated_keys":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.IncludeUntranslatedKeys = &val

		case "include_unverified_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.IncludeUnverifiedTranslations = &val

		case "lsp":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Lsp = &val

		case "message":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Message = &val

		case "priority":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Priority = &val

		case "quality":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Quality = &val

		case "source_locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SourceLocaleID = &val

		case "styleguide_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.StyleguideID = &val

		case "tag":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tag = &val

		case "target_locale_ids":
			ok := false
			params.TargetLocaleIDs, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "translation_type":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.TranslationType = &val

		case "unverify_translations_upon_delivery":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UnverifyTranslationsUponDelivery = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationOrderParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Category != nil && *params.Category != "" {
		queryParams["category"] = *params.Category
	}

	if params.IncludeUntranslatedKeys != nil {
		queryParams["include_untranslated_keys"] = strconv.FormatBool(*params.IncludeUntranslatedKeys)
	}

	if params.IncludeUnverifiedTranslations != nil {
		queryParams["include_unverified_translations"] = strconv.FormatBool(*params.IncludeUnverifiedTranslations)
	}

	if params.Lsp != nil && *params.Lsp != "" {
		queryParams["lsp"] = *params.Lsp
	}

	if params.Message != nil && *params.Message != "" {
		queryParams["message"] = *params.Message
	}

	if params.Priority != nil {
		queryParams["priority"] = strconv.FormatBool(*params.Priority)
	}

	if params.Quality != nil {
		queryParams["quality"] = strconv.FormatBool(*params.Quality)
	}

	if params.Tag != nil && *params.Tag != "" {
		queryParams["tag"] = *params.Tag
	}

	if params.TranslationType != nil && *params.TranslationType != "" {
		queryParams["translation_type"] = *params.TranslationType
	}

	if params.UnverifyTranslationsUponDelivery != nil {
		queryParams["unverify_translations_upon_delivery"] = strconv.FormatBool(*params.UnverifyTranslationsUponDelivery)
	}

	return queryParams
}

type ProjectParams struct {
	AccountID               *string `json:"account_id,omitempty"  cli:"opt --account-id"`
	MainFormat              *string `json:"main_format,omitempty"  cli:"opt --main-format"`
	Name                    *string `json:"name,omitempty"  cli:"opt --name"`
	ProjectImage            *string `json:"project_image,omitempty"  cli:"opt --project-image"`
	RemoveProjectImage      *bool   `json:"remove_project_image,omitempty"  cli:"opt --remove-project-image"`
	SharesTranslationMemory *bool   `json:"shares_translation_memory,omitempty"  cli:"opt --shares-translation-memory"`
}

func (params *ProjectParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "account_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.AccountID = &val

		case "main_format":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.MainFormat = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		case "project_image":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectImage = &val

		case "remove_project_image":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.RemoveProjectImage = &val

		case "shares_translation_memory":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SharesTranslationMemory = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *ProjectParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.AccountID != nil && *params.AccountID != "" {
		queryParams["account_id"] = *params.AccountID
	}

	if params.MainFormat != nil && *params.MainFormat != "" {
		queryParams["main_format"] = *params.MainFormat
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	if params.RemoveProjectImage != nil {
		queryParams["remove_project_image"] = strconv.FormatBool(*params.RemoveProjectImage)
	}

	if params.SharesTranslationMemory != nil {
		queryParams["shares_translation_memory"] = strconv.FormatBool(*params.SharesTranslationMemory)
	}

	return queryParams
}

type ReleasesParams struct {
	Branch      *string  `json:"branch,omitempty"  cli:"opt --branch"`
	Description *string  `json:"description,omitempty"  cli:"opt --description"`
	Platforms   []string `json:"platforms,omitempty"  cli:"opt --platforms"`
}

func (params *ReleasesParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "description":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Description = &val

		case "platforms":
			ok := false
			params.Platforms, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *ReleasesParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Description != nil && *params.Description != "" {
		queryParams["description"] = *params.Description
	}

	return queryParams
}

type ScreenshotMarkerParams struct {
	KeyID        *string `json:"key_id,omitempty"  cli:"opt --key-id"`
	Presentation *string `json:"presentation,omitempty"  cli:"opt --presentation"`
}

func (params *ScreenshotMarkerParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "key_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.KeyID = &val

		case "presentation":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Presentation = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *ScreenshotMarkerParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Presentation != nil && *params.Presentation != "" {
		queryParams["presentation"] = *params.Presentation
	}

	return queryParams
}

type ScreenshotParams struct {
	Description *string `json:"description,omitempty"  cli:"opt --description"`
	Filename    *string `json:"filename,omitempty"  cli:"opt --filename"`
	Name        *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *ScreenshotParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "description":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Description = &val

		case "filename":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Filename = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *ScreenshotParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Description != nil && *params.Description != "" {
		queryParams["description"] = *params.Description
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type SpaceParams struct {
	ID   *string `json:"id,omitempty"  cli:"opt --id"`
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *SpaceParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ID = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *SpaceParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.ID != nil && *params.ID != "" {
		queryParams["id"] = *params.ID
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type StyleguideParams struct {
	Audience           *string `json:"audience,omitempty"  cli:"opt --audience"`
	Business           *string `json:"business,omitempty"  cli:"opt --business"`
	CompanyBranding    *string `json:"company_branding,omitempty"  cli:"opt --company-branding"`
	Formatting         *string `json:"formatting,omitempty"  cli:"opt --formatting"`
	GlossaryTerms      *string `json:"glossary_terms,omitempty"  cli:"opt --glossary-terms"`
	GrammarConsistency *string `json:"grammar_consistency,omitempty"  cli:"opt --grammar-consistency"`
	GrammaticalPerson  *string `json:"grammatical_person,omitempty"  cli:"opt --grammatical-person"`
	LiteralTranslation *string `json:"literal_translation,omitempty"  cli:"opt --literal-translation"`
	OverallTone        *string `json:"overall_tone,omitempty"  cli:"opt --overall-tone"`
	Samples            *string `json:"samples,omitempty"  cli:"opt --samples"`
	TargetAudience     *string `json:"target_audience,omitempty"  cli:"opt --target-audience"`
	Title              *string `json:"title,omitempty"  cli:"opt --title"`
	VocabularyType     *string `json:"vocabulary_type,omitempty"  cli:"opt --vocabulary-type"`
}

func (params *StyleguideParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "audience":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Audience = &val

		case "business":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Business = &val

		case "company_branding":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.CompanyBranding = &val

		case "formatting":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Formatting = &val

		case "glossary_terms":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.GlossaryTerms = &val

		case "grammar_consistency":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.GrammarConsistency = &val

		case "grammatical_person":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.GrammaticalPerson = &val

		case "literal_translation":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LiteralTranslation = &val

		case "overall_tone":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.OverallTone = &val

		case "samples":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Samples = &val

		case "target_audience":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.TargetAudience = &val

		case "title":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Title = &val

		case "vocabulary_type":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.VocabularyType = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *StyleguideParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Audience != nil && *params.Audience != "" {
		queryParams["audience"] = *params.Audience
	}

	if params.Business != nil && *params.Business != "" {
		queryParams["business"] = *params.Business
	}

	if params.CompanyBranding != nil && *params.CompanyBranding != "" {
		queryParams["company_branding"] = *params.CompanyBranding
	}

	if params.Formatting != nil && *params.Formatting != "" {
		queryParams["formatting"] = *params.Formatting
	}

	if params.GlossaryTerms != nil && *params.GlossaryTerms != "" {
		queryParams["glossary_terms"] = *params.GlossaryTerms
	}

	if params.GrammarConsistency != nil && *params.GrammarConsistency != "" {
		queryParams["grammar_consistency"] = *params.GrammarConsistency
	}

	if params.GrammaticalPerson != nil && *params.GrammaticalPerson != "" {
		queryParams["grammatical_person"] = *params.GrammaticalPerson
	}

	if params.LiteralTranslation != nil && *params.LiteralTranslation != "" {
		queryParams["literal_translation"] = *params.LiteralTranslation
	}

	if params.OverallTone != nil && *params.OverallTone != "" {
		queryParams["overall_tone"] = *params.OverallTone
	}

	if params.Samples != nil && *params.Samples != "" {
		queryParams["samples"] = *params.Samples
	}

	if params.TargetAudience != nil && *params.TargetAudience != "" {
		queryParams["target_audience"] = *params.TargetAudience
	}

	if params.Title != nil && *params.Title != "" {
		queryParams["title"] = *params.Title
	}

	if params.VocabularyType != nil && *params.VocabularyType != "" {
		queryParams["vocabulary_type"] = *params.VocabularyType
	}

	return queryParams
}

type TagParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Name   *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *TagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TagParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

type TranslationParams struct {
	Autotranslate *bool   `json:"autotranslate,omitempty"  cli:"opt --autotranslate"`
	Branch        *string `json:"branch,omitempty"  cli:"opt --branch"`
	Content       *string `json:"content,omitempty"  cli:"opt --content"`
	Excluded      *bool   `json:"excluded,omitempty"  cli:"opt --excluded"`
	KeyID         *string `json:"key_id,omitempty"  cli:"opt --key-id"`
	LocaleID      *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	PluralSuffix  *string `json:"plural_suffix,omitempty"  cli:"opt --plural-suffix"`
	Unverified    *bool   `json:"unverified,omitempty"  cli:"opt --unverified"`
}

func (params *TranslationParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "autotranslate":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Autotranslate = &val

		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "content":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Content = &val

		case "excluded":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Excluded = &val

		case "key_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.KeyID = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "plural_suffix":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.PluralSuffix = &val

		case "unverified":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Unverified = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Autotranslate != nil {
		queryParams["autotranslate"] = strconv.FormatBool(*params.Autotranslate)
	}

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Content != nil && *params.Content != "" {
		queryParams["content"] = *params.Content
	}

	if params.Excluded != nil {
		queryParams["excluded"] = strconv.FormatBool(*params.Excluded)
	}

	if params.PluralSuffix != nil && *params.PluralSuffix != "" {
		queryParams["plural_suffix"] = *params.PluralSuffix
	}

	if params.Unverified != nil {
		queryParams["unverified"] = strconv.FormatBool(*params.Unverified)
	}

	return queryParams
}

type UploadParams struct {
	Autotranslate      *bool             `json:"autotranslate,omitempty"  cli:"opt --autotranslate"`
	Branch             *string           `json:"branch,omitempty"  cli:"opt --branch"`
	ConvertEmoji       *bool             `json:"convert_emoji,omitempty"  cli:"opt --convert-emoji"`
	File               *string           `json:"file,omitempty"  cli:"opt --file"`
	FileEncoding       *string           `json:"file_encoding,omitempty"  cli:"opt --file-encoding"`
	FileFormat         *string           `json:"file_format,omitempty"  cli:"opt --file-format"`
	FormatOptions      map[string]string `json:"format_options,omitempty"  cli:"opt --format-options"`
	LocaleID           *string           `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	LocaleMapping      map[string]string `json:"locale_mapping,omitempty"  cli:"opt --locale-mapping"`
	MarkReviewed       *bool             `json:"mark_reviewed,omitempty"  cli:"opt --mark-reviewed"`
	SkipUnverification *bool             `json:"skip_unverification,omitempty"  cli:"opt --skip-unverification"`
	SkipUploadTags     *bool             `json:"skip_upload_tags,omitempty"  cli:"opt --skip-upload-tags"`
	Tags               *string           `json:"tags,omitempty"  cli:"opt --tags"`
	UpdateDescriptions *bool             `json:"update_descriptions,omitempty"  cli:"opt --update-descriptions"`
	UpdateTranslations *bool             `json:"update_translations,omitempty"  cli:"opt --update-translations"`
}

func (params *UploadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "autotranslate":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Autotranslate = &val

		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "convert_emoji":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ConvertEmoji = &val

		case "file":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.File = &val

		case "file_encoding":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FileEncoding = &val

		case "file_format":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FileFormat = &val

		case "format_options":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.FormatOptions = val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "locale_mapping":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.LocaleMapping = val

		case "mark_reviewed":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.MarkReviewed = &val

		case "skip_unverification":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SkipUnverification = &val

		case "skip_upload_tags":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SkipUploadTags = &val

		case "tags":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tags = &val

		case "update_descriptions":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UpdateDescriptions = &val

		case "update_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UpdateTranslations = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *UploadParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Autotranslate != nil {
		queryParams["autotranslate"] = strconv.FormatBool(*params.Autotranslate)
	}

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.ConvertEmoji != nil {
		queryParams["convert_emoji"] = strconv.FormatBool(*params.ConvertEmoji)
	}

	if params.FileEncoding != nil && *params.FileEncoding != "" {
		queryParams["file_encoding"] = *params.FileEncoding
	}

	if params.FileFormat != nil && *params.FileFormat != "" {
		queryParams["file_format"] = *params.FileFormat
	}

	for key, value := range convertMapToQueryParams("format_options", params.FormatOptions) {
		queryParams[key] = value
	}

	for key, value := range convertMapToQueryParams("locale_mapping", params.LocaleMapping) {
		queryParams[key] = value
	}

	if params.MarkReviewed != nil {
		queryParams["mark_reviewed"] = strconv.FormatBool(*params.MarkReviewed)
	}

	if params.SkipUnverification != nil {
		queryParams["skip_unverification"] = strconv.FormatBool(*params.SkipUnverification)
	}

	if params.SkipUploadTags != nil {
		queryParams["skip_upload_tags"] = strconv.FormatBool(*params.SkipUploadTags)
	}

	if params.Tags != nil && *params.Tags != "" {
		queryParams["tags"] = *params.Tags
	}

	if params.UpdateDescriptions != nil {
		queryParams["update_descriptions"] = strconv.FormatBool(*params.UpdateDescriptions)
	}

	if params.UpdateTranslations != nil {
		queryParams["update_translations"] = strconv.FormatBool(*params.UpdateTranslations)
	}

	return queryParams
}

type WebhookParams struct {
	Active      *bool   `json:"active,omitempty"  cli:"opt --active"`
	CallbackUrl *string `json:"callback_url,omitempty"  cli:"opt --callback-url"`
	Description *string `json:"description,omitempty"  cli:"opt --description"`
	Events      *string `json:"events,omitempty"  cli:"opt --events"`
}

func (params *WebhookParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "active":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Active = &val

		case "callback_url":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.CallbackUrl = &val

		case "description":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Description = &val

		case "events":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Events = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *WebhookParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Active != nil {
		queryParams["active"] = strconv.FormatBool(*params.Active)
	}

	if params.CallbackUrl != nil && *params.CallbackUrl != "" {
		queryParams["callback_url"] = *params.CallbackUrl
	}

	if params.Description != nil && *params.Description != "" {
		queryParams["description"] = *params.Description
	}

	if params.Events != nil && *params.Events != "" {
		queryParams["events"] = *params.Events
	}

	return queryParams
}

// Get details on a single account.
func (client *Client) AccountShow(id string) (*AccountDetails, error) {
	retVal := new(AccountDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s", url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all accounts the current user has access to.
func (client *Client) AccountsList(page, perPage int) ([]*Account, error) {
	retVal := []*Account{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts")

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new authorization.
func (client *Client) AuthorizationCreate(params *AuthorizationParams) (*AuthorizationWithToken, error) {
	retVal := new(AuthorizationWithToken)
	err := func() error {

		url := fmt.Sprintf("/v2/authorizations")

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing authorization. API calls using that token will stop working.
func (client *Client) AuthorizationDelete(id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/authorizations/%s", url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single authorization.
func (client *Client) AuthorizationShow(id string) (*Authorization, error) {
	retVal := new(Authorization)
	err := func() error {

		url := fmt.Sprintf("/v2/authorizations/%s", url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing authorization.
func (client *Client) AuthorizationUpdate(id string, params *AuthorizationParams) (*Authorization, error) {
	retVal := new(Authorization)
	err := func() error {

		url := fmt.Sprintf("/v2/authorizations/%s", url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all your authorizations.
func (client *Client) AuthorizationsList(page, perPage int) ([]*Authorization, error) {
	retVal := []*Authorization{}
	err := func() error {

		url := fmt.Sprintf("/v2/authorizations")

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Export translations from Phrase to Bitbucket according to the .phraseapp.yml file within the Bitbucket Repository.
func (client *Client) BitbucketSyncExport(id string, params *BitbucketSyncParams) (*BitbucketSyncExportResponse, error) {
	retVal := new(BitbucketSyncExportResponse)
	err := func() error {

		url := fmt.Sprintf("/v2/bitbucket_syncs/%s/export", url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Import translations from Bitbucket to Phrase according to the .phraseapp.yml file within the Bitbucket repository.
func (client *Client) BitbucketSyncImport(id string, params *BitbucketSyncParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/bitbucket_syncs/%s/import", url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// List all Bitbucket repositories for which synchronisation with Phrase is activated.
func (client *Client) BitbucketSyncsList(page, perPage int, params *BitbucketSyncParams) ([]*BitbucketSync, error) {
	retVal := []*BitbucketSync{}
	err := func() error {

		url := fmt.Sprintf("/v2/bitbucket_syncs")

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new rule for blacklisting keys.
func (client *Client) BlacklistedKeyCreate(project_id string, params *BlacklistedKeyParams) (*BlacklistedKey, error) {
	retVal := new(BlacklistedKey)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing rule for blacklisting keys.
func (client *Client) BlacklistedKeyDelete(project_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single rule for blacklisting keys for a given project.
func (client *Client) BlacklistedKeyShow(project_id, id string) (*BlacklistedKey, error) {
	retVal := new(BlacklistedKey)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing rule for blacklisting keys.
func (client *Client) BlacklistedKeyUpdate(project_id, id string, params *BlacklistedKeyParams) (*BlacklistedKey, error) {
	retVal := new(BlacklistedKey)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all rules for blacklisting keys for the given project.
func (client *Client) BlacklistedKeysList(project_id string, page, perPage int) ([]*BlacklistedKey, error) {
	retVal := []*BlacklistedKey{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys", url.QueryEscape(project_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Compare branch with main branch.
func (client *Client) BranchCompare(project_id, name string, params *BranchParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches/%s/compare", url.QueryEscape(project_id), url.QueryEscape(name))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Create a new branch.
func (client *Client) BranchCreate(project_id string, params *BranchParams) (*Branch, error) {
	retVal := new(Branch)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing branch.
func (client *Client) BranchDelete(project_id, name string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches/%s", url.QueryEscape(project_id), url.QueryEscape(name))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type BranchMergeParams struct {
	Strategy *string `json:"strategy,omitempty"  cli:"opt --strategy"`
}

func (params *BranchMergeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "strategy":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Strategy = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *BranchMergeParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Strategy != nil && *params.Strategy != "" {
		queryParams["strategy"] = *params.Strategy
	}

	return queryParams
}

// Merge an existing branch.
func (client *Client) BranchMerge(project_id, name string, params *BranchMergeParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches/%s/merge", url.QueryEscape(project_id), url.QueryEscape(name))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single branch for a given project.
func (client *Client) BranchShow(project_id, name string) (*Branch, error) {
	retVal := new(Branch)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches/%s", url.QueryEscape(project_id), url.QueryEscape(name))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing branch.
func (client *Client) BranchUpdate(project_id, name string, params *BranchParams) (*Branch, error) {
	retVal := new(Branch)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches/%s", url.QueryEscape(project_id), url.QueryEscape(name))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all branches the of the current project.
func (client *Client) BranchesList(project_id string, page, perPage int) ([]*Branch, error) {
	retVal := []*Branch{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/branches", url.QueryEscape(project_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new comment for a key.
func (client *Client) CommentCreate(project_id, key_id string, params *CommentParams) (*Comment, error) {
	retVal := new(Comment)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments", url.QueryEscape(project_id), url.QueryEscape(key_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type CommentDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing comment.
func (client *Client) CommentDelete(project_id, key_id, id string, params *CommentDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type CommentMarkCheckParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentMarkCheckParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentMarkCheckParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Check if comment was marked as read. Returns 204 if read, 404 if unread.
func (client *Client) CommentMarkCheck(project_id, key_id, id string, params *CommentMarkCheckParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type CommentMarkReadParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentMarkReadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentMarkReadParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a comment as read.
func (client *Client) CommentMarkRead(project_id, key_id, id string, params *CommentMarkReadParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type CommentMarkUnreadParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentMarkUnreadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentMarkUnreadParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a comment as unread.
func (client *Client) CommentMarkUnread(project_id, key_id, id string, params *CommentMarkUnreadParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type CommentShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single comment.
func (client *Client) CommentShow(project_id, key_id, id string, params *CommentShowParams) (*Comment, error) {
	retVal := new(Comment)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing comment.
func (client *Client) CommentUpdate(project_id, key_id, id string, params *CommentParams) (*Comment, error) {
	retVal := new(Comment)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", url.QueryEscape(project_id), url.QueryEscape(key_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type CommentsListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *CommentsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *CommentsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all comments for a key.
func (client *Client) CommentsList(project_id, key_id string, page, perPage int, params *CommentsListParams) ([]*Comment, error) {
	retVal := []*Comment{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments", url.QueryEscape(project_id), url.QueryEscape(key_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new distribution.
func (client *Client) DistributionCreate(account_id string, params *DistributionsParams) (*Distribution, error) {
	retVal := new(Distribution)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions", url.QueryEscape(account_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing distribution.
func (client *Client) DistributionDelete(account_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single distribution.
func (client *Client) DistributionShow(account_id, id string) (*Distribution, error) {
	retVal := new(Distribution)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing distribution.
func (client *Client) DistributionUpdate(account_id, id string, params *DistributionsParams) (*Distribution, error) {
	retVal := new(Distribution)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all distributions for the given account.
func (client *Client) DistributionsList(account_id string, page, perPage int) ([]*DistributionPreview, error) {
	retVal := []*DistributionPreview{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions", url.QueryEscape(account_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get a handy list of all localization file formats supported in Phrase.
func (client *Client) FormatsList(page, perPage int) ([]*Format, error) {
	retVal := []*Format{}
	err := func() error {

		url := fmt.Sprintf("/v2/formats")

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all glossaries the current user has access to.
func (client *Client) GlossariesList(account_id string, page, perPage int) ([]*Glossary, error) {
	retVal := []*Glossary{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries", url.QueryEscape(account_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new glossary.
func (client *Client) GlossaryCreate(account_id string, params *GlossaryParams) (*Glossary, error) {
	retVal := new(Glossary)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries", url.QueryEscape(account_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing glossary.
func (client *Client) GlossaryDelete(account_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single glossary.
func (client *Client) GlossaryShow(account_id, id string) (*Glossary, error) {
	retVal := new(Glossary)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing glossary.
func (client *Client) GlossaryUpdate(account_id, id string, params *GlossaryParams) (*Glossary, error) {
	retVal := new(Glossary)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new glossary term.
func (client *Client) GlossaryTermCreate(account_id, glossary_id string, params *GlossaryTermParams) (*GlossaryTerm, error) {
	retVal := new(GlossaryTerm)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms", url.QueryEscape(account_id), url.QueryEscape(glossary_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing glossary term.
func (client *Client) GlossaryTermDelete(account_id, glossary_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single glossary term.
func (client *Client) GlossaryTermShow(account_id, glossary_id, id string) (*GlossaryTerm, error) {
	retVal := new(GlossaryTerm)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing glossary term.
func (client *Client) GlossaryTermUpdate(account_id, glossary_id, id string, params *GlossaryTermParams) (*GlossaryTerm, error) {
	retVal := new(GlossaryTerm)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new glossary term translation.
func (client *Client) GlossaryTermTranslationCreate(account_id, glossary_id, term_id string, params *GlossaryTermTranslationParams) (*GlossaryTermTranslation, error) {
	retVal := new(GlossaryTermTranslation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s/translations", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(term_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing glossary term translation.
func (client *Client) GlossaryTermTranslationDelete(account_id, glossary_id, term_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s/translations/%s", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(term_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Update an existing glossary term translation.
func (client *Client) GlossaryTermTranslationUpdate(account_id, glossary_id, term_id, id string, params *GlossaryTermTranslationParams) (*GlossaryTermTranslation, error) {
	retVal := new(GlossaryTermTranslation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms/%s/translations/%s", url.QueryEscape(account_id), url.QueryEscape(glossary_id), url.QueryEscape(term_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all glossary terms the current user has access to.
func (client *Client) GlossaryTermsList(account_id, glossary_id string, page, perPage int) ([]*GlossaryTerm, error) {
	retVal := []*GlossaryTerm{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/glossaries/%s/terms", url.QueryEscape(account_id), url.QueryEscape(glossary_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type InvitationCreateParams struct {
	Email       *string           `json:"email,omitempty"  cli:"opt --email"`
	LocaleIDs   *string           `json:"locale_ids,omitempty"  cli:"opt --locale-ids"`
	Permissions map[string]string `json:"permissions,omitempty"  cli:"opt --permissions"`
	ProjectIDs  *string           `json:"project_ids,omitempty"  cli:"opt --project-ids"`
	Role        *string           `json:"role,omitempty"  cli:"opt --role"`
}

func (params *InvitationCreateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "email":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Email = &val

		case "locale_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleIDs = &val

		case "permissions":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.Permissions = val

		case "project_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectIDs = &val

		case "role":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Role = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *InvitationCreateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Email != nil && *params.Email != "" {
		queryParams["email"] = *params.Email
	}

	if params.LocaleIDs != nil && *params.LocaleIDs != "" {
		queryParams["locale_ids"] = *params.LocaleIDs
	}

	for key, value := range convertMapToQueryParams("permissions", params.Permissions) {
		queryParams[key] = value
	}

	if params.ProjectIDs != nil && *params.ProjectIDs != "" {
		queryParams["project_ids"] = *params.ProjectIDs
	}

	if params.Role != nil && *params.Role != "" {
		queryParams["role"] = *params.Role
	}

	return queryParams
}

// Invite a person to an account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationCreate(account_id string, params *InvitationCreateParams) (*Invitation, error) {
	retVal := new(Invitation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations", url.QueryEscape(account_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing invitation (must not be accepted yet). Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationDelete(account_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Resend the invitation email (must not be accepted yet). Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationResend(account_id, id string) (*Invitation, error) {
	retVal := new(Invitation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations/%s/resend", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("POST", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get details on a single invitation. Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationShow(account_id, id string) (*Invitation, error) {
	retVal := new(Invitation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type InvitationUpdateParams struct {
	LocaleIDs   *string           `json:"locale_ids,omitempty"  cli:"opt --locale-ids"`
	Permissions map[string]string `json:"permissions,omitempty"  cli:"opt --permissions"`
	ProjectIDs  *string           `json:"project_ids,omitempty"  cli:"opt --project-ids"`
	Role        *string           `json:"role,omitempty"  cli:"opt --role"`
}

func (params *InvitationUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "locale_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleIDs = &val

		case "permissions":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.Permissions = val

		case "project_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectIDs = &val

		case "role":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Role = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *InvitationUpdateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.LocaleIDs != nil && *params.LocaleIDs != "" {
		queryParams["locale_ids"] = *params.LocaleIDs
	}

	for key, value := range convertMapToQueryParams("permissions", params.Permissions) {
		queryParams[key] = value
	}

	if params.ProjectIDs != nil && *params.ProjectIDs != "" {
		queryParams["project_ids"] = *params.ProjectIDs
	}

	if params.Role != nil && *params.Role != "" {
		queryParams["role"] = *params.Role
	}

	return queryParams
}

// Update an existing invitation (must not be accepted yet). The <code>email</code> cannot be updated. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationUpdate(account_id, id string, params *InvitationUpdateParams) (*Invitation, error) {
	retVal := new(Invitation)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List invitations for an account. It will also list the accessible resources like projects and locales the invited user has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.
func (client *Client) InvitationsList(account_id string, page, perPage int) ([]*Invitation, error) {
	retVal := []*Invitation{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/invitations", url.QueryEscape(account_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobCompleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobCompleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobCompleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a job as completed.
func (client *Client) JobComplete(project_id, id string, params *JobCompleteParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/complete", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new job.
func (client *Client) JobCreate(project_id string, params *JobParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing job.
func (client *Client) JobDelete(project_id, id string, params *JobDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type JobKeysCreateParams struct {
	Branch            *string  `json:"branch,omitempty"  cli:"opt --branch"`
	TranslationKeyIDs []string `json:"translation_key_ids,omitempty"  cli:"opt --translation-key-ids"`
}

func (params *JobKeysCreateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "translation_key_ids":
			ok := false
			params.TranslationKeyIDs, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobKeysCreateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Add multiple keys to a existing job.
func (client *Client) JobKeysCreate(project_id, id string, params *JobKeysCreateParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/keys", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobKeysDeleteParams struct {
	Branch            *string  `json:"branch,omitempty"  cli:"opt --branch"`
	TranslationKeyIDs []string `json:"translation_key_ids,omitempty"  cli:"opt --translation-key-ids"`
}

func (params *JobKeysDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "translation_key_ids":
			ok := false
			params.TranslationKeyIDs, ok = v.([]string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobKeysDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Remove multiple keys from existing job.
func (client *Client) JobKeysDelete(project_id, id string, params *JobKeysDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/keys", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type JobReopenParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobReopenParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobReopenParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a job as uncompleted.
func (client *Client) JobReopen(project_id, id string, params *JobReopenParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/reopen", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single job for a given project.
func (client *Client) JobShow(project_id, id string, params *JobShowParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobStartParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobStartParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobStartParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Starts an existing job in state draft.
func (client *Client) JobStart(project_id, id string, params *JobStartParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/start", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobUpdateParams struct {
	Branch   *string     `json:"branch,omitempty"  cli:"opt --branch"`
	Briefing *string     `json:"briefing,omitempty"  cli:"opt --briefing"`
	DueDate  **time.Time `json:"due_date,omitempty"  cli:"opt --due-date"`
	Name     *string     `json:"name,omitempty"  cli:"opt --name"`
}

func (params *JobUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "briefing":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Briefing = &val

		case "due_date":
			val, ok := v.(*time.Time)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.DueDate = &val

		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobUpdateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Briefing != nil && *params.Briefing != "" {
		queryParams["briefing"] = *params.Briefing
	}

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

// Update an existing job.
func (client *Client) JobUpdate(project_id, id string, params *JobUpdateParams) (*JobDetails, error) {
	retVal := new(JobDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobLocaleCompleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobLocaleCompleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocaleCompleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a job locale as completed.
func (client *Client) JobLocaleComplete(project_id, job_id, id string, params *JobLocaleCompleteParams) (*JobLocale, error) {
	retVal := new(JobLocale)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales/%s/complete", url.QueryEscape(project_id), url.QueryEscape(job_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobLocaleDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobLocaleDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocaleDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing job locale.
func (client *Client) JobLocaleDelete(project_id, job_id, id string, params *JobLocaleDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales/%s", url.QueryEscape(project_id), url.QueryEscape(job_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type JobLocaleReopenParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobLocaleReopenParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocaleReopenParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark a job locale as uncompleted.
func (client *Client) JobLocaleReopen(project_id, job_id, id string, params *JobLocaleReopenParams) (*JobLocale, error) {
	retVal := new(JobLocale)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales/%s/reopen", url.QueryEscape(project_id), url.QueryEscape(job_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobLocaleShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobLocaleShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocaleShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get a single job locale for a given job.
func (client *Client) JobLocaleShow(project_id, job_id, id string, params *JobLocaleShowParams) (*JobLocale, error) {
	retVal := new(JobLocale)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locale/%s", url.QueryEscape(project_id), url.QueryEscape(job_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing job locale.
func (client *Client) JobLocaleUpdate(project_id, job_id, id string, params *JobLocaleParams) (*JobLocale, error) {
	retVal := new(JobLocale)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales/%s", url.QueryEscape(project_id), url.QueryEscape(job_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new job locale.
func (client *Client) JobLocalesCreate(project_id, job_id string, params *JobLocaleParams) (*JobLocale, error) {
	retVal := new(JobLocale)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales", url.QueryEscape(project_id), url.QueryEscape(job_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobLocalesListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *JobLocalesListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobLocalesListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all job locales for a given job.
func (client *Client) JobLocalesList(project_id, job_id string, page, perPage int, params *JobLocalesListParams) ([]*JobLocale, error) {
	retVal := []*JobLocale{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs/%s/locales", url.QueryEscape(project_id), url.QueryEscape(job_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type JobsListParams struct {
	AssignedTo *string `json:"assigned_to,omitempty"  cli:"opt --assigned-to"`
	Branch     *string `json:"branch,omitempty"  cli:"opt --branch"`
	OwnedBy    *string `json:"owned_by,omitempty"  cli:"opt --owned-by"`
	State      *string `json:"state,omitempty"  cli:"opt --state"`
}

func (params *JobsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "assigned_to":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.AssignedTo = &val

		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "owned_by":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.OwnedBy = &val

		case "state":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.State = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *JobsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.AssignedTo != nil && *params.AssignedTo != "" {
		queryParams["assigned_to"] = *params.AssignedTo
	}

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.OwnedBy != nil && *params.OwnedBy != "" {
		queryParams["owned_by"] = *params.OwnedBy
	}

	if params.State != nil && *params.State != "" {
		queryParams["state"] = *params.State
	}

	return queryParams
}

// List all jobs for the given project.
func (client *Client) JobsList(project_id string, page, perPage int, params *JobsListParams) ([]*Job, error) {
	retVal := []*Job{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/jobs", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new key.
func (client *Client) KeyCreate(project_id string, params *TranslationKeyParams) (*TranslationKeyDetails, error) {
	retVal := new(TranslationKeyDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.Branch != nil {
			err := writer.WriteField("branch", *params.Branch)
			if err != nil {
				return err
			}
		}

		if params.DataType != nil {
			err := writer.WriteField("data_type", *params.DataType)
			if err != nil {
				return err
			}
		}

		if params.Description != nil {
			err := writer.WriteField("description", *params.Description)
			if err != nil {
				return err
			}
		}

		if params.LocalizedFormatKey != nil {
			err := writer.WriteField("localized_format_key", *params.LocalizedFormatKey)
			if err != nil {
				return err
			}
		}

		if params.LocalizedFormatString != nil {
			err := writer.WriteField("localized_format_string", *params.LocalizedFormatString)
			if err != nil {
				return err
			}
		}

		if params.MaxCharactersAllowed != nil {
			err := writer.WriteField("max_characters_allowed", strconv.FormatInt(*params.MaxCharactersAllowed, 10))
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}

		if params.NamePlural != nil {
			err := writer.WriteField("name_plural", *params.NamePlural)
			if err != nil {
				return err
			}
		}

		if params.OriginalFile != nil {
			err := writer.WriteField("original_file", *params.OriginalFile)
			if err != nil {
				return err
			}
		}

		if params.Plural != nil {
			err := writer.WriteField("plural", strconv.FormatBool(*params.Plural))
			if err != nil {
				return err
			}
		}

		if params.RemoveScreenshot != nil {
			err := writer.WriteField("remove_screenshot", strconv.FormatBool(*params.RemoveScreenshot))
			if err != nil {
				return err
			}
		}

		if params.Screenshot != nil {
			part, err := writer.CreateFormFile("screenshot", filepath.Base(*params.Screenshot))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.Screenshot)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.Tags != nil {
			err := writer.WriteField("tags", *params.Tags)
			if err != nil {
				return err
			}
		}

		if params.Unformatted != nil {
			err := writer.WriteField("unformatted", strconv.FormatBool(*params.Unformatted))
			if err != nil {
				return err
			}
		}

		if params.XmlSpacePreserve != nil {
			err := writer.WriteField("xml_space_preserve", strconv.FormatBool(*params.XmlSpacePreserve))
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("POST", url, ctype, paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeyDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *KeyDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeyDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing key.
func (client *Client) KeyDelete(project_id, id string, params *KeyDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type KeyShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *KeyShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeyShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single key for a given project.
func (client *Client) KeyShow(project_id, id string, params *KeyShowParams) (*TranslationKeyDetails, error) {
	retVal := new(TranslationKeyDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing key.
func (client *Client) KeyUpdate(project_id, id string, params *TranslationKeyParams) (*TranslationKeyDetails, error) {
	retVal := new(TranslationKeyDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.Branch != nil {
			err := writer.WriteField("branch", *params.Branch)
			if err != nil {
				return err
			}
		}

		if params.DataType != nil {
			err := writer.WriteField("data_type", *params.DataType)
			if err != nil {
				return err
			}
		}

		if params.Description != nil {
			err := writer.WriteField("description", *params.Description)
			if err != nil {
				return err
			}
		}

		if params.LocalizedFormatKey != nil {
			err := writer.WriteField("localized_format_key", *params.LocalizedFormatKey)
			if err != nil {
				return err
			}
		}

		if params.LocalizedFormatString != nil {
			err := writer.WriteField("localized_format_string", *params.LocalizedFormatString)
			if err != nil {
				return err
			}
		}

		if params.MaxCharactersAllowed != nil {
			err := writer.WriteField("max_characters_allowed", strconv.FormatInt(*params.MaxCharactersAllowed, 10))
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}

		if params.NamePlural != nil {
			err := writer.WriteField("name_plural", *params.NamePlural)
			if err != nil {
				return err
			}
		}

		if params.OriginalFile != nil {
			err := writer.WriteField("original_file", *params.OriginalFile)
			if err != nil {
				return err
			}
		}

		if params.Plural != nil {
			err := writer.WriteField("plural", strconv.FormatBool(*params.Plural))
			if err != nil {
				return err
			}
		}

		if params.RemoveScreenshot != nil {
			err := writer.WriteField("remove_screenshot", strconv.FormatBool(*params.RemoveScreenshot))
			if err != nil {
				return err
			}
		}

		if params.Screenshot != nil {
			part, err := writer.CreateFormFile("screenshot", filepath.Base(*params.Screenshot))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.Screenshot)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.Tags != nil {
			err := writer.WriteField("tags", *params.Tags)
			if err != nil {
				return err
			}
		}

		if params.Unformatted != nil {
			err := writer.WriteField("unformatted", strconv.FormatBool(*params.Unformatted))
			if err != nil {
				return err
			}
		}

		if params.XmlSpacePreserve != nil {
			err := writer.WriteField("xml_space_preserve", strconv.FormatBool(*params.XmlSpacePreserve))
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("PATCH", url, ctype, paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysDeleteParams struct {
	Branch   *string `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
}

func (params *KeysDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeysDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	return queryParams
}

// Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.
func (client *Client) KeysDelete(project_id string, params *KeysDeleteParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysListParams struct {
	Branch   *string `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Order    *string `json:"order,omitempty"  cli:"opt --order"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort     *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *KeysListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeysListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// List all keys for the given project. Alternatively you can POST requests to /search.
func (client *Client) KeysList(project_id string, page, perPage int, params *KeysListParams) ([]*TranslationKey, error) {
	retVal := []*TranslationKey{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysSearchParams struct {
	Branch   *string `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Order    *string `json:"order,omitempty"  cli:"opt --order"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort     *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *KeysSearchParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeysSearchParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Search keys for the given project matching query.
func (client *Client) KeysSearch(project_id string, page, perPage int, params *KeysSearchParams) ([]*TranslationKey, error) {
	retVal := []*TranslationKey{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/search", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("POST", url, "application/json", paramsBuf, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysTagParams struct {
	Branch   *string `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Tags     *string `json:"tags,omitempty"  cli:"opt --tags"`
}

func (params *KeysTagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "tags":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tags = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeysTagParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Tags != nil && *params.Tags != "" {
		queryParams["tags"] = *params.Tags
	}

	return queryParams
}

// Tags all keys matching query. Same constraints as list.
func (client *Client) KeysTag(project_id string, params *KeysTagParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/tag", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysUntagParams struct {
	Branch   *string `json:"branch,omitempty"  cli:"opt --branch"`
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Tags     *string `json:"tags,omitempty"  cli:"opt --tags"`
}

func (params *KeysUntagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleID = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "tags":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tags = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *KeysUntagParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Tags != nil && *params.Tags != "" {
		queryParams["tags"] = *params.Tags
	}

	return queryParams
}

// Removes specified tags from keys matching query.
func (client *Client) KeysUntag(project_id string, params *KeysUntagParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/untag", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new locale.
func (client *Client) LocaleCreate(project_id string, params *LocaleParams) (*LocaleDetails, error) {
	retVal := new(LocaleDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type LocaleDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *LocaleDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *LocaleDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing locale.
func (client *Client) LocaleDelete(project_id, id string, params *LocaleDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type LocaleDownloadParams struct {
	Branch                        *string           `json:"branch,omitempty"  cli:"opt --branch"`
	ConvertEmoji                  *bool             `json:"convert_emoji,omitempty"  cli:"opt --convert-emoji"`
	Encoding                      *string           `json:"encoding,omitempty"  cli:"opt --encoding"`
	FallbackLocaleID              *string           `json:"fallback_locale_id,omitempty"  cli:"opt --fallback-locale-id"`
	FileFormat                    *string           `json:"file_format,omitempty"  cli:"opt --file-format"`
	FormatOptions                 map[string]string `json:"format_options,omitempty"  cli:"opt --format-options"`
	IncludeEmptyTranslations      *bool             `json:"include_empty_translations,omitempty"  cli:"opt --include-empty-translations"`
	IncludeTranslatedKeys         *bool             `json:"include_translated_keys,omitempty"  cli:"opt --include-translated-keys"`
	IncludeUnverifiedTranslations *bool             `json:"include_unverified_translations,omitempty"  cli:"opt --include-unverified-translations"`
	KeepNotranslateTags           *bool             `json:"keep_notranslate_tags,omitempty"  cli:"opt --keep-notranslate-tags"`
	SkipUnverifiedTranslations    *bool             `json:"skip_unverified_translations,omitempty"  cli:"opt --skip-unverified-translations"`
	Tag                           *string           `json:"tag,omitempty"  cli:"opt --tag"`
	Tags                          *string           `json:"tags,omitempty"  cli:"opt --tags"`
	UseLastReviewedVersion        *bool             `json:"use_last_reviewed_version,omitempty"  cli:"opt --use-last-reviewed-version"`
}

func (params *LocaleDownloadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "convert_emoji":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ConvertEmoji = &val

		case "encoding":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Encoding = &val

		case "fallback_locale_id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FallbackLocaleID = &val

		case "file_format":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.FileFormat = &val

		case "format_options":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.FormatOptions = val

		case "include_empty_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.IncludeEmptyTranslations = &val

		case "include_translated_keys":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.IncludeTranslatedKeys = &val

		case "include_unverified_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.IncludeUnverifiedTranslations = &val

		case "keep_notranslate_tags":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.KeepNotranslateTags = &val

		case "skip_unverified_translations":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.SkipUnverifiedTranslations = &val

		case "tag":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tag = &val

		case "tags":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tags = &val

		case "use_last_reviewed_version":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.UseLastReviewedVersion = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *LocaleDownloadParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.ConvertEmoji != nil {
		queryParams["convert_emoji"] = strconv.FormatBool(*params.ConvertEmoji)
	}

	if params.Encoding != nil && *params.Encoding != "" {
		queryParams["encoding"] = *params.Encoding
	}

	if params.FallbackLocaleID != nil && *params.FallbackLocaleID != "" {
		queryParams["fallback_locale_id"] = *params.FallbackLocaleID
	}

	if params.FileFormat != nil && *params.FileFormat != "" {
		queryParams["file_format"] = *params.FileFormat
	}

	for key, value := range convertMapToQueryParams("format_options", params.FormatOptions) {
		queryParams[key] = value
	}

	if params.IncludeEmptyTranslations != nil {
		queryParams["include_empty_translations"] = strconv.FormatBool(*params.IncludeEmptyTranslations)
	}

	if params.IncludeTranslatedKeys != nil {
		queryParams["include_translated_keys"] = strconv.FormatBool(*params.IncludeTranslatedKeys)
	}

	if params.IncludeUnverifiedTranslations != nil {
		queryParams["include_unverified_translations"] = strconv.FormatBool(*params.IncludeUnverifiedTranslations)
	}

	if params.KeepNotranslateTags != nil {
		queryParams["keep_notranslate_tags"] = strconv.FormatBool(*params.KeepNotranslateTags)
	}

	if params.SkipUnverifiedTranslations != nil {
		queryParams["skip_unverified_translations"] = strconv.FormatBool(*params.SkipUnverifiedTranslations)
	}

	if params.Tag != nil && *params.Tag != "" {
		queryParams["tag"] = *params.Tag
	}

	if params.Tags != nil && *params.Tags != "" {
		queryParams["tags"] = *params.Tags
	}

	if params.UseLastReviewedVersion != nil {
		queryParams["use_last_reviewed_version"] = strconv.FormatBool(*params.UseLastReviewedVersion)
	}

	return queryParams
}

// Download a locale in a specific file format.
func (client *Client) LocaleDownload(project_id, id string, params *LocaleDownloadParams) ([]byte, error) {
	retVal := []byte{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales/%s/download", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		retVal, err = ioutil.ReadAll(reader)
		return err

	}()
	return retVal, err
}

type LocaleShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *LocaleShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *LocaleShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single locale for a given project.
func (client *Client) LocaleShow(project_id, id string, params *LocaleShowParams) (*LocaleDetails, error) {
	retVal := new(LocaleDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing locale.
func (client *Client) LocaleUpdate(project_id, id string, params *LocaleParams) (*LocaleDetails, error) {
	retVal := new(LocaleDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type LocalesListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *LocalesListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *LocalesListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all locales for the given project.
func (client *Client) LocalesList(project_id string, page, perPage int, params *LocalesListParams) ([]*Locale, error) {
	retVal := []*Locale{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Remove a user from the account. The user will be removed from the account but not deleted from Phrase. Access token scope must include <code>team.manage</code>.
func (client *Client) MemberDelete(account_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/members/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single user in the account. Access token scope must include <code>team.manage</code>.
func (client *Client) MemberShow(account_id, id string) (*Member, error) {
	retVal := new(Member)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/members/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type MemberUpdateParams struct {
	LocaleIDs   *string           `json:"locale_ids,omitempty"  cli:"opt --locale-ids"`
	Permissions map[string]string `json:"permissions,omitempty"  cli:"opt --permissions"`
	ProjectIDs  *string           `json:"project_ids,omitempty"  cli:"opt --project-ids"`
	Role        *string           `json:"role,omitempty"  cli:"opt --role"`
}

func (params *MemberUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "locale_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.LocaleIDs = &val

		case "permissions":
			rval, err := ValidateIsRawMap(k, v)
			if err != nil {
				return err
			}
			val, err := ConvertToStringMap(rval)
			if err != nil {
				return err
			}
			params.Permissions = val

		case "project_ids":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ProjectIDs = &val

		case "role":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Role = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *MemberUpdateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.LocaleIDs != nil && *params.LocaleIDs != "" {
		queryParams["locale_ids"] = *params.LocaleIDs
	}

	for key, value := range convertMapToQueryParams("permissions", params.Permissions) {
		queryParams[key] = value
	}

	if params.ProjectIDs != nil && *params.ProjectIDs != "" {
		queryParams["project_ids"] = *params.ProjectIDs
	}

	if params.Role != nil && *params.Role != "" {
		queryParams["role"] = *params.Role
	}

	return queryParams
}

// Update user permissions in the account. Developers and translators need <code>project_ids</code> and <code>locale_ids</code> assigned to access them. Access token scope must include <code>team.manage</code>.
func (client *Client) MemberUpdate(account_id, id string, params *MemberUpdateParams) (*Member, error) {
	retVal := new(Member)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/members/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get all users active in the account. It also lists resources like projects and locales the member has access to. In case nothing is shown the default access from the role is used. Access token scope must include <code>team.manage</code>.
func (client *Client) MembersList(account_id string, page, perPage int) ([]*Member, error) {
	retVal := []*Member{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/members", url.QueryEscape(account_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type OrderConfirmParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *OrderConfirmParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *OrderConfirmParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Confirm an existing order and send it to the provider for translation. Same constraints as for create.
func (client *Client) OrderConfirm(project_id, id string, params *OrderConfirmParams) (*TranslationOrder, error) {
	retVal := new(TranslationOrder)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/orders/%s/confirm", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new order. Access token scope must include <code>orders.create</code>.
func (client *Client) OrderCreate(project_id string, params *TranslationOrderParams) (*TranslationOrder, error) {
	retVal := new(TranslationOrder)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/orders", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type OrderDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *OrderDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *OrderDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Cancel an existing order. Must not yet be confirmed.
func (client *Client) OrderDelete(project_id, id string, params *OrderDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/orders/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type OrderShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *OrderShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *OrderShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single order.
func (client *Client) OrderShow(project_id, id string, params *OrderShowParams) (*TranslationOrder, error) {
	retVal := new(TranslationOrder)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/orders/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type OrdersListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *OrdersListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *OrdersListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all orders for the given project.
func (client *Client) OrdersList(project_id string, page, perPage int, params *OrdersListParams) ([]*TranslationOrder, error) {
	retVal := []*TranslationOrder{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/orders", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new project.
func (client *Client) ProjectCreate(params *ProjectParams) (*ProjectDetails, error) {
	retVal := new(ProjectDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects")

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.AccountID != nil {
			err := writer.WriteField("account_id", *params.AccountID)
			if err != nil {
				return err
			}
		}

		if params.MainFormat != nil {
			err := writer.WriteField("main_format", *params.MainFormat)
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}

		if params.ProjectImage != nil {
			part, err := writer.CreateFormFile("project_image", filepath.Base(*params.ProjectImage))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.ProjectImage)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.RemoveProjectImage != nil {
			err := writer.WriteField("remove_project_image", strconv.FormatBool(*params.RemoveProjectImage))
			if err != nil {
				return err
			}
		}

		if params.SharesTranslationMemory != nil {
			err := writer.WriteField("shares_translation_memory", strconv.FormatBool(*params.SharesTranslationMemory))
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("POST", url, ctype, paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing project.
func (client *Client) ProjectDelete(id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s", url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single project.
func (client *Client) ProjectShow(id string) (*ProjectDetails, error) {
	retVal := new(ProjectDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s", url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing project.
func (client *Client) ProjectUpdate(id string, params *ProjectParams) (*ProjectDetails, error) {
	retVal := new(ProjectDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s", url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.AccountID != nil {
			err := writer.WriteField("account_id", *params.AccountID)
			if err != nil {
				return err
			}
		}

		if params.MainFormat != nil {
			err := writer.WriteField("main_format", *params.MainFormat)
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}

		if params.ProjectImage != nil {
			part, err := writer.CreateFormFile("project_image", filepath.Base(*params.ProjectImage))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.ProjectImage)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.RemoveProjectImage != nil {
			err := writer.WriteField("remove_project_image", strconv.FormatBool(*params.RemoveProjectImage))
			if err != nil {
				return err
			}
		}

		if params.SharesTranslationMemory != nil {
			err := writer.WriteField("shares_translation_memory", strconv.FormatBool(*params.SharesTranslationMemory))
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("PATCH", url, ctype, paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all projects the current user has access to.
func (client *Client) ProjectsList(page, perPage int) ([]*Project, error) {
	retVal := []*Project{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects")

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new release.
func (client *Client) ReleaseCreate(account_id, distribution_id string, params *ReleasesParams) (*Release, error) {
	retVal := new(Release)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases", url.QueryEscape(account_id), url.QueryEscape(distribution_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing release.
func (client *Client) ReleaseDelete(account_id, distribution_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases/%s", url.QueryEscape(account_id), url.QueryEscape(distribution_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Publish a release for production.
func (client *Client) ReleasePublish(account_id, distribution_id, id string) (*Release, error) {
	retVal := new(Release)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases/%s/publish", url.QueryEscape(account_id), url.QueryEscape(distribution_id), url.QueryEscape(id))

		rc, err := client.sendRequest("POST", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get details on a single release.
func (client *Client) ReleaseShow(account_id, distribution_id, id string) (*Release, error) {
	retVal := new(Release)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases/%s", url.QueryEscape(account_id), url.QueryEscape(distribution_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing release.
func (client *Client) ReleaseUpdate(account_id, distribution_id, id string, params *ReleasesParams) (*Release, error) {
	retVal := new(Release)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases/%s", url.QueryEscape(account_id), url.QueryEscape(distribution_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all releases for the given distribution.
func (client *Client) ReleasesList(account_id, distribution_id string, page, perPage int) ([]*ReleasePreview, error) {
	retVal := []*ReleasePreview{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/distributions/%s/releases", url.QueryEscape(account_id), url.QueryEscape(distribution_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new screenshot.
func (client *Client) ScreenshotCreate(project_id string, params *ScreenshotParams) (*Screenshot, error) {
	retVal := new(Screenshot)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.Description != nil {
			err := writer.WriteField("description", *params.Description)
			if err != nil {
				return err
			}
		}

		if params.Filename != nil {
			part, err := writer.CreateFormFile("filename", filepath.Base(*params.Filename))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.Filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("POST", url, ctype, paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing screenshot.
func (client *Client) ScreenshotDelete(project_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single screenshot for a given project.
func (client *Client) ScreenshotShow(project_id, id string) (*Screenshot, error) {
	retVal := new(Screenshot)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing screenshot.
func (client *Client) ScreenshotUpdate(project_id, id string, params *ScreenshotParams) (*Screenshot, error) {
	retVal := new(Screenshot)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.Description != nil {
			err := writer.WriteField("description", *params.Description)
			if err != nil {
				return err
			}
		}

		if params.Filename != nil {
			part, err := writer.CreateFormFile("filename", filepath.Base(*params.Filename))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.Filename)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.Name != nil {
			err := writer.WriteField("name", *params.Name)
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("PATCH", url, ctype, paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new screenshot marker.
func (client *Client) ScreenshotMarkerCreate(project_id, screenshot_id string, params *ScreenshotMarkerParams) (*ScreenshotMarker, error) {
	retVal := new(ScreenshotMarker)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s/markers", url.QueryEscape(project_id), url.QueryEscape(screenshot_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing screenshot marker.
func (client *Client) ScreenshotMarkerDelete(project_id, screenshot_id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s/markers", url.QueryEscape(project_id), url.QueryEscape(screenshot_id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single screenshot marker for a given project.
func (client *Client) ScreenshotMarkerShow(project_id, screenshot_id, id string) (*ScreenshotMarker, error) {
	retVal := new(ScreenshotMarker)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s/markers/%s", url.QueryEscape(project_id), url.QueryEscape(screenshot_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing screenshot marker.
func (client *Client) ScreenshotMarkerUpdate(project_id, screenshot_id string, params *ScreenshotMarkerParams) (*ScreenshotMarker, error) {
	retVal := new(ScreenshotMarker)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s/markers", url.QueryEscape(project_id), url.QueryEscape(screenshot_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all screenshot markers for the given project.
func (client *Client) ScreenshotMarkersList(project_id, id string, page, perPage int) ([]*ScreenshotMarker, error) {
	retVal := []*ScreenshotMarker{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots/%s/markers", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all screenshots for the given project.
func (client *Client) ScreenshotsList(project_id string, page, perPage int) ([]*Screenshot, error) {
	retVal := []*Screenshot{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/screenshots", url.QueryEscape(project_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Show details for current User.
func (client *Client) ShowUser() (*User, error) {
	retVal := new(User)
	err := func() error {

		url := fmt.Sprintf("/v2/user")

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type SpaceCreateParams struct {
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *SpaceCreateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *SpaceCreateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

// Create a new Space.
func (client *Client) SpaceCreate(account_id string, params *SpaceCreateParams) (*Space, error) {
	retVal := new(Space)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces", url.QueryEscape(account_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete the specified Space.
func (client *Client) SpaceDelete(account_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Show the specified Space.
func (client *Client) SpaceShow(account_id, id string) (*Space, error) {
	retVal := new(Space)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type SpaceUpdateParams struct {
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *SpaceUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "name":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Name = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *SpaceUpdateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Name != nil && *params.Name != "" {
		queryParams["name"] = *params.Name
	}

	return queryParams
}

// Update the specified Space.
func (client *Client) SpaceUpdate(account_id, id string, params *SpaceUpdateParams) (*Space, error) {
	retVal := new(Space)
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s", url.QueryEscape(account_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all Spaces for the given account.
func (client *Client) SpacesList(account_id string, page, perPage int) ([]*Space, error) {
	retVal := []*Space{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces", url.QueryEscape(account_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type SpacesProjectsCreateParams struct {
	ID *string `json:"id,omitempty"  cli:"opt --id"`
}

func (params *SpacesProjectsCreateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "id":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.ID = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *SpacesProjectsCreateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.ID != nil && *params.ID != "" {
		queryParams["id"] = *params.ID
	}

	return queryParams
}

// Adds an existing project to the space.
func (client *Client) SpacesProjectsCreate(account_id, space_id string, params *SpacesProjectsCreateParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s/projects", url.QueryEscape(account_id), url.QueryEscape(space_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Removes a specified project from the specified space.
func (client *Client) SpacesProjectsDelete(account_id, space_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s/projects/%s", url.QueryEscape(account_id), url.QueryEscape(space_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// List all projects for the specified Space.
func (client *Client) SpacesProjectsList(account_id, space_id string, page, perPage int) ([]*Project, error) {
	retVal := []*Project{}
	err := func() error {

		url := fmt.Sprintf("/v2/accounts/%s/spaces/%s/projects", url.QueryEscape(account_id), url.QueryEscape(space_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new style guide.
func (client *Client) StyleguideCreate(project_id string, params *StyleguideParams) (*StyleguideDetails, error) {
	retVal := new(StyleguideDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/styleguides", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing style guide.
func (client *Client) StyleguideDelete(project_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single style guide.
func (client *Client) StyleguideShow(project_id, id string) (*StyleguideDetails, error) {
	retVal := new(StyleguideDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Update an existing style guide.
func (client *Client) StyleguideUpdate(project_id, id string, params *StyleguideParams) (*StyleguideDetails, error) {
	retVal := new(StyleguideDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all styleguides for the given project.
func (client *Client) StyleguidesList(project_id string, page, perPage int) ([]*Styleguide, error) {
	retVal := []*Styleguide{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/styleguides", url.QueryEscape(project_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new tag.
func (client *Client) TagCreate(project_id string, params *TagParams) (*TagWithStats, error) {
	retVal := new(TagWithStats)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/tags", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TagDeleteParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TagDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TagDeleteParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Delete an existing tag.
func (client *Client) TagDelete(project_id, name string, params *TagDeleteParams) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/tags/%s", url.QueryEscape(project_id), url.QueryEscape(name))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("DELETE", url, "application/json", paramsBuf, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type TagShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TagShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TagShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details and progress information on a single tag for a given project.
func (client *Client) TagShow(project_id, name string, params *TagShowParams) (*TagWithStats, error) {
	retVal := new(TagWithStats)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/tags/%s", url.QueryEscape(project_id), url.QueryEscape(name))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TagsListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TagsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TagsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all tags for the given project.
func (client *Client) TagsList(project_id string, page, perPage int, params *TagsListParams) ([]*Tag, error) {
	retVal := []*Tag{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/tags", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a translation.
func (client *Client) TranslationCreate(project_id string, params *TranslationParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationExcludeParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationExcludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationExcludeParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Set exclude from export flag on an existing translation.
func (client *Client) TranslationExclude(project_id, id string, params *TranslationExcludeParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/exclude", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationIncludeParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationIncludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationIncludeParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Remove exclude from export flag from an existing translation.
func (client *Client) TranslationInclude(project_id, id string, params *TranslationIncludeParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/include", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationReviewParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationReviewParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationReviewParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark an existing translation as reviewed.
func (client *Client) TranslationReview(project_id, id string, params *TranslationReviewParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/review", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single translation.
func (client *Client) TranslationShow(project_id, id string, params *TranslationShowParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationUnverifyParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationUnverifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationUnverifyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Mark an existing translation as unverified.
func (client *Client) TranslationUnverify(project_id, id string, params *TranslationUnverifyParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/unverify", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationUpdateParams struct {
	Autotranslate *bool   `json:"autotranslate,omitempty"  cli:"opt --autotranslate"`
	Branch        *string `json:"branch,omitempty"  cli:"opt --branch"`
	Content       *string `json:"content,omitempty"  cli:"opt --content"`
	Excluded      *bool   `json:"excluded,omitempty"  cli:"opt --excluded"`
	PluralSuffix  *string `json:"plural_suffix,omitempty"  cli:"opt --plural-suffix"`
	Unverified    *bool   `json:"unverified,omitempty"  cli:"opt --unverified"`
}

func (params *TranslationUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "autotranslate":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Autotranslate = &val

		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "content":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Content = &val

		case "excluded":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Excluded = &val

		case "plural_suffix":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.PluralSuffix = &val

		case "unverified":
			val, ok := v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Unverified = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationUpdateParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Autotranslate != nil {
		queryParams["autotranslate"] = strconv.FormatBool(*params.Autotranslate)
	}

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Content != nil && *params.Content != "" {
		queryParams["content"] = *params.Content
	}

	if params.Excluded != nil {
		queryParams["excluded"] = strconv.FormatBool(*params.Excluded)
	}

	if params.PluralSuffix != nil && *params.PluralSuffix != "" {
		queryParams["plural_suffix"] = *params.PluralSuffix
	}

	if params.Unverified != nil {
		queryParams["unverified"] = strconv.FormatBool(*params.Unverified)
	}

	return queryParams
}

// Update an existing translation.
func (client *Client) TranslationUpdate(project_id, id string, params *TranslationUpdateParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationVerifyParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *TranslationVerifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationVerifyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Verify an existing translation.
func (client *Client) TranslationVerify(project_id, id string, params *TranslationVerifyParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/verify", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsByKeyParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsByKeyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsByKeyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// List translations for a specific key.
func (client *Client) TranslationsByKey(project_id, key_id string, page, perPage int, params *TranslationsByKeyParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/keys/%s/translations", url.QueryEscape(project_id), url.QueryEscape(key_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsByLocaleParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsByLocaleParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsByLocaleParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsByLocale(project_id, locale_id string, page, perPage int, params *TranslationsByLocaleParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/locales/%s/translations", url.QueryEscape(project_id), url.QueryEscape(locale_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsExcludeParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsExcludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsExcludeParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Exclude translations matching query from locale export.
func (client *Client) TranslationsExclude(project_id string, params *TranslationsExcludeParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/exclude", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsIncludeParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsIncludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsIncludeParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Include translations matching query in locale export.
func (client *Client) TranslationsInclude(project_id string, params *TranslationsIncludeParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/include", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsList(project_id string, page, perPage int, params *TranslationsListParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsReviewParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
}

func (params *TranslationsReviewParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsReviewParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	return queryParams
}

// Review translations matching query.
func (client *Client) TranslationsReview(project_id string, params *TranslationsReviewParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/review", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsSearchParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsSearchParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsSearchParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Search translations for the given project. Provides the same search interface as <code>translations#index</code> but allows POST requests to avoid limitations imposed by GET requests. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsSearch(project_id string, page, perPage int, params *TranslationsSearchParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/search", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("POST", url, "application/json", paramsBuf, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsUnverifyParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsUnverifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsUnverifyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Mark translations matching query as unverified.
func (client *Client) TranslationsUnverify(project_id string, params *TranslationsUnverifyParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/unverify", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsVerifyParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
	Order  *string `json:"order,omitempty"  cli:"opt --order"`
	Q      *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort   *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsVerifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		case "order":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Order = &val

		case "q":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Q = &val

		case "sort":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Sort = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *TranslationsVerifyParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	if params.Order != nil && *params.Order != "" {
		queryParams["order"] = *params.Order
	}

	if params.Q != nil && *params.Q != "" {
		queryParams["q"] = *params.Q
	}

	if params.Sort != nil && *params.Sort != "" {
		queryParams["sort"] = *params.Sort
	}

	return queryParams
}

// Verify translations matching query.
func (client *Client) TranslationsVerify(project_id string, params *TranslationsVerifyParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/verify", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Upload a new language file. Creates necessary resources in your project.
func (client *Client) UploadCreate(project_id string, params *UploadParams) (*Upload, error) {
	retVal := new(Upload)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/uploads", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

		if params.Autotranslate != nil {
			err := writer.WriteField("autotranslate", strconv.FormatBool(*params.Autotranslate))
			if err != nil {
				return err
			}
		}

		if params.Branch != nil {
			err := writer.WriteField("branch", *params.Branch)
			if err != nil {
				return err
			}
		}

		if params.ConvertEmoji != nil {
			err := writer.WriteField("convert_emoji", strconv.FormatBool(*params.ConvertEmoji))
			if err != nil {
				return err
			}
		}

		if params.File != nil {
			part, err := writer.CreateFormFile("file", filepath.Base(*params.File))
			if err != nil {
				return err
			}
			file, err := os.Open(*params.File)
			if err != nil {
				return err
			}
			_, err = io.Copy(part, file)
			if err != nil {
				return err
			}
			err = file.Close()
			if err != nil {
				return err
			}
		}

		if params.FileEncoding != nil {
			err := writer.WriteField("file_encoding", *params.FileEncoding)
			if err != nil {
				return err
			}
		}

		if params.FileFormat != nil {
			err := writer.WriteField("file_format", *params.FileFormat)
			if err != nil {
				return err
			}
		}

		if params.FormatOptions != nil {
			for key, val := range params.FormatOptions {
				err := writer.WriteField("format_options["+key+"]", val)
				if err != nil {
					return err
				}
			}
		}

		if params.LocaleID != nil {
			err := writer.WriteField("locale_id", *params.LocaleID)
			if err != nil {
				return err
			}
		}

		if params.LocaleMapping != nil {
			for key, val := range params.LocaleMapping {
				err := writer.WriteField("locale_mapping["+key+"]", val)
				if err != nil {
					return err
				}
			}
		}

		if params.MarkReviewed != nil {
			err := writer.WriteField("mark_reviewed", strconv.FormatBool(*params.MarkReviewed))
			if err != nil {
				return err
			}
		}

		if params.SkipUnverification != nil {
			err := writer.WriteField("skip_unverification", strconv.FormatBool(*params.SkipUnverification))
			if err != nil {
				return err
			}
		}

		if params.SkipUploadTags != nil {
			err := writer.WriteField("skip_upload_tags", strconv.FormatBool(*params.SkipUploadTags))
			if err != nil {
				return err
			}
		}

		if params.Tags != nil {
			err := writer.WriteField("tags", *params.Tags)
			if err != nil {
				return err
			}
		}

		if params.UpdateDescriptions != nil {
			err := writer.WriteField("update_descriptions", strconv.FormatBool(*params.UpdateDescriptions))
			if err != nil {
				return err
			}
		}

		if params.UpdateTranslations != nil {
			err := writer.WriteField("update_translations", strconv.FormatBool(*params.UpdateTranslations))
			if err != nil {
				return err
			}
		}
		err := writer.WriteField("utf8", "✓")
		writer.Close()

		rc, err := client.sendRequest("POST", url, ctype, paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type UploadShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *UploadShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *UploadShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// View details and summary for a single upload.
func (client *Client) UploadShow(project_id, id string, params *UploadShowParams) (*Upload, error) {
	retVal := new(Upload)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/uploads/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type UploadsListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *UploadsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *UploadsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all uploads for the given project.
func (client *Client) UploadsList(project_id string, page, perPage int, params *UploadsListParams) ([]*Upload, error) {
	retVal := []*Upload{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/uploads", url.QueryEscape(project_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type VersionShowParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *VersionShowParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *VersionShowParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// Get details on a single version.
func (client *Client) VersionShow(project_id, translation_id, id string, params *VersionShowParams) (*TranslationVersionWithUser, error) {
	retVal := new(TranslationVersionWithUser)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/versions/%s", url.QueryEscape(project_id), url.QueryEscape(translation_id), url.QueryEscape(id))

		rc, err := client.sendGetRequest(url, params.QueryParams(), 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type VersionsListParams struct {
	Branch *string `json:"branch,omitempty"  cli:"opt --branch"`
}

func (params *VersionsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "branch":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Branch = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

func (params *VersionsListParams) QueryParams() map[string]string {
	var queryParams = make(map[string]string, 0)

	if params.Branch != nil && *params.Branch != "" {
		queryParams["branch"] = *params.Branch
	}

	return queryParams
}

// List all versions for the given translation.
func (client *Client) VersionsList(project_id, translation_id string, page, perPage int, params *VersionsListParams) ([]*TranslationVersion, error) {
	retVal := []*TranslationVersion{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/translations/%s/versions", url.QueryEscape(project_id), url.QueryEscape(translation_id))

		rc, err := client.sendGetRequestPaginated(url, params.QueryParams(), 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Create a new webhook.
func (client *Client) WebhookCreate(project_id string, params *WebhookParams) (*Webhook, error) {
	retVal := new(Webhook)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks", url.QueryEscape(project_id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("POST", url, "application/json", paramsBuf, 201)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing webhook.
func (client *Client) WebhookDelete(project_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single webhook.
func (client *Client) WebhookShow(project_id, id string) (*Webhook, error) {
	retVal := new(Webhook)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("GET", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Perform a test request for a webhook.
func (client *Client) WebhookTest(project_id, id string) error {

	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s/test", url.QueryEscape(project_id), url.QueryEscape(id))

		rc, err := client.sendRequest("POST", url, "", nil, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Update an existing webhook.
func (client *Client) WebhookUpdate(project_id, id string, params *WebhookParams) (*Webhook, error) {
	retVal := new(Webhook)
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", url.QueryEscape(project_id), url.QueryEscape(id))

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("PATCH", url, "application/json", paramsBuf, 200)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all webhooks for the given project.
func (client *Client) WebhooksList(project_id string, page, perPage int) ([]*Webhook, error) {
	retVal := []*Webhook{}
	err := func() error {

		url := fmt.Sprintf("/v2/projects/%s/webhooks", url.QueryEscape(project_id))

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)

		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if client.debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

func GetUserAgent() string {
	agent := "PhraseApp go (" + ClientVersion + ")"
	if ua := os.Getenv("PHRASEAPP_USER_AGENT"); ua != "" {
		agent = ua + "; " + agent
	}
	return agent
}

func convertMapToQueryParams(paramName string, params map[string]string) map[string]string {
	queryParams := make(map[string]string)
	for key, value := range params {
		paramName := fmt.Sprintf("%s[%s]", paramName, key)
		queryParams[paramName] = value
	}

	return queryParams
}
