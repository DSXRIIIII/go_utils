package functionA

import (
	"github.com/DSXRIIIII/go-utils/go-jaeger/functionB"
	"github.com/DSXRIIIII/go-utils/go-jaeger/tracer"
	"github.com/gin-gonic/gin"
)

func Chain(c *gin.Context) {
	// 从请求上下文中获取链路追踪的相关信息并开启一个新的Span，这里使用了请求的路径作为Span的名称
	ctx, span := tracer.Start(c.Request.Context(), "functionA")
	id := functionB.ChainEnd(ctx)
	defer span.End()
	c.JSON(200, gin.H{"tracingID": id})
}
