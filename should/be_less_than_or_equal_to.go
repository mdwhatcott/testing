package should

import "errors"

func BeLessThanOrEqualTo(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		return nil
	}
	err = BeLessThan(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return failure("%v was not less than or equal to %v", actual, expected)
	}

	if err != nil {
		return err
	}
	return nil
}
