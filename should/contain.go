package should

import (
	"reflect"
	"strings"

	"github.com/mdwhatcott/testing/compare"
)

// Contain determines whether actual is a member of expected[0].
// The expected value may be a map, array, slice, or string:
//   - In the case of maps the actual value is assumed to be a map key.
//   - In the case of slices and arrays the actual value is assumed to be a member.
//   - In the case of strings the actual value may be a rune or substring.
func Contain(actual interface{}, expected ...interface{}) error {
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
		comparer := compare.New().Compare
		for i := 0; i < actualValue.Len(); i++ {
			item := actualValue.Index(i).Interface()
			if comparer(EXPECTED, item).OK() {
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

const reflectRune = reflect.Int32

var containerKinds = []reflect.Kind{reflect.Map, reflect.Array, reflect.Slice, reflect.String}
