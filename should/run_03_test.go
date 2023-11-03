package should_test

import (
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: t}
	should.Run(fixture)
	should.So(t, t.Failed(), should.BeFalse)
}

type Suite03 struct{ *testing.T }

func (this *Suite03) SkipTestThatFails() {
	should.So(this.T, 1, should.Equal, 2)
}
