package should

import "fmt"

func validateExpected(count int, expected ...interface{}) error {
	if len(expected) == count {
		return nil
	}

	return fmt.Errorf(
		"%w: please provide %d expected value%s (not %d)",
		errExpectedCountInvalid,
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
