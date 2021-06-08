package assert_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func TestSo(t *testing.T) {
	assertNil(t, assert.So(1, should.Equal, 1))
	assertErr(t, assert.So(1, should.Equal, 2))
}
