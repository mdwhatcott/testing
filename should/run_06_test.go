package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

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
