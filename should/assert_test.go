package should_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func invalid(t *testing.T, actual, expected error) {
	t.Helper()
	if !errors.Is(actual, expected) {
		t.Errorf("[FAIL]\n"+
			"expected: %v\n"+
			"actual:   %v",
			expected,
			actual,
		)
	} else if testing.Verbose() {
		t.Log("\n", actual, "\n", "(above error report printed for visual inspection)")
	}
}
func fail(t *testing.T, err error) {
	t.Helper()
	if !errors.Is(err, should.ErrAssertionFailure) {
		t.Error("[FAIL] expected assertion failure, got:", err)
	} else if testing.Verbose() {
		t.Log("\n", err, "\n", "(above error report printed for visual inspection)")
	}
}
func pass(t *testing.T, actual error) {
	if actual != nil {
		t.Helper()
		t.Error("[FAIL] expected nil err, got:", actual)
	}
}
