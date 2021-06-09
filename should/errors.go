package should

import (
	"errors"
	"fmt"
)

// TODO: is this the right package for these error declarations?
var (
	ErrExpectedCountInvalid = fmt.Errorf("expected count invalid")
	ErrTypeMismatch         = fmt.Errorf("type mismatch")

	ErrAssertionFailure     = errors.New("assertion failure")
	ErrNilCheck             = fmt.Errorf("%w: 'nil check'", ErrAssertionFailure)      // Deprecated
	ErrBoolCheck            = fmt.Errorf("%w: 'bool check'", ErrAssertionFailure)     // Deprecated
	ErrEqualityCheck        = fmt.Errorf("%w: 'equality check'", ErrAssertionFailure) // Deprecated
)

/*

## TODO

- BeEmpty
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
