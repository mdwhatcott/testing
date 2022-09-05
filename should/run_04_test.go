package should_test

import (
	"testing"
)

func TestLong(t *testing.T) {
	if !testing.Short() {
		t.Skip("This test only to be run in when -test.short flag passed.")
	}
	fixture := &Suite04{T: New(t)}
	Run(fixture)
	fixture.So(t.Failed(), BeFalse)
}

type Suite04 struct{ *T }

func (this *Suite04) LongTestThatWouldFailButShouldBeSkippedInShortMode() {
	this.So(1, Equal, 2)
}
