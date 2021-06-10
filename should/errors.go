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

- BeIn (&NOT)
- Contain (&NOT)
- ContainKey (&NOT)

- ContainSubstring (&NOT)
- StartWith (&NOT)
- EndWith (&NOT)

- BeGreaterThan
- BeGreaterThanOrEqualTo
- BeLessThan
- BeLessThanOrEqualTo

- Panic (&NOT)
- PanicWith
- WrapError

- HappenAfter
- HappenBefore
- HappenBetween
- HappenOnOrAfter
- HappenOnOrBefore
- HappenOnOrBetween
- HappenWithin (&NOT)

*/
