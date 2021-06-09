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

	return fmt.Errorf(
		"%w: want empty %s but len() was %d",
		ErrAssertionFailure,
		reflect.TypeOf(actual).String(),
		length,
	)
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}
