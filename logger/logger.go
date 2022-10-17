package logger

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func New() {
	log.SetFormatter(&log.JSONFormatter{})
	log.SetLevel(log.WarnLevel)
	log.SetOutput(os.Stdout)
}
