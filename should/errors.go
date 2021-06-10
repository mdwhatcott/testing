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

func negatedFailure(format string, args ...interface{}) error {
	args = append([]interface{}{ErrAssertionFailure}, args...)
	return fmt.Errorf("negated %w: "+format, args...)
}
func failure(format string, args ...interface{}) error {
	args = append([]interface{}{ErrAssertionFailure}, args...)
	return fmt.Errorf("%w: "+format, args...)
}

/*

## TODO

- Panic            (&NOT)

- BeIn             (&NOT) for slices, strings, arrays, map keys
- Contain          (&NOT) for slices, strings, arrays, map keys
- StartWith        (&NOT) for slices, strings, arrays
- EndWith          (&NOT) for slices, strings, arrays

- BeGreaterThan           for numerics, time.Time
- BeGreaterThanOrEqualTo  for numerics, time.Time
- BeLessThan              for numerics, time.Time
- BeLessThanOrEqualTo     for numerics, time.Time

*/
