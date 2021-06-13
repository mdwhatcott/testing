package should

import (
	"errors"
	"fmt"
	"reflect"
)

// WrapError uses errors.Is to verify that actual is an error value
// that wraps expected[0] (also an error value).
func WrapError(actual interface{}, expected ...interface{}) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	inner, ok := expected[0].(error)
	if !ok {
		return errTypeMismatch(expected[0])
	}

	outer, ok := actual.(error)
	if !ok {
		return errTypeMismatch(actual)
	}

	if errors.Is(outer, inner) {
		return nil
	}

	return fmt.Errorf("%w:\n"+
		"\touter err:  (%s)\n"+
		"\tdoes not\n"+
		"\twrap inner: (%s)",
		ErrAssertionFailure,
		outer,
		inner,
	)
}

func errTypeMismatch(v interface{}) error {
	return fmt.Errorf(
		"%w: got %s, want error",
		ErrTypeMismatch,
		reflect.TypeOf(v),
	)
}
