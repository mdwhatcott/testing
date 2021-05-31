# github.com/mdwhatcott/testing



	package assert // import "github.com/mdwhatcott/testing/assert"
	
	
	TYPES
	
	type Assertion struct {
		// Has unexported fields.
	}
	    Assertion is an intermediate type, not to be instantiated directly.
	
	func (this *Assertion) Equals(expected interface{})
	    Equals asserts that the value provided is equal to the expected value.
	
	func (this *Assertion) IsFalse()
	    IsFalse asserts that the value provided to That is false.
	
	func (this *Assertion) IsNil()
	    IsNil asserts that the value provided to That is nil.
	
	func (this *Assertion) IsTrue()
	    IsTrue asserts that the value provided to That is true.
	
	type That struct {
		// Has unexported fields.
	}
	    That is an intermediate type, not to be instantiated directly
	
	func With(t testingT) *That
	    With allows assertions as in: assert.With(t).That(actual).Equals(expected)
	
	func (this *That) That(actual interface{}) *Assertion
	    That is an intermediate method call, as in:
	    assert.With(t).That(actual).Equals(expected)
	

---

	package suite // import "github.com/mdwhatcott/testing/suite"
	
	Package suite implements an xUnit-style test runner, aiming for an optimum
	balance between simplicity and utility. It is based on the following
	packages:
	
	    - [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
	    - [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)
	
	For those using GoLand by JetBrains, you may find the following "live
	template" helpful:
	
	    func Test$NAME$Suite(t *testing.T) {
	    	suite.Run(&$NAME$Suite{T: t}, suite.Options.UnitTests())
	    }
	
	    type $NAME$Suite struct {
	    	*testing.T
	    }
	
	    func (this *$NAME$Suite) Setup() {
	    }
	
	    func (this *$NAME$Suite) Test$END$() {
	    }
	
	Happy testing!
	
	FUNCTIONS
	
	func Run(fixture interface{}, options ...Option)
	    Run accepts a fixture with Test* methods and optional setup/teardown methods
	    and executes the suite. Fixtures must be struct types which embed a
	    *testing.T. Assuming a fixture struct with test methods 'Test1' and 'Test2'
	    execution would proceed as follows:
	
	        1. fixture.SetupSuite()
	        2. fixture.Setup()
	        3. fixture.Test1()
	        4. fixture.Teardown()
	        5. fixture.Setup()
	        6. fixture.Test2()
	        7. fixture.Teardown()
	        8. fixture.TeardownSuite()
	
	    The methods provided by Options may be supplied to this function to tweak
	    the execution.
	
	
	TYPES
	
	type Opt struct{}
	
	var Options Opt
	    Options provides the sole entrypoint to the option functions provided by
	    this package.
	
	func (Opt) FreshFixture() Option
	    FreshFixture signals to Run that the new instances of the provided fixture
	    are to be instantiated for each and every test case. The Setup and Teardown
	    methods are also executed on the specifically instantiated fixtures. NOTE:
	    the SetupSuite and TeardownSuite methods are always run on the provided
	    fixture instance, regardless of this options having been provided.
	
	func (Opt) IntegrationTests() Option
	    IntegrationTests is a composite option that signals to Run that the test
	    suite should be treated as an integration test suite, avoiding parallelism
	    and utilizing shared fixtures to allow reuse of potentially expensive
	    resources.
	
	func (Opt) ParallelFixture() Option
	    ParallelFixture signals to Run that the provided fixture instance can be
	    executed in parallel with other go test functions. This option assumes that
	    `go test` was invoked with the -parallel flag.
	
	func (Opt) ParallelTests() Option
	    ParallelTests signals to Run that the test methods on the provided fixture
	    instance can be executed in parallel with each other. This option assumes
	    that `go test` was invoked with the -parallel flag.
	
	func (Opt) SharedFixture() Option
	    SharedFixture signals to Run that the provided fixture instance is to be
	    used to run all test methods. This mode is not compatible with
	    ParallelFixture or ParallelTests and disables them.
	
	func (Opt) UnitTests() Option
	    UnitTests is a composite option that signals to Run that the test suite can
	    be treated as a unit-test suite by employing parallelism and fresh fixtures
	    to maximize the chances of exposing unwanted coupling between tests.
	
	type Option func(*config)
	    Option is a function that modifies a config. See Options for provided
	    behaviors.
	
