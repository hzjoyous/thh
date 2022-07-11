package arms

import (
	"fmt"
	"github.com/spf13/cast"
	"net/http"
	"time"
)

func InArray[vT any](need vT, haystack []vT) bool {
	for _, v := range haystack {
		if cast.ToString(need) == cast.ToString(v) {
			return true
		}
	}
	return false
}

func Date() string {
	return time.Now().UTC().Format(http.TimeFormat)
}

func GetMicroTime() int64 {
	return time.Now().UnixNano() / 1000000
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
