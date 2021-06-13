package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldContain(t *testing.T) {
	assertFail(t, should.Contain("not enough"), should.ErrExpectedCountInvalid)
	assertFail(t, should.Contain("too", "many", "args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.Contain("string", false), should.ErrKindMismatch)
	assertFail(t, should.Contain(1, "hi"), should.ErrKindMismatch)

	// strings:
	assertFail(t, should.Contain("", "no"), should.ErrAssertionFailure)
	assertPass(t, should.Contain("integrate", "rat"))
	assertPass(t, should.Contain("abc", 'b'))

	// slices:
	assertFail(t, should.Contain([]byte("abc"), 'd'), should.ErrAssertionFailure)
	assertPass(t, should.Contain([]byte("abc"), 'b'))
	assertPass(t, should.Contain([]byte("abc"), 98))

	// arrays:
	assertFail(t, should.Contain([3]byte{'a', 'b', 'c'}, 'd'), should.ErrAssertionFailure)
	assertPass(t, should.Contain([3]byte{'a', 'b', 'c'}, 'b'))
	assertPass(t, should.Contain([3]byte{'a', 'b', 'c'}, 98))

	// maps:
	assertFail(t, should.Contain(map[rune]int{'a': 1}, 'b'), should.ErrAssertionFailure)
	assertPass(t, should.Contain(map[rune]int{'a': 1}, 'a'))
}

func TestShouldNotContain(t *testing.T) {
	assertFail(t, should.NOT.Contain("not enough"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.Contain("too", "many", "args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.Contain(false, "string"), should.ErrKindMismatch)
	assertFail(t, should.NOT.Contain("hi", 1), should.ErrKindMismatch)

	// strings:
	assertPass(t, should.NOT.Contain("", "no"))
	assertFail(t, should.NOT.Contain("integrate", "rat"), should.ErrAssertionFailure)
	assertFail(t, should.NOT.Contain("abc", 'b'), should.ErrAssertionFailure)

	// slices:
	assertPass(t, should.NOT.Contain([]byte("abc"), 'd'))
	assertFail(t, should.NOT.Contain([]byte("abc"), 'b'), should.ErrAssertionFailure)
	assertFail(t, should.NOT.Contain([]byte("abc"), 98), should.ErrAssertionFailure)

	// arrays:
	assertPass(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 'd'))
	assertFail(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 'b'), should.ErrAssertionFailure)
	assertFail(t, should.NOT.Contain([3]byte{'a', 'b', 'c'}, 98), should.ErrAssertionFailure)

	// maps:
	assertPass(t, should.NOT.Contain(map[rune]int{'a': 1}, 'b'))
	assertFail(t, should.NOT.Contain(map[rune]int{'a': 1}, 'a'), should.ErrAssertionFailure)
}
