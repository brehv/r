package r

import (
	"reflect"
	"strconv"
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
	if len(s) == 0 {
		return nil
	}
	switch reflect.TypeOf(subj).Kind() {
	case reflect.Struct, reflect.Ptr:
		return fStruct(subj, s)
	case reflect.Map:
		return fMap(subj, s)
	case reflect.Slice, reflect.Array:
		return fSlice(subj, s)
	default:
		return nil
	}
}

func fStruct(subj interface{}, s []string) interface{} {
	if isNil(subj) {
		return nil
	}

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
			if !current.FieldByName(n).CanSet() {
				return nil
			}
			current.FieldByName(n).Set(f)
		}
	}
	if len(s) == 1 {
		field := current.FieldByName(s[0])
		if field.IsValid() {
			return field.Interface()
		}
	} else {
		next := current.FieldByName(s[0])
		updated := F(next.Interface(), strings.Join(s[1:], "."))
		if updated == nil {
			return nil
		}
		if next.IsValid() {
			return reflect.ValueOf(updated).Interface()
		}
	}
	return current.Interface()
}

func fMap(subj interface{}, s []string) interface{} {
	if isNil(subj) {
		return nil
	}
	current := reflect.ValueOf(subj)
	if len(s) == 1 {
		key := reflect.ValueOf(s[0])
		field := current.MapIndex(key)
		if field.IsValid() {
			return field.Interface()
		}
	} else {
		key := reflect.ValueOf(s[0])
		next := current.MapIndex(key)
		updated := F(next.Interface(), strings.Join(s[1:], "."))
		if updated == nil {
			return nil
		}
		if next.IsValid() {
			return reflect.ValueOf(updated).Interface()
		}
	}
	return nil
}

func fSlice(subj interface{}, s []string) interface{} {
	if isNil(subj) {
		return nil
	}
	index, err := strconv.Atoi(s[0])
	if err != nil {
		return nil
	}
	current := reflect.ValueOf(subj)
	if index > current.Len() {
		return nil
	}
	if len(s) == 1 {
		field := current.Index(index)
		if field.IsValid() {
			return field.Interface()
		}
	} else {
		next := current.Index(index)
		updated := F(next.Interface(), strings.Join(s[1:], "."))
		if updated == nil {
			return nil
		}
		if next.IsValid() {
			return reflect.ValueOf(updated).Interface()
		}
	}
	return nil
}