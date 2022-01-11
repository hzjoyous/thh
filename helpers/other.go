package helpers

import (
	"net/http"
	"strings"
	"time"
)

func GetMicroTime() int64 {
	return time.Now().UnixNano() / 1000000
}

func InArrayInt64(need int64, haystack []int64) bool {
	for _, v := range haystack {
		if need == v {
			return true
		}
	}
	return false
}

func InArrayString(need string, haystack []string) bool {
	for _, v := range haystack {
		if strings.Contains(need, v) {
			return true
		}
	}
	return false
}

func Date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}
