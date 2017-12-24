package seaweed

import (
	"os"

	logging "github.com/op/go-logging"
)

// NewLogger returns a *logging.Logger configured for use with seaweed
func NewLogger() *logging.Logger {
	var log = logging.MustGetLogger("example")

	backend := logging.NewLogBackend(os.Stdout, "", 0)

	var format = logging.MustStringFormatter(
		`%{time:15:04:05.000} %{shortfunc} %{level} %{message}`,
	)

	backendFormatter := logging.NewBackendFormatter(backend, format)
	backendLeveled := logging.AddModuleLevel(backendFormatter)

	backendLeveled.SetLevel(logging.INFO, "")
	logging.SetBackend(backendLeveled)

	return log
}
