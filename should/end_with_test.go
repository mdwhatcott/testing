package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldEndWith(t *testing.T) {
	invalid(t, should.EndWith("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.EndWith("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.EndWith("string", false), should.ErrKindMismatch)
	invalid(t, should.EndWith(1, "hi"), should.ErrKindMismatch)

	// strings:
	fail(t, should.EndWith("", "no"))
	pass(t, should.EndWith("abc", 'c'))
	pass(t, should.EndWith("integrate", "ate"))

	// slices:
	fail(t, should.EndWith([]byte{}, 'b'))
	fail(t, should.EndWith([]byte(nil), 'b'))
	fail(t, should.EndWith([]byte("abc"), 'b'))
	pass(t, should.EndWith([]byte("abc"), 'c'))
	pass(t, should.EndWith([]byte("abc"), 99))

	// arrays:
	fail(t, should.EndWith([3]byte{'a', 'b', 'c'}, 'b'))
	pass(t, should.EndWith([3]byte{'a', 'b', 'c'}, 'c'))
	pass(t, should.EndWith([3]byte{'a', 'b', 'c'}, 99))
}
