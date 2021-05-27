package suite_test

import "fmt"

type FakeT struct{ failures []string }

func (this *FakeT) Helper() {}
func (this *FakeT) Errorf(format string, args ...interface{}) {
	this.failures = append(this.failures, fmt.Sprintf(format, args...))
}
