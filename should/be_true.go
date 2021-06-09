package should

import (
	"fmt"
)

// BeTrue verifies that actual is the boolean true value.
func BeTrue(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if !boolean {
		return fmt.Errorf("%w: got <false>, want <true>", ErrBoolCheck)
	}
	return nil
}
