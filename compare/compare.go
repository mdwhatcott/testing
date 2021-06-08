// Package compare facilitates comparisons of any two values according to a set of specifications.
package compare

import (
	"bytes"
	"encoding/json"
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"strings"
	"testing"
	"time"
)

type Comparer interface {
	Compare(a, b interface{}) (result Comparison)
}

type Comparison struct {
	ok     bool
	report string
}

func (this Comparison) OK() bool       { return this.ok }
func (this Comparison) Report() string { return this.report }

type comparer struct {
	config *config
}

func ForTesting(t *testing.T, options ...Option) Comparer {
	return New(append(options, testingT(t))...)
}
func New(options ...Option) Comparer {
	return comparer{config: newConfig(options...)}
}

func (this comparer) Compare(a, b interface{}) (result Comparison) {
	result.ok = this.check(a, b)
	result.report = report(result.OK(), this.resolveFormatter(a), a, b)
	this.config.reportT(result)
	return result
}

func (this comparer) resolveFormatter(a interface{}) Formatter {
	if this.config.formatter != nil {
		return this.config.formatter
	}
	config := new(config)
	defaultFormatterForType(a)(config)
	return config.formatter
}
func (this comparer) check(a, b interface{}) bool {
	for _, spec := range this.config.specs {
		if !spec.IsSatisfiedBy(a, b) {
			continue
		}
		if spec.Compare(a, b) {
			return true
		}
		break
	}
	return false
}

type Option func(*config)

func testingT(t *testing.T) Option {
	return func(this *config) { this.t = t }
}
func With(specs ...Specification) Option {
	return func(this *config) { this.specs = append(this.specs, specs...) }
}
func Format(formatter Formatter) Option {
	return func(this *config) { this.formatter = formatter }
}
func FormatVerb(verb string) Option {
	return Format(func(v interface{}) string { return fmt.Sprintf(verb, v) })
}
func FormatLength() Option {
	return Format(func(v interface{}) string {
		return fmt.Sprintf("Length: %d  Value: %#v", reflect.ValueOf(v).Len(), v)
	})
}
func FormatJSON(indent string) Option {
	return Format(func(v interface{}) string {
		raw, err := json.Marshal(v)
		if err != nil {
			return err.Error()
		}
		if indent == "" {
			return string(raw)
		}
		indented := new(bytes.Buffer)
		_ = json.Indent(indented, raw, "", indent)
		return indented.String()
	})
}

type config struct {
	t         *testing.T
	specs     []Specification
	formatter Formatter
}

func newConfig(options ...Option) *config {
	this := new(config)
	for _, option := range options {
		option(this)
	}
	if len(this.specs) == 0 {
		this.specs = []Specification{NumericEquality{}, TimeEquality{}, DeepEquality{}}
	}
	return this
}
func (this *config) reportT(result Comparison) {
	if result.OK() {
		return
	}
	if this.t == nil {
		return
	}
	this.t.Error(result.Report())
}

type Specification interface {
	IsSatisfiedBy(a, b interface{}) bool
	Compare(a, b interface{}) bool
}

// DeepEquality compares any two values using reflect.DeepEqual.
// https://golang.org/pkg/reflect/#DeepEqual
type DeepEquality struct{}

