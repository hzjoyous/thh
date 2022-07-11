package arms

import (
	"encoding/json"
	"github.com/spf13/cast"
)

func JsonEncode(obj any) string {
	marshal, err := json.Marshal(obj)
	if err != nil {
		return ""
	}
	return cast.ToString(marshal)
}

func JsonDecode[T any](str string) T {
	var obj T
	_ = json.Unmarshal([]byte(str), &obj)
	return obj
}

func JsonDecodeBl[T any](str []byte) T {
	var obj T
	_ = json.Unmarshal(str, &obj)
	return obj
}
