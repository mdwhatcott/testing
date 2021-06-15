// Package compare facilitates comparisons of any two values according to a set of specifications.
package compare

import (
	"fmt"
	"math"
	"reflect"
	"runtime/debug"
	"strings"
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

func New() Comparer {
	return comparer{config: newConfig()}
}

func (this comparer) Compare(a, b interface{}) (result Comparison) {
	result.ok = this.check(a, b)
	result.report = report(result.OK(), this.resolveFormatter(a), a, b)
	return result
}

func (this comparer) resolveFormatter(v interface{}) formatter {
	switch {
	case isNumeric(v) || isTime(v):
		return FormatVerb("%v")
	default:
		return FormatVerb("%#v")
	}
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

func FormatVerb(verb string) formatter {
	return func(v interface{}) string { return fmt.Sprintf(verb, v) }
}

type config struct {
	specs     []Specification
	formatter formatter
}

func newConfig() *config {
	this := new(config)
	if len(this.specs) == 0 {
		this.specs = []Specification{
			NumericEquality{},
			TimeEquality{},
			DeepEquality{},
		}
	}
	return this
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

type formatter func(interface{}) string

func report(equal bool, format formatter, a, b interface{}) string {
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

func diff(a, b string) string {
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
