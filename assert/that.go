package assert

import "github.com/mdwhatcott/testing/should"

// That starts an assertion, as in: err := assert.That(actual).Equals(expected)
func That(actual interface{}) *Assertion {
	return &Assertion{actual: actual}
}

// Assertion is an intermediate type, not to be instantiated directly.
type Assertion struct {
	actual interface{}
}

// IsNil asserts that the value provided to That is nil.
func (this *Assertion) IsNil() error {
	return should.BeNil(this.actual)
}

// IsTrue asserts that the value provided to That is true.
func (this *Assertion) IsTrue() error {
	return should.BeTrue(this.actual)
}

// IsFalse asserts that the value provided to That is false.
func (this *Assertion) IsFalse() error {
	return should.BeFalse(this.actual)
}

// Equals asserts that the value provided is equal to the expected value.
func (this *Assertion) Equals(expected interface{}) error {
	return should.Equal(this.actual, expected)
}
