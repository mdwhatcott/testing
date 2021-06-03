package should

import "errors"

var (
	errNilCheckFailed       = errors.New("nil check failed")
	errExpectedCountInvalid = errors.New("expected count invalid")
	errExpectedTypeMismatch = errors.New("expected type mismatch")
	errActualTypeMismatch   = errors.New("actual type mismatch")
	errEqualityMismatch     = errors.New("equality mismatch")
)

/*

## TODO

- Equal for numerics of differing type: So(8, should.Equal, uint(8))
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

- NotBeNil
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
