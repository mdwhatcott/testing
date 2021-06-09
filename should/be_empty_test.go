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

func TestShouldNotBeEmpty(t *testing.T) {
	assertFail(t, should.NOT.BeEmpty([]string(nil), "extra"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.BeEmpty(42), should.ErrKindMismatch)

	assertFail(t, should.NOT.BeEmpty([]string(nil)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make([]string, 0, 0)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make([]string, 0, 1)), should.ErrAssertionFailure)
	assertPass(t, should.NOT.BeEmpty([]string{""}))

	assertFail(t, should.NOT.BeEmpty([0]string{}), should.ErrAssertionFailure)
	assertPass(t, should.NOT.BeEmpty([1]string{}))

	assertFail(t, should.NOT.BeEmpty(chan string(nil)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make(chan string)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make(chan string, 1)), should.ErrAssertionFailure)
	assertPass(t, should.NOT.BeEmpty(nonEmptyChannel()))

	assertFail(t, should.NOT.BeEmpty(map[string]string(nil)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make(map[string]string)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(make(map[string]string, 1)), should.ErrAssertionFailure)
	assertPass(t, should.NOT.BeEmpty(map[string]string{"": ""}))

	assertFail(t, should.NOT.BeEmpty(""), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeEmpty(*new(string)), should.ErrAssertionFailure)
	assertPass(t, should.NOT.BeEmpty(" "))
}

func nonEmptyChannel() chan string {
	c := make(chan string, 1)
	c <- ""
	return c
}
