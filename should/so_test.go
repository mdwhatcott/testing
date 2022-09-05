package should_test

import (
	"log"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	log.SetFlags(0)
	should.So(t, true, should.BeTrue)
	should.So(nil, true, should.BeFalse)
	should.So(should.Fmt{}, true, should.BeFalse)
	should.So(should.Log{}, nil, should.NOT.BeNil)
}
