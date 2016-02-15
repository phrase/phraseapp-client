package main

import (
	"fmt"
	"path/filepath"
	"strings"
	"testing"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func getBaseSource() *Source {
	source := &Source{
		File:        "./tests/<locale_code>.yml",
		ProjectID:   "project-id",
		AccessToken: "access-token",
		FileFormat:  "yml",
		Extension:   "",
		Params: new(phraseapp.UploadParams),
		RemoteLocales: getBaseLocales(),
	}
	source.Extension = filepath.Ext(source.File)
	return source
}

func TestPushPreconditions(t *testing.T) {
	fmt.Println("Push#Source#CheckPreconditions")
	source := getBaseSource()
	for _, file := range []string{
		"",
		"no_extension",
		"./<locale_code>/<locale_code>.yml",
		"./*/*/en.yml",
	} {
		source.File = file
		if err := source.CheckPreconditions(); err == nil {
			t.Errorf("CheckPrecondition did not fail for pattern: '%s'", file)
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./*/en.yml",
		"./*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		source.File = file
		if err := source.CheckPreconditions(); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}

func TestSourceFields(t *testing.T) {
	fmt.Println("Push#Source#Fields")
	source := getBaseSource()

	if source.File != "./tests/<locale_code>.yml" {
		t.Errorf("Expected File to be %s and not %s", "./tests/<locale_code>.yml", source.File)
	}

	if source.AccessToken != "access-token" {
		t.Errorf("Expected AccesToken to be %s and not %s", "access-token", source.AccessToken)
	}

	if source.ProjectID != "project-id" {
		t.Errorf("Expected ProjectID to be %s and not %s", "project-id", source.ProjectID)
	}

	if source.FileFormat != "yml" {
		t.Errorf("Expected FileFormat to be %s and not %s", "yml", source.FileFormat)
	}

}

func TestSourceLocaleFilesOne(t *testing.T) {
	fmt.Println("Push#Source#LocaleFiles#1")
	source := getBaseSource()
	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "",
			RFC:  "en",
			ID:   "",
			Path: absPath,
		},
	}

	if len(localeFiles) == len(expectedFiles) {
		if err = compareLocaleFiles(localeFiles, expectedFiles); err != nil {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf("LocaleFiles should contain %v and not %v", expectedFiles, localeFiles)
	}
}

func TestSourceLocaleFilesTwo(t *testing.T) {
	fmt.Println("Push#Source#LocaleFiles#2")
	source := getBaseSource()
	source.File = "./**/<locale_name>.yml"
	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "en",
			RFC:  "",
			ID:   "",
			Path: absPath,
		},
	}

	if len(localeFiles) == len(expectedFiles) {
		if err = compareLocaleFiles(localeFiles, expectedFiles); err != nil {
			t.Errorf(err.Error())
		}
	} else {
		t.Errorf("LocaleFiles should contain %v and not %v", expectedFiles, localeFiles)
	}
}

func TestReplacePlaceholderInParams(t *testing.T) {
	fmt.Println("Push#Source#ReplacePlaceholderInParams")
	source := getBaseSource()
	lid := "<locale_code>"
	source.Params.LocaleID = &lid
	localeFile := &LocaleFile{
		Name: "en",
		RFC:  "en",
		ID:   "",
		Path: "",
	}
	s := source.replacePlaceholderInParams(localeFile)
	if s != "en" {
		t.Errorf("Expected LocaleId to equal '%s' but was '%s'", "en", s)
		t.Fail()
	}
}
func TestGetRemoteLocaleForLocaleFile(t *testing.T) {
	fmt.Println("Push#Source#getRemoteLocaleForLocaleFile")
	source := getBaseSource()
	localeFile := &LocaleFile{
		Name: "english",
		RFC:  "en",
		ID:   "",
		Path: "",
	}
	locale := source.getRemoteLocaleForLocaleFile(localeFile)
	if locale.Name != localeFile.Name {
		t.Errorf("Expected LocaleName to equal '%s' but was '%s'", "ennglish", localeFile.Name)
		t.Fail()
	}
	if locale.Code != localeFile.RFC {
		t.Errorf("Expected LocaleId to equal '%s' but was '%s'", "en", localeFile.RFC)
		t.Fail()
	}
}

type Pattern struct {
	File         string
	Ext          string
	TestPath     string
	ExpectedRFC  string
	ExpectedName string
	ExpectedTag  string
}

func TestReducerPatterns(t *testing.T) {
	fmt.Println("Push#Reducer")
	for idx, pattern := range []*Pattern{
		&Pattern{
			File:        "./.abc/<locale_code>.yml",
			Ext:         "yml",
			TestPath:    ".abc/en.yml",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:        "./abc+/defg./}{x/][etc??/<locale_code>.yml",
			Ext:         "yml",
			TestPath:    "abc+/defg./}{x/][etc??/en.yml",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:     "./*.yml",
			Ext:      "yml",
			TestPath: "en.yml",
		},
		&Pattern{
			File:     "./locales/?*.yml",
			Ext:      "yml",
			TestPath: "locales/?en.yml",
		},
		&Pattern{
			File:     "./locales/en.yml",
			Ext:      "yml",
			TestPath: "locales/en.yml",
		},
		&Pattern{
			File:     "./locales/.yml",
			Ext:      "yml",
			TestPath: "locales/en.yml",
		},
		&Pattern{
			File:        "./abc/defg/<locale_code>.lproj/.strings",
			Ext:         "strings",
			TestPath:    "abc/defg/en.lproj/Localizable.strings",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:        "./locales/<locale_code>.yml",
			Ext:         "yml",
			TestPath:    "locales/en.yml",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:        "./config/<tag>/<locale_code>.yml",
			Ext:         "yml",
			TestPath:    "config/abc/en.yml",
			ExpectedRFC: "en",
			ExpectedTag: "abc",
		},
		&Pattern{
			File:         "./config/<locale_name>/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "config/german/de.yml",
			ExpectedRFC:  "de",
			ExpectedTag:  "",
			ExpectedName: "german",
		},
		&Pattern{
			File:        "./config/<locale_code>/*.yml",
			Ext:         "yml",
			TestPath:    "config/en/english.yml",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:        "./no_tag/<tag>/<locale_code>.lproj/Localizable.strings",
			Ext:         "strings",
			TestPath:    "no_tag/abc/en.lproj/Localizable.strings",
			ExpectedRFC: "en",
			ExpectedTag: "abc",
		},
		&Pattern{
			File:        "./abc/<locale_code>.lproj/.strings",
			Ext:         "strings",
			TestPath:    "abc/en.lproj/Localizable.strings",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:        "./abc/<locale_code>.lproj/<tag>.strings",
			Ext:         "strings",
			TestPath:    "abc/en.lproj/MyStoryboard.strings",
			ExpectedRFC: "en",
			ExpectedTag: "MyStoryboard",
		},
		&Pattern{
			File:        "./*/<tag>/<locale_code>-values/Strings.xml",
			Ext:         "xml",
			TestPath:    "no_tag/abc/en-values/Strings.xml",
			ExpectedRFC: "en",
			ExpectedTag: "abc",
		},
		&Pattern{
			File:        "./*/<tag>/play.<locale_code>",
			Ext:         "<locale_code>",
			TestPath:    "no_tag/abc/play.en",
			ExpectedRFC: "en",
			ExpectedTag: "abc",
		},
		&Pattern{
			File:         "./*/<tag>/<locale_name>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "no_tag/abc/play.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "abc",
			ExpectedName: "play",
		},
		&Pattern{
			File:         "./abc/<locale_code>/<tag>.<locale_name>",
			Ext:          "strings",
			TestPath:     "abc/en/MyStoryboard.english",
			ExpectedRFC:  "en",
			ExpectedTag:  "MyStoryboard",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./**/*.<locale_name>",
			Ext:          "<locale_name>",
			TestPath:     "abc/defg/hijk/some_name.english",
			ExpectedName: "english",
		},
		&Pattern{
			File:        "./**/<tag>.<locale_code>",
			Ext:         "<locale_code>",
			TestPath:    "abc/defg/hijk/someTag.en",
			ExpectedRFC: "en",
			ExpectedTag: "someTag",
		},

		&Pattern{
			File:        "./lang/<locale_code>/**/*.php",
			Ext:         "php",
			TestPath:    "lang/en/hijk/bla/bla/someName.php",
			ExpectedRFC: "en",
		},
		&Pattern{
			File:         "./**/xyz/<locale_name>/<tag>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "abc/xyz/english/someTag.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/abc/defg/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./**/<tag>/<locale_name>/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "abc/defg/someTag/english/en.yml",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/abc/defg/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/**/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/someTag/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>_more/**/<tag>no_tag/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english_more/abc/defg/someTagno_tag/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/**/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/someTag/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./**/<locale_name>/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/english/someTag/filename.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/**/*/<tag>/main.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/someTag/main.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./**/<locale_code>/<locale_name>/*/<tag>/main.yml",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/haha/hahah/hahah/hahah/en/english/hahah/someTag/main.yml",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./<locale_name>/*/<tag>/**/main.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/xyz/someTag/haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/main.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		&Pattern{
			File:         "./**/<locale_name>/*/<tag>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/haha/hahah/hahah/hahah/hahah/hahah/english/hahah/someTag.en",
			ExpectedRFC:  "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
	} {

		tokens := splitPathToTokens(pattern.File)

		fmt.Println(strings.Repeat("-", 10))
		fmt.Println(idx+1, "\n  SourceFile:", pattern.File, "\n  TestPath:", pattern.TestPath)

		pathTokens := splitPathToTokens(pattern.TestPath)
		localeFile := extractParamsFromPathTokens(tokens, pathTokens)

		if localeFile.RFC != pattern.ExpectedRFC {
			t.Errorf("Expected RFC to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedRFC, localeFile.RFC, idx+1)
			t.Fail()
		}

		if localeFile.Tag != pattern.ExpectedTag {
			t.Errorf("Expected Tag to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedTag, localeFile.Tag, idx+1)
			t.Fail()
		}

		if localeFile.Name != pattern.ExpectedName {
			t.Errorf("Expected LocaleName to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedName, localeFile.Name, idx+1)
			t.Fail()
		}
	}
	fmt.Println(strings.Repeat("-", 10))
}

func TestSplitString(t *testing.T) {
	tt := []struct {
		str string
		cut string
		exp []string
	}{
		{"a/b/c/d", "/\\", []string{"a", "b", "c", "d"}},
		{"a/b\\c/d", "/\\", []string{"a", "b", "c", "d"}},
		{"/a/b/c/d/", "/\\", []string{"", "a", "b", "c", "d"}},
		{"a/b/日\\本/語/d", "/\\", []string{"a", "b", "日", "本", "語", "d"}},
		{"a/b/日\\c本c/語/d", "/\\本", []string{"a", "b", "日", "c", "c", "語", "d"}},
	}

	for i := range tt {
		got := splitString(tt[i].str, tt[i].cut)
		if len(got) != len(tt[i].exp) {
			t.Errorf("expected %d elements for %q, got %d", len(tt[i].exp), tt[i].str, len(got))
		}
	}
}

func TestSystemFiles(t *testing.T) {
	tt := []struct {
		pattern  string
		expFiles []string
	}{
		{"a/b/c/d.txt", []string{"a/b/c/d.txt"}},
		{"a/b/c/d.*", []string{"a/b/c/d.txt", "a/b/c/d.jpg"}},
		{"a/b/*/d.txt", []string{"a/b/c/d.txt", "a/b/x/d.txt"}},
		{"a/b/c/*.txt", []string{"a/b/c/d.txt", "a/b/c/e.txt"}},
		{"a/*/c/d.txt", []string{"a/b/c/d.txt", "a/y/c/d.txt"}},
		{"a/*/c/*.txt", []string{"a/b/c/d.txt", "a/b/c/e.txt", "a/y/c/d.txt"}},

		{"a/**/d.txt", []string{"a/d.txt", "a/b/d.txt", "a/b/c/d.txt", "a/y/c/d.txt", "a/b/x/d.txt"}},
		{"a/**/*/d.txt", []string{"a/b/d.txt", "a/b/c/d.txt", "a/b/x/d.txt", "a/y/c/d.txt"}},
		{"a/**/c/d.txt", []string{"a/b/c/d.txt", "a/y/c/d.txt"}},
		{"a/**/c/*.txt", []string{"a/b/c/d.txt", "a/b/c/e.txt", "a/y/c/d.txt"}},
		{"a/*/**/c/d.txt", []string{"a/b/c/d.txt", "a/y/c/d.txt"}},
	}

	for _, tti := range tt {
		src := new(Source)
		src.File = filepath.Join("testdata", tti.pattern)
		src.Extension = filepath.Ext(tti.pattern)

		matches, err := src.SystemFiles()
		if err != nil {
			t.Errorf("%s: didn't expect an error, got: %s", src.File, err)
			continue
		}

		exp := map[string]bool{}
		for _, f := range tti.expFiles {
			exp[filepath.Join("testdata", f)] = true
		}

		for _, got := range matches {
			if _, found := exp[got]; !found {
				t.Errorf("%s: got unexpected file %q", src.File, got)
				continue
			}
			delete(exp, got)
		}

		for k, _ := range exp {
			t.Errorf("%s: expected to get file %q, but it didn't appear", src.File, k)
		}
	}
}