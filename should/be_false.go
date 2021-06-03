package should

import (
	"fmt"
	"reflect"
)

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
