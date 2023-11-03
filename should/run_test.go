package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSuiteWithSetupsAndTeardowns(t *testing.T) {
	fixture := &Suite01{T: t}

	should.Run(fixture, should.Options.IntegrationTests())

	should.So(t, fixture.events, should.Equal, []string{
		"Setup",
		"Test",
		"Teardown",
	})
}

type Suite01 struct {
	*testing.T
	events []string
}

func (this *Suite01) Setup()              { this.record("Setup") }
func (this *Suite01) Teardown()           { this.record("Teardown") }
func (this *Suite01) Test()               { this.record("Test") }
func (this *Suite01) record(event string) { this.events = append(this.events, event) }

/////////////////////////////////////

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

/////////////////////////////

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: t}
	should.Run(fixture)
	should.So(t, t.Failed(), should.BeFalse)
}

type Suite03 struct{ *testing.T }

func (this *Suite03) SkipTestThatFails() {
	should.So(this, 1, should.Equal, 2)
}

////////////////////////////////////////////////////////////////////////////////////

func TestSuiteWithSetupsAndTeardownsSkippedEntirelyIfAllTestsSkipped(t *testing.T) {
	fixture := &Suite06{T: t}

	should.Run(fixture, should.Options.SharedFixture())

	should.So(t, fixture.events, should.BeNil)
}

type Suite06 struct {
	*testing.T
	events []string
}

func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) SkipTest()           { this.record("SkipTest") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }

//////////////////////////////////////////////

func TestSuiteWithSkippedTests(t *testing.T) {
	fixture := &Suite07{T: t}

	should.Run(fixture, should.Options.SharedFixture())

	should.So(t, fixture.events, should.Equal, []string{
		"Setup",
		"Test1",
		"Teardown",
	})
}

type Suite07 struct {
	*testing.T
	events []string
}

func (this *Suite07) Setup()              { this.record("Setup") }
func (this *Suite07) Teardown()           { this.record("Teardown") }
func (this *Suite07) Test1()              { this.record("Test1") }
func (this *Suite07) SkipTest2()          { this.record("SkipTest2") }
func (this *Suite07) record(event string) { this.events = append(this.events, event) }
