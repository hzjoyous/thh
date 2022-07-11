package logger

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"thh/arms"
	"thh/arms/config"
	"thh/conf"
)

var log = logrus.StandardLogger()

func Std() *logrus.Logger {
	return log
}

func Info(args ...any) {
	Std().Info(args...)
}

type MyFormatter struct {
}

func (m *MyFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	var b *bytes.Buffer
	if entry.Buffer != nil {
		b = entry.Buffer
	} else {
		b = &bytes.Buffer{}
	}
	timestamp := entry.Time.Format("2006-01-02 15:04:05.999")
	var newLog string
	if entry.HasCaller() {

		newLog = fmt.Sprintf("[%s] [%s] [%s:%d %s] msg=%s\n",
			timestamp, entry.Level,
			entry.Caller.File, entry.Caller.Line, entry.Caller.Function,
			entry.Message,
		)
	} else {
		newLog = fmt.Sprintf("[%s] [%s] msg=%s\n",
			timestamp, entry.Level,
			entry.Message,
		)
	}

	b.WriteString(newLog)
	return b.Bytes(), nil
}

func Init(logPath string) {
	//logrus.SetFormatter(&logrus.JSONFormatter{})
	//formatter := &logrus.TextFormatter{
	//	DisableQuote: true,
	//	//TimestampFormat: "2006-01-02 15:04:05", //时间格式
	//	FullTimestamp: true,
	//	ForceColors:   conf.LogType() == conf.LogTypeStdout,
	//}
	logrus.SetReportCaller(true)

	logrus.SetFormatter(&MyFormatter{})

	debug := config.GetBool("app.debug")

	log.Out = os.Stdout
	if debug {
		log.Level = logrus.TraceLevel
	}

	switch conf.LogType() {
	case conf.LogTypeStdout:
		return
	case conf.LogTypeFile:
		// You could set this to any `io.Writer` such as a file
		if err := arms.FilePutContents(logPath, []byte(""), true); err != nil {
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
