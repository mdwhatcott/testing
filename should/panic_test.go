package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldPanic(t *testing.T) {
	assertFail(t, should.Panic("to", "many"), should.ErrExpectedCountInvalid)
	assertFail(t, should.Panic("wrong type"), should.ErrTypeMismatch)
	assertFail(t, should.Panic(func() {}), should.ErrAssertionFailure)
	assertPass(t, should.Panic(func() { panic("yay") }))
}

func TestShouldNotPanic(t *testing.T) {
	assertFail(t, should.NOT.Panic("to", "many"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.Panic("wrong type"), should.ErrTypeMismatch)
	assertFail(t, should.NOT.Panic(func() { panic("boo") }), should.ErrAssertionFailure)
	assertPass(t, should.NOT.Panic(func() {}))
}
