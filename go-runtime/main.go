package main

import (
	"github.com/DSXRIIIII/go-utils/go-runtime/cond"
)

func main() {
	//var chan1 = make(chan struct{})
	//go func() {
	//	for {
	//		select {
	//		case <-chan1:
	//			fmt.Println("123")
	//		default:
	//			fmt.Println("456")
	//		}
	//	}
	//}()
	//time.Sleep(1 * time.Minute)
	//// 查看当前的goroutine数量
	//fmt.Println(runtime.NumGoroutine())

	//chanel案例
	//channel.ChannelDemo()

	//waitGroup案例
	//waitgroup.WaitGroupCalculate()

	//Cond 案例
	cond.CondDemo()

}
