package should

import "testing"

type Func func(actual any, expected ...any) error

type T struct{ *testing.T }

func New(t *testing.T) *T {
	return &T{T: t}
}
func (this *T) So(actual any, assertion Func, expected ...any) (ok bool) {
	this.Helper()
	err := assertion(actual, expected...)
	if err != nil {
		this.Error(err)
	}
	return err == nil
}
func (this *T) Print(v ...any)            { this.Helper(); this.Log(v...) }
func (this *T) Printf(f string, v ...any) { this.Helper(); this.Logf(f, v...) }
func (this *T) Println(v ...any)          { this.Helper(); this.Log(v...) }
func (this *T) Write(p []byte) (n int, err error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}
