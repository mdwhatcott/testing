package should

import (
	"errors"
)

var (
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

/*

## TODO

- HaveLength
- BeIn
- Contain
- ContainKey

- ContainSubstring
- StartWith
- EndWith

- BeGreaterThan
- BeGreaterThanOrEqualTo
- BeLessThan
- BeLessThanOrEqualTo

- Panic
- PanicWith
- WrapError

- HappenAfter
- HappenBefore
- HappenBetween
- HappenOnOrAfter
- HappenOnOrBefore
- HappenOnOrBetween
- HappenWithin

## Negations:

- NotBeEmpty
- NotHaveLength
- NotBeIn
- NotContain
- NotContainKey

- NotContainSubstring
- NotStartWith
- NotEndWith

- NotPanic

- NotHappenWithin

*/
