package should

import (
	"fmt"
	"reflect"
)

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases.
func Equal(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	if len(EXPECTED) != 1 {
		return fmt.Errorf("%w: please provide a single expected value (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	expected := EXPECTED[0]
	if reflect.DeepEqual(actual, expected) {
		return nil
	}
	return fmt.Errorf("%w:\n"+
		"Expected: %#v\n"+
		"Actual:   %#v",
		errEqualityMismatch,
		expected,
		actual,
	)
}
