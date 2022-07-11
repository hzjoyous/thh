package arms

import (
	"github.com/spf13/cast"
	"sync"
	"time"
)

// Together 并行执行
func Together(job func(goId int), counter int) {
	var wg sync.WaitGroup
	for i := 1; i <= counter; i++ {
		wg.Add(1)
		time.Sleep(cast.ToDuration(RandomNum(2)+1) * time.Second)
		go func(i int) {
			defer wg.Done()
			job(i)
		}(i)
	}
	wg.Wait()
}
