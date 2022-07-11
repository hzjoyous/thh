package demo

import (
	"fmt"
	"github.com/spf13/cobra"
	"sync"
)

func init() {
	appendCommand(&cobra.Command{Use: "spliceTest", Short: "spliceTest", Run: spliceTest})
}

func spliceTest(cmd *cobra.Command, args []string) {

	fmt.Println("spliceTest")
	type user struct {
	}

	type demo struct {
		lock sync.Mutex
		list []*user
	}
	for {
		d := new(demo)
		wg := sync.WaitGroup{}
		for i := 1; i <= 6; i++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				defer d.lock.Unlock()
				d.lock.Lock()
				d.list = append(d.list, new(user))
			}()
		}
		wg.Wait()

		fmt.Println(d.list)
	}
}
