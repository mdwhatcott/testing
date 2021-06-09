package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeTrue(t *testing.T) {
	assertPass(t, should.BeTrue(true))
	assertFail(t, should.BeTrue(1, 2), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeTrue(1), should.ErrTypeMismatch)
	assertFail(t, should.BeTrue(false), should.ErrAssertionFailure)
}
