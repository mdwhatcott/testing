package should

import (
	"fmt"
	"reflect"
	"time"
)

// HappenAt verifies that the actual value is a time.Time that is
// equal to the expected time.Time.
func HappenAt(ACTUAL interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	if len(EXPECTED) != 1 {
		return fmt.Errorf("%w: please provide a single expected value (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	expected, ok := EXPECTED[0].(time.Time)
	if !ok {
		return fmt.Errorf("%w: please provide an expected time.Time (not %d)", errExpectedTypeMismatch, reflect.TypeOf(EXPECTED))
	}
	actual, ok := ACTUAL.(time.Time)
	if !ok {
		return fmt.Errorf("%w: please provide an actual time.Time (not %d)", errActualTypeMismatch, reflect.TypeOf(ACTUAL))
	}
	if actual.Equal(expected) {
		return nil
	}
	return fmt.Errorf("%w:\n"+
		"Expected: %s\n"+
		"Actual:   %s",
		errEqualityMismatch,
		expected.Format(time.RFC3339Nano),
		actual.Format(time.RFC3339Nano),
	)
}
