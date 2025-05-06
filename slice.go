package anyutil

import (
	"fmt"
	"reflect"
)

// MapSlice converts a slice of K to a slice of T.
// This will only work if T is an interface{} type, i contains interfaces to T, or K is convertible to T by a type cast.
func MapSlice[T, K any](i []K) (o []T) { // T,K is in reverse order on purpose to allow K to be inferred
	var emptyT T
	var emptyK K
	var isConvertible bool

	tK := reflect.TypeOf(emptyK)
	tT := reflect.TypeOf(emptyT)
	if tT != nil && tK != nil {
		isConvertible = tK.ConvertibleTo(tT)
	}
	if i == nil {
		return
	}
	v := reflect.ValueOf(i)
	for idx := 0; idx < v.Len(); idx++ {
		var a any
		if isConvertible {
			a = v.Index(idx).Convert(tT).Interface()
		} else {
			a = v.Index(idx).Interface()
		}
		v2 := a.(T)
		o = append(o, v2)
	}
	return o
}

// MapSliceFunc converts a slice of K to a slice of T using the mapping function f.
func MapSliceFunc[K, T any](i []K, f func(j K) T) (o []T) {
	for _, v := range i {
		o = append(o, f(v))
	}
	return o
}

// IsSlice returns true if value is a slice
func IsSlice(value any) bool {
	return reflect.TypeOf(value).Kind() == reflect.Slice
}

func Join[K any](values []K, sep string) (out string) {
	if len(values) == 0 {
		return ""
	}
	out = fmt.Sprint(values[0])
	for _, v := range values[1:] {
		out += sep + fmt.Sprint(v)
	}
	return
}
