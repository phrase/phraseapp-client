// Package to parse a tag.
//
// Tags are annotation with struct fields, that allow for some elegant solutions for different problems. Examples are
// validation and alike. This is an example on what that might look like:
//	type Example struct {
//		AField string `tagName1:"tagValue1" tagName2:"tagValue2"`
//	}
//
// As the syntax is somewhat weird and the tag interface only supports a getter, this tag parser was written. It will
// use go's tag parser to retrieve the tag for a given prefix, parse the value, and return a map of strings to strings.
// Keys and values are separated by an equal sign '=', values might be quoted using single quotes "'", and key-value
// pairs are separated using whitespace.
package tagparse

import (
	"fmt"
	"reflect"
	"strings"
)

func tagSplit(tag string) (fields []string, e error) {
	fields = []string{}
	idxStart := 0
	quoted := false
	for i, c := range tag {
		if i == idxStart && (c == ' ' || c == '\t') {
			idxStart = i + 1
			continue
		}
		if c == '\'' {
			quoted = !quoted
		}

		if (c == ' ' || i+1 == len(tag)) && !quoted {
			fields = append(fields, strings.TrimSpace(tag[idxStart:i+1]))
			idxStart = i + 1
		}
	}
	if quoted {
		return nil, fmt.Errorf("failed to parse tag due to erroneous quotes")
	}
	return fields, nil
}

func parseTag(tagString string, customS CustomSplitter) (result map[string]string, e error) {
	fields, e := tagSplit(tagString)
	if e != nil {
		return nil, e
	}

	result = map[string]string{}
	for fIdx := range fields {
		kvList := strings.SplitN(fields[fIdx], "=", 2)
		var key, value string
		if len(kvList) == 2 {
			key = strings.TrimSpace(kvList[0])
			value = strings.Trim(kvList[1], "'")
		} else {
			key, value, e = customS(fields[fIdx])
			if e != nil {
				return nil, fmt.Errorf("failed to parse annotation (value missing): %q", fields[fIdx])
			}
		}

		if _, found := result[key]; found {
			return nil, fmt.Errorf("key %q set multiple times", key)
		}
		result[key] = value
	}
	return result, nil
}

type CustomSplitter func(string) (string, string, error)

// Parse tag with the given prefix of the given field. Return a map of strings to strings. If errors occur they are
// returned accordingly.
func Parse(field reflect.StructField, prefix string, customHander func(string) (string, string, error)) (result map[string]string, e error) {
	return ParseCustom(field, prefix, func(_ string) (string, string, error) { return "", "", fmt.Errorf("failed") })
}

// Like Parse, but with a custom splitter used for tag values that don't have the form `key=value`.
func ParseCustom(field reflect.StructField, prefix string, customF CustomSplitter) (result map[string]string, e error) {
	tagString := field.Tag.Get(prefix)

	return parseTag(tagString, customF)
}
