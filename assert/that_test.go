package assert_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
)

func TestAssertThat_Passes(t *testing.T) {
	assertNil(t, assert.That(nil).IsNil())
	assertNil(t, assert.That(true).IsTrue())
	assertNil(t, assert.That(false).IsFalse())
	assertNil(t, assert.That(1).Equals(1))
}
func assertNil(t *testing.T, actual error) {
	if actual == nil {
		return
	}
	t.Helper()
	t.Error("Expected <nil> err, got:", actual)
}

func TestAssertThat_Fails(t *testing.T) {
	assertErr(t, assert.That(t).IsNil())
	assertErr(t, assert.That(false).IsTrue())
	assertErr(t, assert.That(true).IsFalse())
	assertErr(t, assert.That(1).Equals(2))
}
func assertErr(t *testing.T, actual error) {
	if actual != nil {
		return
	}
	t.Helper()
	t.Error("Expected non-<nil> err, got:", actual)
}
