package main

import (
	"context"
	"fmt"
	"github.com/DSXRIIIII/go-utils/go-redis/zset"
	"github.com/redis/go-redis/v9"
)

var (
	redisClient *redis.Client
	ctx         context.Context
)

func init() {
	// Options go-redis连接池
	redisClient = redis.NewClient(&redis.Options{
		// TODO 信息
		Addr: "116.198.246.11:6379",
		//Username: "default",
		Password: "lhh0719.", // 没有密码，默认值
		DB:       0,          // 默认DB 0
	})

	// 实现redis连接的第二种方式
	//opt, err := redis.ParseURL("redis://default:lhh0719.@116.198.246.11:6379/0")
	//if err != nil {
	//	panic(err)
	//}
	//
	//redisClient := redis.NewClient(opt)

	ctx = context.Background()
	fmt.Printf("----- redis connect success -----\n")
}

func main() {

	// Hash 数据结构
	// HSet 设置hash表的值 key value为一张数据表
	// HGet 指定field和key字段 获取hash表的value值
	// HGetAll 获取key对应的所有数据
	// Scan(&type) 获取指定类型并填充
	// SAdd(key string, members ...interface{}) *IntCmd  向名称为key的set中添加元素member
	// SCard(key string) *IntCmd 获取集合set元素个数
	// SIsMember(key string, member interface{}) *BoolCmd 判断元素member是否在集合set中
	// SMembers(key string) *StringSliceCmd 返回名称为 key 的 set 的所有元素
	// SDiff(keys ...string) *StringSliceCmd 求差集
	// SDiffStore(destination string, keys ...string) *IntCmd 求差集并将差集保存到 destination 的集合
	// SInter(keys ...string) *StringSliceCmd 求交集
	// SInterStore(destination string, keys ...string) *IntCmd 求交集并将交集保存到 destination 的集合
	// SUnion(keys ...string) *StringSliceCmd 求并集
	// SUnionStore(destination string, keys ...string) *IntCmd
	// SPop(key string) *StringCmd 随机返回集合中的一个元素，并且删除这个元素
	// SPopN(key string, count int64) *StringSliceCmd 随机返回集合中的count个元素，并且删除这些元素
	// SRem(key string, members ...interface{}) *IntCmd 删除名称为 key 的 set 中的元素 member,并返回删除的元素个数
	// SRandMember(key string) *StringCmd 随机返回名称为 key 的 set 的一个元素
	// SRandMemberN(key string, count int64) *StringSliceCmd 随机返回名称为 key 的 set 的count个元素
	// SMembersMap(key string) *StringStructMapCmd 把集合里的元素转换成map的key
	// SMove(source, destination string, member interface{}) *BoolCmd 移动集合source中的一个member元素到集合destination中去
	// hashType.HashFunction(ctx, redisClient)

	// String 数据结构
	// Set 指定key value 过期时间
	// SetNX 加锁命令，重复加锁不会影响已经加锁的键
	// Incr 原子性增加 decr 原子性减少 但是必须key保证存在
	// Expire key续期 指定续期时间
	// Del 删除键key
	//stringFunction := stringType.NewStringFunction(ctx, redisClient)
	//stringFunction.IncrDemo()
	//stringFunction.DelDemo()

	// List 数据类型
	//从列表左边插入数据,list不存在则新建一个继续插入数据
	//LPush(key string, values ...interface{}) *IntCmd
	//跟LPush的区别是，仅当列表存在的时候才插入数据
	//LPushX(key string, value interface{}) *IntCmd
	//返回名称为 key 的 list 中 start 至 end 之间的元素
	//返回从0开始到-1位置之间的数据，意思就是返回全部数据
	//LRange(key string, start, stop int64) *StringSliceCmd
	//返回列表的长度大小
	//LLen(key string) *IntCmd
	//截取名称为key的list的数据，list的数据为截取后的值
	//LTrim(key string, start, stop int64) *StatusCmd
	//根据索引坐标，查询列表中的数据
	//LIndex(key string, index int64) *StringCmd
	//给名称为key的list中index位置的元素赋值
	//LSet(key string, index int64, value interface{}) *StatusCmd
	//在指定位置插入数据。op为"after或者before"
	//LInsert(key, op string, pivot, value interface{}) *IntCmd
	//在指定位置前面插入数据
	//LInsertBefore(key string, pivot, value interface{}) *IntCmd
	//在指定位置后面插入数据
	//LInsertAfter(key string, pivot, value interface{}) *IntCmd
	//从列表左边删除第一个数据，并返回删除的数据
	//LPop(key string) *StringCmd
	//删除列表中的数据。删除count个key的list中值为value 的元素。
	//LRem(key string, count int64, value interface{}) *IntCmd
	//listFunction := list.NewListFunction(ctx, redisClient)
	//listFunction.LPushDemo()

	zSetFunction := zset.NewZSetFunction(ctx, redisClient)
	zSetFunction.ZSetDemo()
	zSetFunction.CountDemo()
}
