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
