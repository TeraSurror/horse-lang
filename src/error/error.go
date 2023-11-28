package error

import (
	"fmt"
	"os"
)

func ReportError(line int, message string) {
	report(line, "", message)
}

func report(line int, where string, message string) {
	fmt.Fprintf(os.Stderr, "[line: %d] Error %s: %s", line, where, message)
}
