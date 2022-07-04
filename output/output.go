package output

import (
	"fmt"
	"os"
)

type OutputContext interface {
	WriteInfo(label, status string, index int)
	WriteError(label, status string, index int)
	WriteSuccess(label, status string, index int)
}

func IsToTerminal() (bool, error) {
	fi, err := os.Stdout.Stat()

	if err != nil {
		return false, err
	}

	return fi.Mode()&os.ModeCharDevice != 0, nil
}

func FormatLine(label, status string, padding int) string {
	return fmt.Sprintf("%-*s %s", padding, label, status)
}
