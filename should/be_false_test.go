package should

import "testing"

func TestShouldBeFalse(t *testing.T) {
	assertPass(t, BeFalse(false))
	assertFail(t, BeFalse(1, 2), errExpectedCountInvalid)
	assertFail(t, BeFalse(1), errTypeMismatch)
	assertFail(t, BeFalse(true), errBoolCheck)
}
