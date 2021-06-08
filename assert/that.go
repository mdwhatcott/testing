package assert

import "github.com/mdwhatcott/testing/should"

// With allows assertions as in: assert.With(t).That(actual).Equals(expected)
func With(t testingT) *That {
	return &That{t: t}
}

// That is an intermediate type, not to be instantiated directly
type That struct{ t testingT }

// That is an intermediate method call, as in: assert.With(t).That(actual).Equals(expected)
func (this *That) That(actual interface{}) *Assertion {
	return &Assertion{
		testingT: this.t,
		actual:   actual,
	}
}

type testingT interface {
	Helper()
	Error(args ...interface{})
}

// Assertion is an intermediate type, not to be instantiated directly.
type Assertion struct {
	testingT
	actual interface{}
}

// IsNil asserts that the value provided to That is nil.
func (this *Assertion) IsNil() {
	err := should.BeNil(this.actual)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// IsTrue asserts that the value provided to That is true.
func (this *Assertion) IsTrue() {
	err := should.BeTrue(this.actual)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// IsFalse asserts that the value provided to That is false.
func (this *Assertion) IsFalse() {
	err := should.BeFalse(this.actual)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// Equals asserts that the value provided is equal to the expected value.
func (this *Assertion) Equals(expected interface{}) {
	err := should.Equal(this.actual, expected)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}
