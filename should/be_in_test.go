package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeIn(t *testing.T) {
	invalid(t, should.BeIn("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.BeIn("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.BeIn(false, "string"), should.ErrKindMismatch)
	invalid(t, should.BeIn("hi", 1), should.ErrKindMismatch)

	// strings:
	fail(t, should.BeIn("no", ""))
	pass(t, should.BeIn("rat", "integrate"))
	pass(t, should.BeIn('b', "abc"))

	// slices:
	fail(t, should.BeIn('d', []byte("abc")))
	pass(t, should.BeIn('b', []byte("abc")))
	pass(t, should.BeIn(98, []byte("abc")))

	// arrays:
	fail(t, should.BeIn('d', [3]byte{'a', 'b', 'c'}))
	pass(t, should.BeIn('b', [3]byte{'a', 'b', 'c'}))
	pass(t, should.BeIn(98, [3]byte{'a', 'b', 'c'}))

	// maps:
	fail(t, should.BeIn('b', map[rune]int{'a': 1}))
	pass(t, should.BeIn('a', map[rune]int{'a': 1}))
}

func TestShouldNotBeIn(t *testing.T) {
	invalid(t, should.NOT.BeIn("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.BeIn("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.BeIn(false, "string"), should.ErrKindMismatch)
	invalid(t, should.NOT.BeIn("hi", 1), should.ErrKindMismatch)

	// strings:
	pass(t, should.NOT.BeIn("no", "yes")) // So("no", should.NOT.BeIn, "")
	fail(t, should.NOT.BeIn("rat", "integrate"))
	fail(t, should.NOT.BeIn('b', "abc"))

	// slices:
	pass(t, should.NOT.BeIn('d', []byte("abc")))
	fail(t, should.NOT.BeIn('b', []byte("abc")))
	fail(t, should.NOT.BeIn(98, []byte("abc")))

	// arrays:
	pass(t, should.NOT.BeIn('d', [3]byte{'a', 'b', 'c'}))
	fail(t, should.NOT.BeIn('b', [3]byte{'a', 'b', 'c'}))
	fail(t, should.NOT.BeIn(98, [3]byte{'a', 'b', 'c'}))

	// maps:
	pass(t, should.NOT.BeIn('b', map[rune]int{'a': 1}))
	fail(t, should.NOT.BeIn('a', map[rune]int{'a': 1}))
}
