package cond

import (
	"fmt"
	"sync"
	"time"
)

func CondDemo() {
	var mu sync.Mutex
	cond := sync.NewCond(&mu)

	for i := 0; i < 10; i++ {
		go func(i int) {
			cond.L.Lock()
			defer cond.L.Unlock()
			cond.Wait()
			fmt.Printf("Goroutine %d is awake\n", i)
		}(i)
	}

	time.Sleep(time.Second)
	for i := 0; i < 10; i++ {
		cond.Signal() // 被唤醒的 goroutine 会依次打印消息
		time.Sleep(100 * time.Millisecond)
	}
}
