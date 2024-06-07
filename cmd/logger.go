/*
Copyright © 2024 digitalhydra <digitalhydra@proton.me>
*/
package cmd

import (
	"os"

	"github.com/charmbracelet/log"
)

func Logger() *log.Logger {
	logger := log.New(os.Stderr)
	logger.SetReportTimestamp(false)
	return logger
}
