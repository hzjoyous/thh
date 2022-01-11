package template

import (
	"context"
	"fmt"
	"time"
)

func init() {
	addConsole("cancel", "cancel",
		func() {
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

		})
}
