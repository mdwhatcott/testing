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
	assertFail(t, BeNil(notNil), errNilCheck)
}

func TestShouldNotBeNil(t *testing.T) {
	assertPass(t, NOT.BeNil(notNil))
	assertFail(t, NOT.BeNil(nil), errNilCheck)
	assertFail(t, NOT.BeNil([]string(nil)), errNilCheck)
	assertFail(t, NOT.BeNil(1, 2), errExpectedCountInvalid)

}

var notNil = errors.New("not nil")
