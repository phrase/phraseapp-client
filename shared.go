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

// PhrasePath replacement for ugly slicing string logic on paths
type PhrasePath struct {
	UserPath    string
	Separator   string
	Components  []string
	GlobPattern string
	IsDir       bool
}

func (p *PhrasePath) RelPath() string {
	return path.Join(p.Separator, p.UserPath, p.Separator)
}

func (p *PhrasePath) AbsPath() (string, error) {
	absBase, err := filepath.Abs("")
	if err != nil {
		return "", err
	}
	return path.Join(absBase, p.Separator, p.RelPath()), nil
}

func (p *PhrasePath) isLocalePatternUsed() bool {
	return p.isLocaleNameTagInPath() || p.isLocaleIdTagInPath()
}

func (p *PhrasePath) isLocaleNameTagInPath() bool {
	return componentContains(p.Components, "<locale_name>")
}

func (p *PhrasePath) isLocaleIdTagInPath() bool {
	return componentContains(p.Components, "<locale_id>")
}

func (p *PhrasePath) isLocaleRFCInPath() bool {
	return componentContains(p.Components, "<locale_code>")
}

func (p *PhrasePath) isFormatTagInPath() bool {
	return componentContains(p.Components, "<file_format>")
}

func (p *PhrasePath) isTagNameInPath() bool {
	return componentContains(p.Components, "<tag>")
}

func (p *PhrasePath) isValidLocalePath(locale *phraseapp.Locale) (bool, error) {
	localePresent := (locale != nil)

	if !localePresent {
		return false, fmt.Errorf("Locale not set")
	}

	if p.isLocaleRFCInPath() && (locale.Code == "") {
		return false, fmt.Errorf("Locale code is not set for Locale with Id: %s but used in file name", locale.Id)
	}
	return true, nil
}

