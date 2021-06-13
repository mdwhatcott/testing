package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeIn(t *testing.T) {
	assertFail(t, should.BeIn("not enough"), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeIn("too", "many", "args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeIn(false, "string"), should.ErrKindMismatch)
	assertFail(t, should.BeIn("hi", 1), should.ErrKindMismatch)

	// strings:
	assertFail(t, should.BeIn("no", ""), should.ErrAssertionFailure)
	assertPass(t, should.BeIn("rat", "integrate"))
	assertPass(t, should.BeIn('b', "abc"))

	// slices:
	assertFail(t, should.BeIn('d', []byte("abc")), should.ErrAssertionFailure)
	assertPass(t, should.BeIn('b', []byte("abc")))
	assertPass(t, should.BeIn(98, []byte("abc")))

	// arrays:
	assertFail(t, should.BeIn('d', [3]byte{'a', 'b', 'c'}), should.ErrAssertionFailure)
	assertPass(t, should.BeIn('b', [3]byte{'a', 'b', 'c'}))
	assertPass(t, should.BeIn(98, [3]byte{'a', 'b', 'c'}))

	// maps:
	assertFail(t, should.BeIn('b', map[rune]int{'a': 1}), should.ErrAssertionFailure)
	assertPass(t, should.BeIn('a', map[rune]int{'a': 1}))
}

func TestShouldNotBeIn(t *testing.T) {
	assertFail(t, should.NOT.BeIn("not enough"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.BeIn("too", "many", "args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.BeIn(false, "string"), should.ErrKindMismatch)
	assertFail(t, should.NOT.BeIn("hi", 1), should.ErrKindMismatch)

	// strings:
	assertPass(t, should.NOT.BeIn("no", "yes")) // So("no", should.NOT.BeIn, "")
	assertFail(t, should.NOT.BeIn("rat", "integrate"), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeIn('b', "abc"), should.ErrAssertionFailure)

	// slices:
	assertPass(t, should.NOT.BeIn('d', []byte("abc")))
	assertFail(t, should.NOT.BeIn('b', []byte("abc")), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeIn(98, []byte("abc")), should.ErrAssertionFailure)

	// arrays:
	assertPass(t, should.NOT.BeIn('d', [3]byte{'a', 'b', 'c'}))
	assertFail(t, should.NOT.BeIn('b', [3]byte{'a', 'b', 'c'}), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeIn(98, [3]byte{'a', 'b', 'c'}), should.ErrAssertionFailure)

	// maps:
	assertPass(t, should.NOT.BeIn('b', map[rune]int{'a': 1}))
	assertFail(t, should.NOT.BeIn('a', map[rune]int{'a': 1}), should.ErrAssertionFailure)
}
