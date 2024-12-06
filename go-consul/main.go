package main

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-consul/client"
	"github.com/DSXRIIIII/go-utils/go-consul/discovery"
	"github.com/DSXRIIIII/go-utils/go-consul/ports"
	"github.com/DSXRIIIII/go-utils/go-consul/proto/file"
	"github.com/DSXRIIIII/go-utils/go-consul/server"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"time"
)

func main() {
	ctx, cancelFunc := context.WithCancel(context.Background())
	defer cancelFunc()
	consulCancel, err := discovery.RegisterToConsul(ctx, "hello-grpc")
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = consulCancel()
	}()

	go func() {
		time.Sleep(5 * time.Second)
		logrus.Info("grpc client build start")
		grpcClient, cancel, fail := client.NewGRPCClient(context.Background())
		if fail != nil {
			panic(fail)
		}
		defer func() {
			_ = cancel()
		}()
		helloGRPC := client.NewClientGRPC(grpcClient)
		_, fail = helloGRPC.SayHello(context.Background(), nil)
		if fail != nil {
			panic(fail)
		}
	}()
	server.RunGRPCServer("hello-grpc", func(server *grpc.Server) {
		svc := ports.NewGRPCServer()
		file.RegisterGreeterServer(server, svc)
	})
}
