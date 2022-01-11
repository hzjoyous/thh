package template

import (
	"fmt"
)

func init() {
	addConsole("readWithWrite", "readWithWrite",
		func() {
			fmt.Println("readWithWrite")
			read := make(chan int)
			readEnd := make(chan int)

			go func() {
				for i := 1; i <= 10; i++ {
					read <- i
				}
				readEnd <- 1
			}()

			write := func() {
				for {
					select {
					case data := <-read:
						fmt.Println(data)
						break
					case <-readEnd:
						return
					}
				}
			}

			write()

		})
}
