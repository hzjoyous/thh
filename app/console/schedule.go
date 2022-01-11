package console

import (
	"fmt"
	"github.com/robfig/cron/v3"
	"thh/helpers/logger"
	"time"
)

var c = cron.New()

func RunJob() {

	//mirai_manager.MiraiClientManager()

	_, err := c.AddFunc("* * * * *", heart)
	if err != nil {
		fmt.Println(err)
	}
	//c.Run()
	Logger.Info("The task is set successfully")

}

// 之后如果需要平滑关闭可以参考如下代码
//var needStop = false
//var needStopL = &sync.Mutex{}
//
//func demo(){
//	ctx:=c.Stop()
//	Stop()
//	<-ctx.Done()
//}
//
//func Stop() {
//	c.Run()
//	needStopL.Lock()
//	defer needStopL.Unlock()
//	needStop = true
//}
//
//func getStop() bool {
//	needStopL.Lock()
//	defer needStopL.Unlock()
//	return needStop
//}

func heart() {
	fmt.Println("HEART_IN_RUN_JOB :", time.Now())
}
