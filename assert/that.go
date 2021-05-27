package assert

import "reflect"

// With allows assertions as in: assert.With(t).That(actual).Equals(expected)
func With(t T) *that {
	return &that{t: t}
}
func (this *that) That(actual interface{}) *assertion {
	return That(this.t, actual)
}

type that struct{ t T }

// That allows assertions as in: assert.That(t, actual).Equals(expected)
func That(t T, actual interface{}) *assertion {
	return &assertion{T: t, actual: actual}
}

type T interface {
	Helper()
	Errorf(format string, args ...interface{})
}

type assertion struct {
	T
	actual interface{}
}

func (this *assertion) IsNil() {
	this.Helper()
	if this.actual != nil && !reflect.ValueOf(this.actual).IsNil() {
		this.Equals(nil)
	}
}
func (this *assertion) IsTrue() {
	this.Helper()
	this.Equals(true)
}
func (this *assertion) IsFalse() {
	this.Helper()
	this.Equals(false)
}
func (this *assertion) Equals(expected interface{}) {
	this.Helper()

	if !reflect.DeepEqual(this.actual, expected) {
		this.Errorf("\n"+
			"Expected: %#v\n"+
			"Actual:   %#v",
			expected,
			this.actual,
		)
	}
}
