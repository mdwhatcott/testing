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
