package should

import (
	"errors"
	"fmt"
)

func BeGreaterThanOrEqualTo(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if err == nil {
		fmt.Println("HIHIHIHI")
		return nil
	}
	err = BeGreaterThan(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return failure("%v was not greater than or equal to %v", actual, expected)
	}

	if err != nil {
		return err
	}
	return nil
}
