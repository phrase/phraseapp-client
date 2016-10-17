package cli

import (
	"fmt"
	"sort"
)

func validateTagMap(tagMap map[string]string, allowedTags ...string) (e error) {
	sort.Strings(allowedTags)
	for usedTag, _ := range tagMap {
		idx := sort.SearchStrings(allowedTags, usedTag)
		if idx >= len(allowedTags) || allowedTags[idx] != usedTag {
			return fmt.Errorf("unknown tag %q", usedTag)
		}
	}
	return nil
}
