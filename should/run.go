package should

import (
	"reflect"
	"strings"
	"testing"
)

type config struct {
	freshFixture    bool
	parallelFixture bool
	parallelTests   bool
}

// Option is a function that modifies a config.
// See Options for provided behaviors.
type Option func(*config)

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
func (Opt) FreshFixture() Option {
	return func(c *config) {
		c.freshFixture = true
	}
}

// SharedFixture signals to Run that the
// provided fixture instance is to be used
// to run all test methods. This mode is
// not compatible with ParallelFixture or
// ParallelTests and disables them.
func (Opt) SharedFixture() Option {
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
func (Opt) ParallelFixture() Option {
	return func(c *config) {
		c.parallelFixture = true
	}
}

// ParallelTests signals to Run that the
// test methods on the provided fixture
// instance can be executed in parallel
// with each other. This option assumes
// that `go test` was invoked with the
// -parallel flag.
func (Opt) ParallelTests() Option {
	return func(c *config) {
		c.parallelTests = true
		c.freshFixture = true
		Options.FreshFixture()(c)
	}
}

// UnitTests is a composite option that
// signals to Run that the test suite can
// be treated as a unit-test suite by
// employing parallelism and fresh fixtures
// to maximize the chances of exposing
// unwanted coupling between tests.
func (Opt) UnitTests() Option {
	return func(c *config) {
		Options.ParallelTests()(c)
		Options.ParallelFixture()(c)
	}
}

// IntegrationTests is a composite option that
// signals to Run that the test suite should be
// treated as an integration test suite, avoiding
// parallelism and utilizing shared fixtures to
// allow reuse of potentially expensive resources.
func (Opt) IntegrationTests() Option {
	return func(c *config) {
		Options.SharedFixture()(c)
	}
}

/*
Run accepts a fixture with Test* methods and
optional setup/teardown methods and executes
the suite. Fixtures must be struct types which
embed a *suite.T. Assuming a fixture struct
with test methods 'Test1' and 'Test2' execution
would proceed as follows:

 2. fixture.Setup()
 3. fixture.Test1()
 4. fixture.Teardown()
 5. fixture.Setup()
 6. fixture.Test2()
 7. fixture.Teardown()

The methods provided by Options may be supplied
to this function to tweak the execution.
*/
func Run(fixture any, options ...Option) {
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

	if len(testNames) == 0 {
		t.Skip("NOT IMPLEMENTED (no test cases defined, or they are all marked as skipped)")
		return
	}

	if config.parallelFixture {
		t.Parallel()
	}

	for _, name := range skippedTestNames {
		testCase{T: t, manualSkip: true, name: name}.Run()
	}

	for _, name := range testNames {
		testCase{T: t, name: name, config: config, fixtureType: fixtureType, fixtureValue: fixtureValue}.Run()
	}
}

type testCase struct {
	*testing.T
	name         string
	config       *config
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
	if this.config.parallelTests {
		t.Parallel()
	}

	fixtureValue := this.fixtureValue
	if this.config.freshFixture {
		fixtureValue = reflect.New(this.fixtureType.Elem())
	}
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
