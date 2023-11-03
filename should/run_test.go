package should_test

import (
	"sync"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

var (
	mutex      sync.Mutex
	testEvents []string
)

func TestRunSuite(t *testing.T) {
	should.Run(&Suite{T: t})

	mutex.Lock()
	defer mutex.Unlock()
	should.So(t, testEvents, should.Equal, []string{
		"Setup", "Test1", "Teardown",
		"Setup", "Test3", "Teardown",
	})
}

type Suite struct{ *testing.T }

func (this *Suite) Setup()     { this.record("Setup") }
func (this *Suite) Teardown()  { this.record("Teardown") }
func (this *Suite) Test1()     { this.record("Test1") }
func (this *Suite) SkipTest2() { this.record("SkipTest2") }
func (this *Suite) Test3()     { this.record("Test3") }

func (this *Suite) record(event string) {
	mutex.Lock()
	defer mutex.Unlock()
	testEvents = append(testEvents, event)
}
