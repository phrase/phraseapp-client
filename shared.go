package main

import (
	"fmt"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var Debug bool

type LocaleFiles []*LocaleFile
type LocaleFile struct {
	Path, Name, ID, Code, Tag, FileFormat string
	ExistsRemote                          bool
}

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

func (localeFile *LocaleFile) RelPath() string {
	callerPath, _ := os.Getwd()
	relativePath, _ := filepath.Rel(callerPath, localeFile.Path)
	return relativePath
}

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
