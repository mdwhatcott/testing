package should

import (
	"errors"
	"io"
	"log"
	"testing"
)

// TODO (future): remove upon moving to v2 module version

// Deprecated
type Reporter interface {
	Helper()
	Report(error)
	io.Writer
}

// Deprecated
func Report(reporters ...Reporter) *T {
	panic(errNotSupported)
}

// Deprecated
type TestingReporter struct{ *testing.T }

// Deprecated
func NewTestingReporter(*testing.T) *TestingReporter {
	panic(errNotSupported)
}

// Deprecated
func (this *TestingReporter) Report(error) {
	panic(errNotSupported)
}

// Deprecated
func (this *TestingReporter) Write([]byte) (n int, err error) {
	panic(errNotSupported)
}

// Deprecated
type CompositeReporter struct{ reporters []Reporter }

// Deprecated
func (this *CompositeReporter) Helper() {
	panic(errNotSupported)
}

// Deprecated
func NewCompositeReporter(reporters ...Reporter) *CompositeReporter {
	panic(errNotSupported)
}

// Deprecated
func (this *CompositeReporter) Report(err error) {
	panic(errNotSupported)
}

// Deprecated
func (this *CompositeReporter) Write(p []byte) (n int, err error) {
	panic(errNotSupported)
}

// Deprecated
type WriterReporter struct{ io.Writer }

// Deprecated
func (this *WriterReporter) Helper() {
	panic(errNotSupported)
}

// Deprecated
func NewWriterReporter(io.Writer) *WriterReporter {
	panic(errNotSupported)
}

// Deprecated
func (this *WriterReporter) Report(err error) {
	panic(errNotSupported)
}

// Deprecated
type LogReporter struct{ logger *log.Logger }

// Deprecated
func NewLogReporter(logger *log.Logger) *LogReporter {
	panic(errNotSupported)
}

// Deprecated
func (this LogReporter) Report(err error) {
	panic(errNotSupported)
}

// Deprecated
func (this LogReporter) Write(p []byte) (n int, err error) {
	panic(errNotSupported)
}

// Deprecated
func (this LogReporter) Helper() {
	panic(errNotSupported)
}

var errNotSupported = errors.New("no longer supported")
