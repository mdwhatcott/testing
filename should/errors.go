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

- ContainSubstring (&NOT)

- BeIn             (&NOT)
- Contain          (&NOT)
- ContainKey       (&NOT)

- StartWith        (&NOT)
- EndWith          (&NOT)

- Panic            (&NOT)

- BeGreaterThan          (w/ support for time.Time)
- BeGreaterThanOrEqualTo (w/ support for time.Time)
- BeLessThan             (w/ support for time.Time)
- BeLessThanOrEqualTo    (w/ support for time.Time)

*/
