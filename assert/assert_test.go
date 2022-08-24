package assert_test

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	log.SetFlags(0)
	assert.So(t, true, should.BeTrue)
	assert.So(nil, true, should.BeFalse)
	assert.So(assert.Fmt{}, true, should.BeFalse)
	assert.So(assert.Log{}, nil, should.NOT.BeNil)
}
