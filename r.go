package r

import (
	"reflect"
	"strings"
)

// R is a function that allows you to replace a value in a struct n levels deep
// returning a new copy containing the required val
// See r_test.go for examples
func R(subj interface{}, fName string, val interface{}) interface{} {
	if isNil(subj) {
		return subj
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
			field.Set(reflect.ValueOf(val))
		}
	} else {
		next := current.FieldByName(s[0])
		updated := R(next.Interface(), strings.Join(s[1:], "."), val)
		if next.IsValid() {
			next.Set(reflect.ValueOf(updated))
		}
	}
	return current.Interface()
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}
	switch reflect.TypeOf(i).Kind() {
	case reflect.Ptr, reflect.Map, reflect.Chan, reflect.Slice:
		return reflect.ValueOf(i).IsNil()
	}
	return false
}
