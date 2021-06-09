package so_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/so"
)

func TestSo(t *testing.T) {
	assertNil(t, so.The(1, should.Equal, 1))
	assertErr(t, so.The(1, should.Equal, 2))
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
