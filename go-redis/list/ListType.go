package list

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ListFunction struct {
	ctx    context.Context
	client *redis.Client
}

func NewListFunction(ctx context.Context, client *redis.Client) *ListFunction {
	return &ListFunction{
		ctx:    ctx,
		client: client,
	}
}

func (l ListFunction) LPushDemo() {
	//仅当列表存在的时候才插入数据,此时列表不存在，无法插入
	l.client.LPushX(l.ctx, "studentList", "tom")

	//此时列表不存在，依然可以插入
	l.client.LPush(l.ctx, "studentList", "jack")

	//此时列表存在的时候才能插入数据
	l.client.LPushX(l.ctx, "studentList", "tom")

	// LPush支持一次插入任意个数据
	err := l.client.LPush(l.ctx, "studentList", "lily", "lilei", "zhangsan", "lisi").Err()
	if err != nil {
		panic(err)
	}
	// 返回从0开始到-1位置之间的数据，意思就是返回全部数据
	vals, err := l.client.LRange(l.ctx, "studentList", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	result, err := l.client.LLen(l.ctx, "studentList").Result()
	fmt.Printf("列表长度为：%v\n", result)
	fmt.Println(vals) //注意列表是有序的，输出结果是[lisi zhangsan lilei lily tom jack]
}
