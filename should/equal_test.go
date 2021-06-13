package should_test

import (
	"testing"
	"time"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldEqual(t *testing.T) {
	invalid(t, should.Equal("not enough args"), should.ErrExpectedCountInvalid)
	invalid(t, should.Equal("too", "many", "args"), should.ErrExpectedCountInvalid)

	fail(t, should.Equal(1, 2))
	pass(t, should.Equal(1, 1))
	pass(t, should.Equal(1, uint(1)))

	now := time.Now()
	pass(t, should.Equal(now.UTC(), now.In(time.Local)))
	fail(t, should.Equal(time.Now(), time.Now()))

	fail(t, should.Equal(struct{ A string }{}, struct{ B string }{}))
	pass(t, should.Equal(struct{ A string }{}, struct{ A string }{}))

	fail(t, should.Equal([]byte("hi"), []byte("bye")))
	pass(t, should.Equal([]byte("hi"), []byte("hi")))
}

func TestShouldNotEqual(t *testing.T) {
	invalid(t, should.NOT.Equal("not enough args"), should.ErrExpectedCountInvalid)
	invalid(t, should.NOT.Equal("too", "many", "args"), should.ErrExpectedCountInvalid)
	fail(t, should.NOT.Equal(1, 1))
	pass(t, should.NOT.Equal(1, 2))
}
