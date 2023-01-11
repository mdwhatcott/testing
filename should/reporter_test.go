package should_test

import (
	"bytes"
	"io"
	"log"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestReporterReportsFailures(t *testing.T) {
	var (
		writer1 bytes.Buffer
		writer2 bytes.Buffer
	)
	reporter := should.Report(
		should.NewWriterReporter(&writer1),
		should.NewLogReporter(log.New(&writer2, "", 0)),
	)
	reporter.So(1, should.Equal, 2) // FAIL

	should.New(t).So(writer1.Len(), should.BeGreaterThan, 0)
	should.New(t).So(writer2.Len(), should.BeGreaterThan, 0)
}

func TestReporterIgnoresPassingTests(t *testing.T) {
	var (
		writer1 bytes.Buffer
		writer2 bytes.Buffer
	)
	reporter := should.Report(
		should.NewWriterReporter(&writer1),
		should.NewLogReporter(log.New(&writer2, "", 0)),
	)
	reporter.So(1, should.Equal, 1) // PASS

	should.New(t).So(writer1.Len(), should.Equal, 0)
	should.New(t).So(writer2.Len(), should.Equal, 0)
}

func TestReporterIsWriter(t *testing.T) {
	var (
		writer1 bytes.Buffer
		writer2 bytes.Buffer
	)
	reporter := should.Report(
		should.NewWriterReporter(&writer1),
		should.NewLogReporter(log.New(&writer2, "", 0)),
	)

	const message = "Hello, world!"
	n, err := io.WriteString(reporter, message)

	should.So(t, err, should.BeNil)
	should.So(t, n, should.Equal, len(message))
	should.So(t, writer1.String(), should.Equal, message)
	should.So(t, writer2.String(), should.Equal, message+"\n")
}
