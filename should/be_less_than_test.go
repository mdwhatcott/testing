package should_test

import (
	"math"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestShouldBeLessThan(t *testing.T) {
	assert := NewAssertion(t)

	assert.ExpectedCountInvalid("actual-but-missing-expected", should.BeLessThan)
	assert.ExpectedCountInvalid("actual", should.BeLessThan, "expected", "required")
	assert.TypeMismatch(true, should.BeLessThan, 1)
	assert.TypeMismatch(1, should.BeLessThan, true)

	assert.Fail("b", should.BeLessThan, "a") // both strings
	assert.Pass("a", should.BeLessThan, "b")

	assert.Fail(1, should.BeLessThan, 1) // both ints
	assert.Pass(1, should.BeLessThan, 2)

	assert.Pass(float32(1.0), should.BeLessThan, float64(2)) // both floats
	assert.Fail(2.0, should.BeLessThan, 1.0)

	assert.Pass(int32(1), should.BeLessThan, int64(2)) // both signed
	assert.Fail(int32(2), should.BeLessThan, int64(1))

	assert.Pass(uint32(1), should.BeLessThan, uint64(2)) // both unsigned
	assert.Fail(uint32(2), should.BeLessThan, uint64(1))

	assert.Pass(int32(1), should.BeLessThan, uint32(2)) // signed and unsigned
	assert.Fail(int32(2), should.BeLessThan, uint32(1))
	// if actual < 0: true
	// (because by definition the expected value, an unsigned value can't be < 0)
	const reallyBig uint64 = math.MaxUint64
	assert.Pass(-1, should.BeLessThan, reallyBig)

	assert.Pass(uint32(1), should.BeLessThan, int32(2)) // unsigned and signed
	assert.Fail(uint32(2), should.BeLessThan, int32(1))
	// if actual > math.MaxInt64: false
	// (because by definition the expected value, a signed value can't be > math.MaxInt64)
	const tooBig uint64 = math.MaxInt64 + 1
	assert.Fail(tooBig, should.BeLessThan, 42)

	assert.Pass(1.0, should.BeLessThan, 2) // float and integer
	assert.Fail(2.0, should.BeLessThan, 1)

	assert.Pass(1, should.BeLessThan, 2.0) // integer and float
	assert.Fail(2, should.BeLessThan, 1.0)
}
