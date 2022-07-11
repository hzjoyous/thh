package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"os/signal"
	"runtime"
	"sync"
)

func init() {
	appendCommand(&cobra.Command{Use: "demo:signal", Short: "demo:signal", Run: signalHandle})
	appendCommand(&cobra.Command{Use: "demo:runtime", Short: "demo:runtime", Run: runtimeDemo})
	appendCommand(&cobra.Command{Use: "demo:syncPool", Short: "demo:syncPool", Run: syncPoolDemo})
}

func signalHandle(cmd *cobra.Command, args []string) {
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	fmt.Println("开始等待")
	<-quit
	fmt.Println("收到信号，结束")
}

func syncPoolDemo(cmd *cobra.Command, args []string) {

	var pool = sync.Pool{
		New: func() any {
			return "123"
		},
	}
	t := pool.Get().(string)
	fmt.Println(t)

	pool.Put("321")

	t2 := pool.Get().(string)
	fmt.Println(t2)
}

type dataStruct struct {
	name string
}

func runtimeDemo(cmd *cobra.Command, args []string) {

	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d B\n", m.Alloc/8)
	var i int64

	for i = 0; i < 100000000000; i++ {
		useMem()
		//start := time.Now()
		//runtime.GC()
		//elapsed := time.Now().Sub(start)
		//fmt.Println("该函数执行完成耗时：", elapsed)

		if i%100000 == 0 {
			runtime.ReadMemStats(&m)
			fmt.Printf("%d B\n", m.Alloc/8)
		}
	}
}
func useMem() {
	data := dataStruct{}
	data.name = "name"
}
