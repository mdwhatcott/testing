package should_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldWrapError(t *testing.T) {
	assertFail(t, should.WrapError(0), should.ErrExpectedCountInvalid)
	assertFail(t, should.WrapError(0, 1, 2), should.ErrExpectedCountInvalid)

	assertFail(t, should.WrapError(inner, 42), should.ErrTypeMismatch)
	assertFail(t, should.WrapError(42, inner), should.ErrTypeMismatch)

	assertPass(t, should.WrapError(outer, inner))
	assertFail(t, should.WrapError(inner, outer), should.ErrAssertionFailure)
}

var (
	inner = errors.New("inner")
	outer = fmt.Errorf("outer(%w)", inner)
)
