package RWMutex

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	data    = make(map[int]int)
	rwMutex sync.RWMutex
	wg      sync.WaitGroup
)

// 写入操作
func write(key, value int) {
	defer wg.Done()
	rwMutex.Lock()
	defer rwMutex.Unlock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	data[key] = value
	fmt.Printf("写入: key=%d, value=%d\n", key, value)
}

// 读取操作
func read(key int) {
	defer wg.Done()
	rwMutex.RLock()
	defer rwMutex.RUnlock()
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
	value, exists := data[key]
	if exists {
		fmt.Printf("读取: key=%d, value=%d\n", key, value)
	} else {
		fmt.Printf("读取: key=%d 不存在\n", key)
	}
}

func Work() {
	rand.Seed(time.Now().UnixNano())

	// 模拟多个写入操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go write(i, i*10)
	}

	// 模拟多个读取操作
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go read(i)
	}

	wg.Wait()
	fmt.Println("所有操作完成")
}
