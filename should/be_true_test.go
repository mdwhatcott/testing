package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeTrue(t *testing.T) {
	invalid(t, should.BeTrue(1, 2), should.ErrExpectedCountInvalid)
	invalid(t, should.BeTrue(1), should.ErrTypeMismatch)
	fail(t, should.BeTrue(false))
	pass(t, should.BeTrue(true))
}
