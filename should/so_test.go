package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSo(t *testing.T) {
	should.So(t, 1, should.Equal, 1)
}
