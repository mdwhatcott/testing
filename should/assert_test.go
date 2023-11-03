package should_test

import (
	"errors"
	"fmt"
	"math"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
	"time"

	"github.com/mdwhatcott/testing/should"
)

type fakeT struct {
	errs []any
}

func (this *fakeT) Helper() {}
func (this *fakeT) Error(a ...any) {
	this.errs = a
}

func TestSoFailure(t *testing.T) {
	fakeT := &fakeT{}
	should.So(fakeT, 1, should.Equal, 2)
	if len(fakeT.errs) != 1 {
		t.Fatal("expected 1 err, got:", len(fakeT.errs))
	}
	err := fakeT.errs[0].(error)
	if !errors.Is(err, should.ErrAssertionFailure) {
		t.Error("expected assertion failure, got:", err)
	}
}

type Assertion struct{ *testing.T }

func NewAssertion(t *testing.T) *Assertion {
	return &Assertion{T: t}
}
func (this *Assertion) ExpectedCountInvalid(actual any, assertion should.Func, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrExpectedCountInvalid)
}
func (this *Assertion) TypeMismatch(actual any, assertion should.Func, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrTypeMismatch)
}
func (this *Assertion) KindMismatch(actual any, assertion should.Func, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrKindMismatch)
}
func (this *Assertion) Fail(actual any, assertion should.Func, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, should.ErrAssertionFailure)
}
func (this *Assertion) Pass(actual any, assertion should.Func, expected ...any) {
	this.Helper()
	this.err(actual, assertion, expected, nil)
}
func (this *Assertion) err(actual any, assertion should.Func, expected []any, expectedErr error) {
	this.Helper()
	_, file, line, _ := runtime.Caller(2)
	subTest := fmt.Sprintf("%s:%d", filepath.Base(file), line)
	this.Run(subTest, func(t *testing.T) {
		t.Helper()
		err := assertion(actual, expected...)
		if !errors.Is(err, expectedErr) {
			t.Errorf("[FAIL]\n"+
				"expected: %v\n"+
				"actual:   %v",
				expected,
				actual,
			)
		} else if testing.Verbose() {
			t.Log(
				"\n", err, "\n",
				"(above error report printed for visual inspection)",
			)
		}
	})
}

func TestShouldBeEmpty(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeEmpty, "EXTRA")

	assert.KindMismatch(42, should.BeEmpty)

	assert.Pass([]string(nil), should.BeEmpty)
	assert.Pass(make([]string, 0, 0), should.BeEmpty)
	assert.Pass(make([]string, 0, 1), should.BeEmpty)
	assert.Fail([]string{""}, should.BeEmpty)

	assert.Pass([0]string{}, should.BeEmpty) // The only possible empty array!
	assert.Fail([1]string{}, should.BeEmpty)

	assert.Pass(chan string(nil), should.BeEmpty)
	assert.Pass(make(chan string), should.BeEmpty)
	assert.Pass(make(chan string, 1), should.BeEmpty)
	assert.Fail(nonEmptyChannel(), should.BeEmpty)

	assert.Pass(map[string]string(nil), should.BeEmpty)
	assert.Pass(make(map[string]string), should.BeEmpty)
	assert.Pass(make(map[string]string, 1), should.BeEmpty)
	assert.Fail(map[string]string{"": ""}, should.BeEmpty)

	assert.Pass("", should.BeEmpty)
	assert.Pass(*new(string), should.BeEmpty)
	assert.Fail(" ", should.BeEmpty)
}

func TestShouldNotBeEmpty(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeEmpty, "EXTRA")
	assert.KindMismatch(42, should.NOT.BeEmpty)

	assert.Fail([]string(nil), should.NOT.BeEmpty)
	assert.Fail(make([]string, 0, 0), should.NOT.BeEmpty)
	assert.Fail(make([]string, 0, 1), should.NOT.BeEmpty)
	assert.Pass([]string{""}, should.NOT.BeEmpty)

	assert.Fail([0]string{}, should.NOT.BeEmpty)
	assert.Pass([1]string{}, should.NOT.BeEmpty)

	assert.Fail(chan string(nil), should.NOT.BeEmpty)
	assert.Fail(make(chan string), should.NOT.BeEmpty)
	assert.Fail(make(chan string, 1), should.NOT.BeEmpty)
	assert.Pass(nonEmptyChannel(), should.NOT.BeEmpty)

	assert.Fail(map[string]string(nil), should.NOT.BeEmpty)
	assert.Fail(make(map[string]string), should.NOT.BeEmpty)
	assert.Fail(make(map[string]string, 1), should.NOT.BeEmpty)
	assert.Pass(map[string]string{"": ""}, should.NOT.BeEmpty)

	assert.Fail("", should.NOT.BeEmpty)
	assert.Fail(*new(string), should.NOT.BeEmpty)
	assert.Pass(" ", should.NOT.BeEmpty)
}

