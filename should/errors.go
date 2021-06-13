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
	format = "%w: " + full
	return fmt.Errorf(format, ErrAssertionFailure)
}

/*

## TODO

- BeIn             (&NOT) for slices, strings, arrays, map keys
- StartWith        (&NOT) for slices, strings, arrays
- EndWith          (&NOT) for slices, strings, arrays

- BeGreaterThan           for numerics, time.Time
- BeGreaterThanOrEqualTo  for numerics, time.Time
- BeLessThan              for numerics, time.Time
- BeLessThanOrEqualTo     for numerics, time.Time

*/
