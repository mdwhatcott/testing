package so_test

import (
	"errors"
	"testing"

	"github.com/mdwhatcott/testing/so"
)

func TestSo(t *testing.T) {
	assertNil(t, so.The(1, shouldPass, 1))
	assertErr(t, so.The(1, shouldFail, 2))
}

func assertErr(t *testing.T, err error) {
	if err != nil {
		return
	}
	t.Helper()
	t.Error("Expected non-<nil> value, got:", err)
}

func assertNil(t *testing.T, err error) {
	if err == nil {
		return
	}
	t.Helper()
	t.Error("Expected <nil> value, got:", err)
}

func shouldPass(actual interface{}, expected ...interface{}) error {
	_ = actual
	_ = expected
	return nil
}
func shouldFail(actual interface{}, expected ...interface{}) error {
	_ = actual
	_ = expected
	return errors.New("failure")
}
