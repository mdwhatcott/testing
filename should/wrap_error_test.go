package should_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldWrapError(t *testing.T) {
	invalid(t, should.WrapError(0), should.ErrExpectedCountInvalid)
	invalid(t, should.WrapError(0, 1, 2), should.ErrExpectedCountInvalid)

	invalid(t, should.WrapError(inner, 42), should.ErrTypeMismatch)
	invalid(t, should.WrapError(42, inner), should.ErrTypeMismatch)

	pass(t, should.WrapError(outer, inner))
	fail(t, should.WrapError(inner, outer))
}

var (
	inner = errors.New("inner")
	outer = fmt.Errorf("outer(%w)", inner)
)
