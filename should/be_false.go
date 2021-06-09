package should

import "fmt"

// BeFalse verifies that actual is the boolean false value.
func BeFalse(actual interface{}, expected ...interface{}) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if boolean {
		return fmt.Errorf("%w: want <false>, got <true>", ErrBoolCheck)
	}

	return nil
}
