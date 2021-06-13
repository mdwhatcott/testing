package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldHaveLength(t *testing.T) {
	invalid(t, should.HaveLength("", 1, "too many"), should.ErrExpectedCountInvalid)
	invalid(t, should.HaveLength(true, 0), should.ErrKindMismatch)
	invalid(t, should.HaveLength(42, 0), should.ErrKindMismatch)
	invalid(t, should.HaveLength("", ""), should.ErrKindMismatch)

	pass(t, should.HaveLength([]string(nil), 0))
	pass(t, should.HaveLength([]string{}, 0))
	pass(t, should.HaveLength([]string{""}, 1))
	fail(t, should.HaveLength([]string{""}, 2))

	pass(t, should.HaveLength([0]string{}, 0)) // The only possible empty array!
	fail(t, should.HaveLength([1]string{}, 2))

	pass(t, should.HaveLength(chan string(nil), 0))
	fail(t, should.HaveLength(nonEmptyChannel(), 2))

	pass(t, should.HaveLength(map[string]string{"": ""}, 1))
	fail(t, should.HaveLength(map[string]string{"": ""}, 2))

	pass(t, should.HaveLength("", 0))
	pass(t, should.HaveLength("123", 3))
	fail(t, should.HaveLength("123", 4))
}
