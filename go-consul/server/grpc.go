package server

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

func init() {
	logger := logrus.New()
	logger.SetLevel(logrus.WarnLevel)
}

func RunGRPCServer(_ string, registerServer func(server *grpc.Server)) {
	addr := "127.0.0.1:5678"
	RunGRPCServerOnAddr(addr, registerServer)
}

func RunGRPCServerOnAddr(addr string, registerServer func(server *grpc.Server)) {
	grpcServer := grpc.NewServer()
	registerServer(grpcServer)
	listen, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Panic(err)
	}
	logrus.Infof("Starting gRPC server, Listening: %s", addr)
	if err := grpcServer.Serve(listen); err != nil {
		logrus.Panic(err)
	}

}
