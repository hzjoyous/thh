package arms

import (
	"fmt"
	"testing"
)

func TestPower(t *testing.T) {
	Together(func(goId int) {
		fmt.Println(goId)
	}, 10)
}
