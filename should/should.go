/*
Package should

This package strives to make it easy, even fun, for
software developers to produce
> a quick, sure, and repeatable proof that every element of the code works as it should.
(See [The Programmer's Oath](http://blog.cleancoder.com/uncle-bob/2015/11/18/TheProgrammersOath.html))

The simplest way is by combining the So function with the many provided assertions, such as should.Equal:

	package whatever

	import (
		"log"
		"testing"

		"github.com/mdwhatcott/testing/should"
	)

	func Test(t *testing.T) {
		should.So(t, 1, should.Equal, 1)
	}

This package also implement an xUnit-style test
runner, which is based on the following packages:

  - [github.com/stretchr/testify/suite](https://pkg.go.dev/github.com/stretchr/testify/suite)
  - [github.com/smartystreets/gunit](https://pkg.go.dev/github.com/smartystreets/gunit)

For those using an IDE by JetBrains, you may
find the following "live template" helpful:

	func Test$NAME$Suite(t *testing.T) {
		should.Run(&$NAME$Suite{T: should.New(t)})
	}

	type $NAME$Suite struct {
		*should.T
	}

	func (this *$NAME$Suite) Setup() {
	}

	func (this *$NAME$Suite) Test$END$() {
	}

From a test method like the one in the template above, simply use the embedded So method:

	func (this TheSuite) TestSomething() {
		this.So(1, should.Equal, 1)
	}

Happy testing!
*/
package should

import (
	"errors"
	"fmt"
	"math"
	"os"
	"reflect"
	"runtime/debug"
	"strconv"
	"strings"
	"testing"
	"time"
)

// T is wrapper over *testing.T.
type T struct{ *testing.T }

func New(t *testing.T) *T { return &T{T: t} }

func (this *T) So(actual any, assertion Func, expected ...any) {
	this.Helper()
	So(this, actual, assertion, expected...)
}

// Run accepts a fixture with Test* methods and
// optional setup/teardown methods and executes
// the suite. Fixtures must be struct types which
// embed a *should.T. Assuming a fixture struct
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
	t := fixtureValue.Elem().FieldByName("T").Interface().(*T)

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
	*T
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
	fixtureValue.Elem().FieldByName("T").Set(reflect.ValueOf(New(t)))

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

type Func func(actual any, expected ...any) error

type testingT interface {
	Helper()
	Error(...any)
}

// So is the basic assertion mechanism in the general form:
// should.So(t, 1, should.Equal, 2)
func So(t testingT, actual any, assertion Func, expected ...any) {
	t.Helper()
	if err := assertion(actual, expected...); err != nil {
		t.Error(err)
	}
}

// NOT (a singleton) constrains all negated assertions to their own namespace.
var NOT negated

type negated struct{}

type specification interface {
	assertable(a, b any) bool
	passes(a, b any) bool
}

var (
	ErrExpectedCountInvalid = errors.New("expected count invalid")
	ErrTypeMismatch         = errors.New("type mismatch")
	ErrKindMismatch         = errors.New("kind mismatch")
	ErrAssertionFailure     = errors.New("assertion failure")
)

func failure(format string, args ...any) error {
	trace := stack()
	if len(trace) > 0 {
		format += "\nStack: (filtered)\n%s"
		args = append(args, trace)
	}
	return wrap(ErrAssertionFailure, format, args...)
}
func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		fileLineRaw := lines[x+1]
		if strings.Contains(fileLineRaw, "_test.go:") {
			filtered = append(filtered, lines[x], fileLineRaw)
			line, ok := readSourceCodeLine(fileLineRaw)
			if ok {
				filtered = append(filtered, "  "+line)
			}

		}
	}
	return "> " + strings.Join(filtered, "\n> ")
}
func readSourceCodeLine(fileLineRaw string) (string, bool) {
	fileLineJoined := strings.Fields(strings.TrimSpace(fileLineRaw))[0]
	fileLine := strings.Split(fileLineJoined, ":")
	sourceCode, _ := os.ReadFile(fileLine[0])
	sourceCodeLines := strings.Split(string(sourceCode), "\n")
	lineNumber, _ := strconv.Atoi(fileLine[1])
	lineNumber--
	lineNumber = max(len(sourceCodeLines)-1, lineNumber)
	return sourceCodeLines[lineNumber], true
}
func wrap(inner error, format string, args ...any) error {
	return fmt.Errorf("%w: "+fmt.Sprintf(format, args...), inner)
}

