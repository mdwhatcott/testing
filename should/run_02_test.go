package should_test

import (
	"testing"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: New(t)}
	Run(fixture, Options.UnitTests())
	fixture.So(fixture.counter, Equal, 0)
}

type Suite02 struct {
	*T
	counter int
}

func (this *Suite02) TestSomething() {
	_, _ = this.Write([]byte("*** this should appear in the test log!"))
	this.counter++
}
