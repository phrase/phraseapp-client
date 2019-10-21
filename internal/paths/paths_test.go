package paths

import (
	"testing"
)

func TestIsPhraseAppYamlConfig(t *testing.T) {
	tests := []struct {
		test   string
		expect bool
	}{
		{
			test:   "abc/*/.phraseapp.yml",
			expect: true,
		}, {
			test:   "abc/*/abc.phraseapp.yml",
			expect: true,
		}, {
			test:   "abc/*/.phraseapp.yml/lol.yml",
			expect: false,
		}, {
			test:   "abc/*/.phraseapp.yml/.yml",
			expect: false,
		},
	}

	for _, test := range tests {
		was := IsPhraseYmlConfig(test.test)
		if was != test.expect {
			t.Fatalf("expected %q to yield %v, but was: %v", test.test, test.expect, was)
		}
	}
}
