## String API
```go
//给数据库中名称为key的string赋予值value,并设置失效时间，0为永久有效
Set(key string, value interface{}, expiration time.Duration) *StatusCmd
//查询数据库中名称为key的value值
Get(key string) *StringCmd
//设置一个key的值，并返回这个key的旧值
GetSet(key string, value interface{}) *StringCmd
//如果key不存在，则设置这个key的值,并设置key的失效时间。如果key存在，则设置不生效
SetNX(key string, value interface{}, expiration time.Duration) *BoolCmd
//批量查询key的值。比如redisDb.MGet("name1","name2","name3")
MGet(keys ...string) *SliceCmd
//批量设置key的值。redisDb.MSet("key1", "value1", "key2", "value2", "key3", "value3")
MSet(pairs ...interface{}) *StatusCmd
//Incr函数每次加一,key对应的值必须是整数或nil
//否则会报错incr key1: ERR value is not an integer or out of range
Incr(key string) *IntCmd
// IncrBy函数,可以指定每次递增多少,key对应的值必须是整数或nil
IncrBy(key string, value int64) *IntCmd
// IncrByFloat函数,可以指定每次递增多少，跟IncrBy的区别是累加的是浮点数
IncrByFloat(key string, value float64) *FloatCmd
// Decr函数每次减一,key对应的值必须是整数或nil.否则会报错
Decr(key string) *IntCmd
//DecrBy,可以指定每次递减多少,key对应的值必须是整数或nil
DecrBy(key string, decrement int64) *IntCmd
//删除key操作,支持批量删除	redisDb.Del("key1","key2","key3")
Del(keys ...string) *IntCmd
//设置key的过期时间,单位秒
Expire(key string, expiration time.Duration) *BoolCmd
//给数据库中名称为key的string值追加value
Append(key, value string) *IntCmd
```



## List API
```go
 //从列表左边插入数据,list不存在则新建一个继续插入数据
LPush(key string, values ...interface{}) *IntCmd
//跟LPush的区别是，仅当列表存在的时候才插入数据
LPushX(key string, value interface{}) *IntCmd
//返回名称为 key 的 list 中 start 至 end 之间的元素
//返回从0开始到-1位置之间的数据，意思就是返回全部数据
LRange(key string, start, stop int64) *StringSliceCmd
//返回列表的长度大小
LLen(key string) *IntCmd
//截取名称为key的list的数据，list的数据为截取后的值
LTrim(key string, start, stop int64) *StatusCmd
//根据索引坐标，查询列表中的数据
LIndex(key string, index int64) *StringCmd
//给名称为key的list中index位置的元素赋值
LSet(key string, index int64, value interface{}) *StatusCmd
//在指定位置插入数据。op为"after或者before"
LInsert(key, op string, pivot, value interface{}) *IntCmd
//在指定位置前面插入数据
LInsertBefore(key string, pivot, value interface{}) *IntCmd
//在指定位置后面插入数据
LInsertAfter(key string, pivot, value interface{}) *IntCmd
//从列表左边删除第一个数据，并返回删除的数据
LPop(key string) *StringCmd
//删除列表中的数据。删除count个key的list中值为value 的元素。
LRem(key string, count int64, value interface{}) *IntCmd
```

## Set API
```go
//向名称为key的set中添加元素member
SAdd(key string, members ...interface{}) *IntCmd
//获取集合set元素个数
SCard(key string) *IntCmd
//判断元素member是否在集合set中
SIsMember(key string, member interface{}) *BoolCmd
//返回名称为 key 的 set 的所有元素
SMembers(key string) *StringSliceCmd
//求差集
SDiff(keys ...string) *StringSliceCmd
//求差集并将差集保存到 destination 的集合
SDiffStore(destination string, keys ...string) *IntCmd
//求交集
SInter(keys ...string) *StringSliceCmd
//求交集并将交集保存到 destination 的集合
SInterStore(destination string, keys ...string) *IntCmd
//求并集
SUnion(keys ...string) *StringSliceCmd
//求并集并将并集保存到 destination 的集合
SUnionStore(destination string, keys ...string) *IntCmd
//随机返回集合中的一个元素，并且删除这个元素
SPop(key string) *StringCmd
// 随机返回集合中的count个元素，并且删除这些元素
SPopN(key string, count int64) *StringSliceCmd
//删除名称为 key 的 set 中的元素 member,并返回删除的元素个数
SRem(key string, members ...interface{}) *IntCmd
//随机返回名称为 key 的 set 的一个元素
SRandMember(key string) *StringCmd
//随机返回名称为 key 的 set 的count个元素
SRandMemberN(key string, count int64) *StringSliceCmd
//把集合里的元素转换成map的key
SMembersMap(key string) *StringStructMapCmd
//移动集合source中的一个member元素到集合destination中去
SMove(source, destination string, member interface{}) *BoolCmd
```

