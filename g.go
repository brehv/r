package r

import (
	"reflect"
	"strings"
)

// G is a function that allows you to get a value in a struct n levels deep.
func G(subj interface{}, fName string) interface{} {
	if isNil(subj) {
		return subj
	}

	s := strings.Split(fName, ".")
	var current reflect.Value

	switch reflect.TypeOf(subj).Kind() {
	case reflect.Ptr:
		current = reflect.Indirect(reflect.ValueOf(subj))
	case reflect.Map:
		if len(s) == 1 {
			return reflect.ValueOf(subj).MapIndex(reflect.ValueOf(s[0])).Interface()
		}
		return G(reflect.ValueOf(subj).MapIndex(reflect.ValueOf(s[0])).Interface(), strings.Join(s[1:], "."))
	default:
		original := reflect.ValueOf(subj)
		current = reflect.New(original.Type()).Elem()
		for j := 0; j < original.NumField(); j++ {
			f := original.Field(j)
			n := original.Type().Field(j).Name
			current.FieldByName(n).Set(f)
		}
	}

	if len(s) == 1 {
		field := current.FieldByName(fName)
		if field.IsValid() {
			return field.Interface()
		}
	} else {
		next := current.FieldByName(s[0])
		updated := G(next.Interface(), strings.Join(s[1:], "."))
		if next.IsValid() {
			return updated
		}
	}
	return current.Interface()
}
