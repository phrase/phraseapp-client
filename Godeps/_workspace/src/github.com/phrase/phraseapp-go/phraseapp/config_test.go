package phraseapp

import (
	"fmt"
	"os"
	"testing"
)

func TestValidateIsType(t *testing.T) {
	var t1 string = "foobar"
	var t2 int = 1
	var t3 bool = true
	expErrT1 := fmt.Sprintf(cfgValueErrStr, "a", t1)
	expErrT2 := fmt.Sprintf(cfgValueErrStr, "a", t2)
	expErrT3 := fmt.Sprintf(cfgValueErrStr, "a", t3)

	switch res, err := ValidateIsString("a", t1); {
	case err != nil:
		t.Errorf("didn't expect an error, got %q", err)
	case res != t1:
		t.Errorf("expected value to be %q, got %q", t1, res)
	}

	switch _, err := ValidateIsString("a", t2); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT2:
		t.Errorf("expected error to be %q, got %q", expErrT2, err)
	}

	switch _, err := ValidateIsString("a", t3); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT3:
		t.Errorf("expected error to be %q, got %q", expErrT3, err)
	}

	switch _, err := ValidateIsInt("a", t1); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT1:
		t.Errorf("expected error to be %q, got %q", expErrT1, err)
	}

	switch res, err := ValidateIsInt("a", t2); {
	case err != nil:
		t.Errorf("didn't expect an error, got %q", err)
	case res != t2:
		t.Errorf("expected value to be %q, got %q", t2, res)
	}

	switch _, err := ValidateIsInt("a", t3); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT3:
		t.Errorf("expected error to be %q, got %q", expErrT3, err)
	}

	switch _, err := ValidateIsBool("a", t1); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT1:
		t.Errorf("expected error to be %q, got %q", expErrT1, err)
	}

	switch _, err := ValidateIsBool("a", t2); {
	case err == nil:
		t.Errorf("expect an error, got none")
	case err.Error() != expErrT2:
		t.Errorf("expected error to be %q, got %q", expErrT2, err)
	}

	switch res, err := ValidateIsBool("a", t3); {
	case err != nil:
		t.Errorf("didn't expect an error, got %q", err)
	case res != t3:
		t.Errorf("expected value to be %t, got %t", t3, res)
	}
}

func TestValidateIsRawMapHappyPath(t *testing.T) {
	m := map[interface{}]interface{}{
		"foo": "bar",
		"fuu": 1,
		"few": true,
	}

	res, err := ValidateIsRawMap("a", m)
	if err != nil {
		t.Errorf("didn't expect an error, got %q", err)
	}

	if len(m) != len(res) {
		t.Errorf("expected %d elements, got %d", len(m), len(res))
	}

	for k, v := range res {
		if value, found := m[k]; !found {
			t.Errorf("expected key %q to be in source set, it wasn't", k)
		} else if value != v {
			t.Errorf("expected value of %q to be %q, got %q", k, value, v)
		}
	}
}

func TestValidateIsRawMapWithErrors(t *testing.T) {
	m := map[interface{}]interface{}{
		4: "should be error",
	}

	expErr := fmt.Sprintf(cfgKeyErrStr, "a.4", 4)
	_, err := ValidateIsRawMap("a", m)
	if err == nil {
		t.Errorf("expect an error, got none")
	} else if err.Error() != expErr {
		t.Errorf("expected error %q, got %q", expErr, err)
	}
}

