package should

import (
	"fmt"
	"reflect"
)

func Equal(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
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