func validateExpected(count int, expected []any) error {
	length := len(expected)
	if length == count {
		return nil
	}

	s := pluralize(length)
	return wrap(ErrExpectedCountInvalid, "got %d value%s, want %d", length, s, count)
}

func pluralize(count int) string {
	if count == 1 {
		return ""
	}
	return "s"
}

func validateType(actual, expected any) error {
	ACTUAL := reflect.TypeOf(actual)
	EXPECTED := reflect.TypeOf(expected)
	if ACTUAL == EXPECTED {
		return nil
	}
	return wrap(ErrTypeMismatch, "got %s, want %s", ACTUAL, EXPECTED)
}

func validateKind(actual any, kinds ...reflect.Kind) error {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	for _, k := range kinds {
		if k == kind {
			return nil
		}
	}
	return wrap(ErrKindMismatch, "got %s, want one of %v", kind, kinds)
}

var unsignedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Uintptr: {},
}

func isUnsignedInteger(v any) bool {
	_, found := unsignedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

var signedIntegerKinds = map[reflect.Kind]struct{}{
	reflect.Int:   {},
	reflect.Int8:  {},
	reflect.Int16: {},
	reflect.Int32: {},
	reflect.Int64: {},
}

func isSignedInteger(v any) bool {
	_, found := signedIntegerKinds[reflect.TypeOf(v).Kind()]
	return found
}

var numericKinds = map[reflect.Kind]struct{}{
	reflect.Int:     {},
	reflect.Int8:    {},
	reflect.Int16:   {},
	reflect.Int32:   {},
	reflect.Int64:   {},
	reflect.Uint:    {},
	reflect.Uint8:   {},
	reflect.Uint16:  {},
	reflect.Uint32:  {},
	reflect.Uint64:  {},
	reflect.Float32: {},
	reflect.Float64: {},
}

func isNumeric(v any) bool {
	of := reflect.TypeOf(v)
	if of == nil {
		return false
	}
	_, found := numericKinds[of.Kind()]
	return found
}

var kindsWithLength = []reflect.Kind{
	reflect.Map,
	reflect.Chan,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

var containerKinds = []reflect.Kind{
	reflect.Map,
	reflect.Array,
	reflect.Slice,
	reflect.String,
}

// BeEmpty uses reflection to verify that len(actual) == 0.
func BeEmpty(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, kindsWithLength...)
	if err != nil {
		return err
	}

	length := reflect.ValueOf(actual).Len()
	if length == 0 {
		return nil
	}

	TYPE := reflect.TypeOf(actual).String()
	return failure("got len(%s) == %d, want empty %s", TYPE, length, TYPE)
}

// BeEmpty (negated!)
func (negated) BeEmpty(actual any, expected ...any) error {
	err := BeEmpty(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}
	if err != nil {
		return err
	}
	TYPE := reflect.TypeOf(actual).String()
	return failure("got empty %s, want non-empty %s", TYPE, TYPE)
}

// BeFalse verifies that actual is the boolean false value.
func BeFalse(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if boolean {
		return failure("got <true>, want <false>")
	}

	return nil
}

// BeIn determines whether actual is a member of expected[0].
// It defers to Contain.
func BeIn(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = Contain(expected[0], actual)
	if err != nil {
		return err
	}

	return nil
}

// BeIn (negated!)
func (negated) BeIn(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	return NOT.Contain(expected[0], actual)
}

// BeNil verifies that actual is the nil value.
func BeNil(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	if actual == nil || interfaceHasNilValue(actual) {
		return nil
	}

	return failure("got %#v, want <nil>", actual)
}
func interfaceHasNilValue(actual any) bool {
	value := reflect.ValueOf(actual)
	kind := value.Kind()
	nillable := kind == reflect.Slice ||
		kind == reflect.Chan ||
		kind == reflect.Func ||
		kind == reflect.Ptr ||
		kind == reflect.Map

	// Careful: reflect.Value.IsNil() will panic unless it's
	// an interface, chan, map, func, slice, or ptr
	// Reference: http://golang.org/pkg/reflect/#Value.IsNil
	return nillable && value.IsNil()
}

// BeNil negated!
func (negated) BeNil(actual any, expected ...any) error {
	err := BeNil(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("got nil, want non-<nil>")
}

// BeTrue verifies that actual is the boolean true value.
func BeTrue(actual any, expected ...any) error {
	err := validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, *new(bool))
	if err != nil {
		return err
	}

	boolean := actual.(bool)
	if !boolean {
		return failure("got <false>, want <true>")
	}
	return nil
}

// Contain determines whether actual contains expected[0].
// The actual value may be a map, array, slice, or string:
//   - In the case of maps the expected value is assumed to be a map key.
//   - In the case of slices and arrays the expected value is assumed to be a member.
//   - In the case of strings the expected value may be a rune or substring.
func Contain(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	err = validateKind(actual, containerKinds...)
	if err != nil {
		return err
	}

	actualValue := reflect.ValueOf(actual)
	EXPECTED := expected[0]

	switch reflect.TypeOf(actual).Kind() {
	case reflect.Map:
		expectedValue := reflect.ValueOf(EXPECTED)
		value := actualValue.MapIndex(expectedValue)
		if value.IsValid() {
			return nil
		}
	case reflect.Array, reflect.Slice:
		for i := 0; i < actualValue.Len(); i++ {
			item := actualValue.Index(i).Interface()
			if Equal(EXPECTED, item) == nil {
				return nil
			}
		}
	case reflect.String:
		err = validateKind(EXPECTED, reflect.String, reflectRune)
		if err != nil {
			return err
		}

		expectedRune, ok := EXPECTED.(rune)
		if ok {
			EXPECTED = string(expectedRune)
		}

		full := actual.(string)
		sub := EXPECTED.(string)
		if strings.Contains(full, sub) {
			return nil
		}
	}

	return failure("\n"+
		"   item absent: %#v\n"+
		"   within:      %#v",
		EXPECTED,
		actual,
	)
}

// Contain (negated!)
func (negated) Contain(actual any, expected ...any) error {
	err := Contain(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"item found: %#v\n"+
		"within:     %#v",
		expected[0],
		actual,
	)
}

const reflectRune = reflect.Int32

// Equal verifies that the actual value is equal to the expected value.
// It uses reflect.DeepEqual in most cases, but also compares numerics
// regardless of specific type and compares time.Time values using the
// time.Equal method.
func Equal(actual any, EXPECTED ...any) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}

	expected := EXPECTED[0]

	for _, spec := range equalitySpecs {
		if !spec.assertable(actual, expected) {
			continue
		}
		if spec.passes(actual, expected) {
			return nil
		}
		break
	}
	return failure(report(actual, expected))
}

