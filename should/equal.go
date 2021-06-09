package should

import (
	"errors"
	"fmt"

	"github.com/mdwhatcott/testing/compare"
)

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases.
func Equal(actual interface{}, EXPECTED ...interface{}) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}
	result := compare.New().Compare(EXPECTED[0], actual)
	if result.OK() {
		return nil
	}
	return fmt.Errorf("%w: %s", errEqualityCheck, result.Report())
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
