package arms

import (
	"fmt"
	"github.com/spf13/cast"
	"testing"
)

func TestArrayMap(t *testing.T) {
	a := []int{1, 2, 3, 4}
	r := ArrayMap(func(item int) string {
		return cast.ToString(item)
	}, a)
	fmt.Println(r)
}

func TestArrayFilter(t *testing.T) {
	a := []int{1, 2, 3, 4}
	a = ArrayFilter(func(item int) bool {
		return item > 2
	}, a)
	fmt.Println(a)
}
