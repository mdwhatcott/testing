package assert

import (
	"fmt"
	"log"
)

func So(t T, actual any, assertion Func, expected ...any) {
	if t == nil {
		t = Fmt{}
	}
	t.Helper()
	err := assertion(actual, expected...)
	if err != nil {
		t.Error(err)
	}
}

type T interface {
	Helper()
	Error(...any)
}

type Fmt struct{}
type Log struct{}

func (Fmt) Helper() {}
func (Log) Helper() {}

func (Fmt) Error(a ...any) { fmt.Println(a...) }
func (Log) Error(a ...any) { log.Println(a...) }

type Func func(actual any, expected ...any) error
