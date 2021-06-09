package should

import (
	"errors"
	"fmt"
)

var (
	errExpectedCountInvalid = fmt.Errorf("expected count invalid")
	errTypeMismatch         = fmt.Errorf("type mismatch")

	errAssertionFailure = errors.New("assertion failure")
	errNilCheck         = fmt.Errorf("%w: 'nil check'", errAssertionFailure)
	errBoolCheck        = fmt.Errorf("%w: 'bool check'", errAssertionFailure)
	errEqualityCheck    = fmt.Errorf("%w: 'equality check'", errAssertionFailure)
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
