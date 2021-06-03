package suite

import "testing"

// TODO: docs

type T struct{ *testing.T }

type assertion func(actual interface{}, expected ...interface{}) error

func (this *T) So(actual interface{}, assertion assertion, expected ...interface{}) {
	err := assertion(actual, expected...)
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}

func (this *T) Write(p []byte) (n int, err error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}
