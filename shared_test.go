package main

import (
	"fmt"
	"testing"
)

func TestSourceCheckPreconditions(t *testing.T) {
	fmt.Println("Shared#CheckPreconditions")
	for _, file := range []string{
		"",
		"no_extension",
		"./<locale_code>/<locale_code>.yml",
		"./**/**/en.yml",
		"./**/*/*/en.yml",
	} {
		if err := CheckPreconditions(file); err == nil {
			t.Errorf("CheckPrecondition did not fail for pattern: '%s'", file)
		}
	}

	for _, file := range []string{
		"./<tag>/<locale_code>.yml",
		"./**/*/en.yml",
		"./**/*/<locale_name>/<locale_code>/<tag>.yml",
	} {
		if err := CheckPreconditions(file); err != nil {
			t.Errorf("CheckPrecondition should not fail with: %s", err.Error())
		}
	}
}
