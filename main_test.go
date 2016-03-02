package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
	"testing"

	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/dynport/dgtk/cli"
	"github.com/phrase/phraseapp-client/Godeps/_workspace/src/github.com/phrase/phraseapp-go/phraseapp"
)

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

type ByPath []*LocaleFile

func (a ByPath) Len() int           { return len(a) }
func (a ByPath) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByPath) Less(i, j int) bool { return a[i].Path < a[j].Path }

func compareLocaleFiles(actualFiles LocaleFiles, expectedFiles LocaleFiles) error {
	sort.Sort(ByPath(actualFiles))
	sort.Sort(ByPath(expectedFiles))
	for idx, actualFile := range actualFiles {
		expected := expectedFiles[idx]
		actual := actualFile
		if expected.Path != actual.Path {
			return fmt.Errorf("Expected Path %s should eql %s", expected.Path, actual.Path)
		}
		if expected.Name != actual.Name {
			return fmt.Errorf("Expected Name %s should eql %s", expected.Name, actual.Name)
		}
		if expected.RFC != actual.RFC {
			return fmt.Errorf("Expected RFC %s should eql %s", expected.RFC, actual.RFC)
		}
		if expected.ID != actual.ID {
			return fmt.Errorf("Expected ID %s should eql %s", expected.ID, actual.ID)
		}
	}
	return nil
}

func captureStderr(f func() error) (string, error) {
	old := cli.DefaultWriter
	defer func() {
		cli.DefaultWriter = old
	}()
	buf := &bytes.Buffer{}
	cli.DefaultWriter = buf
	err := f()
	return buf.String(), err
}

func runWithCfg(cfg *phraseapp.Config, cmd string, additionalOpts ...string) (string, error) {
	r, err := router(cfg)
	if err != nil {
		return "", err
	}

	stderr, err := ioutil.TempFile("", "phraseapp_cli_test")
	if err != nil {
		return "", err
	}
	defer os.Remove(stderr.Name())

	opts := strings.Split(cmd, "/")
	opts = append(opts, additionalOpts...)
	opts = append(opts, "-h")

	out, err := captureStderr(func() error {
		return r.Run(opts...)
	})

	if err != cli.ErrorHelpRequested {
		return "", err
	}
	return out, nil
}

func TestCLIHelp_NoDefaults(t *testing.T) {
	cfg := new(phraseapp.Config)

	out, err := runWithCfg(cfg, "locale/download")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}
	if len(defaults) > 0 {
		t.Errorf("expected no defaults in configuration, found: %#v", defaults)
	}
}

func collectDefaults(out string) (map[string]string, error) {
	res := map[string]string{}

	sc := bufio.NewScanner(bytes.NewBufferString(out))
	for sc.Scan() {
		line := strings.TrimSpace(sc.Text())
		if !strings.Contains(line, "(default:") {
			continue
		}
		fields := strings.Split(line, " ")

		if name, ok := searchFieldWithPrefix(fields, hasPrefix("--")); ok {
			if value, ok := searchFieldWithPredecessor(fields, "(default:"); ok {
				res[name] = value
			}
		}

	}
	return res, sc.Err()
}

func searchFieldWithPredecessor(flds []string, pre string) (string, bool) {
	i := findInStringList(flds, func(s string) bool {
		return s == pre
	})

	switch i {
	case -1:
		return "", false
	case len(flds) - 1:
		return "", false
	default:
		return strings.Trim(flds[i+1], ")"), true
	}

}

func searchFieldWithPrefix(flds []string, pred func(string) bool) (string, bool) {
	switch i := findInStringList(flds, pred); i {
	case -1:
		return "", false
	default:
		return flds[i], true
	}
}

func hasPrefix(p string) func(s string) bool {
	return func(s string) bool {
		return strings.HasPrefix(s, p)
	}
}

func findInStringList(l []string, f func(string) bool) int {
	for i := range l {
		if f(l[i]) {
			return i
		}
	}
	return -1
}

func matchDefaultExpectations(t *testing.T, got, exp map[string]string) {
	for k, v := range exp {
		switch gv, found := got[k]; {
		case found && gv != v:
			t.Errorf("%s: expected value %q, got %q", k, v, gv)
		case !found:
			t.Errorf("%s: expected to be present, it wasn't", k)
		}
		delete(got, k)
	}

	for k, _ := range got {
		t.Errorf("%s: is present, but wasn't expected", k)
	}
}

func TestCLIHelp_FileFormatDefault(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.DefaultFileFormat = "FILE_FORMAT"

	out, err := runWithCfg(cfg, "locale/download")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	matchDefaultExpectations(t, defaults, map[string]string{
		"--file-format": "FILE_FORMAT",
	})
}

func TestCLIHelp_FileFormatDefaultTwice(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.DefaultFileFormat = "FILE_FORMAT"
	cfg.Defaults = map[string]map[string]interface{}{}
	cfg.Defaults["locale/download"] = map[string]interface{}{
		"file_format": "OTHER_FILE_FORMAT",
	}

	out, err := runWithCfg(cfg, "locale/download")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	matchDefaultExpectations(t, defaults, map[string]string{
		"--file-format": "OTHER_FILE_FORMAT",
	})
}

func TestCLIHelp_FileFormatDefaultThrice(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.DefaultFileFormat = "FILE_FORMAT"
	cfg.Defaults = map[string]map[string]interface{}{}
	cfg.Defaults["locale/download"] = map[string]interface{}{
		"file_format": "OTHER_FILE_FORMAT",
	}

	out, err := runWithCfg(cfg, "locale/download", "--file-format", "YET_ANOTHER_FILE_FORMAT")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	matchDefaultExpectations(t, defaults, map[string]string{
		"--file-format": "YET_ANOTHER_FILE_FORMAT",
	})
}

func itop(i int) *int {
	return &i
}

func TestCLIHelp_PerPageSettings(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.Page = itop(2)
	cfg.PerPage = itop(12)

	out, err := runWithCfg(cfg, "locales/list")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	matchDefaultExpectations(t, defaults, map[string]string{
		"--page":     "2",
		"--per-page": "12",
	})
}

func TestCLIHelp_PerPageSettingsOverride(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.Page = itop(2)
	cfg.PerPage = itop(12)

	out, err := runWithCfg(cfg, "locales/list", "--page", "3", "--per-page", "42")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	matchDefaultExpectations(t, defaults, map[string]string{
		"--page":     "3",
		"--per-page": "42",
	})
}

func TestCLIHelp_FormatOptions(t *testing.T) {
	cfg := new(phraseapp.Config)
	cfg.Defaults = map[string]map[string]interface{}{}
	cfg.Defaults["locale/download"] = map[string]interface{}{
		"format_options": map[interface{}]interface{}{
			"convert_placeholders": "true",
		},
	}

	out, err := runWithCfg(cfg, "locale/download")
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	defaults, err := collectDefaults(out)
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}

	// FormatOptions are ignored with regard to defaults!
	matchDefaultExpectations(t, defaults, map[string]string{})
}
