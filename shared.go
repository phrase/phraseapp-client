package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-go/phraseapp"
)

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
	p := &PhrasePath{UserPath: userPath, Separator: string(os.PathSeparator)}

	if strings.HasSuffix(p.UserPath, path.Join("**", "*")) {
		p.Mode = "**/*"
	} else if strings.HasSuffix(p.UserPath, "*") {
		p.Mode = "*"
	} else {
		p.Mode = ""
	}

	p.UserPath = strings.TrimSpace(trimSuffix(p.UserPath, p.Mode))

	split := strings.Split(p.UserPath, p.Separator)
	for _, part := range split {
		if part != p.Separator {
			if !p.LocaleTagInPath {
				p.LocaleTagInPath = strings.Contains(part, "<locale_name>")
			}
			if !p.FormatTagInPath {
				p.FormatTagInPath = strings.Contains(part, "<format_name>")
			}
			p.Components = append(p.Components, part)
		}
	}

	return p
}

// File handling
func RecursiveStrategy(root, fileFormat string) ([]string, error) {
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

func SingleDirectoryStrategy(root, fileFormat string) ([]string, error) {
	files, err := filepath.Glob(root)
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

type LocalePaths []*LocalePath
type LocalePath struct {
	Path     string
	LocaleId string
}

func FilePathsWithLocales(p *PhrasePath, locales []*phraseapp.Locale) (LocalePaths, error) {
	files := []*LocalePath{}
	for _, locale := range locales {
		absPath, err := newLocaleFile(p, locale.Name)
		if err != nil {
			return nil, err
		}
		files = append(files, &LocalePath{Path: absPath, LocaleId: locale.Id})
	}
	return files, nil
}

func newLocaleFile(p *PhrasePath, localeName string) (string, error) {
	absPath, err := p.AbsPath()
	if err != nil {
		return "", err
	}

	realPath := strings.Replace(absPath, "<locale_name>", localeName, -1)

	return realPath, nil
}

func CreateFile(realPath string) error {
	err := fileExists(realPath)
	if err != nil {
		absDir := filepath.Dir(realPath)
		os.MkdirAll(absDir, 0700)

		f, err := os.Create(realPath)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
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

func DefaultFileName(mode, localePath string) string {
	if mode != "" {
		if isDir(localePath) {
			return "phrase.yml"
		}
	}
	return ""
}

func authenticate() error {
	defaultCredentials, err := ConfigDefaultCredentials()
	if err != nil {
		return err
	}

	phraseapp.RegisterAuthCredentials(defaultCredentials, defaultCredentials)

	return nil
}
