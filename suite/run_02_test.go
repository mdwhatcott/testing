package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestFreshFixture(t *testing.T) {
	fixture := &Suite02{T: t}
	suite.Run(fixture, suite.Options.UnitTests())
	assert.With(t).That(fixture.counter).Equals(0)
}

type Suite02 struct {
	*testing.T
	counter int
}

func (this *Suite02) TestSomething() {
	this.counter++
}
