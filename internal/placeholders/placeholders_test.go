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
		{
			"config/locales/en.yml",
			"config/locales/<locale_code>.yml",
		}: {
			"locale_code": "en",
		},
		{
			"abc/defg/en.lproj/Localizable.strings",
			"./abc/defg/<locale_code>.lproj/Localizable.strings",
		}: {
			"locale_code": "en",
		},
		{
			"config/german/de.yml",
			"config/<locale_name>/<locale_code>.yml",
		}: {
			"locale_name": "german",
			"locale_code": "de",
		},
		{
			"config/german/de/de.yml",
			"config/<locale_name>/<locale_code>/*.yml",
		}: {
			"locale_name": "german",
			"locale_code": "de",
		},
		{
			"abc/en.lproj/MyStoryboard.strings",
			"./abc/<locale_code>.lproj/<tag>.strings",
		}: {
			"locale_code": "en",
			"tag":         "MyStoryboard",
		},
		{
			"no_tag/abc/play.en",
			"*/<tag>/<locale_name>.<locale_code>",
		}: {
			"locale_name": "play",
			"locale_code": "en",
			"tag":         "abc",
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

func TestToGlobbing(t *testing.T) {
	tests := []struct {
		path    string
		pattern string
	}{
		{
			path:    "abc/*.lproj/*.strings",
			pattern: "abc/<locale_code>.lproj/.strings",
		}, {
			path:    "abc/defg/*.lproj/*.strings",
			pattern: "abc/defg/<locale_code>.lproj/.strings",
		}, {
			path:    "abc/defg/*.lproj/Localizable.strings",
			pattern: "abc/defg/<locale_code>.lproj/Localizable.strings",
		},
	}
	for _, test := range tests {
		result := ToGlobbingPattern(test.pattern)

		if result != test.path {
			t.Errorf("expected result to be %q, but got %q", test.path, result)
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
