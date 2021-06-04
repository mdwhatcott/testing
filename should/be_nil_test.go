package should

import (
	"errors"
	"testing"
)

func TestShouldBeNil(t *testing.T) {
	assertPass(t, BeNil(nil))
	assertPass(t, BeNil([]string(nil)))
	assertPass(t, BeNil((*string)(nil)))
	assertFail(t, BeNil(1, 2), errExpectedCountInvalid)
	assertFail(t, BeNil(errors.New("not nil")), errNilCheck)
}
