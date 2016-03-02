package cli

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
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
	isMap    bool
	mapValue map[string]string
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
	case reflect.Int, reflect.Int64:
		i, e := strconv.ParseInt(o.value, 0, 64)
		if e != nil {
			return e
		}
		field.SetInt(i)
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
			case reflect.Int, reflect.Int64:
				val, e := strconv.ParseInt(parts[i], 0, 64)
				if e != nil {
					return e
				}

				sl.Index(i).SetInt(val)
			default:
				return fmt.Errorf("invalid type %q for slice", st.String())
			}
			field.Set(sl)
		}
	case reflect.Map:
		ml := reflect.MakeMap(field.Type())

		valueType := field.Type().Elem()

		for k, v := range o.mapValue {
			var val interface{}
			switch valueType.Kind() {
			case reflect.String:
				val = v
			case reflect.Int64:
				ival, err := strconv.ParseInt(v, 10, 64)
				if err != nil {
					return fmt.Errorf("option value is not a valid integer: %s", err)
				}
				val = ival
			default:
				return fmt.Errorf("invalid type %q for slice", valueType.Kind())
			}
			ml.SetMapIndex(reflect.ValueOf(k), reflect.ValueOf(val))
		}

		field.Set(ml)
	default:
		if field.Type().String() == "*time.Time" {
			t := &time.Time{}
			err := t.UnmarshalText([]byte(o.value))
			if err != nil {
				return fmt.Errorf("Expected time in RFC3339 format, e.g. 2016-10-14T01:23:00Z01:00, got: %s", err.Error())
			}
			field.Set(reflect.ValueOf(t))
			return nil
		}

		return fmt.Errorf("%s, invalid type %q", field.Kind(), field.Type().String())
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

	switch field.Type.Kind() {
	case reflect.Bool:
		opt.isFlag = true
	case reflect.Map:
		keyType := field.Type.Key().Kind()
		if keyType != reflect.String {
			return fmt.Errorf("Options with Map type must have string keys, got %T.", keyType)
		}

		valueType := field.Type.Elem().Kind()
		switch valueType {
		case reflect.String, reflect.Int64:
			opt.isMap = true
			opt.mapValue = map[string]string{}
		default:
			return fmt.Errorf("Options with Map type must have string or int64 values, got %s", valueType)
		}
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
