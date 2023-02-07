package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{
		TimestampFormat: "2006/01/02 - 15:04:05",
	})
	src := "./logs/" + time.Now().Format("2006-01-02") + ".txt"
	logFile, err := os.OpenFile(src, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		logrus.Panicln(err)
	}
	logrus.SetOutput(logFile)
}
