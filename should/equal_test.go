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
	//assertPass(t, Equal(1, uint(1))) // TODO: numerics of differing type (but semantically equal value)

	//now := time.Now()
	//assertPass(t, Equal(now.UTC(), now)) // TODO: time.Time instants (in different time zones)
	assertFail(t, Equal(time.Now(), time.Now()), errEqualityCheck)

	assertFail(t, Equal(struct{A string}{}, struct{B string}{}), errEqualityCheck)
	assertPass(t, Equal(struct{A string}{}, struct{A string}{}))

	assertFail(t, Equal([]byte("hi"), []byte("bye")), errEqualityCheck)
	assertPass(t, Equal([]byte("hi"), []byte("hi")))
}