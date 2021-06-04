package should

import (
	"errors"
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
	// TODO: []byte
	// TODO: time.Time
	// TODO: struct pointers?
	expected := EXPECTED[0]
	if reflect.DeepEqual(actual, expected) {
		return nil
	}
	return fmt.Errorf("\n" +
		"%w:\n"+
		"  expected: %#v\n"+
		"  actual:   %#v",
		errEqualityCheck,
		expected,
		actual,
	)
}

// Equal negated!
func (not) Equal(actual interface{}, expected ...interface{}) error {
	// TODO: test
	err := Equal(actual, expected...)
	if errors.Is(err, errAssertionFailed) {
		return nil
	}
	if err != nil {
		return err
	}
	return fmt.Errorf("\n"+
		"negated %w\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		errEqualityCheck,
		expected[0],
		actual,
	)
}
