package placeholders

import "testing"

func TestResolve(t *testing.T) {
	tests := map[struct {
		path    string
		pattern string
	}]map[string]string{
		{
			"english-en_foo.yml",
			"<locale_name>-<locale_code>_<tag>.yml",
		}: {
			"locale_name": "english",
			"locale_code": "en",
			"tag":         "foo",
		},
		{
			"asd-bla.json",
			"*-<tag>*.json",
		}: {
			"tag": "bla",
		},
	}

	for input, expected := range tests {
		result, err := Resolve(input.path, input.pattern)
		if err != nil {
			t.Error(err)
		}

		if !areEqual(result, expected) {
			t.Errorf("got %v, but want %v", result, expected)
		}
	}
}

func TestResolve_errorPlaceholderReuse(t *testing.T) {
	tests := []struct {
		path    string
		pattern string
		err     string
	}{
		{
			path:    "en_foo.yml",
			pattern: "<locale_code>_<locale_code>.yml",
			err:     `string "en_foo.yml" does not match pattern "(?P<locale_code>[^/]+)_(?P<locale_code>[^/]+)\\.yml": placeholder "locale_code" is used twice with different values`,
		},
	}

	for _, test := range tests {
		result, err := Resolve(test.path, test.pattern)

		if result != nil {
			t.Errorf("expected result to be <nil>, but got %v", result)
		}

		if err.Error() != test.err {
			t.Errorf("expected error to be %q, but got %q", test.err, err)
		}
	}
}

func areEqual(got, want map[string]string) bool {
	for kw, vw := range want {
		vg, ok := got[kw]

		if !ok || vg != vw {
			return false
		}
	}

	return true
}
