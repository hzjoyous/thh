package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "build", Short: "build", Run: build})
}
func build(cmd *cobra.Command, args []string) {
	fmt.Println("build")
	times := "2020-09-18 15:04:05"
	tagStartTime, _ := time.Parse("2025-01-01 01:01:01", times)
	tagEndTime, _ := time.Parse("2025-12-31 01:01:01", times)

	fmt.Println(tagStartTime, tagEndTime)
}
