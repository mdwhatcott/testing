package should

import "testing"

type Func func(actual any, expected ...any) error

func So(t *testing.T, actual any, assertion Func, expected ...any) {
	t.Helper()
	if err := assertion(actual, expected...); err != nil {
		t.Error(err)
	}
}
