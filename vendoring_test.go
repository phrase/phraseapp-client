package main

import (
	"os/exec"
	"sort"
	"strings"
	"testing"
)

func TestVendoring(t *testing.T) {
	root, err := loadRoot()
	if err != nil {
		t.Fatal(err)
	}
	b, err := exec.Command("go", "list", "-f", `{{ join .Deps "\n" }}`).CombinedOutput()
	if err != nil {
		t.Fatal(err)
	}
	with := []string{}
	for _, f := range strings.Fields(string(b)) {
		if strings.Contains(f, ".") && !strings.HasPrefix(f, root) {
			with = append(with, "  - "+f)
		}
	}
	sort.Strings(with)
	if len(with) > 0 {
		t.Errorf("dependencies not vendored:\n%s", strings.Join(with, "\n"))
	}
}

func loadRoot() (string, error) {
	root, err := exec.Command("go", "list").CombinedOutput()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(root)), nil
}
