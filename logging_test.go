
package curl

import (
	"testing"
	"bytes"
	"log"
	"os"
	"fmt"
	"regexp"
)

func TestDefaultLogLevel(t *testing.T) {
    if log_level != _DEFAULT_LOG_LEVEL {t.Error("Test failed, expected DEFAULT_LOG_LEVEL level.")}
}

func TestsetLogLevel(t *testing.T) {
    setLogLevel("DEBUG")
    defer setLogLevel("DEFAULT_LOG_LEVEL")
    if log_level != _DEBUG {t.Error("Test failed, expected DEBUG level.")}
    setLogLevel("INFO")
    if log_level != _INFO {t.Error("Test failed, expected INFO level.")}
    setLogLevel("WARN")
    if log_level != _WARN {t.Error("Test failed, expected WARN level.")}
    setLogLevel("ERROR")
    if log_level != _ERROR {t.Error("Test failed, expected ERROR level.")}
}

var (
    testFormat = "test format %s"
    testArgument = "test string 1"
    expectedRegexp = regexp.MustCompile(".*" + fmt.Sprintf(testFormat, testArgument) + "\n$")
)


func TestLogf(t *testing.T) {
    buf := new(bytes.Buffer)
    log.SetOutput(buf)
    defer log.SetOutput(os.Stderr)
    setLogLevel("DEBUG")
    defer setLogLevel("DEFAULT_LOG_LEVEL")

    logf(_DEBUG, testFormat, testArgument)
    line := buf.String()
    matched := expectedRegexp.MatchString(line)
    if !matched {
        t.Errorf("log output should match %q and is %q.", expectedRegexp, line)
    }
}

func TestLogfUsesLogLevel(t *testing.T) {
    buf := new(bytes.Buffer)
    log.SetOutput(buf)
    defer log.SetOutput(os.Stderr)
    setLogLevel("WARN")
    defer setLogLevel("DEFAULT_LOG_LEVEL")

    logf(_DEBUG, testFormat, testArgument)
    line := buf.String()
    expectedLine := ""
    if line != expectedLine {
        t.Errorf("log output should match %q and is %q.", expectedLine, line)
    }
}
