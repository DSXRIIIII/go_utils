package env

import (
	"fmt"
	"sync"
)

func PrintOddEven() {
	wg := sync.WaitGroup{}
	wg.Add(2)

	chan1 := make(chan struct{})
	chan2 := make(chan struct{})

	// 打印奇数的 goroutine
	go func() {
		defer wg.Done()
		for i := 1; i <= 9; i += 2 {
			<-chan1
			fmt.Printf("协程1打印：%d\n ", i)
			chan2 <- struct{}{}
		}
	}()

	// 打印偶数的 goroutine
	go func() {
		defer wg.Done()
		for i := 0; i <= 8; i += 2 {
			fmt.Printf("协程2打印：%d\n ", i)
			chan1 <- struct{}{}
			<-chan2
		}
	}()

	wg.Wait()
}

func PrintOddEvenBYDoneChannel() {
	oddCh := make(chan struct{})
	evenCh := make(chan struct{})
	doneCh := make(chan struct{})

	// 打印奇数
	go func() {
		for i := 1; i < 10; i += 2 {
			<-oddCh
			fmt.Printf("协程1打印：%d\n ", i)
			evenCh <- struct{}{}
		}
	}()

	// 打印偶数
	go func() {
		for i := 0; i < 10; i += 2 {
			<-evenCh
			fmt.Printf("协程2打印：%d\n ", i)
			if i == 8 {
				doneCh <- struct{}{}
				return
			}
			oddCh <- struct{}{}
		}
	}()

	// 启动打印（偶数先打）
	evenCh <- struct{}{}
	<-doneCh
}

func PrintLetterNumber() {
	ch1 := make(chan struct{})
	ch2 := make(chan struct{})
	ch3 := make(chan struct{})
	go func() {
		i := 1
		for {
			select {
			case <-ch1:
				fmt.Printf("goroutine 1 print:%v\n", i)
				i++
				if i == 100 {
					ch3 <- struct{}{}
					return
				}
				ch2 <- struct{}{}
			}
		}
	}()

	go func() {
		i := 'A'
		for {
			if i > 'Z' {
				ch3 <- struct{}{}
				return
			}
			select {
			case <-ch2:
				fmt.Printf("goroutine 2 print:%v\n", string(i))
				i++
				ch1 <- struct{}{}
			}
		}
	}()
	ch1 <- struct{}{}
	<-ch3
}
