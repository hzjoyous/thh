package tool

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "z:waitGroup", Short: "this is a z:waitGroup", Run: waitGroup})
}
func waitGroup(cmd *cobra.Command, args []string) {
	var wg sync.WaitGroup
	workNum := 20
	// 之前由于把i设置为0开始，导致可能会减到负一导致报错
	for i := 1; i <= workNum; i++ {
		// 严谨点这样就不会受计算影响
		wg.Add(1)
		go func(i int) {
			defer func(i int) {
				fmt.Println(i, "end")
				wg.Done()
			}(i)
			time.Sleep(time.Second * 3)
		}(i)
	}
	wg.Wait()
	fmt.Println("end")
}
