package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/phrase/phraseapp-client/internal/shared"
)

func Validate(file, formatName, formatExtension string) error {
	if strings.TrimSpace(file) == "" {
		return fmt.Errorf("File patterns may not be empty!\nFor more information see %s", shared.DocsConfigUrl)
	}

	fileExtension := strings.Trim(filepath.Ext(file), ".")

	if fileExtension == "" {
		return fmt.Errorf("%q has no file extension", file)
	}

	if fileExtension == "<locale_code>" {
		return nil
	}

	if formatExtension != "" && formatExtension != fileExtension {
		return fmt.Errorf(
			"File extension %q does not equal %q (format: %q) for file %q.\nFor more information see %s",
			fileExtension, formatExtension, formatName, file, shared.DocsFormatsUrl(formatName),
		)
	}

	return nil
}

func Exists(absPath string) error {
	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		return fmt.Errorf("no such file or directory: %s", absPath)
	}
	return nil
}

func IsDir(path string) bool {
	stat, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return stat.IsDir()
}

func Segments(s string) []string {
	return strings.FieldsFunc(filepath.Clean(s), func(c rune) bool { return c == filepath.Separator })
}
