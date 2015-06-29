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
	UserPath        string
	Separator       string
	Components      []string
	Mode            string
	LocaleTagInPath bool
	FormatTagInPath bool
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

func (p *PhrasePath) LocaleTagInFile() bool {
	if p.LocaleTagInPath {
		if len(p.Components) > 0 {
			return strings.Contains(p.Components[len(p.Components)-1], "<locale_name>")
		}
	}
	return false
}

func PathComponents(userPath string) *PhrasePath {
	p := &PhrasePath{Separator: string(os.PathSeparator)}

	p.Mode = extractGlobMode(userPath)
	p.UserPath = cleanUserPath(userPath, p.Mode)
	p.Components = splitToParts(p.UserPath, p.Separator)
	p.LocaleTagInPath = componentContains(p.Components, "<locale_name>")
	p.FormatTagInPath = componentContains(p.Components, "<format_name>")

	return p
}

func extractGlobMode(userPath string) string {
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
	return &LocalePath{Path: relPath, LocaleId: l.LocaleId, LocaleName: l.LocaleName, LocaleCode: l.LocaleCode}
}

type LocalePaths []*LocalePath
type LocalePath struct {
	Path       string
	LocaleId   string
	LocaleName string
	LocaleCode string
}

// File expansion for * and **/* and ""
func expandSingleDirectory(p *PhrasePath, paths LocalePaths, fileFormat string) (LocalePaths, error) {
	expandedPaths := []*LocalePath{}
	for _, localePath := range paths {

		asDirectory := fmt.Sprintf("%s/", localePath.Path)
		pathsPerDirectory, err := glob(asDirectory, fileFormat)

		if err != nil {
			return nil, err
		}

		for _, absPath := range pathsPerDirectory {
			localeToPathMapping := CopyLocalePath(absPath, localePath)
			expandedPaths = append(expandedPaths, localeToPathMapping)
		}

	}
	return expandedPaths, nil
}

func recurseDirectory(fileFormat string, paths LocalePaths) (LocalePaths, error) {
	expandedPaths := []*LocalePath{}
	for _, localePath := range paths {
		newPaths, err := walk(localePath.Path, fileFormat)
		if err != nil {
			return nil, err
		}
		for _, newPath := range newPaths {
			localeToPathMapping := &LocalePath{Path: newPath, LocaleId: localePath.LocaleId}
			expandedPaths = append(expandedPaths, localeToPathMapping)
		}
	}
	return expandedPaths, nil
}

func walk(root, fileFormat string) ([]string, error) {
	fileList := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if isLocaleFile(f.Name(), fileFormat) {
			fileList = append(fileList, path)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}

	return fileList, nil
}

func glob(root, fileFormat string) ([]string, error) {
	files, err := filepath.Glob(root + "*")

	if err != nil {
		return nil, err
	}
	localeFiles := []string{}
	for _, f := range files {
		if fileFormat != "" {
			if isLocaleFile(f, fileFormat) {
				localeFiles = append(localeFiles, f)
			}
		} else {
			localeFiles = append(localeFiles, f)
		}
	}
	return localeFiles, nil
}

// Locale placeholder logic <locale_name>
func ExpandPathsWithLocale(p *PhrasePath, localeId string, locales []*phraseapp.Locale) (LocalePaths, error) {
	switch {
	case p.LocaleTagInPath && localeId == "":
		newFiles, err := filePathsWithLocales(p, locales)
		if err != nil {
			return nil, err
		}
		return newFiles, nil

	case !p.LocaleTagInPath:
		absPath, err := filepath.Abs(p.UserPath)
		if err != nil {
			return nil, err
		}

		localePath := []*LocalePath{&LocalePath{Path: absPath, LocaleId: localeId}}
		return localePath, nil

	default:
		return defaultPathWithLocale(p, localeId, locales)
	}
}

