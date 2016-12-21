package paclient

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/jpillora/backoff"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func getUploadResult(client *phraseapp.Client, projectID string, upload *phraseapp.Upload) (result string, err error) {
	b := &backoff.Backoff{
		Min:    500 * time.Millisecond,
		Max:    10 * time.Second,
		Factor: 2,
		Jitter: true,
	}

	for ; result != "success" && result != "error"; result = upload.State {
		time.Sleep(b.Duration())
		upload, err = client.UploadShow(projectID, upload.ID)
		if err != nil {
			break
		}
	}

	return
}

func isDir(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func splitPathIntoSegments(path string) []string {
	segments := []string{}
	start := 0
	for i := range path {
		if os.IsPathSeparator(path[i]) {
			segments = append(segments, path[start:i])
			start = i + 1
		}
	}
	return append(segments, path[start:])
}

func findFilesInPath(root string) ([]string, error) {
	files := []string{}
	err := filepath.Walk(root, func(path string, f os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !f.Mode().IsDir() {
			files = append(files, path)
		}
		return nil
	})

	return files, err
}

func splitString(s string, set string) []string {
	if len(set) == 1 {
		return strings.Split(s, set)
	}

	slist := []string{}
	charSet := map[rune]bool{}

	for _, r := range set {
		charSet[r] = true
	}

	start := 0
	for i, r := range s {
		if _, found := charSet[r]; found {
			slist = append(slist, s[start:i])
			start = i + utf8.RuneLen(r)
		}
	}
	if start < len(s) {
		slist = append(slist, s[start:])
	}

	return slist
}

func validateFileCandidate(tokens []string, ignoreTokenCnt int, cand string) bool {
	candTokens := splitPathIntoSegments(cand)
	candTokenCnt := len(candTokens)

	if candTokenCnt < (len(tokens) + ignoreTokenCnt) {
		return false
	}

	candTokens = candTokens[ignoreTokenCnt:]

	for i := 1; i <= len(tokens); i++ {
		expT, gotT := tokens[len(tokens)-i], candTokens[len(candTokens)-i]
		switch {
		case strings.Contains(expT, "*"):
			matched, err := regexp.MatchString(expT, gotT)
			if err != nil {
				panic(err)
			}
			if !matched {
				return false
			}
		case expT != gotT:
			return false
		}
	}

	return true
}

func Contains(seq []string, str string) bool {
	for _, elem := range seq {
		if str == elem {
			return true
		}
	}
	return false
}

func ValidPath(file, formatName, formatExtension string) error {
	if strings.TrimSpace(file) == "" {
		return fmt.Errorf(
			"File patterns may not be empty!\nFor more information see %s", DocsConfigUrl,
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

func docsFormatsUrl(formatName string) string {
	return fmt.Sprintf("%s/guides/formats/%s", DocsBaseUrl, formatName)
}

func containsAnyPlaceholders(s string) bool {
	return PlaceholderRegexp.MatchString(s)
}

func Exists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory: %s", absPath)
	}
	return nil
}

func getBaseLocales() []*phraseapp.Locale {
	return []*phraseapp.Locale{
		&phraseapp.Locale{
			Code: "en",
			ID:   "en-locale-id",
			Name: "english",
		},
		&phraseapp.Locale{
			Code: "de",
			ID:   "de-locale-id",
			Name: "german",
		},
	}
}
