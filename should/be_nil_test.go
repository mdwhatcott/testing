package should_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeNil(t *testing.T) {
	assertPass(t, should.BeNil(nil))
	assertPass(t, should.BeNil([]string(nil)))
	assertPass(t, should.BeNil((*string)(nil)))
	assertFail(t, should.BeNil(1, 2), should.ErrExpectedCountInvalid)
	assertFail(t, should.BeNil(notNil), should.ErrAssertionFailure)
}

func TestShouldNotBeNil(t *testing.T) {
	assertPass(t, should.NOT.BeNil(notNil))
	assertFail(t, should.NOT.BeNil(nil), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeNil([]string(nil)), should.ErrAssertionFailure)
	assertFail(t, should.NOT.BeNil(1, 2), should.ErrExpectedCountInvalid)
}

var notNil = errors.New("not nil")
