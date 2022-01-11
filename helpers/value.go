package helpers

import "strings"

func Join(sep string, ss []string) string {
	return strings.Join(ss, sep)
}

func Implode(sep string, ss []string) string {
	return Join(sep, ss)
}

func Ltrim(s string) string {
	return strings.TrimLeft(s, " ")
}
func ArrayFill(startIndex int, num uint, value interface{}) map[int]interface{} {
	result := make(map[int]interface{})
	var i uint
	for i = 0; i < num; i++ {
		result[startIndex] = value
		startIndex++
	}
	return result
}

func ArrayKeyExists(key interface{}, arr map[interface{}]interface{}) bool {
	if len(arr) == 0 {
		return false
	}
	for k := range arr {
		if key == k {
			return true
		}
	}
	return false
}

func ArrayMap(f func(item interface{}) interface{}, arr map[string]interface{}) (r []interface{}) {
	for _, v := range arr {
		r = append(r, f(v))
	}
	return r
}

func ArrayMerge(arr ...interface{}) []interface{} {
	r := make([]interface{}, 0)
	return append(r, arr...)
}
