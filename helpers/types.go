package helpers

import (
	"encoding/json"
	"strconv"
)

func ToString(_var interface{}) string {
	switch t := _var.(type) {
	default:
		return ""
	case bool:
		if t {
			return "true"
		} else {
			return "false"
		}
	case nil:
		return ""
	case int:
		return strconv.FormatInt(int64(t), 10)
	case int16:
		return strconv.FormatInt(int64(t), 10)
	case int32:
		return strconv.FormatInt(int64(t), 10)
	case int64:
		return strconv.FormatInt(int64(t), 10)
	case uintptr:
		return strconv.FormatInt(int64(t), 10)
	case byte:
		return strconv.FormatInt(int64(t), 10)
	case float32:
		return strconv.FormatFloat(float64(t), 'g', 1, 64)
	case float64:
		return strconv.FormatFloat(t, 'g', 1, 64)
	case uint:
		return strconv.FormatFloat(float64(t), 'g', 1, 64)
	case uint16:
		return strconv.FormatFloat(float64(t), 'g', 1, 64)
	case uint32:
		return strconv.FormatFloat(float64(t), 'g', 1, 64)
	case uint64:
		return strconv.FormatUint(uint64(t), 10)
	case []byte:
		return string(t)
	case []rune:
		return string(t)
	case json.Number:
		num := t.String()
		return num
	case string:
		return t
	}
}

func ToInt64(_var interface{}) int64 {
	switch t := _var.(type) {
	default:
		return 0
	case bool:
		if t {
			return 1
		} else {
			return 0
		}
	case nil:
		return 0
	case int:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case uint:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case json.Number:
		num, _ := t.Int64()
		return num
	case string:
		num, _ := strconv.ParseInt(t, 10, 64)
		return num
	case uintptr:
		return int64(t)
	case byte:
		return int64(t)
	}
}

func ToInt(_var interface{}) int {
	switch t := _var.(type) {
	default:
		return 0
	case bool:
		if t {
			return 1
		} else {
			return 0
		}
	case nil:
		return 0
	case int:
		return t
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case float32:
		return int(t)
	case float64:
		return int(t)
	case uint:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	case json.Number:
		num, _ := t.Int64()
		return int(num)
	case string:
		num, _ := strconv.Atoi(t)
		return int(num)
	case uintptr:
		return int(t)
	case byte:
		return int(t)
	}
}
func ToDouble(_var interface{}) float64 {
	value := ToFloat64(_var)
	return value
}
func ToFloat32(_var interface{}) float32 {
	return float32(ToFloat64(_var))
}
func ToFloat64(_var interface{}) float64 {
	switch t := _var.(type) {
	default:
		return float64(0)
	case bool:
		if t {
			return float64(1)
		} else {
			return float64(0)
		}
	case nil:
		return float64(0)
	case int:
		return float64(t)
	case int16:
		return float64(t)
	case int32:
		// int32 is rune
		return float64(t)
	case int64:
		return float64(t)
	case float32:
		return float64(t)
	case float64:
		return t
	case uint:
		return float64(t)
	case uint16:
		return float64(t)
	case uint32:
		return float64(t)
	case uint64:
		return float64(t)
	case json.Number:
		num, _ := t.Int64()
		return float64(num)
	case string:
		num, _ := strconv.ParseFloat(t, 64)
		return num
	case uintptr:
		return float64(t)
	case byte:
		return float64(t)
	}
}
