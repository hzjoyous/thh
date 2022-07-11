package cmd

import (
	"fmt"
	"sync"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	appendCommand(&cobra.Command{
		Use:   "apilock",
		Short: "HERE PUTS THE COMMAND DESCRIPTION",
		Run:   runApiLock,
		// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
	})
}

type httpLock struct {
	use  bool
	lock sync.Mutex
}

func (itself *httpLock) canUse(timeout int) bool {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	if itself.use == true {
		return false
	} else {
		itself.use = true
		go func() {
			time.Sleep(time.Second * time.Duration(timeout))
			itself.use = false
		}()
		return true
	}
}

var apiLock httpLock

func runApiLock(cmd *cobra.Command, args []string) {

	for li := 0; li <= 10; li++ {
		if apiLock.canUse(3) {
			fmt.Println("可以访问")
		} else {
			fmt.Println("不可访问")
		}
		time.Sleep(time.Millisecond * 500)
	}
}
