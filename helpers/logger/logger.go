package logger

import (
	"github.com/sirupsen/logrus"
	"os"
	"thh/conf"
	"thh/helpers"
)

var log = logrus.StandardLogger()

func Std() *logrus.Logger {
	return log
}

func Info(args ...interface{}) {
	Std().Info(args...)
}

func Init(logPath string) {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//logrus.SetReportCaller(true)
	formater := &logrus.TextFormatter{
		DisableQuote: true,
		//TimestampFormat: "2006-01-02 15:04:05", //时间格式
		FullTimestamp: true,
		ForceColors:   conf.LogType() == conf.LogTypeStdout,
	}
	logrus.SetFormatter(formater)

	log.Out = os.Stdout
	switch conf.LogType() {
	case conf.LogTypeStdout:
		return
	case conf.LogTypeFile:
		// You could set this to any `io.Writer` such as a file
		if err := helpers.FilePutContents(logPath, []byte(""), true); err != nil {
			log.Info(err)
			return
		}

		file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Info("Failed to log to file, using default stderr")
			return
		}
		log.Out = file
		return
	default:
		log.Info("Unknown Log Type")
		return
	}

}
