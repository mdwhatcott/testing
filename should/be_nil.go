package should

import (
	"fmt"
	"reflect"
)

// BeNil verifies that actual is the nil value.
func BeNil(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	if len(EXPECTED) > 0 {
		return fmt.Errorf("%w: please provide 0 expected values (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	if actual == nil || interfaceHasNilValue(actual) {
		return nil
	}
	return fmt.Errorf("%w: expected <nil>, got: %#v", errNilCheckFailed, actual)
}
func interfaceHasNilValue(actual interface{}) bool {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	nillable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map

	// Careful: reflect.Value.IsNil() will panic unless it's an interface, chan, map, func, slice, or ptr
	// Reference: http://golang.org/pkg/reflect/#Value.IsNil
	return nillable && value.IsNil()
}
