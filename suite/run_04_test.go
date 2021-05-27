package suite_test

import (
	"testing"

	"github.com/mdwhatcott/testing/assert"
	"github.com/mdwhatcott/testing/suite"
)

func TestLong(t *testing.T) {
	if !testing.Short() {
		t.Skip("This test only to be run in when -test.short flag passed.")
	}
	fixture := &Suite04{T: t}
	suite.Run(fixture)
	assert.With(t).That(t.Failed()).IsFalse()
}

type Suite04 struct{ *testing.T }

func (this *Suite04) LongTestThatWouldFailButShouldBeSkippedInShortMode() {
	assert.With(this).That(1).Equals(2)
}
