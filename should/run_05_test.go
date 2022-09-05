package should_test

import (
	"testing"
)

func TestFocus(t *testing.T) {
	fixture := &Suite05{
		T:      New(t),
		events: make(map[string]struct{}),
	}

	Run(fixture, Options.SharedFixture())

	fixture.So(t.Failed(), BeFalse)
	if testing.Short() {
		fixture.So(fixture.events, Equal, map[string]struct{}{"1": {}})
	} else {
		fixture.So(fixture.events, Equal, map[string]struct{}{
			"1": {},
			"2": {},
		})
	}
}

type Suite05 struct {
	*T
	events map[string]struct{}
}

func (this *Suite05) FocusTest1() {
	this.events["1"] = struct{}{}
}
func (this *Suite05) FocusLongTest2() {
	this.events["2"] = struct{}{}
}
func (this *Suite05) TestThatFails() {
	this.So(1, Equal, 2)
}
