package helpers

import "strings"

// Contains returns true if str is an element of seq.
func Contains(seq []string, str string) bool {
	for _, elem := range seq {
		if str == elem {
			return true
		}
	}
	return false
}

// ContainsAnySub returns true if s contains at least one element in subs.
func ContainsAnySub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}

	return false
}
