package channel

import (
	"fmt"
	"time"
)

func ChannelDemo() {
	chan1 := make(chan interface{})
	go func() {
		for {
			select {
			case data := <-chan1:
				fmt.Printf("chan1 收到数据:%v\n", data)
			}
		}
	}()
	// 不加func() 就会死锁
	go func() {
		for i := 0; i < 5; i++ {
			chan1 <- i
			time.Sleep(1 * time.Second)
			fmt.Printf("发送:%v\n", i)
			if i == 3 {
				close(chan1)
				fmt.Printf("chan1 closed,now i = %v\n", i)
			}
		}
	}()
	time.Sleep(10 * time.Second)
}
