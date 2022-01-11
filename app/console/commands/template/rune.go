package template

import (
	"fmt"
	"unicode/utf8"
)

func init() {
	addConsole("rune", "rune",
		func() {
			fmt.Println("rune")
			fmt.Println([]byte("你好"))
			fmt.Println(string([]byte{228, 189, 160, 229, 165, 189}))
			fmt.Println([]rune("你好"))
			fmt.Println([]int32("你好"))
			fmt.Println(string([]rune("你好")))
			fmt.Println(string([]rune{20320, 22909}))
			fmt.Println(utf8.RuneCountInString("你好n"))
			fmt.Println(len([]rune("你好n")))
		})
}
