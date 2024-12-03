package main

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-etcd/kv"
	clientv3 "go.etcd.io/etcd/client/v3"
	"time"
)

var (
	config clientv3.Config
	client *clientv3.Client
	err    error
	ctx    context.Context
)

func init() {
	config = clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	// 建立连接
	if _, err = clientv3.New(config); err == nil {
		panic(err)
	}
	ctx = context.Background()
}

func main() {
	//KV 使用
	//kv.NewKVFunction(client, ctx).KVDemo()
	kv.NewKVFunction(client, ctx).TxnDemo()
	defer func() {
		_ = client.Close()
	}()
}