func (this DeepEquality) IsSatisfiedBy(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (this DeepEquality) Compare(a, b interface{}) bool {
	return reflect.DeepEqual(a, b)
}

// SimpleEquality compares any two values using the built-in equality operator (`==`).
// https://golang.org/ref/spec#Comparison_operators
type SimpleEquality struct{}

func (this SimpleEquality) IsSatisfiedBy(a, b interface{}) bool {
	return reflect.TypeOf(a) == reflect.TypeOf(b)
}
func (this SimpleEquality) Compare(a, b interface{}) bool {
	return a == b
}

// NumericEquality compares numeric values using the built-in equality
// operator (`==`). Values of differing numeric reflect.Kind are each
// converted to the type of the other and are compared with `==` in both
// directions. https://golang.org/pkg/reflect/#Kind
type NumericEquality struct{}

func (this NumericEquality) IsSatisfiedBy(a, b interface{}) bool {
	return isNumeric(a) && isNumeric(b)
}
func (this NumericEquality) Compare(a, b interface{}) bool {
	if a == b {
		return true
	}
	aValue := reflect.ValueOf(a)
	bValue := reflect.ValueOf(b)
	aAsB := aValue.Convert(bValue.Type()).Interface()
	bAsA := bValue.Convert(aValue.Type()).Interface()
	return a == bAsA && b == aAsB
}
func isNumeric(v interface{}) bool {
	kind := reflect.TypeOf(v).Kind()
	return kind == reflect.Int ||
		kind == reflect.Int8 ||
		kind == reflect.Int16 ||
		kind == reflect.Int32 ||
		kind == reflect.Int64 ||
		kind == reflect.Uint ||
		kind == reflect.Uint8 ||
		kind == reflect.Uint16 ||
		kind == reflect.Uint32 ||
		kind == reflect.Uint64 ||
		kind == reflect.Float32 ||
		kind == reflect.Float64
}

// TimeEquality compares values both of type time.Time using their Equal method.
// https://golang.org/pkg/time/#Time.Equal
type TimeEquality struct{}

func (this TimeEquality) IsSatisfiedBy(a, b interface{}) bool {
	return isTime(a) && isTime(b)
}
func (this TimeEquality) Compare(a, b interface{}) bool {
	return a.(time.Time).Equal(b.(time.Time))
}

func isTime(v interface{}) bool {
	_, ok := v.(time.Time)
	return ok
}

// LengthEquality compares values that can serve as valid arguments to the built-in len function
// (with the exception of pointers to arrays, which are not yet supported herein).
// https://golang.org/pkg/builtin/#len
type LengthEquality struct{}

func (this LengthEquality) IsSatisfiedBy(a, b interface{}) bool {
	return hasLen(a) && hasLen(b)
}
func (this LengthEquality) Compare(a, b interface{}) bool {
	return reflect.ValueOf(a).Len() == reflect.ValueOf(b).Len()
}
func hasLen(v interface{}) bool {
	switch reflect.TypeOf(v).Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.Chan, reflect.String:
		return true
	default:
		return false
	}
}

type Formatter func(interface{}) string

func report(equal bool, format Formatter, a, b interface{}) string {
	if equal {
		return fmt.Sprintf("%s == %s", format(a), format(b))
	}
	aType := fmt.Sprintf("(%v)", reflect.TypeOf(a))
	bType := fmt.Sprintf("(%v)", reflect.TypeOf(b))
	longestType := int(math.Max(float64(len(aType)), float64(len(bType))))
	aType += strings.Repeat(" ", longestType-len(aType))
	bType += strings.Repeat(" ", longestType-len(bType))
	aFormat := format(a)
	bFormat := format(b)
	typeDiff := diff(bType, aType)
	valueDiff := diff(bFormat, aFormat)

	builder := new(strings.Builder)
	_, _ = fmt.Fprintf(builder, "\n")
	_, _ = fmt.Fprintf(builder, "A: %s %s\n", aType, aFormat)
	_, _ = fmt.Fprintf(builder, "B: %s %s\n", bType, bFormat)
	_, _ = fmt.Fprintf(builder, "   %s %s\n", typeDiff, valueDiff)
	_, _ = fmt.Fprintf(builder, "Stack (filtered):\n%s\n", stack())

	return builder.String()
}

func defaultFormatterForType(v interface{}) Option {
	switch {
	case isNumeric(v) || isTime(v):
		return FormatVerb("%v")
	default:
		return FormatVerb("%#v")
	}
}

func diff(a, b string) string {
	if strings.Contains(a, "\n") || strings.Contains(b, "\n") {
		return ""
	}
	result := new(strings.Builder)

	for x := 0; ; x++ {
		if x >= len(a) && x >= len(b) {
			break
		}
		if x >= len(a) || x >= len(b) || a[x] != b[x] {
			result.WriteString("^")
		} else {
			result.WriteString(" ")
		}
	}
	return result.String()
}

func stack() string {
	lines := strings.Split(string(debug.Stack()), "\n")
	var filtered []string
	for x := 1; x < len(lines)-1; x += 2 {
		if strings.Contains(lines[x+1], "_test.go:") {
			filtered = append(filtered, lines[x], lines[x+1])
		}
	}
	return "> " + strings.Join(filtered, "\n> ")
}
