package main

import (
	"fmt"
	"github.com/DSXRIIIII/go-utils/go-runtime/RWMutex"
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
	//cond.CondDemo()

	//s := NewSlice()
	//defer s.Add(1).Add(2)
	//s.Add(3)

	RWMutex.Work()

}

type Slice []int

func NewSlice() Slice {
	return make(Slice, 0)
}
func (s *Slice) Add(elem int) *Slice {
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}
