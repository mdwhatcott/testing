package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/assert"
)

func TestPass(t *testing.T) {
	T := new(FakeT)

	assert.With(T).That(nil).IsNil()
	assert.With(T).That((*testing.T)(nil)).IsNil()
	assert.With(T).That(true).IsTrue()
	assert.With(T).That(false).IsFalse()

	if len(T.failures) > 0 {
		t.Error("Unexpected failures:", T.failures)
	}
}
func TestFail(t *testing.T) {
	T := new(FakeT)

	assert.With(T).That(true).IsFalse()
	assert.With(T).That(false).IsTrue()
	assert.With(T).That(errors.New("HI")).IsNil()

	if len(T.failures) != 3 {
		t.Error("Expected a failure!")
	}
}

type FakeT struct{ failures []string }

func (this *FakeT) Helper() {}
func (this *FakeT) Error(args ...interface{}) {
	this.failures = append(this.failures, fmt.Sprint(args...))
}
