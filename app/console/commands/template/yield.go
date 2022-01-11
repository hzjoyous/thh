package template

import "fmt"

func init() {
	addConsole("yield", "yield",
		func() {
			a := func(max int) chan int {
				ch := make(chan int)
				go func() {
					for i := 0; i < max; i++ {
						ch <- i
					}
					close(ch)
				}()
				return ch
			}
			for i := range a(10) {
				fmt.Println(i)
			}
		})
}
