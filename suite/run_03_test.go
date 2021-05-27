package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestSkip(t *testing.T) {
	fixture := &Suite03{T: t}
	suite.Run(fixture)
	assert.With(t).That(t.Failed()).IsFalse()
}

type Suite03 struct{ *testing.T }

func (this *Suite03) SkipTestThatFails() {
	assert.With(this).That(1).Equals(2)
}
func (this *Suite03) SkipLongTestThatFails() {
	assert.With(this).That(1).Equals(2)
}
