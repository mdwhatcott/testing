package so_test

import (
	"fmt"
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/so"
)

func TestPass(t *testing.T) {
	T := new(FakeT)

	so.With(T).The(nil, should.BeNil)
	so.With(T).The(true, should.BeTrue)
	so.With(T).The(true, should.BeTrue)
	so.With(T).The(false, should.BeFalse)

	if len(T.failures) > 0 {
		t.Error("Unexpected failures:", T.failures)
	}
	if T.helps > 0 {
		t.Error("Unexpected helper calls:", T.helps)
	}
}

func TestFail(t *testing.T) {
	T := new(FakeT)

	so.With(T).The(true, should.BeFalse)
	so.With(T).The(false, should.BeTrue)
	so.With(T).The(t, should.BeNil)

	if len(T.failures) != 3 {
		t.Error("Expected 3 failures, got:", T.failures)
	}
	if T.helps != 3 {
		t.Error("Expected 3 calls to t.Helper(), got:", T.helps)
	}
}

type FakeT struct {
	helps    int
	failures []string
}

func (this *FakeT) Helper() { this.helps++ }
func (this *FakeT) Error(args ...interface{}) {
	this.failures = append(this.failures, fmt.Sprint(args...))
}
