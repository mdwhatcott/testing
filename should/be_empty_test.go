package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeEmpty(t *testing.T) {
	invalid(t, should.BeEmpty([]string(nil), "extra"), should.ErrExpectedCountInvalid)
	invalid(t, should.BeEmpty(42), should.ErrKindMismatch)

	pass(t, should.BeEmpty([]string(nil)))
	pass(t, should.BeEmpty(make([]string, 0, 0)))
	pass(t, should.BeEmpty(make([]string, 0, 1)))
	fail(t, should.BeEmpty([]string{""}))

	pass(t, should.BeEmpty([0]string{})) // The only possible empty array!
	fail(t, should.BeEmpty([1]string{}))

	pass(t, should.BeEmpty(chan string(nil)))
	pass(t, should.BeEmpty(make(chan string)))
	pass(t, should.BeEmpty(make(chan string, 1)))
	fail(t, should.BeEmpty(nonEmptyChannel()))

	pass(t, should.BeEmpty(map[string]string(nil)))
	pass(t, should.BeEmpty(make(map[string]string)))
	pass(t, should.BeEmpty(make(map[string]string, 1)))
	fail(t, should.BeEmpty(map[string]string{"": ""}))

	pass(t, should.BeEmpty(""))
	pass(t, should.BeEmpty(*new(string)))
	fail(t, should.BeEmpty(" "))
}

func TestShouldNotBeEmpty(t *testing.T) {
	invalid(t, should.NOT.BeEmpty([]string(nil), "extra"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.BeEmpty(42), should.ErrKindMismatch)

	fail(t, should.NOT.BeEmpty([]string(nil)))
	fail(t, should.NOT.BeEmpty(make([]string, 0, 0)))
	fail(t, should.NOT.BeEmpty(make([]string, 0, 1)))
	pass(t, should.NOT.BeEmpty([]string{""}))

	fail(t, should.NOT.BeEmpty([0]string{}))
	pass(t, should.NOT.BeEmpty([1]string{}))

	fail(t, should.NOT.BeEmpty(chan string(nil)))
	fail(t, should.NOT.BeEmpty(make(chan string)))
	fail(t, should.NOT.BeEmpty(make(chan string, 1)))
	pass(t, should.NOT.BeEmpty(nonEmptyChannel()))

	fail(t, should.NOT.BeEmpty(map[string]string(nil)))
	fail(t, should.NOT.BeEmpty(make(map[string]string)))
	fail(t, should.NOT.BeEmpty(make(map[string]string, 1)))
	pass(t, should.NOT.BeEmpty(map[string]string{"": ""}))

	fail(t, should.NOT.BeEmpty(""))
	fail(t, should.NOT.BeEmpty(*new(string)))
	pass(t, should.NOT.BeEmpty(" "))
}

func nonEmptyChannel() chan string {
	c := make(chan string, 1)
	c <- ""
	return c
}
