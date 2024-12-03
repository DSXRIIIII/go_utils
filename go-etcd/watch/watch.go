package watch

import (
	"context"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type WatchFunction struct {
	client *clientv3.Client
	ctx    context.Context
}

func NewWatchFunction(client *clientv3.Client, ctx context.Context) *WatchFunction {
	return &WatchFunction{client: client, ctx: ctx}
}

// WatchDemo 实现对指定键的异步监视功能
func (f WatchFunction) WatchDemo(keysToWatch []string) {

}
