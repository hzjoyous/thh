package helpers

import (
	"fmt"
	"testing"
)

func TestRand(t *testing.T) {
	str := RandomString(10)
	fmt.Println(str)
}
