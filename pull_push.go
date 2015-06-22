package main

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-api-client/Godeps/_workspace/src/gopkg.in/yaml.v2"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func ConfigPushPull() (*PushPullConfig, error) {
	content, err := ConfigContent()
	if err != nil {
		return nil, err
	}

	return parsePushPullArgs(content)
}

type Path struct {
	UserPath        string
	Separator       string
	AbsPath         string
	Components      []string
	Mode            string
	LocaleSpecified bool
	FormatSpecified bool
}

func (p *Path) RealPath() string {
	return path.Join(p.Separator, p.UserPath, p.Separator)
}

func (p *Path) SubPath(toReplace, replacement string) string {
	return path.Join(p.Separator, strings.Replace(p.UserPath, toReplace, replacement, 0), p.Separator)
}

func PathComponents(userPath string) *Path {
	p := &Path{UserPath: userPath, Separator: string(os.PathSeparator)}

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
			if !p.LocaleSpecified {
				p.LocaleSpecified = strings.Contains(part, "<locale_name>")
			}
			if !p.FormatSpecified {
				p.FormatSpecified = strings.Contains(part, "<format_name>")
			}
			p.Components = append(p.Components, part)
		}
	}

	return p
}

func PullStrategy(p *Path, params *Params) ([]string, error) {
	files := []string{}
	if p.LocaleSpecified {
		locales, err := phraseapp.LocalesList(params.ProjectId, 1, 25)
		if err != nil {
			return nil, err
		}
		for _, locale := range locales {
			absPath, err := NewLocaleFile(p, locale.Name)
			if err != nil {
				return nil, err
			}
			files = append(files, absPath)
		}
	} else {
		absPath, err := filepath.Abs(p.UserPath)
		if err != nil {
			return nil, err
		}
		switch {
		case p.Mode == "":
			return []string{absPath}, nil

		case p.Mode == "*":
			return singleDirectoryStrategy(absPath, "")

		}
	}

	return files, nil
}

// File handling
func recursiveStrategy(root, fileFormat string) ([]string, error) {
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

func singleDirectoryStrategy(root, fileFormat string) ([]string, error) {
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

func NewLocaleFile(p *Path, localeName string) (string, error) {
	newPath := p.SubPath("<locale_name>", localeName)

	absPath, err := filepath.Abs(newPath)
	if err != nil {
		return "", err
	}

	err = fileExists(absPath)
	if err != nil {
		absDir := filepath.Dir(absPath)
		os.MkdirAll(absDir, 0644)

		f, err := os.Create(absPath)
		if err != nil {
			return "", err
		}
		defer f.Close()
	}

	return absPath, nil
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

// @TODO: This is not exactly the specified format. Params only contains params to upload/download. AccessToken, File, ProjectId in top-level struct.
// @TODO: Because of this the naming is bad, see wizard.go for almost desired syntax.
// Parsing
type Params struct {
	File        string
	AccessToken string `yaml:"access_token"`
	ProjectId   string `yaml:"project_id"`
	Format      string
	FormatName  string `yaml:"format_name"`
	LocaleId    string `yaml:"locale_id"`
	Emoji       bool   `yaml:"emoji"`
}

type PushPullConfig struct {
	Phraseapp struct {
		AccessToken string `yaml:"access_token"`
		ProjectId   string `yaml:"project_id"`
		Push        struct {
			Sources []*Params
		}
		Pull struct {
			Targets []*Params
		}
	}
}

func parsePushPullArgs(yml string) (*PushPullConfig, error) {
	var pushPullConfig *PushPullConfig

	err := yaml.Unmarshal([]byte(yml), &pushPullConfig)
	if err != nil {
		return nil, err
	}

	return pushPullConfig, nil
}
