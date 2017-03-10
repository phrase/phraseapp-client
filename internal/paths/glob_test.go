package paths

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"
)

func TestGlob(t *testing.T) {
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

	testGlob(directories, files, tests, t)
}

func TestGlob_specialCharacters(t *testing.T) {
	directories := []string{
		"locales",
		"foo?bar",
		"bar[a-z]",
		"bla/*",
	}

	files := []string{
		"en.yml",
		"?",
		"_[^x]_.yml",
	}

	tests := map[string][]string{
		"**/*.yml": {
			"locales/en.yml",
			"locales/_[^x]_.yml",
			"foo?bar/en.yml",
			"foo?bar/_[^x]_.yml",
			"bar[a-z]/en.yml",
			"bar[a-z]/_[^x]_.yml",
			"bla/*/en.yml",
			"bla/*/_[^x]_.yml",
		},
		"foo?bar/*": {
			"foo?bar/en.yml",
			"foo?bar/?",
			"foo?bar/_[^x]_.yml",
		},
		"bar[a*/_[^x*]_.yml": {
			"bar[a-z]/_[^x]_.yml",
		},
		"bla/\\*/*": {
			"bla/*/en.yml",
			"bla/*/?",
			"bla/*/_[^x]_.yml",
		},
		"**/?": {
			"locales/?",
			"foo?bar/?",
			"bar[a-z]/?",
			"bla/*/?",
		},
	}

	testGlob(directories, files, tests, t)
}

func testGlob(directories, files []string, tests map[string][]string, t *testing.T) {
	base, err := ioutil.TempDir("", "test-glob_")
	defer os.RemoveAll(base)
	if err != nil {
		t.Error(err)
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

func TestSplitAtDirGlobOperator(t *testing.T) {
	path := "/foo/bla/bar/baz/asd/a/b/c"

	tests := map[string]struct {
		pathStart, patternStart, pathEnd, patternEnd string
	}{
		"/foo/**/asd/a/b/c": {
			"/foo",
			"/foo",
			"asd/a/b/c",
			"asd/a/b/c",
		},
		"/foo/*/bar/**/a/*/c": {
			"/foo/bla/bar",
			"/foo/*/bar",
			"a/b/c",
			"a/*/c",
		},
		"/**/bar/baz/*/a/*/c": {
			"",
			"",
			"bar/baz/asd/a/b/c",
			"bar/baz/*/a/*/c",
		},
	}

	for pattern, expected := range tests {
		pathStart, patternStart, pathEnd, patternEnd, err := SplitAtDirGlobOperator(path, pattern)
		if err != nil {
			t.Error(err)
		}

		if pathStart != expected.pathStart {
			t.Errorf("expected path start to be %v, got %v", expected.pathStart, pathStart)
		}

		if patternStart != expected.patternStart {
			t.Errorf("expected pattern start to be %v, got %v", expected.patternStart, patternStart)
		}

		if pathEnd != expected.pathEnd {
			t.Errorf("expected path end to be %v, got %v", expected.pathEnd, pathEnd)
		}

		if patternEnd != expected.patternEnd {
			t.Errorf("expected pattern end to be %v, got %v", expected.patternEnd, patternEnd)
		}
	}
}
