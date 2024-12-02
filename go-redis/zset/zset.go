package zset

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
)

type ZSetFunction struct {
	ctx    context.Context
	client *redis.Client
}

func NewZSetFunction(ctx context.Context, client *redis.Client) *ZSetFunction {
	return &ZSetFunction{
		ctx:    ctx,
		client: client,
	}
}

func (z ZSetFunction) ZSetDemo() {
	//统计开发语言排行榜
	zsetKey := "language_rank"
	languages := []redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 92.0, Member: "C/C++"},
	}

	// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
	num, err := z.client.ZAdd(z.ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("zadd failed, err:%v\n", err)
		return
	}
	fmt.Printf("ZAdd添加成功 %d 元素\n", num)
	// 添加一个元素到集合
	z.client.ZAdd(z.ctx, zsetKey, redis.Z{Score: 87, Member: "Vue"}).Err()

	//给元素Vue加上8分，最终vue得分95分
	z.client.ZIncrBy(z.ctx, zsetKey, 8, "Vue")
	// 返回从0到-1位置的集合元素， 元素按分数从小到大排序 0到-1代表则返回全部数据
	values, err := z.client.ZRange(z.ctx, zsetKey, 0, -1).Result()
	if err != nil {
		panic(err)
	}
	for _, val := range values {
		fmt.Println(val)
	}
}

func (z ZSetFunction) CountDemo() {
	//统计开发语言排行榜
	zsetKey := "language_rank"
	//返回集合元素的个数
	size, err := z.client.ZCard(z.ctx, zsetKey).Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(size)
	//统计某个分数段内的元素个数，这里是查询的95<分数<100的元素个数
	count, err := z.client.ZCount(z.ctx, zsetKey, "95", "100").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(count)
}
