package should_test

import (
	"testing"
	"time"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldEqual(t *testing.T) {
	assertFail(t, should.Equal("not enough args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.Equal("too", "many", "args"), should.ErrExpectedCountInvalid)

	assertFail(t, should.Equal(1, 2), should.ErrEqualityCheck)
	assertPass(t, should.Equal(1, 1))
	assertPass(t, should.Equal(1, uint(1)))

	now := time.Now()
	assertPass(t, should.Equal(now.UTC(), now.In(time.Local)))
	assertFail(t, should.Equal(time.Now(), time.Now()), should.ErrEqualityCheck)

	assertFail(t, should.Equal(struct{ A string }{}, struct{ B string }{}), should.ErrEqualityCheck)
	assertPass(t, should.Equal(struct{ A string }{}, struct{ A string }{}))

	assertFail(t, should.Equal([]byte("hi"), []byte("bye")), should.ErrEqualityCheck)
	assertPass(t, should.Equal([]byte("hi"), []byte("hi")))
}

func TestShouldNotEqual(t *testing.T) {
	assertFail(t, should.NOT.Equal("not enough args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.Equal("too", "many", "args"), should.ErrExpectedCountInvalid)
	assertFail(t, should.NOT.Equal(1, 1), should.ErrAssertionFailure)
	assertPass(t, should.NOT.Equal(1, 2))
}
