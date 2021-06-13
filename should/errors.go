package should

import (
	"errors"
	"fmt"
)

var (
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

func failure(format string, args ...interface{}) error {
	full := fmt.Sprintf(format, args...)
	return fmt.Errorf("%w: "+full, ErrAssertionFailure)
}

/*

## TODO

- StartWith        (&NOT) for slices, strings, arrays
- EndWith          (&NOT) for slices, strings, arrays

*/
