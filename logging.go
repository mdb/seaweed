package seaweed

import (
	"os"

	logging "github.com/op/go-logging"
)

// NewLogger returns a *logging.Logger configured for use with seaweed
// and set to the logging.Level it's passed.
func NewLogger(level logging.Level) *logging.Logger {
	var log = logging.MustGetLogger("seaweed")

	backend := logging.NewLogBackend(os.Stdout, "", 0)

	var format = logging.MustStringFormatter(
		`%{time:15:04:05.000} %{shortfunc} %{level} %{message}`,
	)

	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	backendLeveled.SetLevel(level, "")
	logging.SetBackend(backendLeveled)

	return log
}
