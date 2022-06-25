package should

import (
	"math"
	"reflect"
)

// BeLessThan verifies that actual is greater than expected. Both actual and expected must be numeric in type.
func BeLessThan(actual any, EXPECTED ...any) error {
	err := validateExpected(1, EXPECTED)
	if err != nil {
		return err
	}

	expected := EXPECTED[0]
	failed := false

	for _, spec := range lessThanSpecs {
		if !spec.assertable(actual, expected) {
			continue
		}
		if spec.passes(actual, expected) {
			return nil
		}
		failed = true
		break
	}

	if failed {
		return failure("%v was not greater than %v", actual, expected)
	}
	return wrap(ErrTypeMismatch, "could not compare [%v] and [%v]",
		reflect.TypeOf(actual), reflect.TypeOf(expected))
}

var lessThanSpecs = []specification{
	bothStringsLessThan{},
	bothSignedIntegersLessThan{},
	bothUnsignedIntegersLessThan{},
	bothFloatsLessThan{},
	signedAndUnsignedLessThan{},
	unsignedAndSignedLessThan{},
	floatAndIntegerLessThan{},
	integerAndFloatLessThan{},
}

type bothStringsLessThan struct{}

func (bothStringsLessThan) assertable(a, b any) bool {
	return reflect.ValueOf(a).Kind() == reflect.String && reflect.ValueOf(b).Kind() == reflect.String
}
func (bothStringsLessThan) passes(a, b any) bool {
	return reflect.ValueOf(a).String() < reflect.ValueOf(b).String()
}

type bothSignedIntegersLessThan struct{}

func (bothSignedIntegersLessThan) assertable(a, b any) bool {
	return isSignedInteger(a) && isSignedInteger(b)
}
func (bothSignedIntegersLessThan) passes(a, b any) bool {
	return reflect.ValueOf(a).Int() < reflect.ValueOf(b).Int()
}

type bothUnsignedIntegersLessThan struct{}

func (bothUnsignedIntegersLessThan) assertable(a, b any) bool {
	return isUnsignedInteger(a) && isUnsignedInteger(b)
}
func (bothUnsignedIntegersLessThan) passes(a, b any) bool {
	return reflect.ValueOf(a).Uint() < reflect.ValueOf(b).Uint()
}

type bothFloatsLessThan struct{}

func (bothFloatsLessThan) assertable(a, b any) bool {
	return isFloat(a) && isFloat(b)
}
func (bothFloatsLessThan) passes(a, b any) bool {
	return reflect.ValueOf(a).Float() < reflect.ValueOf(b).Float()
}

type signedAndUnsignedLessThan struct{}

func (signedAndUnsignedLessThan) assertable(a, b any) bool {
	return isSignedInteger(a) && isUnsignedInteger(b)
}
func (signedAndUnsignedLessThan) passes(a, b any) bool {
	A := reflect.ValueOf(a)
	B := reflect.ValueOf(b)
	if A.Int() < 0 {
		return true
	}
	return uint64(A.Int()) < B.Uint()
}

type unsignedAndSignedLessThan struct{}

func (unsignedAndSignedLessThan) assertable(a, b any) bool {
	return isUnsignedInteger(a) && isSignedInteger(b)
}
func (unsignedAndSignedLessThan) passes(a, b any) bool {
	A := reflect.ValueOf(a)
	B := reflect.ValueOf(b)
	if A.Uint() > math.MaxInt64 {
		return false
	}
	return int64(A.Uint()) < B.Int()
}

type floatAndIntegerLessThan struct{}

func (floatAndIntegerLessThan) assertable(a, b any) bool {
	return isFloat(a) && isInteger(b)
}
func (floatAndIntegerLessThan) passes(a, b any) bool {
	return asFloat(a) < asFloat(b)
}

type integerAndFloatLessThan struct{}

func (integerAndFloatLessThan) assertable(a, b any) bool {
	return isInteger(a) && isFloat(b)
}
func (integerAndFloatLessThan) passes(a, b any) bool {
	return asFloat(a) < asFloat(b)
}
