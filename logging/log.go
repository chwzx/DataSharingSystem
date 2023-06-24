package logging

import (
	"io"
	"os"

	log "github.com/sirupsen/logrus"
)

func LogSetUp(logLevel, logFile string) {
	switch logLevel {
	case "debug":
		log.SetLevel(log.DebugLevel)
	case "info":
		log.SetLevel(log.InfoLevel)
	case "error":
		log.SetLevel(log.ErrorLevel)
	}

	log.SetOutput(io.MultiWriter(os.Stderr))
}
