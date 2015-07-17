package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"bytes"

	"github.com/mgutz/ansi"
	"github.com/phrase/phraseapp-go/phraseapp"
)

type Locales []*Locale
type Locale struct {
	Path, Name, Id, RFC, Tag, FileFormat string
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
	pc.Parts = splitToParts(userPath, pc.Separator)

	isDir, err := isDir(userPath)
	if err != nil {
		return nil, err
	}

	pc.IsDir = isDir

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

func splitToParts(userPath, separator string) []string {
	split := strings.Split(userPath, separator)
	parts := []string{}
	for _, part := range split {
		if part != separator {
			parts = append(parts, part)
		}
	}
	return parts
}

// Locale to Path mapping
func CopyLocale(relPath string, locale *Locale) *Locale {
	newLocale := &Locale{Path: relPath, Id: locale.Id, Name: locale.Name, Tag: locale.Tag, FileFormat: locale.FileFormat}
	return newLocale
}

func (locale *Locale) Message() string {
	str := ""
	if locale.Name != "" {
		str = fmt.Sprintf("Name: %s", locale.Name)
	}
	if locale.Id != "" {
		str = fmt.Sprintf("%s Id: %s", str, locale.Id)
	}
	if locale.RFC != "" {
		str = fmt.Sprintf("%s RFC5646: %s", str, locale.RFC)
	}
	if locale.Tag != "" {
		str = fmt.Sprintf("%s Tag: %s", str, locale.Tag)
	}
	if locale.FileFormat != "" {
		str = fmt.Sprintf("%s Format: %s", str, locale.FileFormat)
	}
	return strings.TrimSpace(str)
}

// Locale placeholder logic <locale_name>
func (pc *PathComponents) ExpandPathsWithLocale(locales []*phraseapp.Locale, locale *Locale) (Locales, error) {
	return pc.pathsForRemoteLocales(locales, locale)
}

func (pc *PathComponents) singlePathWithLocale(locales []*phraseapp.Locale, localeFile *Locale) (Locales, error) {
	newPaths := []*Locale{}
	newLocaleFile := localeForLocaleId(localeFile.Id, locales)
	valid, err := pc.isValidLocale(newLocaleFile)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("Could not find remote locale with Id:", localeFile.Id)
	}
	locale := &Locale{Id: localeFile.Id, Name: localeFile.Name, Tag: localeFile.Tag, FileFormat: localeFile.FileFormat}
	absPath, err := pc.filePath(locale)
	if err != nil {
		return nil, err
	}
	locale.Path = absPath
	newPaths = append(newPaths, locale)
	return newPaths, nil
}

func (pc *PathComponents) pathsForRemoteLocales(locales []*phraseapp.Locale, info *Locale) (Locales, error) {
	files := []*Locale{}
	for _, remoteLocale := range locales {
		valid, err := pc.isValidLocale(remoteLocale)
		if err != nil {
			fmt.Println(err.Error())
		}
		if !valid {
			continue
		}

		locale := &Locale{Name: remoteLocale.Name, Id: remoteLocale.Id, RFC: remoteLocale.Code, Tag: info.Tag, FileFormat: info.FileFormat}
		absPath, err := pc.filePath(locale)
		if err != nil {
			return nil, err
		}
		locale.Path = absPath
		files = append(files, locale)
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

func (pc *PathComponents) filePath(locale *Locale) (string, error) {
	absPath, err := filepath.Abs(pc.Path)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", locale.Name, -1)
	path = strings.Replace(path, "<locale_code>", locale.RFC, -1)
	path = strings.Replace(path, "<tag>", locale.Tag, -1)

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

func sharedMessage(method string, locale *Locale) {
	yellow := ansi.ColorCode("yellow+b:black")
	reset := ansi.ColorCode("reset")

	localPath := locale.Path
	callerPath, err := os.Getwd()
	if err == nil {
		localPath = "." + strings.Replace(locale.Path, callerPath, "", 1)
	}

	remote := fmt.Sprint(yellow, locale.Message(), reset)
	local := fmt.Sprint(yellow, localPath, reset)

	from, to := "", ""
	if method == "pull" {
		from, to = remote, local
	} else {
		from, to = local, remote
	}

	var buffer bytes.Buffer
	buffer.WriteString("From: ")
	buffer.WriteString(from)
	buffer.WriteString(" To: ")
	buffer.WriteString(to)
	fmt.Println(buffer.String())
}

func contains(pathes []string, str string) bool {
	for _, item := range pathes {
		if str == item {
			return true
		}
	}
	return false
}

func RemoteLocales(projectId string) ([]*phraseapp.Locale, error) {
	locales, err := phraseapp.LocalesList(projectId, 1, 25)
	if err != nil {
		return nil, err
	}
	result := locales
	for len(locales) == 25 {
		locales, err = phraseapp.LocalesList(projectId, 1, 25)
		if err != nil {
			return nil, err
		}
		result = append(result, locales...)
	}
	return locales, nil
}
