package string

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

type StringFunction struct {
	ctx    context.Context
	client *redis.Client
}

func NewStringFunction(ctx context.Context, client *redis.Client) *StringFunction {
	return &StringFunction{
		ctx:    ctx,
		client: client,
	}
}

func (s StringFunction) GetValueDemo() {
	// 简单设置key value字符串 如果时间为-1/0则代表不会过期
	err := s.client.Set(s.ctx, "string-set", "string-set-value", -1).Err()
	if err != nil {
		panic(err)
	}

	// Result函数返回两个值，第一个是key的值，第二个是错误信息
	val, err := s.client.Get(s.ctx, "string-set").Result()
	// 判断查询是否出错
	if err != nil {
		panic(err)
	}
	fmt.Println("string-val：", val)
}

func (s StringFunction) SetNXDemo() {
	_, err := s.client.SetNX(s.ctx, "string-set-nx", "lock", -1).Result()
	// 如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
	err = s.client.SetNX(s.ctx, "string-set-nx", "lock_again", -1).Err()
	if err != nil {
		panic(err)
	}
}

func (s StringFunction) DelDemo() {
	s.client.Del(s.ctx, "string-set", "string-set-nx")
	fmt.Println("del demo success")
}

func (s StringFunction) IncrDemo() {
	//设置一个age测试自增、自减
	err := s.client.Set(s.ctx, "age", "20", 10*time.Minute).Err()
	if err != nil {
		panic(err)
	}
	s.client.Incr(s.ctx, "age")      // 自增
	s.client.IncrBy(s.ctx, "age", 5) //+5
	s.client.Decr(s.ctx, "age")      // 自减
	s.client.DecrBy(s.ctx, "age", 3) //-3 此时age的值是22

	var val string
	val, err = s.client.Get(s.ctx, "age").Result()
	fmt.Println("age=", val) //22
}
