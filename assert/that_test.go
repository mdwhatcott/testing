package assert_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/assert"
)

func TestFlexibleNumericEquality(t *testing.T) {
	t.Skip("TODO?")
	assert.With(t).That(int32(1)).Equals(int64(1))
}

func TestPass(t *testing.T) {
	tt := new(FakeT)

	assert.With(tt).That(nil).IsNil()
	assert.With(tt).That((*testing.T)(nil)).IsNil()
	assert.With(tt).That(true).IsTrue()
	assert.With(tt).That(false).IsFalse()

	if len(tt.failures) > 0 {
		t.Error("Unexpected failures:", tt.failures)
	}
}
func TestFail(t *testing.T) {
	tt := new(FakeT)

	assert.With(tt).That(true).IsFalse()
	assert.With(tt).That(false).IsTrue()
	assert.With(tt).That(errors.New("HI")).IsNil()

	if len(tt.failures) != 3 {
		t.Error("Expected a failure!")
	}
}

type FakeT struct{ failures []string }

func (this *FakeT) Helper() {}
func (this *FakeT) Errorf(format string, args ...interface{}) {
	this.failures = append(this.failures, fmt.Sprintf(format, args...))
}
