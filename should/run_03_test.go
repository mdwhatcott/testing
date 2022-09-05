package should_test

import (
	"testing"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: New(t)}
	Run(fixture)
	fixture.So(t.Failed(), BeFalse)
}

type Suite03 struct{ *T }

func (this *Suite03) SkipTestThatFails() {
	this.So(1, Equal, 2)
}
func (this *Suite03) SkipLongTestThatFails() {
	this.So(1, Equal, 2)
}
