package paths

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

var dirGlobOperator = "**"

// dirGlobOperatorUseValid returns false if '**' occurs, but '/**/' doesn't and pattern does not start with '**/'.
func dirGlobOperatorUseValid(pattern string) bool {
	containsOperator := strings.Contains(pattern, dirGlobOperator)
	operatorIsOwnPathSegment := strings.Contains(pattern, string(filepath.Separator)+dirGlobOperator+string(filepath.Separator))
	startsWithOperator := strings.HasPrefix(pattern, dirGlobOperator+string(filepath.Separator))

	return !containsOperator || (operatorIsOwnPathSegment || startsWithOperator)
}

// Glob supports * and ** globbing according to https://phraseapp.com/docs/developers/cli/configuration/#globbing
func Glob(pattern string) (matches []string, err error) {
	pattern = filepath.Clean(pattern)
	pattern = escape(pattern)

	if strings.Count(pattern, dirGlobOperator) > 1 {
		return nil, fmt.Errorf("invalid pattern '%s': the ** globbing operator may only be used once in a pattern", pattern)
	}

	if !dirGlobOperatorUseValid(pattern) {
		return nil, fmt.Errorf("invalid pattern '%s': the ** globbing operator may only be used as path segment on its own, i.e. …/**/… or **/…", pattern)
	}

	if strings.Contains(pattern, dirGlobOperator) {
		parts := strings.Split(pattern, dirGlobOperator)
		basePattern, endPattern := filepath.Clean(parts[0]), filepath.Clean(parts[1])

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

// escape escapes characters which filepath.Glob would otherwise handle in a special way (except on Windows...)
func escape(s string) string {
	if runtime.GOOS == "windows" {
		return s
	}

	s = strings.Replace(s, "?", "\\?", -1)
	return strings.Replace(s, "[", "\\[", -1)
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
