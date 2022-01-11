package template

import (
	"fmt"
	"time"
)

func init() {
	addConsole("goChannelTest", "goChannelTest",
		func() {
			fmt.Println("goChannelTest")
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
						time.Sleep(time.Millisecond * 500)
						break
					case <-readEnd:
						return
					}
				}
			}

			write()
		})
}