func TestParseYAMLToMap(t *testing.T) {
	var a string
	var b int
	var c bool
	var d []byte
	e := map[string]interface{}{}

	err := ParseYAMLToMap(func(raw interface{}) error {
		m, ok := raw.(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid type received")
		}
		m["a"] = "foo"
		m["b"] = 1
		m["c"] = true
		m["d"] = &struct {
			A string
			B int
		}{A: "bar", B: 2}
		m["e"] = map[interface{}]interface{}{"c": "baz", "d": 4}
		return nil
	}, map[string]interface{}{
		"a": &a,
		"b": &b,
		"c": &c,
		"d": &d,
		"e": &e,
	})
	if err != nil {
		t.Fatalf("didn't expect an error, got %q", err)
	}

	if a != "foo" {
		t.Errorf("expected %q, got %q", "foo", a)
	}

	if b != 1 {
		t.Errorf("expected %d, got %d", 1, b)
	}

	if c != true {
		t.Errorf("expected %t, got %t", true, c)
	}

	if string(d) != "a: bar\nb: 2\n" {
		t.Errorf("expected %s, got %s", "a: bar\nb: 2\n", string(d))
	}

	if val, found := e["c"]; !found {
		t.Errorf("expected e to contain key %q, it didn't", "c")
	} else if val != "baz" {
		t.Errorf("expected e['c'] to have value %q, got %q", "baz", val)
	}

	if val, found := e["d"]; !found {
		t.Errorf("expected e to contain key %q, it didn't", "d")
	} else if val != 4 {
		t.Errorf("expected e['d'] to have value %d, got %d", 4, val)
	}
}

func TestConfigPath_ConfigFromEnv(t *testing.T) {
	// The phraseapp.yml file without the leading '.' so not hidden. Any file can be used from the environment!
	p := os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata/phraseapp.yml")

	os.Setenv("PHRASEAPP_CONFIG", p)
	defer os.Unsetenv("PHRASEAPP_CONFIG")

	path, err := configPath()
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	} else if path != p {
		t.Errorf("expected path to be %q, got %q", p, path)
	}
}

func TestConfigPath_ConfigFromEnvButNotExisting(t *testing.T) {
	os.Setenv("PHRASEAPP_CONFIG", "phraseapp_does_not_exist.yml")
	defer os.Unsetenv("PHRASEAPP_CONFIG")

	_, err := configPath()
	if err == nil {
		t.Fatalf("expect an error, got none")
	}

	expErr := `file "phraseapp_does_not_exist.yml" (given in PHRASEAPP_CONFIG) doesn't exist`
	if err.Error() != expErr {
		t.Errorf("expected error to be %q, got %q", expErr, err)
	}
}

func TestConfigPath_ConfigInCWD(t *testing.T) {
	cwd := os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata")

	oldDir, _ := os.Getwd()
	err := os.Chdir(cwd)
	if err != nil {
		t.Fatalf("didn't expect an error changing the working directory, got: %s", err)
	}
	defer os.Chdir(oldDir)

	path, err := configPath()
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}
	expPath := cwd + "/.phraseapp.yml"
	if path != expPath {
		t.Errorf("expected path to be %q, got %q", expPath, path)
	}
}

func TestConfigPath_ConfigInHomeDir(t *testing.T) {
	cwd := os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata/empty")
	oldDir, _ := os.Getwd()
	err := os.Chdir(cwd)
	if err != nil {
		t.Fatalf("didn't expect an error changing the working directory, got: %s", err)
	}
	defer os.Chdir(oldDir)

	newHome := os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata")
	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", newHome)
	defer os.Setenv("HOME", oldHome)

	path, err := configPath()
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}
	expPath := newHome + "/.phraseapp.yml"
	if path != expPath {
		t.Errorf("expected path to be %q, got %q", expPath, path)
	}
}

func TestConfigPath_NoConfigAvailable(t *testing.T) {
	// For this to work the configuration of the user running the test
	// must be obfuscated (changing the CWD and HOME env variable), so
	// user's files do not inflict the test environment.

	cwd := os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata/empty")
	oldDir, _ := os.Getwd()
	err := os.Chdir(cwd)
	if err != nil {
		t.Fatalf("didn't expect an error changing the working directory, got: %s", err)
	}
	defer os.Chdir(oldDir)

	oldHome := os.Getenv("HOME")
	os.Setenv("HOME", os.ExpandEnv("$GOPATH/src/github.com/phrase/phraseapp-go/testdata/empty2"))
	defer os.Setenv("HOME", oldHome)

	path, err := configPath()
	if err != nil {
		t.Fatalf("didn't expect an error, got: %s", err)
	}
	expPath := ""
	if path != expPath {
		t.Errorf("expected path to be %q, got %q", expPath, path)
	}
}
