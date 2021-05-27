/*
Package suite implements an xUnit-style test
runner, aiming for an optimum balance between
simplicity and utility. It is based on the
following libraries:

	- github.com/stretchr/testify/suite
	- github.com/smartystreets/gunit
 */
package suite

import (
	"reflect"
	"strings"
	"testing"
)

/*
Run accepts a fixture with Test* methods and
optional setup/teardown methods and executes
the suite. Fixtures must be struct types which
embed a *testing.T. Assuming a fixture struct
with test methods 'Test1' and 'Test2' execution
would proceed as follows:

	1. fixture.SetupSuite()
	2. fixture.Setup()
	3. fixture.Test1()
	4. fixture.Teardown()
	5. fixture.Setup()
	6. fixture.Test2()
	7. fixture.Teardown()
	8. fixture.TeardownSuite()

The methods provided by Options may be supplied
to this function to tweak the execution.
*/
func Run(fixture interface{}, options ...option) {
	config := new(config)
	for _, option := range options {
		option(config)
	}
	fixtureValue := reflect.ValueOf(fixture)
	fixtureType := reflect.TypeOf(fixture)
	t := fixtureValue.Elem().FieldByName("T").Interface().(*testing.T)

	var (
		testNames        []string
		skippedTestNames []string
		focusedTestNames []string
	)
	for x := 0; x < fixtureType.NumMethod(); x++ {
		name := fixtureType.Method(x).Name
		if strings.HasPrefix(name, "Test") || strings.HasPrefix(name, "LongTest") {
			testNames = append(testNames, name)
		} else if name == "SkipNow" { // from embedded *testing.T
			continue
		} else if strings.HasPrefix(name, "Skip") {
			skippedTestNames = append(skippedTestNames, name)
		} else if strings.HasPrefix(name, "Focus") {
			focusedTestNames = append(focusedTestNames, name)
		}
	}

	if len(focusedTestNames) > 0 {
		testNames = focusedTestNames
	}

	if len(testNames) == 0 {
		t.Skip("NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)")
		return
	}

	setup, hasSetup := fixture.(setupSuite)
	if hasSetup {
		setup.SetupSuite()
	}

	teardown, hasTeardown := fixture.(teardownSuite)
	if hasTeardown {
		defer teardown.TeardownSuite()
	}

	for _, testMethodName := range skippedTestNames {
		testMethod := fixtureValue.MethodByName(testMethodName)
		_, isNiladic := testMethod.Interface().(func())
		if isNiladic {
			t.Run(testMethodName, func(t *testing.T) {
				t.Skip("Skipping:", testMethodName)
			})
		}
	}
	for _, testMethodName := range testNames {
		if config.freshFixture {
			fixtureValue = reflect.New(fixtureType.Elem())
		}

		testMethod := fixtureValue.MethodByName(testMethodName)
		test, isNiladic := testMethod.Interface().(func())
		if isNiladic {
			if (strings.HasPrefix(testMethodName, "Long") || strings.HasPrefix(testMethodName, "FocusLong")) && testing.Short() {
				t.Run(testMethodName, func(t *testing.T) {
					t.Skip("Skipping long-running test in -test.short mode.")
				})
			} else {
				t.Run(testMethodName, func(t *testing.T) {
					fixtureValue.Elem().FieldByName("T").Set(reflect.ValueOf(t))
					setup, hasSetup := fixtureValue.Interface().(setupTest)
					if hasSetup {
						setup.Setup()
					}

					teardown, hasTeardown := fixtureValue.Interface().(teardownTest)
					if hasTeardown {
						defer teardown.Teardown()
					}

					test()
				})
			}
		}
	}
}

type config struct {
	freshFixture    bool
	parallelFixture bool
	parallelTests   bool
}

type option func(*config)

type Opt struct{}

// Options provides the sole entrypoint
// to the option functions provided by
// this package.
var Options Opt

// FreshFixture signals to Run that the
// new instances of the provided fixture
// are to be instantiated for each and
// every test case. The Setup and Teardown
// methods are also executed on the
// specifically instantiated fixtures.
// NOTE: the SetupSuite and TeardownSuite
// methods are always run on the provided
// fixture instance, regardless of this
// options having been provided.
func (Opt) FreshFixture() option {
	return func(c *config) {
		c.freshFixture = true
	}
}

// SharedFixture signals to Run that the
// provided fixture instance is to be used
// to run all test methods. This mode is
// not compatible with ParallelFixture or
// ParallelTests and disables them.
func (Opt) SharedFixture() option {
	return func(c *config) {
		c.freshFixture = false
		c.parallelTests = false
		c.parallelFixture = false
	}
}

// ParallelFixture signals to Run that the
// provided fixture instance can be executed
// in parallel with other go test functions.
// This option assumes that `go test` was
// invoked with the -parallel flag.
func (Opt) ParallelFixture() option {
	return func(c *config) {
		c.parallelFixture = true
		c.freshFixture = false
	}
}

// ParallelTests signals to Run that the
// test methods on the provided fixture
// instance can be executed in parallel
// with each other. This option assumes
// that `go test` was invoked with the
// -parallel flag.
func (Opt) ParallelTests() option {
	return func(c *config) {
		c.parallelTests = true
		c.freshFixture = false
	}
}

type (
	setupSuite    interface{ SetupSuite() }
	setupTest     interface{ Setup() }
	teardownTest  interface{ Teardown() }
	teardownSuite interface{ TeardownSuite() }
)
