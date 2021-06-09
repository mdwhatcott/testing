package should

import (
	"testing"
	"time"
)

func TestShouldEqual(t *testing.T) {
	assertFail(t, Equal("not enough args"), errExpectedCountInvalid)
	assertFail(t, Equal("too", "many", "args"), errExpectedCountInvalid)

	assertFail(t, Equal(1, 2), errEqualityCheck)
	assertPass(t, Equal(1, 1))
	assertPass(t, Equal(1, uint(1)))

	now := time.Now()
	assertPass(t, Equal(now.UTC(), now.In(time.Local)))
	assertFail(t, Equal(time.Now(), time.Now()), errEqualityCheck)

	assertFail(t, Equal(struct{ A string }{}, struct{ B string }{}), errEqualityCheck)
	assertPass(t, Equal(struct{ A string }{}, struct{ A string }{}))

	assertFail(t, Equal([]byte("hi"), []byte("bye")), errEqualityCheck)
	assertPass(t, Equal([]byte("hi"), []byte("hi")))
}

func TestShouldNotEqual(t *testing.T) {
	assertFail(t, NOT.Equal("not enough args"), errExpectedCountInvalid)
	assertFail(t, NOT.Equal("too", "many", "args"), errExpectedCountInvalid)
	assertFail(t, NOT.Equal(1, 1), errAssertionFailure)
	assertPass(t, NOT.Equal(1, 2))
}
