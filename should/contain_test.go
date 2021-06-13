package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldContain(t *testing.T) {
	invalid(t, should.Contain("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.Contain("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.Contain("string", false), should.ErrKindMismatch)
	invalid(t, should.Contain(1, "hi"), should.ErrKindMismatch)

	// strings:
	fail(t, should.Contain("", "no"))
	pass(t, should.Contain("integrate", "rat"))
	pass(t, should.Contain("abc", 'b'))

	// slices:
	fail(t, should.Contain([]byte("abc"), 'd'))
	pass(t, should.Contain([]byte("abc"), 'b'))
	pass(t, should.Contain([]byte("abc"), 98))

	// arrays:
	fail(t, should.Contain([3]byte{'a', 'b', 'c'}, 'd'))
	pass(t, should.Contain([3]byte{'a', 'b', 'c'}, 'b'))
	pass(t, should.Contain([3]byte{'a', 'b', 'c'}, 98))

	// maps:
	fail(t, should.Contain(map[rune]int{'a': 1}, 'b'))
	pass(t, should.Contain(map[rune]int{'a': 1}, 'a'))
}

func TestShouldNotContain(t *testing.T) {
	invalid(t, should.NOT.Contain("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.Contain("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.Contain(false, "string"), should.ErrKindMismatch)
	invalid(t, should.NOT.Contain("hi", 1), should.ErrKindMismatch)

	// strings:
	pass(t, should.NOT.Contain("", "no"))
	fail(t, should.NOT.Contain("integrate", "rat"))
	fail(t, should.NOT.Contain("abc", 'b'))

	// slices:
	pass(t, should.NOT.Contain([]byte("abc"), 'd'))
	fail(t, should.NOT.Contain([]byte("abc"), 'b'))
	fail(t, should.NOT.Contain([]byte("abc"), 98))

	// arrays:
	pass(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 'd'))
	fail(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 'b'))
	fail(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 98))

	// maps:
	pass(t, should.NOT.Contain(map[rune]int{'a': 1}, 'b'))
	fail(t, should.NOT.Contain(map[rune]int{'a': 1}, 'a'))
}
