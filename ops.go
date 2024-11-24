// Package any has general purpose utility functions for working with interfaces and generic types.
package any

import (
	"cmp"
	"reflect"
)

// If returns the first item if cond is not a zero value (like false), or the second item if it is the zero value.
func If[T1 comparable, T any](cond T1, i1 T, i2 T) T {
	if cond != Zero[T1]() {
		return i1
	} else {
		return i2
	}
}

// Or returns the first of its arguments that is not equal to the zero value.
// If no argument is non-zero, it returns the zero value.
func Or[T comparable](vals ...T) T {
	return cmp.Or(vals...)
}

// Zero returns the zero value of a type.
func Zero[T any]() T {
	var v T
	return v
}

// IsNil is a safe test for nil for any kind of variable, and will not panic
// If i points to a nil object, IsNil will return true, as opposed to i==nil which will return false
func IsNil(i any) bool {
	if i == nil {
		return true
	}
	v := reflect.ValueOf(i)
	k := v.Kind()
	switch k {
	case reflect.Chan:
		fallthrough
	case reflect.Func:
		fallthrough
	case reflect.Interface:
		fallthrough
	case reflect.Map:
		fallthrough
	case reflect.Ptr:
		fallthrough
	case reflect.Slice:
		return v.IsNil()
	}
	return false
}
