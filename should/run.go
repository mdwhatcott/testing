package should

import (
	"reflect"
	"strings"
	"testing"
)

// Run accepts a fixture with Test* methods and
// optional setup/teardown methods and executes
// the suite. Fixtures must be struct types which
// embed a *testing.T. Assuming a fixture struct
// with test methods 'Test1' and 'Test2' execution
// would proceed as follows:
//
//  1. fixture.Setup()
//  2. fixture.Test1()
//  3. fixture.Teardown()
//  4. fixture.Setup()
//  5. fixture.Test2()
//  6. fixture.Teardown()
func Run(fixture any) {
	fixtureValue := reflect.ValueOf(fixture)
	fixtureType := reflect.TypeOf(fixture)
	t := fixtureValue.Elem().FieldByName("T").Interface().(*testing.T)

	var (
		testNames        []string
		skippedTestNames []string
	)
	for x := 0; x < fixtureType.NumMethod(); x++ {
		name := fixtureType.Method(x).Name
		method := fixtureValue.MethodByName(name)
		_, isNiladic := method.Interface().(func())
		if !isNiladic {
			continue
		}
		if strings.HasPrefix(name, "Test") {
			testNames = append(testNames, name)
		} else if strings.HasPrefix(name, "SkipTest") {
			skippedTestNames = append(skippedTestNames, name)
		}
	}
	for _, name := range skippedTestNames {
		testCase{T: t, manualSkip: true, name: name}.Run()
	}
	for _, name := range testNames {
		testCase{T: t, name: name, fixtureType: fixtureType, fixtureValue: fixtureValue}.Run()
	}
}

type testCase struct {
	*testing.T
	name         string
	manualSkip   bool
	fixtureType  reflect.Type
	fixtureValue reflect.Value
}

func (this testCase) Run() {
	_ = this.T.Run(this.name, this.decideRun())
}
func (this testCase) decideRun() func(*testing.T) {
	if this.manualSkip {
		return this.skipFunc("Skipping: " + this.name)
	}
	return this.runTest
}
func (this testCase) skipFunc(message string) func(*testing.T) {
	return func(t *testing.T) { t.Skip(message) }
}
func (this testCase) runTest(t *testing.T) {
	fixtureValue := this.fixtureValue
	fixtureValue = reflect.New(this.fixtureType.Elem())
	fixtureValue.Elem().FieldByName("T").Set(reflect.ValueOf(t))

	setup, hasSetup := fixtureValue.Interface().(setupTest)
	if hasSetup {
		setup.Setup()
	}
	teardown, hasTeardown := fixtureValue.Interface().(teardownTest)
	if hasTeardown {
		defer teardown.Teardown()
	}
	fixtureValue.MethodByName(this.name).Call(nil)
}

type (
	setupTest    interface{ Setup() }
	teardownTest interface{ Teardown() }
)
