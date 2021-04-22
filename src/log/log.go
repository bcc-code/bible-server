package log

import (
	"github.com/sirupsen/logrus"
	"os"

	stackdriver "github.com/TV4/logrus-stackdriver-formatter"
)

// L is the main exposed logger
var L = logrus.New()

// ConfigureGlobalLogger with the correct formatter and debug level
func ConfigureGlobalLogger(logLevel logrus.Level) {

	// Automatically detect if we are in GCR and apply Stackdriver log format
	// https://cloud.google.com/run/docs/reference/container-contract#env-vars
	serviceName := os.Getenv("K_SERVICE")
	if serviceName != "" {
		L.Formatter = stackdriver.NewFormatter(
			stackdriver.WithService(serviceName),
			stackdriver.WithVersion(os.Getenv("K_REVISION")),
		)
	}

	L.Level = logLevel
}
