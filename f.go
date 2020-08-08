package r

import (
	"reflect"
	"strings"
)

// F is a function that allows you to find a value in a struct n levels deep
// returning the value
// See f_test.go for examples
func F(subj interface{}, fName string) interface{} {
	if isNil(subj) {
		return nil
	}

	s := strings.Split(fName, ".")
	var current reflect.Value

	switch reflect.TypeOf(subj).Kind() {
	case reflect.Ptr:
		current = reflect.Indirect(reflect.ValueOf(subj))

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
		updated := F(next.Interface(), strings.Join(s[1:], "."))
		if next.IsValid() {
			return reflect.ValueOf(updated).Interface()
		}
	}
	return current.Interface()
}