// Equal negated!
func (negated) Equal(actual any, expected ...any) error {
	err := Equal(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("\n"+
		"  expected:     %#v\n"+
		"  to not equal: %#v\n"+
		"  (but it did)",
		expected[0],
		actual,
	)
}

var equalitySpecs = []specification{
	numericEquality{},
	timeEquality{},
	deepEquality{},
}

func report(a, b any) string {
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := fmt.Sprintf(format(a), a)
	bFormat := fmt.Sprintf(format(b), b)
	typeDiff := diff(bType, aType)
	valueDiff := diff(bFormat, aFormat)

	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "Expected: %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "Actual:   %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "          %s %s", typeDiff, valueDiff)

	if firstDiffIndex := strings.Index(valueDiff, "^"); firstDiffIndex > 40 {
		start := firstDiffIndex - 20
		_, _ = fmt.Fprintf(builder, "\nInitial discrepancy at index %d:\n", firstDiffIndex)
		_, _ = fmt.Fprintf(builder, "... %s\n", bFormat[start:])
		_, _ = fmt.Fprintf(builder, "... %s\n", aFormat[start:])
		_, _ = fmt.Fprintf(builder, "    %s", valueDiff[start:])
	}

	return builder.String()
}
func format(v any) string {
	if isNumeric(v) || isTime(v) {
		return "%v"
	} else {
		return "%#v"
	}
}
func diff(a, b string) string {
	result := new(strings.Builder)
	for x := 0; x < len(a) && x < len(b); x++ {
		if x >= len(a) || x >= len(b) || a[x] != b[x] {
			result.WriteString("^")
		} else {
			result.WriteString(" ")
		}
	}
	return result.String()
}

