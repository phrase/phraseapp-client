package paclient

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"strings"

	"github.com/phrase/phraseapp-go/phraseapp"
)

type Targets []*Target

func (targets Targets) ProjectIds() []string {
	projectIds := []string{}
	for _, target := range targets {
		projectIds = append(projectIds, target.ProjectID)
	}
	return projectIds
}

type Target struct {
	File          string
	ProjectID     string
	AccessToken   string
	FileFormat    string
	Params        *PullParams
	RemoteLocales []*phraseapp.Locale
}

type PullParams struct {
	phraseapp.LocaleDownloadParams
	LocaleID string
}

func (tgt *Target) UnmarshalYAML(unmarshal func(interface{}) error) error {
	m := map[string]interface{}{}
	err := phraseapp.ParseYAMLToMap(unmarshal, map[string]interface{}{
		"file":         &tgt.File,
		"project_id":   &tgt.ProjectID,
		"access_token": &tgt.AccessToken,
		"file_format":  &tgt.FileFormat,
		"params":       &m,
	})
	if err != nil {
		return err
	}

	tgt.Params = new(PullParams)
	if v, found := m["locale_id"]; found {
		if tgt.Params.LocaleID, err = phraseapp.ValidateIsString("params.locale_id", v); err != nil {
			return err
		}
		// Must delete the param from the map as the LocaleDownloadParams type
		// doesn't support this one and the apply method would return an error.
		delete(m, "locale_id")
	}
	return tgt.Params.ApplyValuesFromMap(m)

}

func (target *Target) CheckPreconditions() error {
	if err := ValidPath(target.File, target.FileFormat, ""); err != nil {
		return err
	}

	if strings.Count(target.File, "*") > 0 {
		return fmt.Errorf(
			"File pattern for 'pull' cannot include any 'stars' *. Please specify direct and valid paths with file name!\n %s#targets", DocsConfigUrl,
		)
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(target.File, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
	}

	if target.GetLocaleID() == "" && !containsAnyPlaceholders(target.File) {
		return fmt.Errorf("Could not find any locale information. Please specify a 'locale_id' in your params or provide a placeholder!")
	}

	return nil
}

func (target *Target) Pull(client *phraseapp.Client) error {
	if err := target.CheckPreconditions(); err != nil {
		return err
	}

	localeFiles, err := target.LocaleFiles(NewLocalGlobber())
	if err != nil {
		return err
	}

	localeIdToFileIsDistinct := (target.GetLocaleID() != "" && len(localeFiles) == 1)

	for _, localeFile := range localeFiles {
		err := createFile(localeFile.Path)
		if err != nil {
			return err
		}

		if localeIdToFileIsDistinct {
			if target.GetLocaleID() != "" {
				localeFile.ID = target.GetLocaleID()
			}
		}

		err = target.DownloadAndWriteToFile(client, localeFile)
		if err != nil {
			return fmt.Errorf("%s for %s", err, localeFile.Path)
		} else {
			sharedMessage("pull", localeFile)
		}
		if Debug {
			fmt.Fprintln(os.Stderr, strings.Repeat("-", 10))
		}
	}

	return nil
}

func (target *Target) DownloadAndWriteToFile(client *phraseapp.Client, localeFile *LocaleFile) error {
	downloadParams := new(phraseapp.LocaleDownloadParams)
	if target.Params != nil {
		*downloadParams = target.Params.LocaleDownloadParams
	}

	if downloadParams.FileFormat == nil {
		downloadParams.FileFormat = &localeFile.FileFormat
	}

	if Debug {
		fmt.Fprintln(os.Stderr, "Target file pattern:", target.File)
		fmt.Fprintln(os.Stderr, "Actual file path", localeFile.Path)
		fmt.Fprintln(os.Stderr, "LocaleID", localeFile.ID)
		fmt.Fprintln(os.Stderr, "ProjectID", target.ProjectID)
		fmt.Fprintln(os.Stderr, "FileFormat", downloadParams.FileFormat)
		fmt.Fprintln(os.Stderr, "ConvertEmoji", downloadParams.ConvertEmoji)
		fmt.Fprintln(os.Stderr, "IncludeEmptyTranslations", downloadParams.IncludeEmptyTranslations)
		fmt.Fprintln(os.Stderr, "KeepNotranslateTags", downloadParams.KeepNotranslateTags)
		fmt.Fprintln(os.Stderr, "Tag", downloadParams.Tag)
		fmt.Fprintln(os.Stderr, "FormatOptions", downloadParams.FormatOptions)
	}

	res, err := client.LocaleDownload(target.ProjectID, localeFile.ID, downloadParams)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(localeFile.Path, res, 0700)
	if err != nil {
		return err
	}
	return nil
}

func (target *Target) LocaleFiles(globber *GlobFinder) (LocaleFiles, error) {
	localeID := target.GetLocaleID()

	files := []*LocaleFile{}
	for _, remoteLocale := range target.RemoteLocales {
		if localeID != "" && !(remoteLocale.ID == localeID || remoteLocale.Name == localeID) {
			continue
		}
		err := target.IsValidLocale(remoteLocale, target.File)
		if err != nil {
			return nil, err
		}

		localeFile := &LocaleFile{
			Name:       remoteLocale.Name,
			ID:         remoteLocale.ID,
			Code:       remoteLocale.Code,
			Tag:        target.GetTag(),
			FileFormat: target.GetFormat(),
			Path:       target.File,
		}

		absPath, err := target.ReplacePlaceholders(globber, localeFile)
		if err != nil {
			return nil, err
		}
		localeFile.Path = absPath

		files = append(files, localeFile)
	}

	return files, nil
}

func (target *Target) IsValidLocale(locale *phraseapp.Locale, localPath string) error {
	if locale == nil {
		return fmt.Errorf("Remote locale could not be downloaded correctly!")
	}

	if strings.Contains(localPath, "<locale_code>") && locale.Code == "" {
		return fmt.Errorf("Locale code is not set for Locale with ID: %s but locale_code is used in file name", locale.ID)
	}
	return nil
}

func (target *Target) ReplacePlaceholders(globber *GlobFinder, localeFile *LocaleFile) (string, error) {
	absPath, err := globber.Abs(target.File)
	if err != nil {
		return "", err
	}

	path := strings.Replace(absPath, "<locale_name>", localeFile.Name, -1)
	path = strings.Replace(path, "<locale_code>", localeFile.Code, -1)
	path = strings.Replace(path, "<tag>", localeFile.Tag, -1)

	return path, nil
}

func (t *Target) GetFormat() string {
	if t.Params != nil && t.Params.FileFormat != nil {
		return *t.Params.FileFormat
	}
	if t.FileFormat != "" {
		return t.FileFormat
	}
	return ""
}

func (t *Target) GetLocaleID() string {
	if t.Params != nil {
		return t.Params.LocaleID
	}
	return ""
}

func (t *Target) GetTag() string {
	if t.Params != nil && t.Params.Tag != nil {
		return *t.Params.Tag
	}
	return ""
}

func createFile(path string) error {
	err := Exists(path)
	if err != nil {
		absDir := filepath.Dir(path)
		err := Exists(absDir)
		if err != nil {
			os.MkdirAll(absDir, 0700)
		}

		f, err := os.Create(path)
		if err != nil {
			return err
		}
		defer f.Close()
	}
	return nil
}
