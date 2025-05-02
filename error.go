package all

import "errors"

// As is a generic version of errors.As.
// It returns the first error in err’s chain that matches type T, return true if found.
// T should be a pointer to a custom error struct type.
// For example: *MyCustomError
func As[T error](err error) (T, bool) {
	var target T
	ok := errors.As(err, &target)
	return target, ok
}
