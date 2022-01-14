package template

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "readWithWrite", Short: "readWithWrite", Run: readWithWrite})
}
func readWithWrite(cmd *cobra.Command, args []string) {
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

}
