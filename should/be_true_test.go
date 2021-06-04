package should

import "testing"

func TestShouldBeTrue(t *testing.T) {
	assertPass(t, BeTrue(true))
	assertFail(t, BeTrue(1, 2), errExpectedCountInvalid)
	assertFail(t, BeTrue(1), errTypeMismatch)
	assertFail(t, BeTrue(false), errBoolCheck)
}
