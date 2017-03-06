package main

import (
	"testing"

	"github.com/phrase/phraseapp-client/internal/paths"
)

type Patterns []*Pattern

type Pattern struct {
	File         string
	Ext          string
	TestPath     string
	ExpectedCode string
	ExpectedName string
	ExpectedTag  string
}

func TestSpecialCharacters(t *testing.T) {
	patterns := Patterns{
		{
			File:     "./locales/?*.yml",
			Ext:      "yml",
			TestPath: "locales/?en.yml",
		},
		{
			File:         "./abc+/defg./}{x/][etc??/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "abc+/defg./}{x/][etc??/en.yml",
			ExpectedCode: "en",
		},
	}
	patterns.TestPatterns(t)
}

func TestSinglePlaceholders(t *testing.T) {
	patterns := Patterns{
		{
			File:         "./.abc/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     ".abc/en.yml",
			ExpectedCode: "en",
		},
		{
			File:     "./*.yml",
			Ext:      "yml",
			TestPath: "en.yml",
		},
		{
			File:     "./locales/en.yml",
			Ext:      "yml",
			TestPath: "locales/en.yml",
		},
		{
			File:     "./locales/.yml",
			Ext:      "yml",
			TestPath: "locales/en.yml",
		},
		{
			File:         "./abc/defg/<locale_code>.lproj/.strings",
			Ext:          "strings",
			TestPath:     "abc/defg/en.lproj/Localizable.strings",
			ExpectedCode: "en",
		},
		{
			File:         "./locales/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "locales/en.yml",
			ExpectedCode: "en",
		},
		{
			File:         "./abc/<locale_code>.lproj/.strings",
			Ext:          "strings",
			TestPath:     "abc/en.lproj/Localizable.strings",
			ExpectedCode: "en",
		},
	}
	patterns.TestPatterns(t)
}

func TestMultiplePlaceholders(t *testing.T) {
	patterns := Patterns{
		{
			File:         "./config/<tag>/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "config/abc/en.yml",
			ExpectedCode: "en",
			ExpectedTag:  "abc",
		},
		{
			File:         "./config/<locale_name>/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "config/german/de.yml",
			ExpectedCode: "de",
			ExpectedTag:  "",
			ExpectedName: "german",
		},
		{
			File:         "./config/<locale_code>/*.yml",
			Ext:          "yml",
			TestPath:     "config/en/english.yml",
			ExpectedCode: "en",
		},
		{
			File:         "./no_tag/<tag>/<locale_code>.lproj/Localizable.strings",
			Ext:          "strings",
			TestPath:     "no_tag/abc/en.lproj/Localizable.strings",
			ExpectedCode: "en",
			ExpectedTag:  "abc",
		},
		{
			File:         "./abc/<locale_code>.lproj/<tag>.strings",
			Ext:          "strings",
			TestPath:     "abc/en.lproj/MyStoryboard.strings",
			ExpectedCode: "en",
			ExpectedTag:  "MyStoryboard",
		},
		{
			File:         "./*/<tag>/<locale_code>-values/Strings.xml",
			Ext:          "xml",
			TestPath:     "no_tag/abc/en-values/Strings.xml",
			ExpectedCode: "en",
			ExpectedTag:  "abc",
		},
		{
			File:         "./lang/<locale_code>/**/*.php",
			Ext:          "php",
			TestPath:     "lang/en/hijk/bla/bla/someName.php",
			ExpectedCode: "en",
		},
		{
			File:         "./**/<tag>/<locale_name>/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "abc/defg/someTag/english/en.yml",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./**/<locale_name>/<tag>/*.yml",
			Ext:          "yml",
			TestPath:     "haha/haha/haha/haha/english/someTag/filename.yml",
			ExpectedCode: "",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/**/*/<tag>/*.yml",
			Ext:          "yml",
			TestPath:     "english/abc/defg/haha/hahah/hahah/someTag/main.yml",
			ExpectedCode: "",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./locales/**/<tag>.*.yml",
			Ext:          "yml",
			TestPath:     "./locales/a/b/my-tag.anything.yml",
			ExpectedCode: "",
			ExpectedTag:  "my-tag",
			ExpectedName: "",
		},
		{
			File:         "./locales/**/<tag>.*.yml",
			Ext:          "yml",
			TestPath:     "./locales/a/b/my-tag.anything.yml",
			ExpectedCode: "",
			ExpectedTag:  "my-tag",
			ExpectedName: "",
		},
		{
			File:         "./**/<tag>.*/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "./locales/a/my-tag.anything/en.yml",
			ExpectedCode: "en",
			ExpectedTag:  "my-tag",
			ExpectedName: "",
		},
	}
	patterns.TestPatterns(t)
}

