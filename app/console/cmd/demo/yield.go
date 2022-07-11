package demo

import (
	"fmt"
	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{Use: "yieldDemo", Short: "yield", Run: yieldDemo})
}
func yieldDemo(cmd *cobra.Command, args []string) {
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
}
