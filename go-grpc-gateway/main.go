package main

import (
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	net "net"
	"net/http"

	hello "github.com/DSXRIIIII/go-utils/go-grpc-gateway/proto"
	"google.golang.org/grpc"
)

type Server struct {
	hello.UnimplementedGreeterServer
}

func (s Server) SayHello(ctx context.Context, request *hello.HelloRequest) (*hello.HelloReply, error) {
	return &hello.HelloReply{Message: request.Name + " world"}, nil
}

func NewServer() *Server {
	return &Server{}
}

func main() {
	// Create a listener on TCP port
	lis, err := net.Listen("tcp", ":6891")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	// Create a gRPC server object
	s := grpc.NewServer()
	// Attach the Greeter service to the server
	hello.RegisterGreeterServer(s, &Server{})
	// Serve gRPC Server
	log.Println("Serving gRPC on 0.0.0.0:6891")
	go func() {
		log.Fatalln(s.Serve(lis))
	}()

	// 创建一个连接到我们刚刚启动的 gRPC 服务器的客户端连接
	// gRPC-Gateway 就是通过它来代理请求（将HTTP请求转为RPC请求）
	conn, err := grpc.NewClient(
		"0.0.0.0:6891",
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalln("Failed to dial server:", err)
	}

	gwmux := runtime.NewServeMux()
	// 注册Greeter
	err = hello.RegisterGreeterHandler(context.Background(), gwmux, conn)
	if err != nil {
		log.Fatalln("Failed to register gateway:", err)
	}

	gwServer := &http.Server{
		Addr:    ":8090",
		Handler: gwmux,
	}
	// 8090端口提供gRPC-Gateway服务
	log.Println("Serving gRPC-Gateway on http://0.0.0.0:8090")
	log.Fatalln(gwServer.ListenAndServe())
}
