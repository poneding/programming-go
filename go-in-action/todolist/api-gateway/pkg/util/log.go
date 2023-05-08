package util

import (
	"log"
	"os"
	"path"
	"time"

	"github.com/sirupsen/logrus"
)

var LogrusObject *logrus.Logger

func init() {
	if LogrusObject != nil {
		src, _ := setOutputFile()
		LogrusObject.SetOutput(src)
		return
	}

	logger := logrus.New()
	src, _ := setOutputFile()
	logger.SetOutput(src)
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
}

func setOutputFile() (*os.File, error) {
	now := time.Now()
	logFilePath := ""
	if dir, err := os.Getwd(); err == nil {
		logFilePath = path.Join(dir, "logs")
	}

	_, err := os.Stat(logFilePath)
	if os.IsNotExist(err) {
		if err := os.Mkdir(logFilePath, os.ModePerm); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	logFileName := now.Format("2006-01-02") + ".log"
	fileName := path.Join(logFilePath, logFileName)
	if _, err := os.Stat(fileName); err != nil {
		if _, err := os.Create(fileName); err != nil {
			log.Println(err)
			return nil, err
		}
	}

	src, err := os.OpenFile(fileName, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return src, nil
}
