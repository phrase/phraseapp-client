package cli

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type option struct {
	field    string
	isFlag   bool
	desc     string
	short    string
	long     string
	required bool
	value    string
	given    bool
}

// Reflect the gathered information into the concrete action instance.
func (o *option) reflectTo(value reflect.Value) (e error) {
	if !o.given && o.value == "" {
		if o.required {
			return fmt.Errorf("option %q is required but not set", o.field)
		}
		return nil
	}

	field := value.FieldByName(o.field)
	if field.Kind() == reflect.Ptr {
		n := reflect.New(field.Type().Elem())
		field.Set(n)
		field = field.Elem()
	}

	switch field.Kind() {
	case reflect.String:
		field.SetString(o.value)
	case reflect.Int:
		i, e := strconv.Atoi(o.value)
		if e != nil {
			return e
		}
		field.SetInt(int64(i))
	case reflect.Bool:
		field.SetBool(o.value == "true")
	case reflect.Slice:
		parts := strings.Split(o.value, ",")
		st := field.Type().Elem()
		sl := reflect.MakeSlice(field.Type(), len(parts), len(parts))
		for i := range parts {
			switch st.Kind() {
			case reflect.String:
				sl.Index(i).SetString(parts[i])
			case reflect.Int:
				val, e := strconv.Atoi(parts[i])
				if e != nil {
					return e
				}
				sl.Index(i).SetInt(int64(val))
			default:
				return fmt.Errorf("invalid type %q for slice", st.String())
			}
			field.Set(sl)
		}
	default:
		return fmt.Errorf("invalid type %q", field.Type().String())
	}
	return nil
}

func (o *option) description() string {
	desc := "    "
	desc += o.shortDescription(" ")
	desc += fmt.Sprintf("%-*s", 30-len(desc), " ") + o.desc
	if o.value != "" {
		if o.desc != "" {
			desc += " "
		}
		desc += "(default: " + o.value + ")"
	}
	return desc
}

func (o *option) shortDescription(sep string) (desc string) {
	if o.short != "" {
		desc += "-" + o.short
	}
	if o.long != "" {
		if o.short != "" {
			desc += sep
		}
		desc += "--" + o.long
	}

	if !o.isFlag {
		desc += " <" + o.field + ">"
	}

	return desc
}

func (a *action) createOption(field reflect.StructField, value reflect.Value, tagMap map[string]string) (e error) {
	if e := validateTagMap(tagMap, "type", "desc", "short", "long", "required", "default"); e != nil {
		return fmt.Errorf("[option:%s] %s", field.Name, e.Error())
	}
	opt := &option{field: field.Name}

	if field.Type.Kind() == reflect.Bool {
		opt.isFlag = true
	}

	opt.short, e = handleShortIdentifier(tagMap)
	if e != nil {
		return e
	}

	opt.long, e = handleLongIdentifier(tagMap)
	if e != nil {
		return e
	}

	opt.required, e = handleRequired(tagMap)
	if e != nil {
		return e
	}

	if opt.required && opt.isFlag {
		return fmt.Errorf(`field %q is a flag and required, that doesn't make much sense`, field.Name)
	}

	opt.value = handlePresetValue(field, value)

	// Do that for error checking (not necessary if value preset otherwise).
	defaultValue, e := handleDefault(field, tagMap)
	if e != nil {
		return fmt.Errorf(`wrong value for "default" tag: %s`, e.Error())
	}

	if opt.value == "" {
		opt.value = defaultValue
	} else { // With a preset value that value is not required any more.
		opt.required = false
	}

	opt.desc = handleDescription(tagMap)

	if opt.short == "" && opt.long == "" {
		return fmt.Errorf("option %q has neither long nor short accessor set", field.Name)
	}

	if opt.short != "" {
		if o, found := a.params[opt.short]; found {
			return fmt.Errorf("short option %q already used (option %q)", opt.short, o.field)
		}
		a.params[opt.short] = opt
	}
	if opt.long != "" {
		if o, found := a.params[opt.long]; found {
			return fmt.Errorf("long option %q already used (option %q)", opt.long, o.field)
		}
		a.params[opt.long] = opt
	}
	a.opts = append(a.opts, opt)
	return nil
}
