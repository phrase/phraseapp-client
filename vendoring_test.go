package main

import (
	"os"
	"os/exec"
	"sort"
	"strings"
	"testing"
)

var notVendored = map[string]struct{}{
	"vendor/golang.org/x/": struct{}{},
}

func TestVendoring(t *testing.T) {
	root, err := loadRoot()
	if err != nil {
		t.Fatal(err)
	}
	b, err := exec.Command("go", "list", "-f", `{{ join .Deps "\n" }}`, "./...").CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	with := map[string]struct{}{}
	for _, f := range strings.Fields(string(b)) {
		if strings.HasPrefix(f, "vendor/golang.org/x") {
			continue
		}
		_, ok := notVendored[f]
		if strings.Contains(f, ".") && !strings.HasPrefix(f, root) && !ok {
			with[f] = struct{}{}
		}
	}
	if len(with) > 0 {
		keys := []string{}
		for k := range with {
			if s, err := os.Stat("vendor/" + k); err != nil || !s.IsDir() {
				keys = append(keys, k)
			}
		}
		if len(keys) > 0 {
			sort.Strings(keys)
			t.Errorf("dependencies not vendored:\n%s", strings.Join(keys, "\n"))
		}
	}
}

func loadRoot() (string, error) {
	root, err := exec.Command("go", "list").CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(root)), nil
}
