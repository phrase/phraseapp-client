package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"path/filepath"
	"strings"
	"testing"

	"github.com/phrase/phraseapp-client/internal/paths"
	"github.com/phrase/phraseapp-client/internal/placeholders"
	"github.com/phrase/phraseapp-go/phraseapp"
)

func getBaseSource() *Source {
	source := &Source{
		File:          "./tests/<locale_code>.yml",
		ProjectID:     "project-id",
		AccessToken:   "access-token",
		FileFormat:    "yml",
		Params:        new(phraseapp.UploadParams),
		RemoteLocales: getBaseLocales(),
	}
	return source
}

func TestPushPreconditions(t *testing.T) {
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
	source := getBaseSource()
	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "english",
			Code: "en",
			ID:   "en-locale-id",
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
	source := getBaseSource()
	source.File = "./tests/<locale_name>.yml"
	localeFiles, err := source.LocaleFiles()

	if err != nil {
		t.Errorf("Should not fail with: %s", err.Error())
	}

	absPath, _ := filepath.Abs("./tests/en.yml")
	expectedFiles := []*LocaleFile{
		&LocaleFile{
			Name: "en",
			Code: "",
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
	source := getBaseSource()
	lid := "<locale_code>"
	source.Params.LocaleID = &lid
	localeFile := &LocaleFile{
		Name: "en",
		Code: "en",
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
	source := getBaseSource()
	localeFile := &LocaleFile{
		Name: "english",
		Code: "en",
		ID:   "",
		Path: "",
	}
	locale := source.getRemoteLocaleForLocaleFile(localeFile)
	if locale.Name != localeFile.Name {
		t.Errorf("Expected LocaleName to equal '%s' but was '%s'", "ennglish", localeFile.Name)
		t.Fail()
	}
	if locale.Code != localeFile.Code {
		t.Errorf("Expected LocaleId to equal '%s' but was '%s'", "en", localeFile.Code)
		t.Fail()
	}
}

func TestCheckPreconditions(t *testing.T) {
	tt := []struct {
		pattern    string
		fileformat string
		expError   string
	}{
		{"", "", "File patterns may not be empty!"},
		{"a/b/c", ".foo", "\"a/b/c\" has no file extension"},
		{"<locale_name>/<locale_name>.foo", ".foo", "<locale_name> can only occur once in a file pattern!"},
		{"<locale_code>/<locale_code>.foo", ".foo", "<locale_code> can only occur once in a file pattern!"},
		{"<tag>/<tag>.foo", ".foo", "<tag> can only occur once in a file pattern!"},
		{"a/**/b/**/c.t", ".t", "** can only occur once in a file pattern!"},
		{"a/*/b/**/d/*/c.t", ".t", "* can only occur once in a file pattern!"},
		{"<locale_name>/<locale_code>/**/a/<tag>/*/c.t", ".t", ""},
	}

	for _, tti := range tt {
		src := new(Source)
		src.File = tti.pattern
		src.FileFormat = tti.fileformat

		err := src.CheckPreconditions()
		switch {
		case tti.expError == "" && err != nil:
			t.Errorf("%s: didn't expect an error, but got: %s", tti.pattern, err)
		case tti.expError != "" && err == nil:
			t.Errorf("%s: expected an error, but got none", tti.pattern)
		case err != nil && !strings.HasPrefix(err.Error(), tti.expError):
			t.Errorf("%s: expected error to have prefix %q, got: %q", tti.pattern, tti.expError, err)
		}
	}
}

type localeFile struct {
	path         string
	code         string
	name         string
	tag          string
	localeExists bool
}

func setupFiles(t *testing.T, files ...string) (dir string) {
	d, err := ioutil.TempDir("/tmp", "phrase-push-test")
	if err != nil {
		t.Fatal(err)
	}

	for _, n := range files {
		p := filepath.Join(d, n)
		if err := os.MkdirAll(filepath.Dir(p), 0755); err != nil {
			t.Fatal(err)
		}
		if err := ioutil.WriteFile(p, nil, 0644); err != nil {
			t.Fatal(err)
		}
	}
	return d
}

func pushd(t *testing.T, dir string) func() {
	wd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	err = os.Chdir(dir)
	if err != nil {
		t.Fatal(err)
	}
	return func() {
		os.Chdir(wd)
	}

}

func setupLocalesFiles(t *testing.T) (dir string) {
	files := []string{
		"a/b/c/d.jpg",
		"a/b/c/d.txt",
		"a/b/c/e.txt",
		"a/b/d.txt",
		"a/b/x/d.txt",
		"a/d.txt",
		"a/y/c/d.txt",
		"b/YY/c/d.txt",
		"b/YY/foo.bar.json",
		"b/YY/foo.json",
	}
	files = append(files, []string{
		"config/locales/application.en.yml",
		"config/locales/attributes/course.en.yml",
		"config/locales/devise.en.yml",
		"config/locales/landing.en.yml",
		"config/locales/layouts.en.yml",
	}...)
	return setupFiles(t, files...)
}

func TestSystemFiles(t *testing.T) {
	d := setupLocalesFiles(t)
	defer os.RemoveAll(d)
	defer pushd(t, d)()

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
		{"./config/locales/**/*.en.yml", []string{
			"config/locales/application.en.yml",
			"config/locales/attributes/course.en.yml",
			"config/locales/devise.en.yml",
			"config/locales/landing.en.yml",
			"config/locales/layouts.en.yml",
		}},
	}

	for _, tti := range tt {
		src := new(Source)
		src.File = tti.pattern

		matches, err := paths.Glob(placeholders.ToGlobbingPattern(src.File))
		if err != nil {
			t.Errorf("%s: didn't expect an error, got: %s", src.File, err)
			continue
		}

		exp := map[string]bool{}
		for _, f := range tti.expFiles {
			exp[f] = true
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

func TestLocaleFiles(t *testing.T) {
	d := setupLocalesFiles(t)
	defer os.RemoveAll(d)
	defer pushd(t, d)()

	tt := []struct {
		pattern string
		files   []localeFile
	}{
		{"a/b/c/d.txt", []localeFile{{"a/b/c/d.txt", "", "", "", false}}},
		{"a/<locale_code>/c/d.txt", []localeFile{
			{"a/b/c/d.txt", "b", "", "", false},
			{"a/y/c/d.txt", "y", "YY", "", true}}}, // local_code sets the name
		{"a/<locale_name>/c/d.txt", []localeFile{
			{"a/b/c/d.txt", "", "b", "", false},
			{"a/y/c/d.txt", "", "y", "", false}}},
		{"a/<tag>/c/d.txt", []localeFile{
			{"a/b/c/d.txt", "", "", "b", false},
			{"a/y/c/d.txt", "", "", "y", false}}},
		{"a/<locale_code>/<locale_name>/<tag>.txt", []localeFile{
			{"a/b/c/d.txt", "b", "c", "d", false},
			{"a/b/c/e.txt", "b", "c", "e", false},
			{"a/b/x/d.txt", "b", "x", "d", false},
			{"a/y/c/d.txt", "y", "c", "d", false}}},
		{"b/<locale_name>/c/<tag>.txt", []localeFile{
			{"b/YY/c/d.txt", "y", "YY", "d", true}}},

		// This shows a toxic example of the pattern mechanism!
		{"b/YY/foo.<locale_name>.json", []localeFile{
			{"b/YY/foo.bar.json", "", "bar", "", false},
		}},
		{"b/YY/<locale_name>.json", []localeFile{
			{"b/YY/foo.json", "", "foo", "", false},
			{"b/YY/foo.bar.json", "", "foo.bar", "", false}, // weird locale!
		}},
	}

	for _, tti := range tt {
		src := new(Source)
		src.File = tti.pattern

		src.RemoteLocales = append(src.RemoteLocales, &phraseapp.Locale{ID: "random", Code: "y", Name: "YY"})
		files, err := src.LocaleFiles()
		if err != nil {
			t.Errorf("%s: didn't expect an error, got: %s", src.File, err)
			continue
		}

		if len(files) != len(tti.files) {
			t.Errorf("expected %d files, got %d", len(tti.files), len(files))
		}

		pathFileMap := map[string]localeFile{}
		for i := range tti.files {
			abs, err := filepath.Abs(tti.files[i].path)
			if err != nil {
				t.Fatalf("didn't expect error, got: %s", err)
			}
			pathFileMap[abs] = tti.files[i]
		}

		for _, lf := range files {
			expFile, found := pathFileMap[lf.Path]
			if !found {
				t.Errorf("%s: file at path %s not expected", src.File, lf.Path)
				continue
			}

			delete(pathFileMap, lf.Path)
			if lf.Code != expFile.code {
				t.Errorf("%s: expected code %q for %q, got %q", src.File, expFile.code, expFile.path, lf.Code)
			}
			if lf.Name != expFile.name {
				t.Errorf("%s: expected name %q for %q, got %q", src.File, expFile.name, expFile.path, lf.Name)
			}
			if lf.Tag != expFile.tag {
				t.Errorf("%s: expected tag %q for %q, got %q", src.File, expFile.tag, expFile.path, lf.Tag)
			}
			if lf.ExistsRemote != expFile.localeExists {
				if expFile.localeExists {
					t.Errorf("%s: expected locale to exist remote, it didn't", src.File)
				} else {
					t.Errorf("%s: expected locale to not exist remote, it does", src.File)
				}
			}
		}

		for k, _ := range pathFileMap {
			t.Errorf("%s: didn't see expected file at %s", src.File, k)
		}
	}
}

type testHandler struct {
	lastFilename string
	lastLocaleID string
	lastTag      string
}

func (th *testHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(64)

	getVal := func(key string) string {
		val := req.MultipartForm.Value[key]
		if len(val) > 0 {
			return val[0]
		}
		return ""
	}

	th.lastFilename = req.MultipartForm.File["file"][0].Filename
	th.lastLocaleID = getVal("locale_id")
	th.lastTag = getVal("tags")

	resp.WriteHeader(http.StatusCreated)
	io.WriteString(resp, `{}`)
}

func TestUploadFile(t *testing.T) {
	d := setupFiles(t, "a/b/c/d.txt")
	defer os.RemoveAll(d)
	th := new(testHandler)

	srv := httptest.NewServer(th)

	c := new(phraseapp.Client)
	c.Credentials.Host = srv.URL
	c.Credentials.Token = "some_token"

	src := new(Source)
	src.Params = new(phraseapp.UploadParams)
	tags := "a,b"
	src.Params.Tags = &tags

	file := new(LocaleFile)
	file.Path = filepath.Join(d, "a/b/c/d.txt")
	file.ID = "locale_id"
	file.Code = "locale-code"
	file.Tag = "sometag"

	_, err := src.uploadFile(c, file)
	if err != nil {
		t.Errorf("didn't expect an error, got: %s", err)
	}

	expPath := path.Base(file.Path)
	expTags := "a,b,sometag"

	if th.lastFilename != expPath {
		t.Errorf("expected file path %q, got %q", expPath, th.lastFilename)
	}
	if th.lastLocaleID != file.ID {
		t.Errorf("expected locale id %q, got %q", file.ID, th.lastLocaleID)
	}
	if th.lastTag != expTags {
		t.Errorf("expected tag %q, got %q", expTags, th.lastTag)
	}
}

func TestRemoteLocaleForLocaleFile(t *testing.T) {
	rlEN := &phraseapp.Locale{ID: "en-locale-id", Name: "english", Code: "en"}
	rlDE := &phraseapp.Locale{ID: "de-locale-id", Name: "deutsch", Code: "de"}
	tt := []struct {
		srcLocaleID string
		remotes     []*phraseapp.Locale
		code        string
		name        string
		expLocales  *phraseapp.Locale
	}{
		{"", nil, "en", "english", nil},
		{"", []*phraseapp.Locale{rlEN, rlDE}, "", "", nil},
		{"en-locale-id", []*phraseapp.Locale{rlEN, rlDE}, "", "", rlEN},
		{"", []*phraseapp.Locale{rlEN, rlDE}, "en", "", rlEN},
		{"", []*phraseapp.Locale{rlEN, rlDE}, "", "english", rlEN},
		{"", []*phraseapp.Locale{rlEN, rlDE}, "en", "english", rlEN},
		{"en-locale-id", []*phraseapp.Locale{rlEN, rlDE}, "en", "english", rlEN},
		{"", []*phraseapp.Locale{rlEN, rlDE}, "", "deutsch", rlDE},

		{"<locale_code>glish", []*phraseapp.Locale{rlEN, rlDE}, "", "", nil},
		{"<locale_code>glish", []*phraseapp.Locale{rlEN, rlDE}, "en", "english", rlEN},
		{"<locale_code>glish", []*phraseapp.Locale{rlEN, rlDE}, "de", "deutsch", nil},

		// let's leave the happy path and create obscure situations
		{"", []*phraseapp.Locale{rlEN, rlDE}, "en", "deutsch", nil},
		{"de-locale-id", []*phraseapp.Locale{rlEN, rlDE}, "en", "", nil},
		{"de-locale-id", []*phraseapp.Locale{rlEN, rlDE}, "", "english", nil},
	}

	for i, tti := range tt {
		src := new(Source)
		src.Params = new(phraseapp.UploadParams)
		src.Params.LocaleID = &tti.srcLocaleID
		src.RemoteLocales = tti.remotes
		lf := new(LocaleFile)
		lf.Name = tti.name
		lf.Code = tti.code
		r := src.getRemoteLocaleForLocaleFile(lf)
		switch {
		case tti.expLocales == nil && r != nil:
			t.Errorf("%d: didn't expect an locale, got %q", i, r.ID)
		case tti.expLocales != nil && r == nil:
			t.Errorf("%d: expected locale %q, but got none", i, tti.expLocales.ID)
		case tti.expLocales != nil && r != nil && tti.expLocales.ID != r.ID:
			t.Errorf("%d: expected locale %q, but got %q", i, tti.expLocales.ID, r.ID)
		}
	}
}

type patternShouldCreateLocale struct {
	Name         string
	Code         string
	ExistsRemote bool
	Expected     bool
	Source       *Source
}

func TestShouldCreateLocale(t *testing.T) {
	source := &Source{
		Format: &phraseapp.Format{
			IncludesLocaleInformation: false,
		},
	}
	testPatterns := []*patternShouldCreateLocale{
		{
			Name:         "",
			Code:         "",
			ExistsRemote: false,
			Expected:     false,
			Source:       source,
		},
		{

			Name:         "",
			Code:         "",
			ExistsRemote: true,
			Expected:     false,
			Source:       source,
		},
		{

			Name:         "English",
			Code:         "",
			ExistsRemote: false,
			Expected:     true,
			Source:       source,
		},
		{

			Name:         "",
			Code:         "en",
			ExistsRemote: false,
			Expected:     true,
			Source:       source,
		},
		{
			Name:         "English",
			Code:         "en",
			ExistsRemote: false,
			Expected:     true,
			Source:       source,
		},
		{
			Name:         "English",
			Code:         "en",
			ExistsRemote: false,
			Expected:     false,
			Source: &Source{
				Format: &phraseapp.Format{
					IncludesLocaleInformation: true,
				},
			},
		},
	}

	for _, pattern := range testPatterns {
		localeFile := &LocaleFile{
			Name:         pattern.Name,
			Code:         pattern.Code,
			ExistsRemote: pattern.ExistsRemote,
		}
		was := localeFile.shouldCreateLocale(pattern.Source)
		if was != pattern.Expected {
			t.Errorf("expected the locale should be created to be: %t, but was %t", pattern.Expected, was)
		}
	}
}
