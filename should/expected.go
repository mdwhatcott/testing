package should

import (
	"fmt"
	"reflect"
)

func validateExpected(count int, expected []interface{}) error {
	if len(expected) == count {
		return nil
	}

	return fmt.Errorf(
		"%w: please provide %d expected value%s (not %d)",
		ErrExpectedCountInvalid,
		count,
		pluralize(count),
		len(expected),
	)
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

func validateType(actual, expected interface{}) error {
	ACTUAL := reflect.TypeOf(actual)
	EXPECTED := reflect.TypeOf(expected)
	if ACTUAL == EXPECTED {
		return nil
	}
	return fmt.Errorf("%w: want %s got: %s", ErrTypeMismatch, ACTUAL.String(), EXPECTED.String())
}
