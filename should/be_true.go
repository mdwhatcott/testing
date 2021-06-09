package should

import (
	"fmt"
	"reflect"
)

// BeTrue verifies that actual is the boolean true value.
func BeTrue(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}
	boolean, ok := actual.(bool)
	if !ok {
		return fmt.Errorf("%w: actual is %s (bool required)", ErrTypeMismatch, reflect.TypeOf(actual))
	}
	if !boolean {
		return fmt.Errorf("%w: expected <true>, got <false> instead", ErrBoolCheck)
	}
	return nil
}
