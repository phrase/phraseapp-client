package paths

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestGlob(t *testing.T) {
	base, err := ioutil.TempDir("", "test-glob_")
	defer os.RemoveAll(base)
	if err != nil {
		t.Error(err)
	}

	directories := []string{
		"foo/bar/baz/asd",
		"foo/bar/xyz/asd",
		"foo/bar/baz/xyz/asd",
	}

	files := []string{
		"en.yml",
		"en.json",
		"de.docx",
		"nanana",
	}

	for _, dir := range directories {
		err := os.MkdirAll(filepath.Join(base, dir), 0755)
		defer os.RemoveAll(filepath.Join(base, dir))
		if err != nil {
			t.Error(err)
		}

		for _, file := range files {
			_, err := os.Create(filepath.Join(base, dir, file))
			if err != nil {
				t.Error(err)
			}
		}
	}

	tests := map[string][]string{
		"**/*.mp3":               {},
		"foo/*/baz/**/asd/*.yml": {"foo/bar/baz/asd/en.yml", "foo/bar/baz/xyz/asd/en.yml"},
		"foo/**/*.yml":           {"foo/bar/baz/asd/en.yml", "foo/bar/xyz/asd/en.yml", "foo/bar/baz/xyz/asd/en.yml"},
		"foo/bar/xyz/asd/*":      {"foo/bar/xyz/asd/en.yml", "foo/bar/xyz/asd/en.json", "foo/bar/xyz/asd/de.docx", "foo/bar/xyz/asd/nanana"},
		"**/asd/*": {
			"foo/bar/baz/asd/en.yml",
			"foo/bar/baz/asd/en.json",
			"foo/bar/baz/asd/de.docx",
			"foo/bar/baz/asd/nanana",

			"foo/bar/xyz/asd/en.yml",
			"foo/bar/xyz/asd/en.json",
			"foo/bar/xyz/asd/de.docx",
			"foo/bar/xyz/asd/nanana",

			"foo/bar/baz/xyz/asd/en.yml",
			"foo/bar/baz/xyz/asd/en.json",
			"foo/bar/baz/xyz/asd/de.docx",
			"foo/bar/baz/xyz/asd/nanana",
		},
	}

	for pattern, expected := range tests {
		matches, err := Glob(filepath.Join(base, pattern))
		if err != nil {
			t.Error(err)
		}

		for idx, match := range matches {
			matches[idx], _ = filepath.Rel(base, match)
		}

		if !areEqual(matches, expected) {
			t.Errorf("expected %v, got %v", expected, matches)
		}
	}
}

func areEqual(s, t []string) bool {
	sort.Strings(s)
	sort.Strings(t)

	for idx := range s {
		if s[idx] != t[idx] {
			return false
		}
	}

	return true
}
