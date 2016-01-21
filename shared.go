package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
)

var Debug bool

type LocaleFiles []*LocaleFile
type LocaleFile struct {
	Path, Name, ID, RFC, Tag, FileFormat string
	ExistsRemote                         bool
}

type BugsnagError struct {
	App        string `json:"app"`
	AppVersion string `json:"app_version"`
	ErrorData  `json:"data"`
	Message    string `json:"message"`
	Name       string `json:"name"`
}

type ErrorData struct {
	Arch string `json:"arch"`
	Os   string `json:"os"`
}

var placeholderRegexp = regexp.MustCompile("<(locale_name|tag|locale_code)>")

func ValidPath(file, format, fileExtension string) error {
	if strings.TrimSpace(file) == "" {
		return fmt.Errorf(
			"File patterns may not be empty!\nFor more information see http://docs.phraseapp.com/developers/cli/configuration/",
		)
	}

	extension := strings.Trim(strings.TrimSpace(filepath.Ext(file)), ".")
	if extension == "" || (fileExtension != "" && extension != fileExtension) {
		extensionInfo := ""
		if fileExtension != "" {
			extensionInfo = fmt.Sprintf(" %q", fileExtension)
		}

		return fmt.Errorf(
			"'%s' does not have the required extension%s.\nFor more information see http://docs.phraseapp.com/guides/formats/%s",
			file, extensionInfo, format,
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

func printErr(err error) {
	ct.Foreground(ct.Red, true)
	fmt.Fprintf(os.Stderr, "\nERROR: %s\n", err)
	ct.ResetColor()
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

func Exists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory: %s", absPath)
	}
	return nil
}

func ReportError(name string, message string) {
	bs := &BugsnagError{
		App:        "phraseapp-client",
		AppVersion: PHRASEAPP_CLIENT_VERSION,
		ErrorData: ErrorData{
			Arch: runtime.GOARCH,
			Os:   runtime.GOOS,
		},
		Name:    name,
		Message: message,
	}

	body, err := json.Marshal(bs)
	if err != nil {
		return
	}

	req, err := http.NewRequest("POST", "https://phraseapp.com/errors", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return
	}
	resp.Body.Close()

}
