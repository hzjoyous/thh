package tool

import (
	"fmt"
	"github.com/spf13/cobra"
	"thh/arms"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "z:waitGroup", Short: "this is a z:waitGroup", Run: waitGroup})
}
func waitGroup(cmd *cobra.Command, args []string) {
	arms.Together(func(i int) {
		defer func(i int) {
			fmt.Println(i, "end")
		}(i)
		time.Sleep(time.Second * 3)
	}, 10)
}
