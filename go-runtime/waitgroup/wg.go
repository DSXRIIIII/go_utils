package waitgroup

import (
	"fmt"
	"sync"
)

func WaitGroupDemo() {
	var wg sync.WaitGroup
	// 添加3个协程任务
	wg.Add(3)
	// 无法控制打印顺序
	for i := 1; i <= 3; i++ {
		go func(num int) {
			defer func() {
				wg.Done()
			}()
			fmt.Printf("打印:%v\n", i)
		}(i)
	}
	// 等待所有协程完成
	wg.Wait()
	fmt.Println("所有协程都完成了任务")
}

func WaitGroupCalculate() {
	var wg sync.WaitGroup
	var result int
	var lock sync.Mutex
	// 分成10个协程，每个协程计算10个数的和
	numPerGoroutine := 10
	numGoroutines := 100 / numPerGoroutine
	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		start := i * numPerGoroutine
		end := start + numPerGoroutine
		go func() {
			defer wg.Done()
			sum := 0
			for j := start; j <= end; j++ {
				sum += j
			}
			lock.Lock()
			result += sum
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("1到100的整数之和为: %d\n", result)
}
