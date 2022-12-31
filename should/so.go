package should

import (
	"fmt"
	"log"
)

func So(T, actual any, assertion Func, expected ...any) {
	err := assertion(actual, expected...)
	if err == nil {
		return
	}
	if tt, ok := T.(t); ok {
		tt.Helper()
		tt.Error(err)
	} else if l, ok := T.(*log.Logger); ok {
		l.Println(err)
	} else {
		fmt.Println(err)
	}
}

type t interface {
	Helper()
	Error(...any)
}

// Deprecated: dead code
type Fmt struct{}

// Deprecated: dead code
type Log struct{}

// Deprecated: dead code
func (Fmt) Helper() { panic("deprecated code") }

// Deprecated: dead code
func (Log) Helper() { panic("deprecated code") }

// Deprecated: dead code
func (Fmt) Error(a ...any) { panic("deprecated code") }

// Deprecated: dead code
func (Log) Error(a ...any) { panic("deprecated code") }

type Func func(actual any, expected ...any) error
