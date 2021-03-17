package log

import (
	"../model"
	log "github.com/sirupsen/logrus"
	"io"
	"os"
)

const PATH_ERROR = "./log/files/error"

func LogError(model model.LogResponse) {
	file, err := os.OpenFile(PATH_ERROR, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Error(err)
	}
	defer file.Close()
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFormatter(&log.TextFormatter{})

	log.SetLevel(log.ErrorLevel)
	log.WithFields(log.Fields{
		"file": model.FileLocation,
		"line": model.Line,
	}).Error(model.Message)
}
