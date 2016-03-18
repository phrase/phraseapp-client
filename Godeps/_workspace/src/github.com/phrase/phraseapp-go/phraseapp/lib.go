package phraseapp

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path/filepath"
	"strconv"
	"time"
)

const (
	RevisionDocs      = "7a0b5ac5b1cf1958754347da2d6c6ce0bd564564"
	RevisionGenerator = "edf10b6ff26ed4eca0a57c138ab168f0665bee5a"
)

type Account struct {
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type AccountPreview struct {
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
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

type BlacklistedKey struct {
	CreatedAt *time.Time `json:"created_at"`
	ID        string     `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
}

type Comment struct {
	CreatedAt *time.Time   `json:"created_at"`
	ID        string       `json:"id"`
	Message   string       `json:"message"`
	UpdatedAt *time.Time   `json:"updated_at"`
	User      *UserPreview `json:"user"`
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

type Project struct {
	Account    *AccountPreview `json:"account"`
	CreatedAt  *time.Time      `json:"created_at"`
	ID         string          `json:"id"`
	MainFormat string          `json:"main_format"`
	Name       string          `json:"name"`
	UpdatedAt  *time.Time      `json:"updated_at"`
}

type ProjectDetails struct {
	Project

	SharesTranslationMemory bool `json:"shares_translation_memory"`
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
	LocalesCreated         int64 `json:"locales_created"`
	TagsCreated            int64 `json:"tags_created"`
	TranslationKeysCreated int64 `json:"translation_keys_created"`
	TranslationsCreated    int64 `json:"translations_created"`
	TranslationsUpdated    int64 `json:"translations_updated"`
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

	CommentsCount        int64  `json:"comments_count"`
	FormatValueType      string `json:"format_value_type"`
	MaxCharactersAllowed int64  `json:"max_characters_allowed"`
	NamePlural           string `json:"name_plural"`
	OriginalFile         string `json:"original_file"`
	ScreenshotUrl        string `json:"screenshot_url"`
	Unformatted          bool   `json:"unformatted"`
	XmlSpacePreserve     bool   `json:"xml_space_preserve"`
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

type CommentParams struct {
	Message *string `json:"message,omitempty"  cli:"opt --message"`
}

func (params *CommentParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

type TranslationKeyParams struct {
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

type LocaleParams struct {
	Code           *string `json:"code,omitempty"  cli:"opt --code"`
	Default        *bool   `json:"default,omitempty"  cli:"opt --default"`
	Main           *bool   `json:"main,omitempty"  cli:"opt --main"`
	Name           *string `json:"name,omitempty"  cli:"opt --name"`
	Rtl            *bool   `json:"rtl,omitempty"  cli:"opt --rtl"`
	SourceLocaleID *string `json:"source_locale_id,omitempty"  cli:"opt --source-locale-id"`
}

func (params *LocaleParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

type TranslationOrderParams struct {
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

type ProjectParams struct {
	AccountID               *string `json:"account_id,omitempty"  cli:"opt --account-id"`
	MainFormat              *string `json:"main_format,omitempty"  cli:"opt --main-format"`
	Name                    *string `json:"name,omitempty"  cli:"opt --name"`
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

type TagParams struct {
	Name *string `json:"name,omitempty"  cli:"opt --name"`
}

func (params *TagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
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

type TranslationParams struct {
	Content      *string `json:"content,omitempty"  cli:"opt --content"`
	Excluded     *bool   `json:"excluded,omitempty"  cli:"opt --excluded"`
	KeyID        *string `json:"key_id,omitempty"  cli:"opt --key-id"`
	LocaleID     *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	PluralSuffix *string `json:"plural_suffix,omitempty"  cli:"opt --plural-suffix"`
	Unverified   *bool   `json:"unverified,omitempty"  cli:"opt --unverified"`
}

func (params *TranslationParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

type UploadParams struct {
	ConvertEmoji       *bool             `json:"convert_emoji,omitempty"  cli:"opt --convert-emoji"`
	File               *string           `json:"file,omitempty"  cli:"opt --file"`
	FileEncoding       *string           `json:"file_encoding,omitempty"  cli:"opt --file-encoding"`
	FileFormat         *string           `json:"file_format,omitempty"  cli:"opt --file-format"`
	FormatOptions      map[string]string `json:"format_options,omitempty"  cli:"opt --format-options"`
	LocaleID           *string           `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	SkipUnverification *bool             `json:"skip_unverification,omitempty"  cli:"opt --skip-unverification"`
	SkipUploadTags     *bool             `json:"skip_upload_tags,omitempty"  cli:"opt --skip-upload-tags"`
	Tags               *string           `json:"tags,omitempty"  cli:"opt --tags"`
	UpdateTranslations *bool             `json:"update_translations,omitempty"  cli:"opt --update-translations"`
}

func (params *UploadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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
		if Debug {
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
		url := fmt.Sprintf("/v2/authorizations/%s", id)

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
		url := fmt.Sprintf("/v2/authorizations/%s", id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/authorizations/%s", id)

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
		if Debug {
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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys", project_id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", project_id, id)

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
		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys/%s", project_id, id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/blacklisted_keys", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments", project_id, key_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing comment.
func (client *Client) CommentDelete(project_id, key_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", project_id, key_id, id)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Check if comment was marked as read. Returns 204 if read, 404 if unread.
func (client *Client) CommentMarkCheck(project_id, key_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", project_id, key_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Mark a comment as read.
func (client *Client) CommentMarkRead(project_id, key_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", project_id, key_id, id)

		rc, err := client.sendRequest("PATCH", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Mark a comment as unread.
func (client *Client) CommentMarkUnread(project_id, key_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s/read", project_id, key_id, id)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single comment.
func (client *Client) CommentShow(project_id, key_id, id string) (*Comment, error) {
	retVal := new(Comment)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", project_id, key_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments/%s", project_id, key_id, id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all comments for a key.
func (client *Client) CommentsList(project_id, key_id string, page, perPage int) ([]*Comment, error) {
	retVal := []*Comment{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/comments", project_id, key_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get a handy list of all localization file formats supported in PhraseApp.
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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/keys", project_id)

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing key.
func (client *Client) KeyDelete(project_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s", project_id, id)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single key for a given project.
func (client *Client) KeyShow(project_id, id string) (*TranslationKeyDetails, error) {
	retVal := new(TranslationKeyDetails)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/keys/%s", project_id, id)

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysDeleteParams struct {
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
}

func (params *KeysDeleteParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Delete all keys matching query. Same constraints as list. Please limit the number of affected keys to about 1,000 as you might experience timeouts otherwise.
func (client *Client) KeysDelete(project_id string, params *KeysDeleteParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysListParams struct {
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Order    *string `json:"order,omitempty"  cli:"opt --order"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort     *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *KeysListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// List all keys for the given project. Alternatively you can POST requests to /search.
func (client *Client) KeysList(project_id string, page, perPage int, params *KeysListParams) ([]*TranslationKey, error) {
	retVal := []*TranslationKey{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys", project_id)

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("GET", url, "application/json", paramsBuf, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysSearchParams struct {
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Order    *string `json:"order,omitempty"  cli:"opt --order"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort     *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *KeysSearchParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Search keys for the given project matching query.
func (client *Client) KeysSearch(project_id string, page, perPage int, params *KeysSearchParams) ([]*TranslationKey, error) {
	retVal := []*TranslationKey{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/search", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysTagParams struct {
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Tags     *string `json:"tags,omitempty"  cli:"opt --tags"`
}

func (params *KeysTagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Tags all keys matching query. Same constraints as list.
func (client *Client) KeysTag(project_id string, params *KeysTagParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/tag", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type KeysUntagParams struct {
	LocaleID *string `json:"locale_id,omitempty"  cli:"opt --locale-id"`
	Q        *string `json:"q,omitempty"  cli:"opt --query -q"`
	Tags     *string `json:"tags,omitempty"  cli:"opt --tags"`
}

func (params *KeysUntagParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Removes specified tags from keys matching query.
func (client *Client) KeysUntag(project_id string, params *KeysUntagParams) (*AffectedResources, error) {
	retVal := new(AffectedResources)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/untag", project_id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/locales", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing locale.
func (client *Client) LocaleDelete(project_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/locales/%s", project_id, id)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

type LocaleDownloadParams struct {
	ConvertEmoji               bool              `json:"convert_emoji,omitempty"  cli:"opt --convert-emoji"`
	Encoding                   *string           `json:"encoding,omitempty"  cli:"opt --encoding"`
	FallbackLocaleID           *string           `json:"fallback_locale_id,omitempty"  cli:"opt --fallback-locale-id"`
	FileFormat                 *string           `json:"file_format,omitempty"  cli:"opt --file-format"`
	FormatOptions              map[string]string `json:"format_options,omitempty"  cli:"opt --format-options"`
	IncludeEmptyTranslations   bool              `json:"include_empty_translations,omitempty"  cli:"opt --include-empty-translations"`
	KeepNotranslateTags        bool              `json:"keep_notranslate_tags,omitempty"  cli:"opt --keep-notranslate-tags"`
	SkipUnverifiedTranslations bool              `json:"skip_unverified_translations,omitempty"  cli:"opt --skip-unverified-translations"`
	Tag                        *string           `json:"tag,omitempty"  cli:"opt --tag"`
}

func (params *LocaleDownloadParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
		case "convert_emoji":
			ok := false
			params.ConvertEmoji, ok = v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
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
			ok := false
			params.IncludeEmptyTranslations, ok = v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "keep_notranslate_tags":
			ok := false
			params.KeepNotranslateTags, ok = v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "skip_unverified_translations":
			ok := false
			params.SkipUnverifiedTranslations, ok = v.(bool)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
		case "tag":
			val, ok := v.(string)
			if !ok {
				return fmt.Errorf(cfgValueErrStr, k, v)
			}
			params.Tag = &val

		default:
			return fmt.Errorf(cfgInvalidKeyErrStr, k)
		}
	}

	return nil
}

// Download a locale in a specific file format.
func (client *Client) LocaleDownload(project_id, id string, params *LocaleDownloadParams) ([]byte, error) {
	retVal := []byte{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/locales/%s/download", project_id, id)

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequest("GET", url, "application/json", paramsBuf, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		retVal, err = ioutil.ReadAll(reader)
		return err

	}()
	return retVal, err
}

// Get details on a single locale for a given project.
func (client *Client) LocaleShow(project_id, id string) (*LocaleDetails, error) {
	retVal := new(LocaleDetails)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/locales/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/locales/%s", project_id, id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all locales for the given project.
func (client *Client) LocalesList(project_id string, page, perPage int) ([]*Locale, error) {
	retVal := []*Locale{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/locales", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Confirm an existing order and send it to the provider for translation. Same constraints as for create.
func (client *Client) OrderConfirm(project_id, id string) (*TranslationOrder, error) {
	retVal := new(TranslationOrder)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/orders/%s/confirm", project_id, id)

		rc, err := client.sendRequest("PATCH", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/orders", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Cancel an existing order. Must not yet be confirmed.
func (client *Client) OrderDelete(project_id, id string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/orders/%s", project_id, id)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details on a single order.
func (client *Client) OrderShow(project_id, id string) (*TranslationOrder, error) {
	retVal := new(TranslationOrder)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/orders/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all orders for the given project.
func (client *Client) OrdersList(project_id string, page, perPage int) ([]*TranslationOrder, error) {
	retVal := []*TranslationOrder{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/orders", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s", id)

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
		url := fmt.Sprintf("/v2/projects/%s", id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s", id)

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
		if Debug {
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
		if Debug {
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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/styleguides", project_id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", project_id, id)

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
		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/styleguides/%s", project_id, id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/styleguides", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/tags", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Delete an existing tag.
func (client *Client) TagDelete(project_id, name string) error {

	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/tags/%s", project_id, name)

		rc, err := client.sendRequest("DELETE", url, "", nil, 204)
		if err != nil {
			return err
		}
		defer rc.Close()

		return nil
	}()
	return err
}

// Get details and progress information on a single tag for a given project.
func (client *Client) TagShow(project_id, name string) (*TagWithStats, error) {
	retVal := new(TagWithStats)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/tags/%s", project_id, name)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all tags for the given project.
func (client *Client) TagsList(project_id string, page, perPage int) ([]*Tag, error) {
	retVal := []*Tag{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/tags", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/translations", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get details on a single translation.
func (client *Client) TranslationShow(project_id, id string) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationUpdateParams struct {
	Content      *string `json:"content,omitempty"  cli:"opt --content"`
	Excluded     *bool   `json:"excluded,omitempty"  cli:"opt --excluded"`
	PluralSuffix *string `json:"plural_suffix,omitempty"  cli:"opt --plural-suffix"`
	Unverified   *bool   `json:"unverified,omitempty"  cli:"opt --unverified"`
}

func (params *TranslationUpdateParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Update an existing translation.
func (client *Client) TranslationUpdate(project_id, id string, params *TranslationUpdateParams) (*TranslationDetails, error) {
	retVal := new(TranslationDetails)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/%s", project_id, id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsByKeyParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsByKeyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// List translations for a specific key.
func (client *Client) TranslationsByKey(project_id, key_id string, page, perPage int, params *TranslationsByKeyParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/keys/%s/translations", project_id, key_id)

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("GET", url, "application/json", paramsBuf, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsByLocaleParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsByLocaleParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// List translations for a specific locale. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsByLocale(project_id, locale_id string, page, perPage int, params *TranslationsByLocaleParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/locales/%s/translations", project_id, locale_id)

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("GET", url, "application/json", paramsBuf, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsExcludeParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsExcludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Exclude translations matching query from locale export.
func (client *Client) TranslationsExclude(project_id string, params *TranslationsExcludeParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/exclude", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsIncludeParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsIncludeParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Include translations matching query in locale export.
func (client *Client) TranslationsInclude(project_id string, params *TranslationsIncludeParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/include", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsListParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsListParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// List translations for the given project. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsList(project_id string, page, perPage int, params *TranslationsListParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations", project_id)

		paramsBuf := bytes.NewBuffer(nil)
		err := json.NewEncoder(paramsBuf).Encode(&params)
		if err != nil {
			return err
		}

		rc, err := client.sendRequestPaginated("GET", url, "application/json", paramsBuf, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsSearchParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsSearchParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// List translations for the given project if you exceed GET request limitations on translations list. If you want to download all translations for one locale we recommend to use the <code>locales#download</code> endpoint.
func (client *Client) TranslationsSearch(project_id string, page, perPage int, params *TranslationsSearchParams) ([]*Translation, error) {
	retVal := []*Translation{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/search", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsUnverifyParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsUnverifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Mark translations matching query as unverified.
func (client *Client) TranslationsUnverify(project_id string, params *TranslationsUnverifyParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/unverify", project_id)

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

type TranslationsVerifyParams struct {
	Order *string `json:"order,omitempty"  cli:"opt --order"`
	Q     *string `json:"q,omitempty"  cli:"opt --query -q"`
	Sort  *string `json:"sort,omitempty"  cli:"opt --sort"`
}

func (params *TranslationsVerifyParams) ApplyValuesFromMap(defaults map[string]interface{}) error {
	for k, v := range defaults {
		switch k {
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

// Verify translations matching query.
func (client *Client) TranslationsVerify(project_id string, params *TranslationsVerifyParams) (*AffectedCount, error) {
	retVal := new(AffectedCount)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/verify", project_id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/uploads", project_id)

		paramsBuf := bytes.NewBuffer(nil)
		writer := multipart.NewWriter(paramsBuf)
		ctype := writer.FormDataContentType()

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
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// View details and summary for a single upload.
func (client *Client) UploadShow(project_id, id string) (*Upload, error) {
	retVal := new(Upload)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/uploads/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all uploads for the given project.
func (client *Client) UploadsList(project_id string, page, perPage int) ([]*Upload, error) {
	retVal := []*Upload{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/uploads", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// Get details on a single version.
func (client *Client) VersionShow(project_id, translation_id, id string) (*TranslationVersionWithUser, error) {
	retVal := new(TranslationVersionWithUser)
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/%s/versions/%s", project_id, translation_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
			reader = io.TeeReader(rc, os.Stderr)
		} else {
			reader = rc
		}

		return json.NewDecoder(reader).Decode(&retVal)

	}()
	return retVal, err
}

// List all versions for the given translation.
func (client *Client) VersionsList(project_id, translation_id string, page, perPage int) ([]*TranslationVersion, error) {
	retVal := []*TranslationVersion{}
	err := func() error {
		url := fmt.Sprintf("/v2/projects/%s/translations/%s/versions", project_id, translation_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/webhooks", project_id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", project_id, id)

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
		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", project_id, id)

		rc, err := client.sendRequest("GET", url, "", nil, 200)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s/test", project_id, id)

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
		url := fmt.Sprintf("/v2/projects/%s/webhooks/%s", project_id, id)

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
		if Debug {
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
		url := fmt.Sprintf("/v2/projects/%s/webhooks", project_id)

		rc, err := client.sendRequestPaginated("GET", url, "", nil, 200, page, perPage)
		if err != nil {
			return err
		}
		defer rc.Close()

		var reader io.Reader
		if Debug {
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
