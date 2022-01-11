package template

import (
	"fmt"
	"thh/helpers"
)

func init() {
	addConsole("readCsv", "readCsv",
		func() {
			g := func() chan int {
				r := make(chan int)
				go func() {
					for i := 0; i < 10; i++ {
						r <- i
					}
					close(r)
				}()
				return r
			}
			for r := range g() {
				fmt.Println(r)
			}
			fmt.Println("readCsv")
			for r := range helpers.ReadCsv("./tmp/t.csv") {
				fmt.Println(r)
			}
		})
}
