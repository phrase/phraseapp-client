package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var dirGlobOperator = string(filepath.Separator) + "**" + string(filepath.Separator)

// Glob supports * and ** globbing according to https://phraseapp.com/docs/developers/cli/configuration/#globbing
func Glob(pattern string) (matches []string, err error) {
	pattern = filepath.Clean(pattern)

	if strings.Count(pattern, dirGlobOperator) > 1 {
		return nil, fmt.Errorf("invalid pattern '%s': the ** globbing operator may only be used once in a pattern", pattern)
	}

	if strings.Contains(pattern, dirGlobOperator) {
		parts := strings.Split(pattern, dirGlobOperator)
		basePattern, endPattern := parts[0], parts[1]

		baseCandidates, err := filepath.Glob(basePattern)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern '%s': %s", pattern, err)
		}

		for _, base := range directoriesOnly(baseCandidates) {
			err = filepath.Walk(filepath.Clean(base), func(path string, info os.FileInfo, err error) error {
				if !info.IsDir() {
					return nil
				}

				matchesInBase, err := Glob(filepath.Join(path, endPattern))
				if err != nil {
					return err
				}

				matches = append(matches, matchesInBase...)
				return nil
			})
		}

	} else {
		candidates, err := filepath.Glob(pattern)
		if err != nil {
			return nil, fmt.Errorf("invalid pattern '%s': %s", pattern, err)
		}
		matches = filesOnly(candidates)
	}

	return matches, nil
}

func filter(candidates []string, f func(os.FileInfo) bool) []string {
	matches := []string{}
	for _, candidate := range candidates {
		fi, err := os.Stat(candidate)
		if err != nil {
			continue
		}

		if f(fi) {
			matches = append(matches, candidate)
		}
	}

	return matches
}

func filesOnly(candidates []string) []string {
	return filter(candidates, func(fi os.FileInfo) bool {
		return !fi.IsDir()
	})
}

func directoriesOnly(candidates []string) []string {
	return filter(candidates, func(fi os.FileInfo) bool {
		return fi.IsDir()
	})
}
