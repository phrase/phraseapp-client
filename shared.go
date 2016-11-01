package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	ct "github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
)

var Debug bool

var separator = string(os.PathSeparator)

var placeholderRegexp = regexp.MustCompile("<(locale_name|tag|locale_code)>")

const docsBaseUrl = "https://phraseapp.com/docs"
const docsConfigUrl = docsBaseUrl + "/developers/cli/configuration"

func docsFormatsUrl(formatName string) string {
	return fmt.Sprintf("%s/guides/formats/%s", docsBaseUrl, formatName)
}

func ValidPath(file, formatName, formatExtension string) error {
	if strings.TrimSpace(file) == "" {
		return fmt.Errorf(
			"File patterns may not be empty!\nFor more information see %s", docsConfigUrl,
		)
	}

	fileExtension := strings.Trim(filepath.Ext(file), ".")

	if fileExtension == "<locale_code>" {
		return nil
	}

	if fileExtension == "" {
		return fmt.Errorf("%q has no file extension", file)
	}

	if formatExtension != "" && formatExtension != fileExtension {
		return fmt.Errorf(
			"File extension %q does not equal %q (format: %q) for file %q.\nFor more information see %s",
			fileExtension, formatExtension, formatName, file, docsFormatsUrl(formatName),
		)
	}

	return nil
}

type ProjectLocales interface {
	ProjectIds() []string
}

func LocalesForProjects(client *phraseapp.Client, projectLocales ProjectLocales) (map[string][]*phraseapp.Locale, error) {
	projectIdToLocales := map[string][]*phraseapp.Locale{}
	for _, pid := range projectLocales.ProjectIds() {
		if _, ok := projectIdToLocales[pid]; !ok {
			remoteLocales, err := RemoteLocales(client, pid)
			if err != nil {
				return nil, err
			}
			projectIdToLocales[pid] = remoteLocales
		}
	}
	return projectIdToLocales, nil
}

func RemoteLocales(client *phraseapp.Client, projectId string) ([]*phraseapp.Locale, error) {
	page := 1
	locales, err := client.LocalesList(projectId, page, 25)
	if err != nil {
		return nil, err
	}
	result := locales
	for len(locales) == 25 {
		page = page + 1
		locales, err = client.LocalesList(projectId, page, 25)
		if err != nil {
			return nil, err
		}
		result = append(result, locales...)
	}
	return result, nil
}

func GetFormats(client *phraseapp.Client) (map[string]*phraseapp.Format, error) {
	formats, err := client.FormatsList(1, 25)
	if err != nil {
		return nil, err
	}
	formatMap := map[string]*phraseapp.Format{}
	for _, format := range formats {
		formatMap[format.ApiName] = format
	}
	return formatMap, nil
}

func Contains(seq []string, str string) bool {
	for _, elem := range seq {
		if str == elem {
			return true
		}
	}
	return false
}

func Exists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory: %s", absPath)
	}
	return nil
}

func isDir(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func sharedMessage(method string, localeFile *LocaleFile) {
	local := localeFile.RelPath()

	if method == "pull" {
		remote := localeFile.Message()
		fmt.Print("Downloaded ")
		ct.Foreground(ct.Green, true)
		fmt.Print(remote)
		ct.ResetColor()
		fmt.Print(" to ")
		ct.Foreground(ct.Green, true)
		fmt.Print(local, "\n")
		ct.ResetColor()
	} else {
		fmt.Print("Uploaded ")
		ct.Foreground(ct.Green, true)
		fmt.Print(local)
		ct.ResetColor()
		fmt.Println(" successfully.")
	}
}

func isNotFound(err error) bool {
	return (err != nil && strings.Contains(err.Error(), "404"))
}
