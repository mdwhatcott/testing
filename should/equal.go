package should

import (
	"errors"

	"github.com/mdwhatcott/testing/compare"
)

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases.
func Equal(actual interface{}, EXPECTED ...interface{}) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}

	err = compare.Equal(actual, EXPECTED[0])
	if err == nil {
		return nil
	}

	return failure(err.Error())
}

// Equal negated!
func (negated) Equal(actual interface{}, expected ...interface{}) error {
	err := Equal(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		expected[0],
		actual,
	)
}
