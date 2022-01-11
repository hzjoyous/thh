package template

import (
	"fmt"
	"gorm.io/gorm/utils"
	"runtime"
	"sync"
	"thh/app/service/remote_service"
	"thh/helpers"
)

func init() {
	addConsole("bigtx", "bigtx",
		func() {
			fmt.Println("bigtx")
			r := remote_service.NewBigTXClient()

			job := func(item remote_service.BigTXClient) func(goId int) {
				return func(goId int) {
					//errNum := 0
					//maxNum := 5 + rand.Intn(30)
					for {
						resp, err := r.RegNew("wu"+helpers.ToString(idMaker.get()), "111111")
						//resp, err := item.SendEx()
						if err != nil {
							fmt.Println(err)
						}

						fmt.Println(resp.String())
						//if strings.Index(resp.String(), "物品发送成功！") != -1 {
						//	counter.add()
						//	fmt.Println(counter.get())
						//	errNum = 0
						//} else {
						//	fmt.Println(resp.String())
						//	errNum += 1
						//	sleepNum := maxNum
						//	if maxNum > errNum {
						//		sleepNum = errNum
						//	}
						//	fmt.Println("返回错误:sleep" + helpers.ToString(sleepNum))
						//	time.Sleep(time.Duration(sleepNum) * time.Second)
						//}
					}
				}
			}(r)
			helpers.Together(job, 20)
		})
}

type IdMakerInOnP struct {
	id   uint64
	lock sync.Mutex
}

var idMaker IdMakerInOnP

func (itself *IdMakerInOnP) get() uint64 {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	itself.id += 1
	return itself.id
}

var counter Counter

type Counter struct {
	lock   sync.Mutex
	number int
}

func (itself *Counter) add() {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	itself.number += 1
}
func (itself *Counter) get() int {
	itself.lock.Lock()
	defer itself.lock.Unlock()
	if itself.number%1000 == 0 {
		var m runtime.MemStats
		runtime.ReadMemStats(&m)
		fmt.Println(utils.ToString(m.Alloc/1024/8) + "kb")
	}
	return itself.number
}
