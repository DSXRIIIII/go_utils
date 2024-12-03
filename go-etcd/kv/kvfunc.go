package kv

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type KVFunction struct {
	client *clientv3.Client
	ctx    context.Context
}

func NewKVFunction(client *clientv3.Client, ctx context.Context) *KVFunction {
	return &KVFunction{client: client, ctx: ctx}
}

func (f KVFunction) KVDemo() {
	_, err := f.client.Put(f.ctx, "test_key", "hello world")
	if err != nil {
		fmt.Println(err)
	}
	res, err := f.client.Get(f.ctx, "test_key")
	if res != nil {
		for _, val := range res.Kvs {
			fmt.Printf("val:%v", val)
		}
	}
}

// TxnDemo 使用Txn实现原子操作的示例函数
func (f KVFunction) TxnDemo() {
	key := "test_txn_key"
	value := "txn_value"

	// 创建一个事务
	txn := f.client.Txn(f.ctx)

	// 事务操作步骤
	// 先获取指定键的值，检查键是否存在
	_ = clientv3.OpGet(key)
	// 如果键不存在，执行设置键值的操作
	putOp := clientv3.OpPut(key, value)

	// 定义条件和操作
	txn.If(clientv3.Compare(clientv3.Version(key), "=", 0)).
		Then(putOp).
		Else()

	// 提交事务并获取响应
	txnResponse, err := txn.Commit()
	if err != nil {
		fmt.Println(err)
		return
	}

	// 根据事务响应判断操作结果
	if txnResponse.Succeeded {
		fmt.Printf("事务成功，已设置键值对：%s -> %s\n", key, value)
	} else {
		fmt.Printf("事务失败，键 %s 可能已存在，当前值为：%s\n", key, string(txnResponse.Responses[0].GetResponseRange().Kvs[0].Value))
	}
}
