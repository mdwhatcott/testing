package should_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeNil(t *testing.T) {
	invalid(t, should.BeNil(1, 2), should.ErrExpectedCountInvalid)
	fail(t, should.BeNil(notNil))
	pass(t, should.BeNil(nil))
	pass(t, should.BeNil([]string(nil)))
	pass(t, should.BeNil((*string)(nil)))
}

func TestShouldNotBeNil(t *testing.T) {
	invalid(t, should.NOT.BeNil(1, 2), should.ErrExpectedCountInvalid)
	fail(t, should.NOT.BeNil(nil))
	fail(t, should.NOT.BeNil([]string(nil)))
	pass(t, should.NOT.BeNil(notNil))
}

var notNil = errors.New("not nil")
