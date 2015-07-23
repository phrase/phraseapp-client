package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/mgutz/ansi"
	"github.com/phrase/phraseapp-go/phraseapp"
)

var Debug bool

type LocaleFiles []*LocaleFile
type LocaleFile struct {
	Path, Name, Id, RFC, Tag, FileFormat string
	ExistsRemote                         bool
}

func (localeFile *LocaleFile) RelPath() string {
	callerPath, _ := os.Getwd()
	relativePath, _ := filepath.Rel(callerPath, localeFile.Path)
	return relativePath
}

// PathComponents replacement for ugly slicing string logic on paths
type PathComponents struct {
	Path        string
	Separator   string
	Parts       []string
	GlobPattern string
	IsDir       bool
}

func (pc *PathComponents) isLocalePatternUsed() bool {
	return pc.isLocaleNameInPath() || pc.isLocaleCodeInPath()
}

func (pc *PathComponents) isLocaleNameInPath() bool {
	return strings.Contains(pc.Path, "<locale_name>")
}

func (pc *PathComponents) isLocaleCodeInPath() bool {
	return strings.Contains(pc.Path, "<locale_code>")
}

func (pc *PathComponents) isTagInPath() bool {
	return strings.Contains(pc.Path, "<tag>")
}

func (pc *PathComponents) isValidLocale(locale *phraseapp.Locale) (bool, error) {
	localePresent := (locale != nil)

	if !localePresent {
		return false, fmt.Errorf("Locale not set")
	}

	if pc.isLocaleCodeInPath() && (locale.Code == "") {
		return false, fmt.Errorf("Locale code is not set for Locale with Id: %s but locale_code is used in file name", locale.Id)
	}
	return true, nil
}

func ExtractPathComponents(userPath string) (*PathComponents, error) {
	pc := &PathComponents{Separator: string(os.PathSeparator)}
	pc.GlobPattern = extractGlobPattern(userPath)
	pc.Path = strings.TrimSpace(strings.TrimSuffix(userPath, pc.GlobPattern))
	pc.Parts = strings.Split(userPath, pc.Separator)

	isDirectory, err := isDir(userPath)
	if err != nil {
		return nil, err
	}
	pc.IsDir = isDirectory
	return pc, nil
}

func isDir(path string) (bool, error) {
	if strings.Contains(path, "<") {
		return false, nil
	}

	file, err := os.Open(path)
	if err != nil {
		return false, err
	}
	defer file.Close()
	stat, err := file.Stat()
	if err != nil {
		return false, err
	}
	switch mode := stat.Mode(); {
	case mode.IsDir():
		return true, nil
	}
	return false, nil
}

func extractGlobPattern(userPath string) string {
	if strings.HasSuffix(userPath, path.Join("**", "*")) {
		return "**/*"
	} else if strings.HasSuffix(userPath, "*") {
		return "*"
	} else {
		return ""
	}
}

// Locale to Path mapping
func (localeFile *LocaleFile) Message() string {
	str := ""
	if Debug {
		if localeFile.Name != "" {
			str = fmt.Sprintf("%s Name: %s", str, localeFile.Name)
		}
		if localeFile.Id != "" {
			str = fmt.Sprintf("%s Id: %s", str, localeFile.Id)
		}
		if localeFile.RFC != "" {
			str = fmt.Sprintf("%s RFC5646: %s", str, localeFile.RFC)
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

// Locale placeholder logic <locale_name>
func (pc *PathComponents) ExpandPathsWithLocale(locales []*phraseapp.Locale, localeFile *LocaleFile) (LocaleFiles, error) {
	files := []*LocaleFile{}
	for _, remoteLocale := range locales {
		if localeFile.Id != "" && !(remoteLocale.Id == localeFile.Id || remoteLocale.Name == localeFile.Id) {
			continue
		}
		valid, err := pc.isValidLocale(remoteLocale)
		if err != nil {
			fmt.Println(err.Error())
		}
		if !valid {
			continue
		}

		localeFile := &LocaleFile{Name: remoteLocale.Name, Id: remoteLocale.Id, RFC: remoteLocale.Code, Tag: localeFile.Tag, FileFormat: localeFile.FileFormat}
		absPath, err := pc.filePath(localeFile)
		if err != nil {
			return nil, err
		}
		localeFile.Path = absPath
		files = append(files, localeFile)
	}
	return files, nil
}

// Locale logic
func localeForLocaleId(localeId string, locales []*phraseapp.Locale) *phraseapp.Locale {
	for _, locale := range locales {
		if locale.Id == localeId {
			return locale
		}
	}
	return nil
}

func (pc *PathComponents) filePath(localeFile *LocaleFile) (string, error) {
	absPath, err := filepath.Abs(pc.Path)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.RFC, -1)
	path = strings.Replace(path, "<tag>", localeFile.Tag, -1)

	return path, nil
}

func Authenticate() error {
	defaultCredentials, err := ConfigDefaultCredentials()
	if err != nil {
		return err
	}
	phraseapp.RegisterAuthCredentials(defaultCredentials, defaultCredentials)
	return nil
}

func printErr(err error, msg string) {
	red := ansi.ColorCode("red+b:black")
	reset := ansi.ColorCode("reset")
	fmt.Fprintf(os.Stderr, "%sERROR: %s %s%s\n", red, err, msg, reset)
}

func sharedMessage(method string, localeFile *LocaleFile) {
	green := ansi.ColorCode("green+b:black")
	reset := ansi.ColorCode("reset")

	local := fmt.Sprint(green, localeFile.RelPath(), reset)

	if method == "pull" {
		remote := fmt.Sprint(green, localeFile.Message(), reset)
		fmt.Println("Downloaded", remote, "to", local)
	} else {
		fmt.Println("Uploaded", local, "successfully.")
	}
}

func RemoteLocales(projectId string) ([]*phraseapp.Locale, error) {
	page := 1
	locales, err := phraseapp.LocalesList(projectId, page, 25)
	if err != nil {
		return nil, err
	}
	result := locales
	for len(locales) == 25 {
		page = page + 1
		locales, err = phraseapp.LocalesList(projectId, page, 25)
		if err != nil {
			return nil, err
		}
		result = append(result, locales...)
	}
	return result, nil
}

func Contains(seq []string, str string) bool {
	for _, elem := range seq {
		if str == elem {
			return true
		}
	}
	return false
}

func TakeWhile(seq []string, predicate func(string) bool) []string {
	take := []string{}
	for _, elem := range seq {
		if !predicate(elem) {
			break
		}
		take = append(take, elem)
	}
	return take
}
