package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/helpers"
)

func init() {
	appendCommand(&cobra.Command{Use: "readCsv", Short: "readCsv", Run: readCsv})
}
func readCsv(cmd *cobra.Command, args []string) {
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
}