## HASH API
```go
//根据key和字段名，删除hash字段，支持批量删除hash字段
HDel(key string, fields ...string) *IntCmd
//检测hash字段名是否存在。
HExists(key, field string) *BoolCmd
//根据key和field字段，查询field字段的值
HGet(key, field string) *StringCmd
//根据key查询所有字段和值
HGetAll(key string) *StringStringMapCmd
//根据key和field字段，累加数值。
HIncrBy(key, field string, incr int64) *IntCmd
//根据key和field字段，累加数值。
HIncrByFloat(key, field string, incr float64) *FloatCmd
//根据key返回所有字段名
HKeys(key string) *StringSliceCmd
//根据key，查询hash的字段数量
HLen(key string) *IntCmd
//根据key和多个字段名，批量查询多个hash字段值
HMGet(key string, fields ...string) *SliceCmd
//根据key和多个字段名和字段值，批量设置hash字段值
HMSet(key string, fields map[string]interface{}) *StatusCmd
//根据key和field字段设置，field字段的值
HSet(key, field string, value interface{}) *BoolCmd
//根据key和field字段，查询field字段的值
HSetNX(key, field string, value interface{}) *BoolCmd
```

## ZSet API
```go
//添加一个或者多个元素到集合，如果元素已经存在则更新分数
ZAdd(key string, members ...Z) *IntCmd
ZAddNX(key string, members ...Z) *IntCmd
ZAddXX(key string, members ...Z) *IntCmd
ZAddCh(key string, members ...Z) *IntCmd
ZAddNXCh(key string, members ...Z) *IntCmd
// 添加一个或者多个元素到集合，如果元素已经存在则更新分数
ZAddXXCh(key string, members ...Z) *IntCmd
//增加元素的分数
ZIncr(key string, member Z) *FloatCmd
ZIncrNX(key string, member Z) *FloatCmd
ZIncrXX(key string, member Z) *FloatCmd
//增加元素的分数，增加的分数必须是float64类型
ZIncrBy(key string, increment float64, member string) *FloatCmd
// 存储增加分数的元素到destination集合
ZInterStore(destination string, store ZStore, keys ...string) *IntCmd
//返回集合元素个数
ZCard(key string) *IntCmd
//统计某个分数范围内的元素个数
ZCount(key, min, max string) *IntCmd
//返回集合中某个索引范围的元素，根据分数从小到大排序
ZRange(key string, start, stop int64) *StringSliceCmd
//ZRevRange的结果是按分数从大到小排序。
ZRevRange(key string, start, stop int64) *StringSliceCmd
//根据分数范围返回集合元素，元素根据分数从小到大排序，支持分页。
ZRangeByScore(key string, opt ZRangeBy) *StringSliceCmd
//根据分数范围返回集合元素，用法类似ZRangeByScore，区别是元素根据分数从大到小排序。
ZRemRangeByScore(key, min, max string) *IntCmd
//用法跟ZRangeByScore一样，区别是除了返回集合元素，同时也返回元素对应的分数
ZRangeWithScores(key string, start, stop int64) *ZSliceCmd
//根据元素名，查询集合元素在集合中的排名，从0开始算，集合元素按分数从小到大排序
ZRank(key, member string) *IntCmd
//ZRevRank的作用跟ZRank一样，区别是ZRevRank是按分数从大到小排序。
ZRevRank(key, member string) *IntCmd 
//查询元素对应的分数
ZScore(key, member string) *FloatCmd
//删除集合元素
ZRem(key string, members ...interface{}) *IntCmd
//根据索引范围删除元素。从最低分到高分的（stop-start）个元素
ZRemRangeByRank(key string, start, stop int64) *IntCmd
```

