package arms

import (
	"errors"
	"sync"
)

var mutex = &sync.Mutex{}

var queueList = make(map[string][]string)

func QueueRPush(key string, data ...string) {
	mutex.Lock()
	defer mutex.Unlock()
	queue, _ := queueList[key]
	queue = append(queue, data...)
	queueList[key] = queue
}

func QueueLPop(key string) (string, error) {
	mutex.Lock()
	defer mutex.Unlock()
	queue, _ := queueList[key]
	if len(queue) > 0 {
		result := queue[0]
		queue = queue[1:]
		queueList[key] = queue
		return result, nil
	}
	return "", errors.New("queue is null")
}

func QueueLen(key string) int {
	mutex.Lock()
	defer mutex.Unlock()
	queue, ok := queueList[key]
	if ok {

		return len(queue)
	}
	return 0
}
