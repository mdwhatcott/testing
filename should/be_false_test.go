package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeFalse(t *testing.T) {
	invalid(t, should.BeFalse(1, 2), should.ErrExpectedCountInvalid)
	invalid(t, should.BeFalse(1), should.ErrTypeMismatch)

	fail(t, should.BeFalse(true))
	pass(t, should.BeFalse(false))
}
