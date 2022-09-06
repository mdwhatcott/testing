# github.com/mdwhatcott/testing



	package should // import "github.com/mdwhatcott/testing/should"
	
	Package should
	
	This package strives to make it easy, even fun, for software
	developers to produce > a quick, sure, and repeatable proof that
	every element of the code works as it should. (See [The Programmer's
	Oath](http://blog.cleancoder.com/uncle-bob/2015/11/18/TheProgrammersOath.html))
	
	The simplest way is by combining the So function with the many provided
	assertions, such as should.Equal:
	
	    package whatever
	
	    import (
	    	"log"
	    	"testing"
	
	    	"github.com/mdwhatcott/testing/should"
	    )
	
	    func Test(t *testing.T) {
	    	should.So(t, 1, should.Equal, 1)
	    }
	
	This package also implement an xUnit-style test runner, which is based on the
	following packages:
	
	  - github.com/stretchr/testify/suite(https://pkg.go.dev/github.com/stretchr/testify/suite)
	  - github.com/smartystreets/gunit(https://pkg.go.dev/github.com/smartystreets/gunit)
	
	For those using an IDE by JetBrains, you may find the following "live template"
	helpful:
	
	    func Test$NAME$Suite(t *testing.T) {
	    	suite.Run(&$NAME$Suite{T: suite.New(t)}, suite.Options.UnitTests())
	    }
	
	    type $NAME$Suite struct {
	    	*suite.T
	    }
	
	    func (this *$NAME$Suite) Setup() {
	    }
	
	    func (this *$NAME$Suite) Test$END$() {
	    }
	
	From a test method like the one in the template above, simply use the embedded
	So method:
	
	    func (this TheSuite) TestSomething() {
	    	this.So(1, should.Equal, 1)
	    }
	
	Happy testing!
	
	VARIABLES
	
	var (
		ErrExpectedCountInvalid = errors.New("expected count invalid")
		ErrTypeMismatch         = errors.New("type mismatch")
		ErrKindMismatch         = errors.New("kind mismatch")
		ErrAssertionFailure     = errors.New("assertion failure")
	)
	var NOT negated
	    NOT (a singleton) constrains all negated assertions to their own namespace.
	
	
	FUNCTIONS
	
	func BeEmpty(actual any, expected ...any) error
	    BeEmpty uses reflection to verify that len(actual) == 0.
	
	func BeFalse(actual any, expected ...any) error
	    BeFalse verifies that actual is the boolean false value.
	
	func BeGreaterThan(actual any, EXPECTED ...any) error
	    BeGreaterThan verifies that actual is greater than expected. Both actual and
	    expected must be strings or numeric in type.
	
	func BeGreaterThanOrEqualTo(actual any, expected ...any) error
	    BeGreaterThanOrEqualTo verifies that actual is less than or equal to
	    expected. Both actual and expected must be strings or numeric in type.
	
	func BeIn(actual any, expected ...any) error
	    BeIn determines whether actual is a member of expected[0]. It defers to
	    Contain.
	
	func BeLessThan(actual any, EXPECTED ...any) error
	    BeLessThan verifies that actual is less than expected. Both actual and
	    expected must be strings or numeric in type.
	
	func BeLessThanOrEqualTo(actual any, expected ...any) error
	    BeLessThanOrEqualTo verifies that actual is less than or equal to expected.
	    Both actual and expected must be strings or numeric in type.
	
	func BeNil(actual any, expected ...any) error
	    BeNil verifies that actual is the nil value.
	
	func BeTrue(actual any, expected ...any) error
	    BeTrue verifies that actual is the boolean true value.
	
	func Contain(actual any, expected ...any) error
	    Contain determines whether actual contains expected[0]. The actual value may
	    be a map, array, slice, or string:
	      - In the case of maps the expected value is assumed to be a map key.
	      - In the case of slices and arrays the expected value is assumed to be a
	        member.
	      - In the case of strings the expected value may be a rune or substring.
	
	func EndWith(actual any, expected ...any) error
	    EndWith verifies that actual ends with expected[0]. The actual value may be
	    an array, slice, or string.
	
	func Equal(actual any, EXPECTED ...any) error
	    Equal verifies that the actual value is equal to the expected value. It uses
	    reflect.DeepEqual in most cases, but also compares numerics regardless of
	    specific type and compares time.Time values using the time.Equal method.
	
	func HaveLength(actual any, expected ...any) error
	    HaveLength uses reflection to verify that len(actual) == 0.
	
	func Panic(actual any, expected ...any) (err error)
	    Panic invokes the func() provided as actual and recovers from any panic.
	    It returns an error if actual() does not result in a panic.
	
	func Run(fixture any, options ...Option)
	    Run accepts a fixture with Test* methods and optional setup/teardown
	    methods and executes the suite. Fixtures must be struct types which embed a
	    *suite.T. Assuming a fixture struct with test methods 'Test1' and 'Test2'
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
	
	func So(t t, actual any, assertion Func, expected ...any)
	func StartWith(actual any, expected ...any) error
	    StartWith verified that actual starts with expected[0]. The actual value may
	    be an array, slice, or string.
	
	func WrapError(actual any, expected ...any) error
	    WrapError uses errors.Is to verify that actual is an error value that wraps
	    expected[0] (also an error value).
	
	
	TYPES
	
	type Fmt struct{}
	
	func (Fmt) Error(a ...any)
	
	func (Fmt) Helper()
	
	type Func func(actual any, expected ...any) error
	
	type Log struct{}
	
	func (Log) Error(a ...any)
	
	func (Log) Helper()
	
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
	    SharedFixture signals to Run that the provided fixture instance is
	    to be used to run all test methods. This mode is not compatible with
	    ParallelFixture or ParallelTests and disables them.
	
	func (Opt) UnitTests() Option
	    UnitTests is a composite option that signals to Run that the test suite can
	    be treated as a unit-test suite by employing parallelism and fresh fixtures
	    to maximize the chances of exposing unwanted coupling between tests.
	
	type Option func(*config)
	    Option is a function that modifies a config. See Options for provided
	    behaviors.
	
	type T struct{ *testing.T }
	    T embeds *testing.T and provides convenient hooks for making assertions and
	    other operations.
	
	func New(t *testing.T) *T
	    New prepares a *T for use with the fixture passed to Run.
	
	func (this *T) FatalSo(actual any, assertion assertion, expected ...any) bool
	    FatalSo is like So but in the event of an assertion failure it calls
	    *testing.T.Fatal.
	
	func (this *T) So(actual any, assertion assertion, expected ...any) bool
	    So invokes the provided assertion with the provided args. In the event of an
	    assertion failure it calls *testing.T.Error.
	
	func (this *T) Write(p []byte) (n int, err error)
	    Write implements io.Writer allowing for the suite to serve as a convenient
	    log target, among other use cases.
	

