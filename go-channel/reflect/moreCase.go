package reflect

import (
	"fmt"
	"reflect"
)

func MoreChooseEnv() {
	var ch1 = make(chan int, 10) // 创建一个带缓冲区大小为10的int通道ch1
	var ch2 = make(chan int, 10) // 创建另一个带缓冲区大小为10的int通道ch2

	var cases = createCases(ch1, ch2) // 生成包含ch1和ch2的接收和发送操作的SelectCase切片

	for i := 0; i < 10; i++ { // 执行10次select操作
		chosen, receive, ok := reflect.Select(cases) // 使用reflect.Select随机选择一个可用的case
		if receive.IsValid() {                       // 如果接收到的值有效，说明是接收操作
			fmt.Println("receive:", cases[chosen].Dir, receive, ok) // 打印接收信息：方向、值、是否成功
		} else { // 否则是发送操作
			fmt.Println("send:", cases[chosen].Dir, ok) // 打印发送信息：方向、是否成功
		}
	}
}

func createCases(chs ...chan int) []reflect.SelectCase {
	var cases []reflect.SelectCase // 创建一个SelectCase切片用于存放所有的case

	// 创建接收case
	for _, ch := range chs {
		cases = append(cases, reflect.SelectCase{ // 每个通道添加一个接收操作的case
			Dir:  reflect.SelectRecv,  // 设置为接收方向
			Chan: reflect.ValueOf(ch), // 设置通道
		})
	}

	// 创建发送case
	for i, ch := range chs {
		v := reflect.ValueOf(i)                   // 要发送的值是当前索引值
		cases = append(cases, reflect.SelectCase{ // 每个通道添加一个发送操作的case
			Dir:  reflect.SelectSend,  // 设置为发送方向
			Chan: reflect.ValueOf(ch), // 设置通道
			Send: v,                   // 要发送的值
		})
	}

	return cases // 返回所有构造好的case
}
