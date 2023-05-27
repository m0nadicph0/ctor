package util

import (
	"bytes"
	"github.com/logrusorgru/aurora/v4"
	"os"
	"testing"
)

func TestWarnExitF(t *testing.T) {
	oldStderr := os.Stderr
	r, w, _ := os.Pipe()
	os.Stderr = w

	exitCalled := false
	osExit = func(code int) {
		exitCalled = true
	}

	WarnExitF("Warning: %s", "something")

	w.Close()
	os.Stderr = oldStderr

	var buf bytes.Buffer
	_, _ = buf.ReadFrom(r)

	expected := aurora.Yellow("Warning: something").String()
	actual := buf.String()

	if expected != actual {
		t.Errorf("Unexpected output:\nExpected: %s\nActual: %s", expected, actual)
	}

	if !exitCalled {
		t.Error("os.Exit(0) was not called as expected")
	}
}
