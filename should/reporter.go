package should

import (
	"io"
	"log"
	"os"
	"testing"
)

type Func func(actual any, expected ...any) error

type Reporter interface {
	Report(error)
	io.Writer
}

type T struct{ Reporter }

func New(t *testing.T) *T {
	return &T{Reporter: NewTestingReporter(t)}
}
func Report(reporters ...Reporter) *T {
	if len(reporters) == 0 {
		reporters = append(reporters, NewWriterReporter(os.Stdout))
	}
	return &T{Reporter: NewCompositeReporter(reporters...)}
}
func (this *T) So(actual any, assertion Func, expected ...any) (ok bool) {
	err := assertion(actual, expected...)
	this.Reporter.Report(err)
	return err == nil
}

type TestingReporter struct{ *testing.T }

func NewTestingReporter(t *testing.T) *TestingReporter {
	return &TestingReporter{T: t}
}
func (this *TestingReporter) Report(err error) {
	if err != nil {
		this.Helper()
		this.Error(err)
	}
}
func (this *TestingReporter) Write(p []byte) (n int, err error) {
	this.Helper()
	this.Log(string(p))
	return len(p), nil
}

type CompositeReporter struct{ reporters []Reporter }

func NewCompositeReporter(reporters ...Reporter) *CompositeReporter {
	return &CompositeReporter{reporters: reporters}
}
func (this *CompositeReporter) Report(err error) {
	for _, reporter := range this.reporters {
		reporter.Report(err)
	}
}
func (this *CompositeReporter) Write(p []byte) (n int, err error) {
	for _, reporter := range this.reporters {
		n, err = reporter.Write(p)
		if err != nil {
			break
		}
	}
	return n, err
}

type WriterReporter struct{ io.Writer }

func NewWriterReporter(writer io.Writer) *WriterReporter {
	return &WriterReporter{Writer: writer}
}
func (this *WriterReporter) Report(err error) {
	if err != nil {
		_, _ = io.WriteString(this, err.Error())
	}
}

type LogReporter struct{ logger *log.Logger }

func NewLogReporter(logger *log.Logger) *LogReporter {
	return &LogReporter{logger: logger}
}
func (this LogReporter) Report(err error) {
	if err != nil {
		this.logger.Print(err.Error())
	}
}
func (this LogReporter) Write(p []byte) (n int, err error) {
	this.logger.Print(string(p))
	return len(p), nil
}
