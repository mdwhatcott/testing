package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
	"github.com/mdwhatcott/testing/suite"
)

func TestSuiteWithSkippedTests(t *testing.T) {
	fixture := &Suite07{T: &suite.T{T: t}}

	suite.Run(fixture, suite.Options.SharedFixture())

	fixture.So(fixture.events, should.Equal, []string{
		"SetupSuite",
		"Setup",
		"Test1",
		"Teardown",
		"TeardownSuite",
	})
}

func TestAssertionMethods(t *testing.T) {
	fixture := &Suite07{T: &suite.T{T: t}}

	suite.Run(fixture, suite.Options.SharedFixture())

	ok := fixture.CheckSo(1, should.Equal, 2)
	if ok {
		t.Error("want false, got true")
	}
	err := fixture.VerifySo(1, should.Equal, 2)
	if err != nil {
		t.Error("want err, got nil")
	}
}

type Suite07 struct {
	*suite.T
	events []string
}

func (this *Suite07) SetupSuite()         { this.record("SetupSuite") }
func (this *Suite07) TeardownSuite()      { this.record("TeardownSuite") }
func (this *Suite07) Setup()              { this.record("Setup") }
func (this *Suite07) Teardown()           { this.record("Teardown") }
func (this *Suite07) Test1()              { this.record("Test1") }
func (this *Suite07) SkipTest2()          { this.record("SkipTest2") }
func (this *Suite07) record(event string) { this.events = append(this.events, event) }
