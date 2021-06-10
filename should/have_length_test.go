package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldHaveLength(t *testing.T) {
	assertFail(t, should.HaveLength("", 1, "too many"), should.ErrExpectedCountInvalid)
	assertFail(t, should.HaveLength(true, 0), should.ErrKindMismatch)
	assertFail(t, should.HaveLength(42, 0), should.ErrKindMismatch)

	assertPass(t, should.HaveLength([]string(nil), 0))
	assertPass(t, should.HaveLength([]string{}, 0))
	assertPass(t, should.HaveLength([]string{""}, 1))
	assertFail(t, should.HaveLength([]string{""}, 2), should.ErrAssertionFailure)

	assertPass(t, should.HaveLength([0]string{}, 0)) // The only possible empty array!
	assertFail(t, should.HaveLength([1]string{}, 2), should.ErrAssertionFailure)

	assertPass(t, should.HaveLength(chan string(nil), 0))
	assertFail(t, should.HaveLength(nonEmptyChannel(), 2), should.ErrAssertionFailure)

	assertPass(t, should.HaveLength(map[string]string{"": ""}, 1))
	assertFail(t, should.HaveLength(map[string]string{"": ""}, 2), should.ErrAssertionFailure)

	assertPass(t, should.HaveLength("", 0))
	assertPass(t, should.HaveLength("123", 3))
	assertFail(t, should.HaveLength("123", 4), should.ErrAssertionFailure)
}
