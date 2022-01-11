package helpers

import (
	"sync"
)

// Together 并行执行
func Together(job func(goId int), counter int) {
	var wg sync.WaitGroup
	for i := 0; i <= counter; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			job(i)
		}(i)
	}
	wg.Wait()
}
