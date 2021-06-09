package should

import (
	"fmt"
	"reflect"
)

// BeFalse verifies that actual is the boolean false value.
func BeFalse(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}
	boolean, ok := actual.(bool)
	if !ok {
		return fmt.Errorf("%w: actual is %s (bool required)", ErrTypeMismatch, reflect.TypeOf(actual))
	}
	if boolean {
		return fmt.Errorf("%w: expected <false>, got <true> instead", ErrBoolCheck)
	}
	return nil
}