func nonEmptyChannel() chan string {
	c := make(chan string, 1)
	c <- ""
	return c
}

func TestShouldBeFalse(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeFalse, "EXTRA")
	assert.TypeMismatch(1, should.BeFalse)

	assert.Fail(true, should.BeFalse)
	assert.Pass(false, should.BeFalse)
}

func TestShouldBeIn(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeIn)
	assert.ExpectedCountInvalid("actual", should.BeIn, "EXPECTED", "EXTRA")

	assert.KindMismatch(false, should.BeIn, "string")
	assert.KindMismatch("hi", should.BeIn, 1)

	// strings:
	assert.Fail("no", should.BeIn, "")
	assert.Pass("rat", should.BeIn, "integrate")
	assert.Pass('b', should.BeIn, "abc")

	// slices:
	assert.Fail('d', should.BeIn, []byte("abc"))
	assert.Pass('b', should.BeIn, []byte("abc"))
	assert.Pass(98, should.BeIn, []byte("abc"))

	// arrays:
	assert.Fail('d', should.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Pass('b', should.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Pass(98, should.BeIn, [3]byte{'a', 'b', 'c'})

	// maps:
	assert.Fail('b', should.BeIn, map[rune]int{'a': 1})
	assert.Pass('a', should.BeIn, map[rune]int{'a': 1})
}

func TestShouldNotBeIn(t *testing.T) {
	assert := NewAssertion(t)
	assert.ExpectedCountInvalid("actual", should.NOT.BeIn)
	assert.ExpectedCountInvalid("actual", should.NOT.BeIn, "EXPECTED", "EXTRA")
	assert.KindMismatch(false, should.NOT.BeIn, "string")
	assert.KindMismatch("hi", should.NOT.BeIn, 1)

	// strings:
	assert.Pass("no", should.NOT.BeIn, "yes")
	assert.Fail("rat", should.NOT.BeIn, "integrate")
	assert.Fail('b', should.NOT.BeIn, "abc")

	// slices:
	assert.Pass('d', should.NOT.BeIn, []byte("abc"))
	assert.Fail('b', should.NOT.BeIn, []byte("abc"))
	assert.Fail(98, should.NOT.BeIn, []byte("abc"))

	// arrays:
	assert.Pass('d', should.NOT.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Fail('b', should.NOT.BeIn, [3]byte{'a', 'b', 'c'})
	assert.Fail(98, should.NOT.BeIn, [3]byte{'a', 'b', 'c'})

	// maps:
	assert.Pass('b', should.NOT.BeIn, map[rune]int{'a': 1})
	assert.Fail('a', should.NOT.BeIn, map[rune]int{'a': 1})
}

func TestShouldBeNil(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeNil, "EXTRA")

	assert.Pass(nil, should.BeNil)
	assert.Pass([]string(nil), should.BeNil)
	assert.Pass((*string)(nil), should.BeNil)
	assert.Fail(notNil, should.BeNil)
}

func TestShouldNotBeNil(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.BeNil, "EXTRA")

	assert.Fail(nil, should.NOT.BeNil)
	assert.Fail([]string(nil), should.NOT.BeNil)
	assert.Pass(notNil, should.NOT.BeNil)
}

var notNil = errors.New("not nil")

func TestShouldBeTrue(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.BeTrue, "EXTRA")

	assert.TypeMismatch(1, should.BeTrue)

	assert.Fail(false, should.BeTrue)
	assert.Pass(true, should.BeTrue)
}

