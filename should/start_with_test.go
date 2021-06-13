package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldStartWith(t *testing.T) {
	invalid(t, should.StartWith("not enough"), should.ErrExpectedCountInvalid)
	invalid(t, should.StartWith("too", "many", "args"), should.ErrExpectedCountInvalid)
	invalid(t, should.StartWith("string", false), should.ErrKindMismatch)
	invalid(t, should.StartWith(1, "hi"), should.ErrKindMismatch)

	// strings:
	fail(t, should.StartWith("", "no"))
	pass(t, should.StartWith("abc", 'a'))
	pass(t, should.StartWith("integrate", "in"))

	// slices:
	fail(t, should.StartWith([]byte{}, 'b'))
	fail(t, should.StartWith([]byte(nil), 'b'))
	fail(t, should.StartWith([]byte("abc"), 'b'))
	pass(t, should.StartWith([]byte("abc"), 'a'))
	pass(t, should.StartWith([]byte("abc"), 97))

	// arrays:
	fail(t, should.StartWith([3]byte{'a', 'b', 'c'}, 'b'))
	pass(t, should.StartWith([3]byte{'a', 'b', 'c'}, 'a'))
	pass(t, should.StartWith([3]byte{'a', 'b', 'c'}, 97))
}
