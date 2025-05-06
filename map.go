package anyutil

import (
	"cmp"
	"fmt"
	"slices"
)

// StringMap converts a map[K]any to a map[K]string.
// It will attempt to cast to a string, and if that fails, will
// use fmt.Print to do the conversion to string.
func StringMap[K cmp.Ordered](m map[K]any) (o map[K]string) {
	o = make(map[K]string, len(m))
	for k, v := range m {
		if s, ok := v.(string); ok {
			o[k] = s
		} else {
			o[k] = fmt.Sprintf("%v", v)
		}
	}
	return
}

// AnyMap converts a map[K]V to a map[K]any.
func AnyMap[K cmp.Ordered, V any](m map[K]V) (o map[K]any) {
	o = make(map[K]any, len(m))
	for k, v := range m {
		o[k] = v
	}
	return
}

// SortedKeys returns the keys of a map in sort order.
// The keys must be sortable, of course.
func SortedKeys[K cmp.Ordered, V any](m map[K]V) []K {
	keys := make([]K, len(m), len(m))
	idx := 0
	for k := range m {
		keys[idx] = k
		idx++
	}
	slices.Sort(keys)
	return keys
}
