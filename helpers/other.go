package helpers

import (
	"fmt"
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

// MicrosecondsStr 将 time.Duration 类型（nano seconds 为单位）
// 输出为小数点后 3 位的 ms （microsecond 毫秒，千分之一秒）
func MicrosecondsStr(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Nanoseconds())/1e6)
}

func FirstElement(args []string) string {
	if len(args) > 0 {
		return args[0]
	}
	return ""
}