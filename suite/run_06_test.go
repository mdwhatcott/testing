package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestSuiteWithSetupsAndTeardownsSkipped(t *testing.T) {
	fixture := &Suite06{T: t}

	suite.Run(fixture, suite.Options.SharedFixture())

	assert.With(t).That(fixture.events).IsNil()
}

type Suite06 struct {
	*testing.T
	events []string
}

func (this *Suite06) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite06) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite06) Setup()              { this.record("Setup") }
func (this *Suite06) Teardown()           { this.record("Teardown") }
func (this *Suite06) SkipTest()           { this.record("SkipTest") }
func (this *Suite06) record(event string) { this.events = append(this.events, event) }
