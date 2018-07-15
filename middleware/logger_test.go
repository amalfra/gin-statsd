package middleware

import (
	"bytes"
	"io"
	"os"
	"strings"
	"testing"
)

var b bytes.Buffer
var writer io.Writer

func init() {
	writer = io.Writer(&b)
}
func TestPrintLogWithNotExistingLevel(t *testing.T) {
	setupLogger(writer, os.Stdout, os.Stderr)
	printLog("log msg", 9)
	if !strings.Contains(b.String(), "INFO") || !strings.Contains(b.String(), "log msg") {
		t.Error("Incorrect log generated")
	}
}

func TestPrintLogWithInfoLevel(t *testing.T) {
	setupLogger(writer, os.Stdout, os.Stderr)
	printLog("log msg", infoLevel)
	if !strings.Contains(b.String(), "INFO") || !strings.Contains(b.String(), "log msg") {
		t.Error("Incorrect log generated")
	}
}

func TestPrintLogWithWarningLevel(t *testing.T) {
	setupLogger(os.Stdout, writer, os.Stderr)
	printLog("log msg", warningLevel)
	if !strings.Contains(b.String(), "WARNING") || !strings.Contains(b.String(), "log msg") {
		t.Error("Incorrect log generated")
	}
}

func TestPrintLogWithErrorLevel(t *testing.T) {
	setupLogger(os.Stdout, os.Stderr, writer)
	printLog("log msg", errorLevel)
	if !strings.Contains(b.String(), "ERROR") || !strings.Contains(b.String(), "log msg") {
		t.Error("Incorrect log generated")
	}
}
