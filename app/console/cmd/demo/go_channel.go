package demo

import (
	"context"
	"fmt"
	"github.com/spf13/cast"
	"github.com/spf13/cobra"
	"math/rand"
	"runtime"
	"time"
)

func init() {
	appendCommand(&cobra.Command{Use: "channel:test", Short: "goChannelTest", Run: goChannelTest})
	appendCommand(&cobra.Command{Use: "channel:cancel", Short: "cancel", Run: cancelByContext})
	appendCommand(&cobra.Command{Use: "channel:manyChannel", Short: "cancel", Run: goManyChannelTest})
	appendCommand(&cobra.Command{Use: "channel:judgeChannelClose", Short: "judgeChannelClose", Run: judgeChannelClose})
	appendCommand(&cobra.Command{Use: "channel:selectDefault", Short: "selectDefault", Run: selectDefault})
	appendCommand(&cobra.Command{Use: "channel:closeChannel", Short: "closeChannel", Run: closeChannel})
	appendCommand(&cobra.Command{Use: "channel:channelManagerMoreGoFinish", Short: "channelManagerMoreGoFinish", Run: channelManagerMoreGoFinish})
	appendCommand(&cobra.Command{Use: "channel:channelSpider", Short: "channelSpider", Run: channelSpider})
}

func channelSpider(cmd *cobra.Command, args []string) {
	cInt := make(chan int, 10)
	// 模拟爬取的数据
	cInt <- 1
	spider := func(i int) {
		// todo spider
		// push next job
		nestJobNum := rand.Intn(3)
		for i := 0; i <= nestJobNum; i++ {
			cInt <- i
		}
	}
	for i := 1; i <= 3; i++ {
		go func() {
			for cData := range cInt {
				fmt.Println("开始消费")
				spider(cData)
				fmt.Println("消费结束")
			}
		}()
	}
	time.Sleep(time.Second * 100)

}

type workJob struct {
	done chan bool
}

// 可以为了实现功能而使用 channel ，不要为了使用 channel 而实现
func channelManagerMoreGoFinish(cmd *cobra.Command, args []string) {
	wList := []workJob{}
	for i := 1; i <= 10; i++ {
		// 如果不初始化会阻塞
		newJob := workJob{make(chan bool)}
		go func(job workJob, i int) {
			t := rand.Int31n(3)
			fmt.Println("will sleep ", t, "second")
			time.Sleep(time.Second * cast.ToDuration(t))
			fmt.Println(i, "job will end")
			job.done <- true
			fmt.Println(i, "job end")
		}(newJob, i)
		wList = append(wList, newJob)
	}

	for _, workJobItem := range wList {
		<-workJobItem.done
	}
	fmt.Println("all end")
}

// 无关闭无泄漏，会回收
func goManyChannelTest(cmd *cobra.Command, args []string) {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("%d B\n", m.Alloc/8)
	var i int64

	for i = 0; i < 100000000000; i++ {
		useChan()

		if i%1000000 == 0 {
			runtime.ReadMemStats(&m)
			fmt.Printf("Thousand-hand:useMem %d KB\n", m.Alloc/1024/8)
		}
	}
}
func useChan() {
	cInt := make(chan int, 0)
	//defer close(cInt)
	//fmt.Println(cInt)
	if false {
		fmt.Println(cInt)
	}
}

func selectDefault(cmd *cobra.Command, args []string) {
	isClose := func(a chan int) bool {
		select {
		case <-a:
			return true
		default:
			return false
		}
	}
	cInt := make(chan int, 1000)
	go func(cInt chan int) {
		for i := 1; i <= 1000; i++ {
			fmt.Println(i, "wait into channel")
			cInt <- i
			if i%100 == 0 {
				time.Sleep(time.Second)
				fmt.Println("sleep 1 second")
			}
		}
		fmt.Println("end")
	}(cInt)
	go func(readCInt chan int) {
		for {

			select {
			case i := <-readCInt:
				fmt.Println(i)
				break
			default:
				// 个人感觉此处的default作用有限，如果不进行一些监控统计或者超时主动关闭的话，此处的任何处理意义不大。
				// 作为一个阻塞分支的非阻塞处理，使用场景有限。大多数场景可以选择不加入default操作，让阻塞分支正常等待即可。
				fmt.Println("no data")
				time.Sleep(time.Second * 1)
			}
		}
	}(cInt)
	time.Sleep(time.Second * 3)
	fmt.Println(isClose(cInt))
}

func judgeChannelClose(cmd *cobra.Command, args []string) {
	cInt := make(chan int, 0)
	isClose := func(a chan int) bool {
		select {
		case <-a:
			return true
		default:
			return false
		}
	}
	fmt.Println(isClose(cInt))
}

func closeChannel(cmd *cobra.Command, args []string) {
	cInt := make(chan int, 0)

	go func(writeCInt chan int) {
		for i := 1; i <= 100; i++ {
			writeCInt <- i
		}
		close(writeCInt)
	}(cInt)
	for data := range cInt {
		fmt.Println(data)
	}
}

// 开辟大量 chan 不进行关闭是否会导致内存泄漏？
// 不会
// 如何判断一个 chan 是否关闭
// 场景1问题、
// 任务分发是否适合使用 channel
// 场景2问题、
// 多任务执行用 channel 判断所有任务结束是否合适
// goChannelTest 简单无缓冲读写 channel
func goChannelTest(cmd *cobra.Command, args []string) {
	cInt := make(chan int, 20)
	endInt := make(chan int)

	go func() {
		for i := 1; i <= 10; i++ {
			fmt.Println(i, "wait into channel")
			cInt <- i
		}
		close(cInt)
		endInt <- 1
	}()

	write := func() {
		for {
			select {
			case data := <-cInt:
				fmt.Println("get ", data, " from channel")
				time.Sleep(time.Millisecond * 500)
				break
			case <-endInt:
				return
			}
		}
	}

	write()
	data2 := <-cInt
	fmt.Println(data2)

	data3 := <-cInt
	fmt.Println(data3)
}

func cancelByContext(cmd *cobra.Command, args []string) {
	// gin 的平滑关闭可以这么理解
	// gin 的关闭操作同时在检查两件事情，一件是当前是否没有链接了，另一件事情是你传入了一个最大超时时间是否已到，如果到了也停止
	// 结束在这个函数里面的标志，表现在 srv.Shutdown() 返回了结果。

	// gin 平滑关闭原理
	// src.Shutdown 关闭外部连接，
	// 之后gin创建了一个定时器
	// 检查当前是否还有运行着的连接，如果有进行select操作
	// select 中同时检测两个channel 第一个是 context.WithTimeout ，他会在5s后才发送
	// 如果他收到了结果会直接return
	// 另一个是 就是定时器的channel，
	// 如果收到的是定时器的 channel ，并不会return，而是进行再一次检查。如果还是有链接存在，就会继续下一轮的select

	fmt.Println(time.Now().String())
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	<-ctx.Done()
	fmt.Println(time.Now().String())

	ch := make(chan int)
	ch2 := make(chan int)

	go func() {
		for i := 5; i > 0; i-- {
			time.Sleep(time.Second * 1)
		}
		ch <- 1
	}()
	go func() {
		time.Sleep(time.Second * 10)
		ch2 <- 1
	}()

	select {
	case <-ch:
		fmt.Println("收到ch")
		break
	case <-ch2:
		fmt.Println("收到ch2")
		break
	}
	fmt.Println("准备关闭了")

}
