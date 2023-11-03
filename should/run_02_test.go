package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: t}
	should.Run(fixture, should.Options.UnitTests())
	should.So(t, fixture.counter, should.Equal, 0)
}

type Suite02 struct {
	*testing.T
	counter int
}

func (this *Suite02) TestSomething() {
	this.T.Log("*** this should appear in the test log!")
	this.counter++
}
