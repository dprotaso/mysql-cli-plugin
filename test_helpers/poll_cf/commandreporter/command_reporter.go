package commandreporter

import (
	"fmt"
	"io"
	"time"

	"github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/config"
)

const timeFormat = "2006-01-02 15:04:05.00 (MST)"

type CommandReporter struct {
	Writer io.Writer
}

func NewCommandReporter(writers ...io.Writer) *CommandReporter {
	var writer io.Writer
	switch len(writers) {
	case 0:
		writer = ginkgo.GinkgoWriter
	case 1:
		writer = writers[0]
	default:
		panic("NewCommandReporter should only take one writer")
	}

	return &CommandReporter{
		Writer: writer,
	}
}

func (r *CommandReporter) Report(startTime time.Time, command string) {
	startColor := ""
	endColor := ""
	if !config.DefaultReporterConfig.NoColor {
		startColor = "\x1b[32m"
		endColor = "\x1b[0m"
	}

	fmt.Fprintf(
		r.Writer,
		"\n%s[%s]> %s %s",
		startColor,
		startTime.UTC().Format(timeFormat),
		command,
		endColor,
	)
}

func (r *CommandReporter) Polling() {
	startColor := ""
	endColor := ""
	if !config.DefaultReporterConfig.NoColor {
		startColor = "\x1b[32m"
		endColor = "\x1b[0m"
	}

	fmt.Fprintf(
		r.Writer,
		"%s.%s",
		startColor,
		endColor,
	)
}