package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/daviddengcn/go-colortext"
	"github.com/phrase/phraseapp-go/phraseapp"
	"regexp"
)

var Debug bool

type LocaleFiles []*LocaleFile
type LocaleFile struct {
	Path, Name, ID, RFC, Tag, FileFormat string
	ExistsRemote                         bool
}

var placeholderRegexp = regexp.MustCompile("<(locale_name|tag|locale_code)>")

func CheckPreconditions(file string) error {
	if strings.TrimSpace(file) == "" {
		return fmt.Errorf(
			"File patterns of a source may not be empty! Please use a valid file pattern: %s",
			"http://docs.phraseapp.com/developers/cli/configuration/",
		)
	}

	extension := filepath.Ext(file)
	if strings.TrimSpace(extension) == "" {
		abs, err := filepath.Abs(file)
		if err != nil {
			return err
		}
		return fmt.Errorf(
			"'%s' does not have a valid extension. Please use a valid extension: %s",
			abs, "http://docs.phraseapp.com/guides/formats/",
		)
	}

	duplicatedPlaceholders := []string{}
	for _, name := range []string{"<locale_name>", "<locale_code>", "<tag>"} {
		if strings.Count(file, name) > 1 {
			duplicatedPlaceholders = append(duplicatedPlaceholders, name)
		}
	}

	starCount := strings.Count(file, "*")
	recCount := strings.Count(file, "**")

	if recCount == 0 && starCount > 1 || starCount-(recCount*2) > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "*")
	}

	if recCount > 1 {
		duplicatedPlaceholders = append(duplicatedPlaceholders, "**")
	}

	if len(duplicatedPlaceholders) > 0 {
		dups := strings.Join(duplicatedPlaceholders, ", ")
		return fmt.Errorf(fmt.Sprintf("%s can only occur once in a file pattern!", dups))
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
		return fmt.Errorf("no such file or directory:", absPath)
	}
	return nil
}
