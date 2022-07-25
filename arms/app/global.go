package app

import (
	"time"
)

var start time.Time

func InitStart() {
	start = time.Now()
}

func GetRunTime() time.Duration {
	return time.Now().Sub(start)
}
