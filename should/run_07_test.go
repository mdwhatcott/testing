package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

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
