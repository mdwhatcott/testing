package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldPanic(t *testing.T) {
	invalid(t, should.Panic("to", "many"), should.ErrExpectedCountInvalid)
	invalid(t, should.Panic("wrong type"), should.ErrTypeMismatch)
	fail(t, should.Panic(func() {}))
	pass(t, should.Panic(func() { panic("yay") }))
}

func TestShouldNotPanic(t *testing.T) {
	invalid(t, should.NOT.Panic("to", "many"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.Panic("wrong type"), should.ErrTypeMismatch)
	fail(t, should.NOT.Panic(func() { panic("boo") }))
	pass(t, should.NOT.Panic(func() {}))
}
