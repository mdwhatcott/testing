package assert

// With allows assertions as in: assert.With(t).That(actual).Equals(expected)
func With(t testingT) *TestingThat {
	return &TestingThat{t: t}
}

// TestingThat is an intermediate type, not to be instantiated directly
type TestingThat struct{ t testingT }

// That is an intermediate method call, as in: assert.With(t).That(actual).Equals(expected)
func (this *TestingThat) That(actual interface{}) *TestingAssertion {
	return &TestingAssertion{
		testingT:  this.t,
		Assertion: That(actual),
	}
}

type testingT interface {
	Helper()
	Error(args ...interface{})
}

// TestingAssertion is an intermediate type, not to be instantiated directly.
type TestingAssertion struct {
	testingT
	*Assertion
}

// IsNil asserts that the value provided to That is nil.
func (this *TestingAssertion) IsNil() {
	err := this.Assertion.IsNil()
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// IsTrue asserts that the value provided to That is true.
func (this *TestingAssertion) IsTrue() {
	err := this.Assertion.IsTrue()
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// IsFalse asserts that the value provided to That is false.
func (this *TestingAssertion) IsFalse() {
	err := this.Assertion.IsFalse()
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// Equals asserts that the value provided is equal to the expected value.
func (this *TestingAssertion) Equals(expected interface{}) {
	err := this.Assertion.Equals(expected)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}
