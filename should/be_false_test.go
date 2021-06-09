package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeFalse(t *testing.T) {
	assertPass(t, should.BeFalse(false))
	assertFail(t, should.BeFalse(1, 2), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeFalse(1), should.ErrTypeMismatch)
	assertFail(t, should.BeFalse(true), should.ErrAssertionFailure)
}