func TestShouldContain(t *testing.T) {
	assert := NewAssertion(t)
	assert.ExpectedCountInvalid("actual", should.Contain)
	assert.ExpectedCountInvalid("actual", should.Contain, "EXPECTED", "EXTRA")

	assert.KindMismatch("string", should.Contain, false)
	assert.KindMismatch(1, should.Contain, "hi")

	// strings:
	assert.Fail("", should.Contain, "no")
	assert.Pass("integrate", should.Contain, "rat")
	assert.Pass("abc", should.Contain, 'b')

	// slices:
	assert.Fail([]byte("abc"), should.Contain, 'd')
	assert.Pass([]byte("abc"), should.Contain, 'b')
	assert.Pass([]byte("abc"), should.Contain, 98)

	// arrays:
	assert.Fail([3]byte{'a', 'b', 'c'}, should.Contain, 'd')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.Contain, 'b')
	assert.Pass([3]byte{'a', 'b', 'c'}, should.Contain, 98)

	// maps:
	assert.Fail(map[rune]int{'a': 1}, should.Contain, 'b')
	assert.Pass(map[rune]int{'a': 1}, should.Contain, 'a')
}

func TestShouldNotContain(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Contain)
	assert.ExpectedCountInvalid("actual", should.NOT.Contain, "EXPECTED", "EXTRA")

	assert.KindMismatch(false, should.NOT.Contain, "string")
	assert.KindMismatch("hi", should.NOT.Contain, 1)

	// strings:
	assert.Pass("", should.NOT.Contain, "no")
	assert.Fail("integrate", should.NOT.Contain, "rat")
	assert.Fail("abc", should.NOT.Contain, 'b')

	// slices:
	assert.Pass([]byte("abc"), should.NOT.Contain, 'd')
	assert.Fail([]byte("abc"), should.NOT.Contain, 'b')
	assert.Fail([]byte("abc"), should.NOT.Contain, 98)

	// arrays:
	assert.Pass([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 'd')
	assert.Fail([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 'b')
	assert.Fail([3]byte{'a', 'b', 'c'}, should.NOT.Contain, 98)

	// maps:
	assert.Pass(map[rune]int{'a': 1}, should.NOT.Contain, 'b')
	assert.Fail(map[rune]int{'a': 1}, should.NOT.Contain, 'a')
}

func TestShouldEqual(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.Equal)
	assert.ExpectedCountInvalid("actual", should.Equal, "EXPECTED", "EXTRA")

	assert.Fail(1, should.Equal, 2)
	assert.Pass(1, should.Equal, 1)
	assert.Pass(1, should.Equal, uint(1))

	now := time.Now()
	assert.Pass(now.UTC(), should.Equal, now.In(time.Local))
	assert.Fail(time.Now(), should.Equal, time.Now())

	assert.Fail(struct{ A string }{}, should.Equal, struct{ B string }{})
	assert.Pass(struct{ A string }{}, should.Equal, struct{ A string }{})

	assert.Fail([]byte("hi"), should.Equal, []byte("bye"))
	assert.Pass([]byte("hi"), should.Equal, []byte("hi"))

	const MAX uint64 = math.MaxUint64
	assert.Fail(-1, should.Equal, MAX)
	assert.Fail(MAX, should.Equal, -1)

	assert.Pass(returnsNilInterface(), should.Equal, nil)

	a := strings.Repeat("a", 55)
	b := strings.Repeat("a", 45) + "XXXX" + "aaaaaa"
	assert.Fail(a, should.Equal, b)
}

func TestShouldNotEqual(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Equal)
	assert.ExpectedCountInvalid("actual", should.NOT.Equal, "EXPECTED", "EXTRA")

	assert.Fail(1, should.NOT.Equal, 1)
	assert.Pass(1, should.NOT.Equal, 2)
}

func returnsNilInterface() any { return nil }

func TestShouldPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.Panic)

	assert.Fail(func() {}, should.Panic)
	assert.Pass(func() { panic("yay") }, should.Panic)
	assert.Pass(func() { panic(nil) }, should.Panic) // tricky!
}

func TestShouldNotPanic(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.NOT.Panic, "EXPECTED", "EXTRA")
	assert.TypeMismatch("wrong type", should.NOT.Panic)

	assert.Fail(func() { panic("boo") }, should.NOT.Panic)
	assert.Pass(func() {}, should.NOT.Panic)
}

func TestShouldWrapError(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual", should.WrapError)
	assert.ExpectedCountInvalid("actual", should.WrapError, "EXPECTED", "EXTRA")

	assert.TypeMismatch(inner, should.WrapError, 42)
	assert.TypeMismatch(42, should.WrapError, inner)

	assert.Pass(outer, should.WrapError, inner)
	assert.Fail(inner, should.WrapError, outer)
}

var (
	inner = errors.New("inner")
	outer = fmt.Errorf("outer(%w)", inner)
)
