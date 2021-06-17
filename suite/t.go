package suite

import "testing"

// T embeds *testing.T and provides convenient
// hooks for making assertions and other operations.
type T struct{ *testing.T }

// So invokes the provided assertion with the provided args.
// In the event of an assertion failure it calls *testing.T.Error.
func (this *T) So(actual interface{}, assertion assertion, expected ...interface{}) {
	err := assertion(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

// FatalSo is like So but in the event of an assertion failure it calls *testing.T.Fatal.
func (this *T) FatalSo(actual interface{}, assertion assertion, expected ...interface{}) {
	err := assertion(actual, expected...)
	if err != nil {
		this.Helper()
		this.Fatal(err)
	}
}

// VerifySo simply invokes the provided assertion and returns the resulting error.
func (this *T) VerifySo(actual interface{}, assertion assertion, expected ...interface{}) error {
	return assertion(actual, expected...)
}

// CheckSo simply invokes the provided assertion and returns whether the resulting error is nil.
func (this *T) CheckSo(actual interface{}, assertion assertion, expected ...interface{}) bool {
	return assertion(actual, expected...) == nil
}

// Write implements io.Writer allowing for the
// suite to serve as a convenient log target,
// among other use cases.
func (this *T) Write(p []byte) (n int, err error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}

type assertion func(actual interface{}, expected ...interface{}) error
