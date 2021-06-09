package should

import (
	"errors"
	"fmt"
)

var (
	errExpectedCountInvalid = fmt.Errorf("expected count invalid")
	errTypeMismatch         = fmt.Errorf("type mismatch")

	errAssertionFailed = errors.New("assertion failed")
	errNilCheck        = fmt.Errorf("%w: 'nil check'", errAssertionFailed)
	errBoolCheck       = fmt.Errorf("%w: 'bool check'", errAssertionFailed)
	errEqualityCheck   = fmt.Errorf("%w: 'equality check'", errAssertionFailed)
)

/*

## TODO

- BeEmpty
- HaveLength
- Contain
- ContainKey
- ContainSubstring
- StartWith
- EndWith
- BeIn
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

## Negations: need to figure out an elegant way (use result of corresponding positive assertion)

- NotBeEmpty
- NotEqual
- NotHaveLength
- NotContain
- NotContainKey
- NotContainSubstring
- NotStartWith
- NotEndWith
- NotBeIn
- NotPanic
- NotHappenWithin

*/
