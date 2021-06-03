package should

import (
	"errors"
	"fmt"
	"reflect"
	"time"
)

func BeNil(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
	if len(EXPECTED) > 0 {
		return fmt.Errorf("%w: please provide 0 expected values (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	if actual == nil || interfaceHasNilValue(actual) {
		return nil
	}
	return fmt.Errorf("%w: expected <nil>, got: %#v", errNilCheckFailed, actual)
}
func interfaceHasNilValue(actual interface{}) bool {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	nillable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map

	// Careful: reflect.Value.IsNil() will panic unless it's an interface, chan, map, func, slice, or ptr
	// Reference: http://golang.org/pkg/reflect/#Value.IsNil
	return nillable && value.IsNil()
}

func BeFalse(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
	if len(EXPECTED) > 0 {
		return fmt.Errorf("%w: please provide 0 expected values (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	boolean, ok := actual.(bool)
	if !ok {
		return fmt.Errorf("%w: please provide an actual bool (not %d)", errActualTypeMismatch, reflect.TypeOf(actual))
	}
	if boolean {
		return fmt.Errorf("%w: expected <false>, got <true> instead", errEqualityMismatch)
	}
	return nil
}

func BeTrue(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
	if len(EXPECTED) > 0 {
		return fmt.Errorf("%w: please provide 0 expected values (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	boolean, ok := actual.(bool)
	if !ok {
		return fmt.Errorf("%w: please provide an actual bool (not %d)", errActualTypeMismatch, reflect.TypeOf(actual))
	}
	if !boolean {
		return fmt.Errorf("%w: expected <true>, got <false> instead", errEqualityMismatch)
	}
	return nil
}

func Equal(actual interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
	if len(EXPECTED) != 1 {
		return fmt.Errorf("%w: please provide a single expected value (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	expected := EXPECTED[0]
	if reflect.DeepEqual(actual, expected) {
		return nil
	}
	return fmt.Errorf("%w:\n"+
		"Expected: %#v\n"+
		"Actual:   %#v",
		errEqualityMismatch,
		expected,
		actual,
	)
}

func HappenAt(ACTUAL interface{}, EXPECTED ...interface{}) error {
	// TODO: test
	// TODO: docs
	if len(EXPECTED) != 1 {
		return fmt.Errorf("%w: please provide a single expected value (not %d)", errExpectedCountInvalid, len(EXPECTED))
	}
	expected, ok := EXPECTED[0].(time.Time)
	if !ok {
		return fmt.Errorf("%w: please provide an expected time.Time (not %d)", errExpectedTypeMismatch, reflect.TypeOf(EXPECTED))
	}
	actual, ok := ACTUAL.(time.Time)
	if !ok {
		return fmt.Errorf("%w: please provide an actual time.Time (not %d)", errActualTypeMismatch, reflect.TypeOf(ACTUAL))
	}
	if actual.Equal(expected) {
		return nil
	}
	return fmt.Errorf("%w:\n"+
		"Expected: %s\n"+
		"Actual:   %s",
		errEqualityMismatch,
		expected.Format(time.RFC3339Nano),
		actual.Format(time.RFC3339Nano),
	)
}

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