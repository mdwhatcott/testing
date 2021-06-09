package so

// With allows assertions as in: so.With(t).So(actual, should.Equal, expected)
func With(t testingT) *TestingThat {
	return &TestingThat{t: t}
}

// TestingThat is an intermediate type, not to be instantiated directly
type TestingThat struct{ t testingT }

// The as a method is analogous to the standalone The function and reports
// errors directly to the previously provided *testing.T as in:
// so.With(t).The(actual, should.Equal, expected)
func (this *TestingThat) The(actual interface{}, assertion assertion, expected ...interface{}) {
	err := The(actual, assertion, expected...)
	if err != nil {
		this.t.Helper()
		this.t.Error(err)
	}
}

type testingT interface {
	Helper()
	Error(args ...interface{})
}
