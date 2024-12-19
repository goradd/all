package all

import (
	"reflect"
)

// ConvertSlice converts a slice of values to a slice of T.
// This will only work if T is an interface{} type,
// or if i contains interfaces to T.
func ConvertSlice[T, K any](i []K) (o []T) {
	if i == nil {
		return
	}
	v := reflect.ValueOf(i)
	for idx := 0; idx < v.Len(); idx++ {
		a := v.Index(idx).Interface()
		v2 := a.(T)
		o = append(o, v2)
	}
	return o
}

func IsSlice(in any) bool {
	return reflect.TypeOf(in).Kind() == reflect.Slice
}