func TestMultipleWithPlaceholderExtension(t *testing.T) {
	patterns := Patterns{
		{
			File:         "./abc/<locale_code>/<tag>.<locale_name>",
			Ext:          "<locale_name>",
			TestPath:     "abc/en/MyStoryboard.english",
			ExpectedCode: "en",
			ExpectedTag:  "MyStoryboard",
			ExpectedName: "english",
		},
		{
			File:         "./*/<tag>/play.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "no_tag/abc/play.en",
			ExpectedCode: "en",
			ExpectedTag:  "abc",
		},
		{
			File:         "./*/<tag>/<locale_name>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "no_tag/abc/play.en",
			ExpectedCode: "en",
			ExpectedTag:  "abc",
			ExpectedName: "play",
		},
		{
			File:         "./**/*.<locale_name>",
			Ext:          "<locale_name>",
			TestPath:     "abc/defg/hijk/some_name.english",
			ExpectedName: "english",
		},
		{
			File:         "./**/<tag>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "abc/defg/hijk/someTag.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
		},
		{
			File:         "./**/xyz/<locale_name>/<tag>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "abc/xyz/english/someTag.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/abc/defg/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/abc/defg/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/**/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/someTag/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>_more/**/<tag>no_tag/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english_more/abc/defg/someTagno_tag/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/**/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/haha/hahah/hahah/someTag/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/<tag>/**/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/someTag/haha/hahah/hahah/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./**/<locale_name>/<tag>/*.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/haha/english/someTag/filename.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/**/*/<tag>/main.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/abc/defg/haha/hahah/hahah/someTag/main.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./**/<locale_code>/<locale_name>/*/<tag>/main.yml",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/en/english/hahah/someTag/main.yml",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./<locale_name>/*/<tag>/**/main.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "english/xyz/someTag/haha/hahah/hahah/main.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./**/<locale_name>/*/<tag>.<locale_code>",
			Ext:          "<locale_code>",
			TestPath:     "haha/haha/haha/hahah/english/hahah/someTag.en",
			ExpectedCode: "en",
			ExpectedTag:  "someTag",
			ExpectedName: "english",
		},
		{
			File:         "./**/<locale_name>/abc.*/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "./locales/a/english/abc.lol/en.yml",
			ExpectedCode: "en",
			ExpectedTag:  "",
			ExpectedName: "english",
		},
		{
			File:         "./a/b/<locale_name>/abc.*/**/<locale_code>.yml",
			Ext:          "yml",
			TestPath:     "./a/b/english/abc.lol/a/b/c/en.yml",
			ExpectedCode: "en",
			ExpectedTag:  "",
			ExpectedName: "english",
		},
	}
	patterns.TestPatterns(t)
}

func (patterns Patterns) TestPatterns(t *testing.T) {
	for idx, pattern := range patterns {
		pattern.TestPattern(t, idx)
	}
}

func (pattern *Pattern) TestPattern(t *testing.T, idx int) {
	tokens := paths.Segments(pattern.File)
	pathTokens := paths.Segments(pattern.TestPath)

	localeFile := extractParamsFromPathTokens(tokens, pathTokens)

	if localeFile.Code != pattern.ExpectedCode {
		t.Errorf("Expected Code to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedCode, localeFile.Code, idx+1)
	}

	if localeFile.Tag != pattern.ExpectedTag {
		t.Errorf("Expected Tag to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedTag, localeFile.Tag, idx+1)
	}

	if localeFile.Name != pattern.ExpectedName {
		t.Errorf("Expected LocaleName to equal '%s' but was '%s' Pattern: %d", pattern.ExpectedName, localeFile.Name, idx+1)
	}
}
