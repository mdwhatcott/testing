package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestFocus(t *testing.T) {
	fixture := &Suite05{
		T:      t,
		events: make(map[string]struct{}),
	}

	suite.Run(fixture, suite.Options.SharedFixture())

	assert.With(t).That(t.Failed()).IsFalse()
	if testing.Short() {
		assert.With(t).That(fixture.events).Equals(map[string]struct{}{"1": {}})
	} else {
		assert.With(t).That(fixture.events).Equals(map[string]struct{}{
			"1": {},
			"2": {},
		})
	}
}

type Suite05 struct {
	*testing.T
	events map[string]struct{}
}

func (this *Suite05) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite05) FocusLongTest2() {
	this.events["2"] = struct{}{}
}
func (this *Suite05) TestThatFails() {
	assert.With(this).That(1).Equals(2)
}
