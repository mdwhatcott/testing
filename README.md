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
	    	should.Run(&$NAME$Suite{T: t})
	    }
	
	    type $NAME$Suite struct {
	    	*testing.T
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
	
	func BeIn(actual any, expected ...any) error
	    BeIn determines whether actual is a member of expected[0]. It defers to
	    Contain.
	
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
	
	func Equal(actual any, EXPECTED ...any) error
	    Equal verifies that the actual value is equal to the expected value. It uses
	    reflect.DeepEqual in most cases, but also compares numerics regardless of
	    specific type and compares time.Time values using the time.Equal method.
	
	func Panic(actual any, expected ...any) (err error)
	    Panic invokes the func() provided as actual and recovers from any panic.
	    It returns an error if actual() does not result in a panic.
	
	func Run(fixture any)
	    Run accepts a fixture with Test* methods and optional setup/teardown
	    methods and executes the suite. Fixtures must be struct types which embed a
	    *testing.T. Assuming a fixture struct with test methods 'Test1' and 'Test2'
	    execution would proceed as follows:
	
	     1. fixture.Setup()
	     2. fixture.Test1()
	     3. fixture.Teardown()
	     4. fixture.Setup()
	     5. fixture.Test2()
	     6. fixture.Teardown()
	
	func So(t testingT, actual any, assertion Func, expected ...any)
	func WrapError(actual any, expected ...any) error
	    WrapError uses errors.Is to verify that actual is an error value that wraps
	    expected[0] (also an error value).
	
	
	TYPES
	
	type Func func(actual any, expected ...any) error
	

