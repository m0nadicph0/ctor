package util

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"os"
)

// ErrExitF writes a formatted error message to the standard error
// and exits with the provided exit code
func ErrExitF(code int, format string, a ...any) {
	fmt.Fprint(os.Stderr, aurora.Red(fmt.Sprintf(format, a)))
	os.Exit(code)
}

// WarnExitF writes a formatted error message to the standard error
// and exits with exit code 0
func WarnExitF(format string, a ...any) {
	fmt.Fprint(os.Stderr, aurora.Yellow(fmt.Sprintf(format, a)))
	os.Exit(0)
}

// WarnExit writes a message to the standard error
// and exits with exit code 0
func WarnExit(message string) {
	fmt.Fprint(os.Stderr, aurora.Yellow(message))
	os.Exit(0)
}
