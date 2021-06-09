package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeEmpty(t *testing.T) {
	assertFail(t, should.BeEmpty([]string(nil), "extra"), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeEmpty(42), should.ErrKindMismatch)

	assertPass(t, should.BeEmpty([]string(nil)))
	assertPass(t, should.BeEmpty(make([]string, 0, 0)))
	assertPass(t, should.BeEmpty(make([]string, 0, 1)))
	assertFail(t, should.BeEmpty([]string{""}), should.ErrAssertionFailure)

	assertPass(t, should.BeEmpty([0]string{})) // The only possible empty array!
	assertFail(t, should.BeEmpty([1]string{}), should.ErrAssertionFailure)

	assertPass(t, should.BeEmpty(chan string(nil)))
	assertPass(t, should.BeEmpty(make(chan string)))
	assertPass(t, should.BeEmpty(make(chan string, 1)))
	assertFail(t, should.BeEmpty(nonEmptyChannel()), should.ErrAssertionFailure)

	assertPass(t, should.BeEmpty(map[string]string(nil)))
	assertPass(t, should.BeEmpty(make(map[string]string)))
	assertPass(t, should.BeEmpty(make(map[string]string, 1)))
	assertFail(t, should.BeEmpty(map[string]string{"": ""}), should.ErrAssertionFailure)

	assertPass(t, should.BeEmpty(""))
	assertPass(t, should.BeEmpty(*new(string)))
	assertFail(t, should.BeEmpty(" "), should.ErrAssertionFailure)
}

func nonEmptyChannel() chan string {
	c := make(chan string, 1)
	c <- ""
	return c
}
