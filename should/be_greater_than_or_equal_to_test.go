package should_test

import (
	"math"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeGreaterThanOrEqualTo(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeGreaterThanOrEqualTo)
	assert.ExpectedCountInvalid("actual", should.BeGreaterThanOrEqualTo, "expected", "required")
	assert.TypeMismatch(true, should.BeGreaterThanOrEqualTo, 1)
	assert.TypeMismatch(1, should.BeGreaterThanOrEqualTo, true)

	assert.Fail("a", should.BeGreaterThanOrEqualTo, "b") // both strings
	assert.Pass("b", should.BeGreaterThanOrEqualTo, "a")

	assert.Pass(2, should.BeGreaterThanOrEqualTo, 1) // both ints
	assert.Pass(1, should.BeGreaterThanOrEqualTo, 1)
	assert.Fail(1, should.BeGreaterThanOrEqualTo, 2)

	assert.Pass(float32(2.0), should.BeGreaterThanOrEqualTo, float64(1)) // both floats
	assert.Pass(float32(2.0), should.BeGreaterThanOrEqualTo, float64(2))
	assert.Fail(1.0, should.BeGreaterThanOrEqualTo, 2.0)

	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, int64(1)) // both signed
	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, int64(2))
	assert.Fail(int32(1), should.BeGreaterThanOrEqualTo, int64(2))

	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, uint64(1)) // both unsigned
	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, uint64(2))
	assert.Fail(uint32(1), should.BeGreaterThanOrEqualTo, uint64(2))

	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, uint32(1)) // signed and unsigned
	assert.Pass(int32(2), should.BeGreaterThanOrEqualTo, uint32(2))
	assert.Fail(int32(1), should.BeGreaterThanOrEqualTo, uint32(2))
	// if actual < 0: false
	// (because by definition the expected value, an unsigned value must be >= 0)
	const reallyBig uint64 = math.MaxUint64 - 1 // TODO: remove decrement (see issue #5: https://github.com/mdwhatcott/testing/issues/5)
	assert.Fail(-1, should.BeGreaterThanOrEqualTo, reallyBig)

	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, int32(1)) // unsigned and signed
	assert.Pass(uint32(2), should.BeGreaterThanOrEqualTo, int32(2))
	assert.Fail(uint32(1), should.BeGreaterThanOrEqualTo, int32(2))
	// if actual > math.MaxInt64: true
	// (because by definition the expected value, a signed value must be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Pass(tooBig, should.BeGreaterThanOrEqualTo, 42)

	assert.Pass(2.0, should.BeGreaterThanOrEqualTo, 1) // float and integer
	assert.Pass(2.0, should.BeGreaterThanOrEqualTo, 2)
	assert.Fail(1.0, should.BeGreaterThanOrEqualTo, 2)

	assert.Pass(2, should.BeGreaterThanOrEqualTo, 1.0) // integer and float
	assert.Pass(2, should.BeGreaterThanOrEqualTo, 2.0)
	assert.Fail(1, should.BeGreaterThanOrEqualTo, 2.0)
}
