package should

import "testing"

func So(t *testing.T, actual any, assertion Func, expected ...any) {
	_ = New(t).So(actual, assertion, expected...)
}
