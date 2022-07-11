package cmd

import (
	"context"
	"fmt"
	"github.com/spf13/cobra"
	"time"
)

func init() {
	appendCommand(CmdSonPainc)
}

var CmdSonPainc = &cobra.Command{
	Use:   "son_painc",
	Short: "HERE PUTS THE COMMAND DESCRIPTION",
	Run:   runSonPainc,
	// Args:  cobra.ExactArgs(1), // 只允许且必须传 1 个参数
}

func runSonPainc(cmd *cobra.Command, args []string) {

	var reStart = make(chan bool, 1)
	for {
		fmt.Println("启动")
		go Son(reStart)
		<-reStart
	}
}

func Son3() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	ctx.Done()
}

func Son2() (result bool) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("运行结束", err)
			result = false
			return
		}
	}()
	result = true
	for {
		fmt.Println("sleep1")
		time.Sleep(time.Second)
	}

}

func Son(reStart chan bool) bool {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("运行结束", err)
			reStart <- false
		}
	}()
	i := 0
	for {
		if i > 5 {
			panic("i is 6 panic")
		}
		i += 1
		fmt.Println(i)
		time.Sleep(time.Second * 1)
	}
}
