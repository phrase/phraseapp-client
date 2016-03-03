package version

import "testing"

func TestNewFromString(t *testing.T) {
	v, err := NewFromString("1.2.3")
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		Name     string
		Expected interface{}
		Value    interface{}
	}{
		{"Major", v.Major, 1},
		{"Minor", v.Minor, 2},
		{"Patch", v.Patch, 3},
	}
	for i, tst := range tests {
		if tst.Expected != tst.Value {
			t.Errorf("test %d (%s): expected %d, was %d", i+1, tst.Name, tst.Expected, tst.Value)
		}
	}
}

func TestParseVersion(t *testing.T) {
	v := &Version{}
	err := v.Parse("1.2.3")
	if err != nil {
		t.Fatal("error parsing version", err)
	}
	if v.Major != 1 {
		t.Errorf("expected Major to be %v, was %v", 1, v.Major)
	}
	if v.Minor != 2 {
		t.Errorf("expected Minor to be %v, was %v", 2, v.Minor)
	}
	if v.Patch != 3 {
		t.Errorf("expected Patch to be %v, was %v", 3, v.Patch)
	}

	a, err := Parse("0.1.2")
	if err != nil {
		t.Fatal("error parsing version", err)
	}
	versions := []string{"0.0.0", "0.0.9", "0.1.1"}
	for i, _ := range versions {
		v := versions[i]
		b, err := Parse(v)
		if err != nil {
			t.Fatalf("error parsing version %v: %v", v, err)
		}
		if a.Less(b) {
			t.Errorf("expected %v to not be less than %v", a, b)
		}
	}

	versions = []string{"1.0.0", "0.2.0", "0.1.3"}
	for i := range versions {
		v := versions[i]
		b, err := Parse(v)
		if err != nil {
			t.Errorf("error parsing version %v: %v", v, err)
		}
		if !a.Less(b) {
			t.Errorf("expected %v to be less than %v", a, b)
		}
	}
}
