package stringz

import "strings"

// Contains returns true if needle is an element of haystack.
func Contains(haystack []string, needle string) bool {
	for _, elem := range haystack {
		if elem == needle {
			return true
		}
	}
	return false
}

// ContainsAnySub returns true if s contains at least one element of subs.
func ContainsAnySub(s string, subs []string) bool {
	for _, sub := range subs {
		if strings.Contains(s, sub) {
			return true
		}
	}

	return false
}
