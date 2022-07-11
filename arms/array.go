package arms

func ArrayMap[T, T2 any](f func(T) T2, list []T) (result []T2) {
	for _, item := range list {
		result = append(result, f(item))
	}
	return
}

func ArrayFilter[T any](f func(T) bool, list []T) (result []T) {
	for _, item := range list {
		if f(item) {
			result = append(result, item)
		}
	}
	return
}

func ArrayFill[t any](startIndex int, num uint, value t) map[int]t {
	result := make(map[int]t)
	var i uint
	for i = 0; i < num; i++ {
		result[startIndex] = value
		startIndex++
	}
	return result
}

func ArrayKeyExists(key any, arr map[any]any) bool {
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