// deepEquality compares any two values using reflect.DeepEqual.
// https://golang.org/pkg/reflect/#DeepEqual
type deepEquality struct{}

func (deepEquality) assertable(a, b any) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (deepEquality) passes(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// numericEquality compares numeric values using the built-in equality
// operator (`==`). Values of differing numeric reflect.Kind are each
// converted to the type of the other and are compared with `==` in both
// directions, with one exception: two mixed integers (one signed and one
// unsigned) are always unequal in the case that the unsigned value is
// greater than math.MaxInt64. https://golang.org/pkg/reflect/#Kind
type numericEquality struct{}

func (numericEquality) assertable(a, b any) bool {
	return isNumeric(a) && isNumeric(b)
}
func (numericEquality) passes(a, b any) bool {
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	if isUnsignedInteger(a) && isSignedInteger(b) && aValue.Uint() >= math.MaxInt64 {
		return false
	}
	if isSignedInteger(a) && isUnsignedInteger(b) && bValue.Uint() >= math.MaxInt64 {
		return false
	}
	aAsB := aValue.Convert(bValue.Type()).Interface()
	bAsA := bValue.Convert(aValue.Type()).Interface()
	return a == bAsA && b == aAsB
}

// timeEquality compares values both of type time.Time using their Equal method.
// https://golang.org/pkg/time/#Time.Equal
type timeEquality struct{}

func (timeEquality) assertable(a, b any) bool {
	return isTime(a) && isTime(b)
}
func (timeEquality) passes(a, b any) bool {
	return a.(time.Time).Equal(b.(time.Time))
}
func isTime(v any) bool {
	_, ok := v.(time.Time)
	return ok
}

// Panic invokes the func() provided as actual and recovers from any
// panic. It returns an error if actual() does not result in a panic.
func Panic(actual any, expected ...any) (err error) {
	err = NOT.Panic(actual, expected...)
	if errors.Is(err, ErrAssertionFailure) {
		return nil
	}

	if err != nil {
		return err
	}

	return failure("provided func did not panic as expected")
}

// Panic (negated!) expects the func() provided as actual to run without panicking.
func (negated) Panic(actual any, expected ...any) (err error) {
	err = validateExpected(0, expected)
	if err != nil {
		return err
	}

	err = validateType(actual, func() {})
	if err != nil {
		return err
	}

	panicked := true
	defer func() {
		r := recover()
		if panicked {
			err = failure(""+
				"provided func should not have"+
				"panicked but it did with: %s", r,
			)
		}
	}()

	actual.(func())()
	panicked = false
	return nil
}

// WrapError uses errors.Is to verify that actual is an error value
// that wraps expected[0] (also an error value).
func WrapError(actual any, expected ...any) error {
	err := validateExpected(1, expected)
	if err != nil {
		return err
	}

	inner, ok := expected[0].(error)
	if !ok {
		return errTypeMismatch(expected[0])
	}

	outer, ok := actual.(error)
	if !ok {
		return errTypeMismatch(actual)
	}

	if errors.Is(outer, inner) {
		return nil
	}

	return fmt.Errorf("%w:\n"+
		"\t            outer err: (%s)\n"+
		"\tshould wrap inner err: (%s)",
		ErrAssertionFailure,
		outer,
		inner,
	)
}

func errTypeMismatch(v any) error {
	return fmt.Errorf(
		"%w: got %s, want error",
		ErrTypeMismatch,
		reflect.TypeOf(v),
	)
}