func filePathsWithLocales(p *PhrasePath, locales []*phraseapp.Locale) (LocalePaths, error) {
	files := []*LocalePath{}
	for _, locale := range locales {
		absPath, err := newLocaleFile(p, locale.Name)
		if err != nil {
			return nil, err
		}
		localePath := &LocalePath{Path: absPath, LocaleId: locale.Id, LocaleName: locale.Name, LocaleCode: locale.Code}
		files = append(files, localePath)
	}
	return files, nil
}

func defaultPathWithLocale(p *PhrasePath, localeId string, locales []*phraseapp.Locale) (LocalePaths, error) {
	if localeId == "" {
		return nil, fmt.Errorf("no target locale id specified")
	}

	absPath, err := filepath.Abs(p.UserPath)
	if err != nil {
		return nil, err
	}

	matchedLocale := localeIdForPath(localeId, locales)
	if matchedLocale == "" {
		return nil, fmt.Errorf("locale specified in your path did not match any remote locales")
	}

	localePath := []*LocalePath{&LocalePath{Path: absPath, LocaleId: matchedLocale}}
	return localePath, nil
}

// Locale logic
func localeIdForPath(localeId string, locales []*phraseapp.Locale) string {
	for _, locale := range locales {
		if localeId == locale.Id {
			return locale.Id
		}
	}
	return ""
}

func newLocaleFile(p *PhrasePath, localeName string) (string, error) {
	absPath, err := p.AbsPath()
	if err != nil {
		return "", err
	}

	realPath := strings.Replace(absPath, "<locale_name>", localeName, -1)

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
func CreateFiles(p *PhrasePath, virtualPaths LocalePaths, fileFormat string) error {
	for _, localePath := range virtualPaths {
		defaultName := defaultFileName(p.Mode, localePath.Path, fileFormat)
		if defaultName != "" {
			localePath.Path = path.Join(localePath.Path, p.Separator, defaultName)
		}
		err := createFile(localePath.Path)
		if err != nil {
			return err
		}
	}
	return nil
}

func defaultFileName(mode, localePath, fileFormat string) string {
	if mode != "" {
		if isDir(localePath) {
			return fmt.Sprintf("phrase.%s", fileFormat)
		}
	}
	return ""
}

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

func isDir(absPath string) bool {
	f, err := os.Open(absPath)
	if err != nil {
		return false
	}

	defer f.Close()
	fi, err := f.Stat()
	if err != nil {
		return false
	}
	switch mode := fi.Mode(); {
	case mode.IsDir():
		return true
	}
	return false
}

func Authenticate() error {
	defaultCredentials, err := ConfigDefaultCredentials()
	if err != nil {
		return err
	}
	phraseapp.RegisterAuthCredentials(defaultCredentials, defaultCredentials)
	return nil
}

func printError(err error, msg string) {
	red := ansi.ColorCode("red+b:black")
	reset := ansi.ColorCode("reset")
	fmt.Fprintf(os.Stderr, "%sERROR: %s %s%s\n", red, err, msg, reset)
}

func sharedMessage(method string, localePath *LocalePath) {
	yellow := ansi.ColorCode("yellow+b:black")
	reset := ansi.ColorCode("reset")

	localPath := localePath.Path
	localeName := localePath.LocaleName
	localeCode := localePath.LocaleCode

	fmt1 := []string{}

	fmt1 = append(fmt1, yellow)
	if localeName != "" {
		fmt1 = append(fmt1, localeName)
	}
	if localeCode != "" {
		fmt1 = append(fmt1, fmt.Sprintf(" (%s)", localeCode))
	}
	fmt1 = append(fmt1, reset)

	if len(fmt1) <= 2 {
		fmt1 = []string{yellow, "?", reset}
	}

	fmt2 := []string{}
	fmt2 = append(fmt2, yellow)
	fmt2 = append(fmt2, localPath)
	fmt2 = append(fmt2, reset)

	remote := strings.Join(fmt1, "")
	local := strings.Join(fmt2, "")

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
