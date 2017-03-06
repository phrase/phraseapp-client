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

// RemoveDuplicates returns s with all duplicate elements removed.
func RemoveDuplicates(s []string) []string {
	seen := map[string]bool{}
	clean := []string{}

	for _, elem := range s {
		if !seen[elem] {
			seen[elem] = true
			clean = append(clean, elem)
		}
	}

	return clean
}
