package helpers

import (
	"fmt"
	"testing"
)

func TestUnidecode(t *testing.T) {
	fmt.Println(Unidecode(`"汉字\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587\u4e2d\u6587"`))
}
