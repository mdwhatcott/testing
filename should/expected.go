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
		"%w: got %d value%s, want %d",
		ErrExpectedCountInvalid,
		len(expected),
		pluralize(len(expected)),
		count,
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
	return fmt.Errorf(
		"%w: got %s, want %s",
		ErrTypeMismatch,
		ACTUAL.String(),
		EXPECTED.String(),
	)
}

func validateKind(actual interface{}, kinds ...reflect.Kind) error {
	ACTUAL := reflect.ValueOf(actual)
	kind := ACTUAL.Kind()
	for _, k := range kinds {
		if k == kind {
			return nil
		}
	}
	return fmt.Errorf(
		"%w: got %s, want one of %v",
		ErrKindMismatch,
		kind.String(),
		kinds,
	)
}
