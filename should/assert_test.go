package should

import (
	"errors"
	"testing"
)

func assertPass(t *testing.T, actual error) {
	t.Helper()
	if actual != nil {
		t.Error("expected nil err, got:", actual)
	}
}

func assertFail(t *testing.T, actual, expected error) {
	t.Helper()

	if !errors.Is(actual, expected) {
		t.Errorf("\n"+
			"expected: %v\n"+
			"actual:   %v",
			expected,
			actual,
		)
	}
}