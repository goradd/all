package any

import (
	"fmt"
	"golang.org/x/exp/constraints"
)

// StringMap converts a map[K]any to a map[K]string.
// It will attempt to cast to a string, and if that fails, will
// use fmt.Print to do the conversion to string.
func StringMap[K constraints.Ordered](m map[K]any) (o map[K]string) {
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

// Map converts a map[K]V to a map[K]any.
func Map[K constraints.Ordered, V any](m map[K]V) (o map[K]any) {
	o = make(map[K]any, len(m))
	for k, v := range m {
		o[k] = v
	}
	return
}