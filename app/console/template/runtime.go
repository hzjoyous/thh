package template

import (
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "runTimeTest", Short: "runTime", Run: runTimeTest})
}

func runTimeTest(cmd *cobra.Command, args []string) {
	fmt.Println("runTime")
	t := time.Now()
	fmt.Println("hello")
	time.Sleep(time.Second * 3)
	elapsed := time.Since(t)
	fmt.Println("app run time", elapsed)
}
