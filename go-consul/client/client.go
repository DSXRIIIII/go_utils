package client

import (
	"context"
	"fmt"
	"github.com/DSXRIIIII/go-utils/go-consul/discovery/consul"
	"github.com/DSXRIIIII/go-utils/go-consul/proto/file"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"math/rand"
)

type ClientGRPC struct {
	client file.GreeterClient
}

func NewClientGRPC(client file.GreeterClient) *ClientGRPC {
	return &ClientGRPC{client: client}
}

func (c ClientGRPC) SayHello(ctx context.Context, _ *interface{}) (*interface{}, error) {
	response, err := c.client.SayHello(ctx, &file.GRPCRequest{
		Request: "request",
	})
	logrus.Info("hello grpc service get response:", response)
	return nil, err
}

// NewGRPCClient 封装grpc客户端
func NewGRPCClient(ctx context.Context) (client file.GreeterClient, close func() error, err error) {
	grpcAddr, err := getAddr(ctx, "hello-grpc")
	if err != nil {
		panic(err)
	}
	opts, err := grpcDialOption(grpcAddr)
	if err != nil {
		return nil, func() error { return nil }, err
	}

	conn, err := grpc.NewClient(grpcAddr, opts...)
	if err != nil {
		return nil, func() error { return nil }, err
	}
	return file.NewGreeterClient(conn), conn.Close, nil
}

func getAddr(ctx context.Context, serviceName string) (string, error) {
	registry, err := consul.New("116.198.246.11:8500")
	if err != nil {
		return "", err
	}
	addr, err := registry.Discover(ctx, serviceName)
	if err != nil {
		return "", err
	}
	if len(addr) == 0 {
		return "", fmt.Errorf("got empty %s addrs from consul", serviceName)
	}
	i := rand.Intn(len(addr))
	logrus.Infof("Discovered %d instance of %s, addrs=%v", len(addr), serviceName, addr)
	return addr[i], nil
}

func grpcDialOption(addr string) ([]grpc.DialOption, error) {
	return []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}, nil
}
