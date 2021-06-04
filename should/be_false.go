package should

import (
	"errors"
	"fmt"
)

// BeFalse verifies that actual is the boolean false value.
func BeFalse(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	err := BeTrue(actual, EXPECTED...)
	if errors.Is(err, errEqualityCheck) {
		return nil
	}
	if err != nil {
		return err
	}
	return fmt.Errorf("%w: expected <false>, got <true> instead", errEqualityCheck)
}
