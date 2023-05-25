package util

import (
	"fmt"
	"github.com/logrusorgru/aurora/v4"
	"os"
)

// ErrExit writes a formatted error message to the standard error
// and exits with the provided exit code
func ErrExit(code int, format string, a ...any) {
	fmt.Fprint(os.Stderr, aurora.Red(fmt.Sprintf(format, a)))
	os.Exit(code)
}
