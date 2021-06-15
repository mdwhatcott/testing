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
			Skip:     false,
			Expected: 0,
			Actual:   0,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: 0,
			Actual:   1,
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: 0.0,
			Actual:   0.0,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: Thing{},
			Actual:   Thing{},
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: Thing{},
			Actual:   Thing{Integer: 1},
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: &Thing{Integer: 2},
			Actual:   &Thing{Integer: 2},
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: []int{1, 2, 3},
			Actual:   []int{1, 2, 3},
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: [3]int{1, 2, 3},
			Actual:   [3]int{1, 2, 3},
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: map[int]int{1: 2},
			Actual:   map[int]int{1: 2},
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: true,
			Actual:   true,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: make(chan int),
			Actual:   make(chan int),
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: "hi",
			Actual:   "hi",
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: "hi\nbye",
			Actual:   "hi\nbyu",
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: now.In(notUTC),
			Actual:   now.In(time.UTC),
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: now.In(notUTC),
			Actual:   now.UTC().Add(time.Nanosecond),
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: int32(0),
			Actual:   0,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: int32(4),
			Actual:   4.0,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: uint32(4),
			Actual:   4.0,
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: complex128(4),
			Actual:   4.0,
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: int32(0),
			Actual:   1,
			AreEqual: false,
		},
		{
			Skip:     false,
			Expected: (func())(nil),
			Actual:   (func())(nil),
			AreEqual: true,
		},
		{
			Skip:     false,
			Expected: func() {},
			Actual:   (func())(nil),
			AreEqual: false,
		},
		{
			Skip:     false,
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
	if this.AreEqual {
		comparer := compare.New()
		result := comparer.Compare(this.Expected, this.Actual)
		report := result.Report()
		if result.OK() {
			t.Log(report)
		} else {
			t.Error(report)
		}
	} else {
		comparer := compare.New()
		result := comparer.Compare(this.Expected, this.Actual)
		if !result.OK() {
			t.Log(
				"(report printed below for visual inspection)",
				result.Report(),
			)
		} else {
			t.Errorf(
				"unequal values %v and %v erroneously deemed equal",
				this.Expected,
				this.Actual,
			)
		}
	}
}
