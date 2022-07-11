package ehandle

import "fmt"

var logManager loggerManager

type loggerManager struct {
	setLogger bool
	logger    Logger
}

func (itself *loggerManager) log(err error) {
	if itself.setLogger {
		itself.logger.Error(err)
	} else {
		fmt.Println(err)
	}
}

type Logger interface {
	Error(any)
}

func InitLogger(logger Logger) {
	logManager.logger = logger
	logManager.setLogger = true
}

func PrIF(err error) bool {
	if err != nil {
		logManager.log(err)
		return true
	}
	return false
}
