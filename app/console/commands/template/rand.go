package template

import (
	"fmt"
	"math/rand"
)

func init() {
	addConsole("rand", "rand",
		func() {
			fmt.Println(rand.Intn(2))
			fmt.Println("rand")
		})
}
