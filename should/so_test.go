package should_test

import (
	"bytes"
	"log"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func Test(t *testing.T) {
	should.So(t, true, should.BeTrue)
	buffer := bytes.Buffer{}
	should.So(log.New(&buffer, "", 0), true, should.BeFalse)
	should.So(t, buffer.Len(), should.BeGreaterThan, 0)
}
