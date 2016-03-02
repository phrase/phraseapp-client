package cli

import (
	"fmt"
	"reflect"
	"strconv"
)

// Arguments are the strings given at the end of the command line.
type argument struct {
	field    string
	desc     string
	position int
	variadic bool
	required bool
	value    string
	values   []string
}

// Reflect the gathered information into the concrete action instance.
func (arg *argument) reflectTo(value reflect.Value) (e error) {
	if arg.value == "" && len(arg.values) == 0 {
		if arg.required && isFieldBlank(value, arg.field) {
			return fmt.Errorf("required argument not set")
		}
		return nil
	}

	switch {
	case arg.variadic:
		field := value.FieldByName(arg.field)
		return setSlice(field, arg.values)
	default:
		field := value.FieldByName(arg.field)
		return arg.setField(field, arg.value)
	}
}

func (arg *argument) setField(target reflect.Value, source string) (e error) {
	switch target.Kind() {
	case reflect.String:
		target.SetString(source)
	case reflect.Int, reflect.Int64:
		i, e := strconv.ParseInt(source, 0, 64)
		if e != nil {
			return fmt.Errorf(`argument %q at index "%d" has wrong type`, arg.field, arg.position)
		}
		target.SetInt(i)
	default:
		return fmt.Errorf("invalid type %q", target.Type().String())
	}
	return nil
}

func isFieldBlank(value reflect.Value, name string) bool {
	if value.Type().Kind() == reflect.Ptr {
		return true
	}
	switch c := value.FieldByName(name).Interface().(type) {
	case string:
		return c == ""
	case int:
		return c == 0
	default:
		return true
	}
}

func setSlice(target reflect.Value, source []string) (e error) {
	switch target.Type().Elem().Kind() {
	case reflect.String:
		target.Set(reflect.ValueOf(source))
	case reflect.Int, reflect.Int64:
		result := make([]int, 0, len(source))
		for _, i := range source {
			j, e := strconv.Atoi(i)
			if e != nil {
				return e
			}
			result = append(result, j)
		}
		target.Set(reflect.ValueOf(result))
	default:
		return fmt.Errorf("..invalid type %q", target.Type().String())
	}
	return nil
}

func (arg *argument) setValue(value string) {
	if arg.variadic {
		arg.values = append(arg.values, value)
	} else {
		arg.value = value
	}
}

func (a *argument) description() string {
	desc := fmt.Sprintf("    %s", a.shortDescription())
	desc += fmt.Sprintf("%-*s", 30-len(desc), " ") + a.desc
	return desc
}

func (a *argument) shortDescription() (desc string) {
	desc = "<" + a.field + ">"
	if !a.required {
		desc += "?"
	}
	if a.variadic {
		desc += "..."
	}
	return desc
}

func (a *action) argumentForPosition(argIdx int) *argument {
	for idx := range a.args {
		arg := a.args[idx]
		if arg.position == argIdx || (arg.variadic && arg.position <= argIdx) {
			return arg
		}
	}
	return nil
}

func (a *action) createArgument(field reflect.StructField, value reflect.Value, tagMap map[string]string) (e error) {
	if e := validateTagMap(tagMap, "type", "desc", "required"); e != nil {
		return fmt.Errorf("[argument:%s] %s", field.Name, e.Error())
	}

	arg := &argument{field: field.Name, position: 0}

	arg.required, e = handleRequired(tagMap)
	if e != nil {
		return e
	}

	arg.variadic = field.Type.Kind() == reflect.Slice

	arg.desc = handleDescription(tagMap)

	if len(a.args) == 0 {
		arg.position = 0
	} else {
		prevArg := a.args[len(a.args)-1]
		if prevArg.variadic {
			return fmt.Errorf("only last argument can be variadic")
		}
		arg.position = prevArg.position + 1
	}

	a.args = append(a.args, arg)
	return nil
}
