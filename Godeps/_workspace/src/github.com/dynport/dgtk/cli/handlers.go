package cli

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"unicode/utf8"
)

func handleRequired(tagMap map[string]string) (required bool, e error) {
	if value, found := tagMap["required"]; found {
		switch value {
		case "true":
			return true, nil
		case "false":
			// ignore; default value is false anyway.
		default:
			return false, fmt.Errorf(`wrong value for "required" tag: %q`, value)
		}
	}
	return false, nil
}

func handleVariadic(tagMap map[string]string) (required bool, e error) {
	if value, found := tagMap["variadic"]; found {
		switch value {
		case "true":
			return true, nil
		case "false":
			// ignore; default value is false anyway.
		default:
			return false, fmt.Errorf(`wrong value for "variadic" tag: %q`, value)
		}
	}
	return false, nil
}

func handlePresetValue(field reflect.StructField, value reflect.Value) string {
	switch field.Type.Kind() {
	case reflect.String:
		if v := value.String(); v != "" {
			return v
		}
	case reflect.Int:
		if v := value.Int(); v != 0 {
			return strconv.Itoa(int(v))
		}
	case reflect.Bool:
		if v := value.Bool(); v {
			return "true"
		}
	}

	return ""
}

func handleDefault(field reflect.StructField, tagMap map[string]string) (value string, e error) {
	if value, found := tagMap["default"]; found {
		switch field.Type.Kind() {
		case reflect.String:
			return value, nil
		case reflect.Int:
			_, e := strconv.Atoi(value) // just error checking
			if e != nil {
				return "", e
			}
			return value, nil
		case reflect.Bool:
			if value == "true" || value == "false" {
				return value, nil
			}
			return "", fmt.Errorf(`value of tag "default" for field %q must be "true" or "false" (not %q)"`, field.Name, value)
		default:
			return "", fmt.Errorf("unknown type: %q", field.Type.String())
		}
	}
	return "", nil
}

func handleShortIdentifier(tagMap map[string]string) (short string, e error) {
	if value, found := tagMap["short"]; found {
		if len(value) != 1 {
			return "", fmt.Errorf("short identifier must be a single character")
		}
		r, _ := utf8.DecodeRuneInString(value)
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') {
			return value, nil
		}
		return "", fmt.Errorf("%q not a valid short identifier (use chars from [A-Za-z])", value)
	}
	return "", nil
}

func handleDescription(tagMap map[string]string) (desc string) {
	if value, found := tagMap["desc"]; found {
		return value
	}
	return ""
}

var longOptionRE = regexp.MustCompile("^[a-zA-Z][a-zA-Z0-9-]+$")

func handleLongIdentifier(tagMap map[string]string) (long string, e error) {
	if value, found := tagMap["long"]; found {
		e := validateLongOption(value)
		if e != nil {
			return "", e
		}
		return value, nil
	}
	return "", nil
}

func validateLongOption(value string) error {
	switch {
	case len(value) <= 1:
		return fmt.Errorf("long options must have at least two characters")
	case !longOptionRE.MatchString(value):
		return fmt.Errorf("long options must only have characters from " + longOptionRE.String())
	}
	return nil
}
