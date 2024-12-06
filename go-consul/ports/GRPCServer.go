package ports

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-consul/proto/file"
	"github.com/sirupsen/logrus"
)

type GRPCServer struct {
}

func (G *GRPCServer) SayHello(ctx context.Context, request *file.GRPCRequest) (*file.GRPCResponse, error) {
	logrus.Info("hello server get success")
	return &file.GRPCResponse{
		Response: "response",
	}, nil
}

func NewGRPCServer() *GRPCServer {
	return &GRPCServer{}
}
