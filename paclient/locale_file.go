package paclient

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type LocaleFiles []*LocaleFile

type LocaleFile struct {
	Path, Name, ID, Code, Tag, FileFormat string
	ExistsRemote                          bool
}

func (localeFile *LocaleFile) RelPath() string {
	callerPath, _ := os.Getwd()
	relativePath, _ := filepath.Rel(callerPath, localeFile.Path)
	return relativePath
}

func (localeFile *LocaleFile) shouldCreateLocale(source *Source) bool {
	if localeFile.ExistsRemote {
		return false
	}

	if source.Format.IncludesLocaleInformation {
		return false
	}

	// we could not find an existing locale in PhraseApp
	// if a locale_name or locale_code was provided by the placeholder logic
	// we assume that it should be created
	// every other source should be uploaded and validated in uploads#create
	return (localeFile.Name != "" || localeFile.Code != "")
}

func (localeFile *LocaleFile) extractParamFromPathToken(globber GlobFinder, patternToken, pathToken string) {
	groups := PlaceholderRegexp.FindAllString(patternToken, -1)
	if len(groups) <= 0 {
		return
	}

	match := strings.Replace(patternToken, ".", "[.]", -1)
	if strings.Contains(match, "*") {
		match = strings.Replace(match, "*", ".*", -1)
	}

	for _, group := range groups {
		replacer := fmt.Sprintf("(?P%s.+)", group)
		match = strings.Replace(match, group, replacer, 1)
	}

	if Debug {
		fmt.Println("  expanded: ", match)
	}

	if match == "" {
		return
	}

	tmpRegexp, err := regexp.Compile(match)
	if err != nil {
		return
	}

	namedMatches := tmpRegexp.SubexpNames()
	subMatches := tmpRegexp.FindStringSubmatch(pathToken)

	if Debug {
		fmt.Println("  namedMatches: ", namedMatches)
		fmt.Println("  subMatches: ", subMatches)
		fmt.Println()
	}

	for i, subMatch := range subMatches {
		value := strings.Trim(subMatch, string(globber.Separator()))
		switch namedMatches[i] {
		case "locale_code":
			localeFile.Code = value
		case "locale_name":
			localeFile.Name = value
		case "tag":
			localeFile.Tag = value
		default:
			// ignore
		}
	}
}

var Debug = os.Getenv("DEBUG") == "true" // fix me

// Locale to Path mapping
func (localeFile *LocaleFile) Message() string {
	str := ""
	if Debug {
		if localeFile.Name != "" {
			str = fmt.Sprintf("%s Name: %s", str, localeFile.Name)
		}
		if localeFile.ID != "" {
			str = fmt.Sprintf("%s Id: %s", str, localeFile.ID)
		}
		if localeFile.Code != "" {
			str = fmt.Sprintf("%s Code: %s", str, localeFile.Code)
		}
		if localeFile.Tag != "" {
			str = fmt.Sprintf("%s Tag: %s", str, localeFile.Tag)
		}
		if localeFile.FileFormat != "" {
			str = fmt.Sprintf("%s Format: %s", str, localeFile.FileFormat)
		}
	} else {
		str = fmt.Sprintf("%s", localeFile.Name)
	}
	return strings.TrimSpace(str)
}
