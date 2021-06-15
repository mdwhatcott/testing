package compare_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/mdwhatcott/testing/compare"
)

func Test(t *testing.T) {
	runCases(t, []TestCase{
		{
			Expected: 0,
			Actual:   0,
			AreEqual: true,
		},
		{
			Expected: 0,
			Actual:   1,
			AreEqual: false,
		},
		{
			Expected: 0.0,
			Actual:   0.0,
			AreEqual: true,
		},
		{
			Expected: Thing{},
			Actual:   Thing{},
			AreEqual: true,
		},
		{
			Expected: Thing{},
			Actual:   Thing{Integer: 1},
			AreEqual: false,
		},
		{
			Expected: &Thing{Integer: 2},
			Actual:   &Thing{Integer: 2},
			AreEqual: true,
		},
		{
			Expected: []int{1, 2, 3},
			Actual:   []int{1, 2, 3},
			AreEqual: true,
		},
		{
			Expected: [3]int{1, 2, 3},
			Actual:   [3]int{1, 2, 3},
			AreEqual: true,
		},
		{
			Expected: map[int]int{1: 2},
			Actual:   map[int]int{1: 2},
			AreEqual: true,
		},
		{
			Expected: true,
			Actual:   true,
			AreEqual: true,
		},
		{
			Expected: make(chan int),
			Actual:   make(chan int),
			AreEqual: false,
		},
		{
			Expected: "hi",
			Actual:   "hi",
			AreEqual: true,
		},
		{
			Expected: "hi\nbye",
			Actual:   "hi\nbyu",
			AreEqual: false,
		},
		{
			Expected: now.In(notUTC),
			Actual:   now.In(time.UTC),
			AreEqual: true,
		},
		{
			Expected: now.In(notUTC),
			Actual:   now.UTC().Add(time.Nanosecond),
			AreEqual: false,
		},
		{
			Expected: int32(0),
			Actual:   0,
			AreEqual: true,
		},
		{
			Expected: int32(4),
			Actual:   4.0,
			AreEqual: true,
		},
		{
			Expected: uint32(4),
			Actual:   4.0,
			AreEqual: true,
		},
		{
			Expected: complex128(4),
			Actual:   4.0,
			AreEqual: false,
		},
		{
			Expected: int32(0),
			Actual:   1,
			AreEqual: false,
		},
		{
			Expected: (func())(nil),
			Actual:   (func())(nil),
			AreEqual: true,
		},
		{
			Expected: func() {},
			Actual:   (func())(nil),
			AreEqual: false,
		},
		{
			Expected: func() {},
			Actual:   func() {},
			AreEqual: false,
		},
	})
}

func runCases(t *testing.T, cases []TestCase) {
	for x, test := range cases {
		t.Run(test.Title(x), test.Run)
	}
}

var now = time.Now()

var notUTC, _ = time.LoadLocation("America/Los_Angeles")

type Thing struct {
	Integer int
}

type TestCase struct {
	Skip     bool
	Expected interface{}
	Actual   interface{}
	AreEqual bool
}

func (this TestCase) Title(x int) string {
	return fmt.Sprintf(
		"%d.Equal(%+v,%+v)==%t",
		x,
		this.Expected,
		this.Actual,
		this.AreEqual,
	)
}

func (this TestCase) Run(t *testing.T) {
	if this.Skip {
		t.Skip()
	}
	err := compare.Compare(this.Expected, this.Actual)
	if this.AreEqual && err != nil {
		t.Fatal("[FAIL]", err)
	} else if !this.AreEqual && err == nil {
		t.Fatalf("[FAIL] unequal values %v and %v erroneously deemed equal", this.Expected, this.Actual)
	} else {
		t.Log("[PASS] (report printed below for visual inspection)", err)
	}
}
