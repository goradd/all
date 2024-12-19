package all

import (
	"fmt"
	"reflect"
)

// FieldMap returns a map of names and values of the top-level exported fields in struct s.
// Anonymous fields are ignored, as well as non-exported fields.
// s can be a pointer or an interface to a struct as well.
func FieldMap(s any) map[string]any {
	ret := make(map[string]interface{})
	sVal := reflect.ValueOf(s)
	typ := sVal.Type()
	k := sVal.Kind()
	if k == reflect.Ptr || k == reflect.Interface {
		sVal = sVal.Elem()
		typ = sVal.Type()
		k = sVal.Kind()
	}
	numFields := typ.NumField()
	for i := 0; i < numFields; i++ {
		field := typ.Field(i)
		if field.Anonymous ||
			field.PkgPath != "" { // non-exported fields have a package path
			continue // Do not worry about anonymous fields, since they are really part of the struct itself
		}
		ret[field.Name] = sVal.FieldByName(field.Name).Interface()
	}
	return ret
}

// SetFields will restore to s values that were extracted using FieldMap
func SetFields(s any, values map[string]any) error {
	sVal := reflect.ValueOf(s)
	k := sVal.Kind()
	if k == reflect.Ptr || k == reflect.Interface {
		sVal = sVal.Elem()
		k = sVal.Kind()
	}

	if sVal.Kind() != reflect.Struct {
		return fmt.Errorf("Not a struct")
	}
	for name, val := range values {
		field := sVal.FieldByName(name)
		if field.IsValid() {
			field.Set(reflect.ValueOf(val))
		}
	}
	return nil
}
