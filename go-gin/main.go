package main

import (
	"github.com/DSXRIIIII/go-utils/go-gin/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	engine := gin.Default()
	engine.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"message": "yes",
		})
	})

	engine.GET("/user/search", func(c *gin.Context) {
		username := c.DefaultQuery("username", "default")
		//username := c.Query("username")
		address := c.Query("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	engine.POST("/user/search", func(c *gin.Context) {
		// DefaultPostForm取不到值时会返回指定的默认值
		//username := c.DefaultPostForm("username", "小王子")
		username := c.PostForm("username")
		address := c.PostForm("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 参数绑定
	engine.GET("/user/search/:username/:address", func(c *gin.Context) {
		username := c.Param("username")
		address := c.Param("address")
		//输出json结果给调用方
		c.JSON(http.StatusOK, gin.H{
			"message":  "ok",
			"username": username,
			"address":  address,
		})
	})

	// 绑定QueryString示例 (/loginQuery?user=q1mi&password=123456)
	engine.GET("/loginForm", func(c *gin.Context) {
		var login Login
		// ShouldBind()会根据请求的Content-Type自行选择绑定器
		if err := c.ShouldBind(&login); err == nil {
			c.JSON(http.StatusOK, gin.H{
				"user":     login.User,
				"password": login.Password,
			})
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	//重定向
	engine.GET("/sogo", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.sogo.com/")
	})

	//engine 构建路由组
	engine.Group("/shop")

	engine.Use(middleware.StatCost())

	_ = engine.Run(":8080")

}

type Login struct {
	User     string
	Password string
}
