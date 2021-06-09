package should

import (
	"fmt"
	"reflect"
)

// BeEmpty uses reflection to verify that len(actual) == 0.
func BeEmpty(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, kindsWithLength...)
	if err != nil {
		return err
	}

	length := reflect.ValueOf(actual).Len()
	if length == 0 {
		return nil
	}

	TYPE := reflect.TypeOf(actual).String()
	return fmt.Errorf(
		"%w: got len(%s) == %d, want empty %s",
		ErrAssertionFailure,
		TYPE,
		length,
		TYPE,
	)
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}
