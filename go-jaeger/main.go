package main

import (
	"context"
	"github.com/DSXRIIIII/go-utils/go-jaeger/functionA"
	"github.com/DSXRIIIII/go-utils/go-jaeger/tracer"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer func() {
		if err := cancel; err != nil {
			logrus.Errorf("Error when canceling context: %v", err)
		}
	}()
	// 初始化Jaeger链路追踪相关配置，获取关闭函数和处理可能的初始化错误
	shutdown, err := tracer.InitJaegerProvider("http://116.198.246.11:14268/api/traces", "jaeger_test_chain")
	if err != nil {
		logrus.Fatal(err)
	}
	defer func() {
		if err := shutdown(ctx); err != nil {
			logrus.Errorf("Error when shutting down tracer: %v", err)
		}
	}()

	// 创建Gin框架实例
	c := gin.New()
	// 为 /jaeger/chain 路径添加GET请求的处理函数，并关联链路追踪相关逻辑
	c.GET("/jaeger/chain", functionA.Chain)

	// 启动Gin服务器，监听在本地的8080端口，可根据实际需求修改端口号
	if err := c.Run(":1234"); err != nil {
		logrus.Fatal(err)
	}
}