func PathComponents(userPath string) (*PhrasePath, error) {
	p := &PhrasePath{Separator: string(os.PathSeparator)}

	p.GlobPattern = extractGlobPattern(userPath)
	p.UserPath = cleanUserPath(userPath, p.GlobPattern)
	p.Components = splitToParts(userPath, p.Separator)
	isDir, err := isDir(userPath)
	if err != nil {
		return nil, err
	}

	p.IsDir = isDir

	return p, nil
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

func cleanUserPath(userPath, mode string) string {
	pathWithoutMode := trimSuffix(userPath, mode)
	return strings.TrimSpace(pathWithoutMode)
}

func splitToParts(userPath, separator string) []string {
	split := strings.Split(userPath, separator)
	components := []string{}
	for _, part := range split {
		if part != separator {
			components = append(components, part)
		}
	}
	return components
}

func componentContains(components []string, pattern string) bool {
	for _, part := range components {
		if strings.Contains(part, pattern) {
			return true
		}
	}
	return false
}

// Locale to Path mapping
func CopyLocalePath(relPath string, l *LocalePath) *LocalePath {
	info := &LocaleFileNameInfo{LocaleId: l.Info.LocaleId, LocaleName: l.Info.LocaleName, Tag: l.Info.Tag, FileFormat: l.Info.FileFormat}
	return &LocalePath{Path: relPath, Info: info}
}

type LocaleFileNameInfo struct {
	LocaleName, LocaleId, LocaleRFC5646, Tag, FileFormat string
}

func (info *LocaleFileNameInfo) Message() string {
	str := ""
	if info.LocaleName != "" {
		str = fmt.Sprintf("Name: %s", info.LocaleName)
	}
	if info.LocaleId != "" {
		str = fmt.Sprintf("%s Id: %s", str, info.LocaleId)
	}
	if info.LocaleRFC5646 != "" {
		str = fmt.Sprintf("%s RFC5646: %s", str, info.LocaleRFC5646)
	}
	if info.Tag != "" {
		str = fmt.Sprintf("%s Tag: %s", str, info.Tag)
	}
	if info.FileFormat != "" {
		str = fmt.Sprintf("%s Format: %s", str, info.FileFormat)
	}
	return strings.TrimSpace(str)
}

type LocalePaths []*LocalePath
type LocalePath struct {
	Path string
	Info *LocaleFileNameInfo
}

// Locale placeholder logic <locale_name>
func ExpandPathsWithLocale(p *PhrasePath, locales []*phraseapp.Locale, info *LocaleFileNameInfo) (LocalePaths, error) {
	switch {
	case p.isLocalePatternUsed() && info.LocaleId != "":
		return singlePathWithLocale(p, locales, info)

	case p.isLocalePatternUsed() && info.LocaleId == "":
		return multiplePathsWithLocales(p, locales, info)
	case !p.isLocalePatternUsed():
		return singlePathWithoutLocale(p, info)

	default:
		if info.LocaleId == "" {
			return nil, fmt.Errorf("no target locale id specified")
		}
		return defaultPathWithLocale(p, locales, info)
	}
}

func singlePathWithLocale(p *PhrasePath, locales []*phraseapp.Locale, info *LocaleFileNameInfo) (LocalePaths, error) {
	newPaths := []*LocalePath{}
	locale := localeForLocaleId(info.LocaleId, locales)
	valid, err := p.isValidLocalePath(locale)
	if err != nil {
		return nil, err
	}
	if !valid {
		return nil, fmt.Errorf("Could not find remote locale with Id:", info.LocaleId)
	}
	permutedInfo := &LocaleFileNameInfo{LocaleId: locale.Id, LocaleName: locale.Name, Tag: info.Tag, FileFormat: info.FileFormat}
	absPath, err := newLocaleFile(p, permutedInfo)
	if err != nil {
		return nil, err
	}
	newPaths = append(newPaths, &LocalePath{Path: absPath, Info: permutedInfo})
	return newPaths, nil
}

func multiplePathsWithLocales(p *PhrasePath, locales []*phraseapp.Locale, info *LocaleFileNameInfo) (LocalePaths, error) {
	files := []*LocalePath{}
	for _, locale := range locales {
		valid, err := p.isValidLocalePath(locale)
		if err != nil {
			fmt.Println(err.Error())
		}
		if !valid {
			continue
		}

		permutedInfo := &LocaleFileNameInfo{LocaleName: locale.Name, LocaleId: locale.Id, Tag: info.Tag, FileFormat: info.FileFormat}

		absPath, err := newLocaleFile(p, permutedInfo)
		if err != nil {
			return nil, err
		}
		localePath := &LocalePath{Path: absPath, Info: permutedInfo}
		files = append(files, localePath)
	}
	return files, nil
}

func singlePathWithoutLocale(p *PhrasePath, info *LocaleFileNameInfo) (LocalePaths, error) {
	absPath, err := filepath.Abs(p.UserPath)
	if err != nil {
		return nil, err
	}

	path := &LocalePath{Path: absPath, Info: info}
	localePath := []*LocalePath{path}
	return localePath, nil
}

func defaultPathWithLocale(p *PhrasePath, locales []*phraseapp.Locale, info *LocaleFileNameInfo) (LocalePaths, error) {
	absPath, err := filepath.Abs(p.UserPath)
	if err != nil {
		return nil, err
	}

	matchedLocale := localeForLocaleId(info.LocaleId, locales)
	if matchedLocale == nil {
		return nil, fmt.Errorf("locale specified in your path did not match any remote locales")
	}

	permutedInfo := &LocaleFileNameInfo{
		LocaleId:      info.LocaleId,
		LocaleName:    matchedLocale.Name,
		LocaleRFC5646: matchedLocale.Code,
		FileFormat:    info.FileFormat,
		Tag:           info.Tag,
	}
	path := &LocalePath{Path: absPath, Info: permutedInfo}
	localePath := []*LocalePath{path}
	return localePath, nil
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

func newLocaleFile(p *PhrasePath, info *LocaleFileNameInfo) (string, error) {
	absPath, err := p.AbsPath()
	if err != nil {
		return "", err
	}

	realPath := strings.Replace(absPath, "<locale_name>", info.LocaleName, -1)
	realPath = strings.Replace(realPath, "<locale_code>", info.LocaleId, -1)
	realPath = strings.Replace(realPath, "<tag>", info.Tag, -1)

	return realPath, nil
}

func isLocaleFile(file, extension string) bool {
	fileExtension := fmt.Sprintf(".%s", extension)
	return strings.HasSuffix(file, fileExtension)
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}

// File creation
func createFile(realPath string) error {
	err := fileExists(realPath)
	if err != nil {
		absDir := filepath.Dir(realPath)
		err := fileExists(absDir)
		if err != nil {
			os.MkdirAll(absDir, 0700)
		}

		f, err := os.Create(realPath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}

func fileExists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory:", absPath)
	}
	return nil
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

func sharedMessage(method string, localePath *LocalePath) {
	yellow := ansi.ColorCode("yellow+b:black")
	reset := ansi.ColorCode("reset")

	localPath := localePath.Path
	callerPath, err := os.Getwd()
	if err == nil {
		localPath = "." + strings.Replace(localePath.Path, callerPath, "", 1)
	}

	remote := fmt.Sprint(yellow, localePath.Info.Message(), reset)
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

func wasAlreadySeen(alreadySeen []string, maybeSeen string) bool {
	for _, seen := range alreadySeen {
		if maybeSeen == seen {
			return true
		}
	}
	return false
}
