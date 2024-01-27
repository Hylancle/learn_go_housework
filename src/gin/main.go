package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/hello", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello，world") //http://localhost:8080/hello
	})

	// 参数路由，路径参数
	r.GET("/users/:name", func(ctx *gin.Context) {
		name := ctx.Param("name")
		ctx.String(http.StatusOK, "hello, "+name)

	})

	// 通配符
	r.GET("/views/*.html", func(ctx *gin.Context) {
		view := ctx.Param(".html") // 带'.'
		ctx.String(http.StatusOK, "view是： "+view)

	})
	// 查询参数
	// GET /order?id=123
	r.GET("/order", func(ctx *gin.Context) {
		id := ctx.Query("id")
		ctx.String(http.StatusOK, "订单ID是： "+id)

	})

	r.POST("/login", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "hello, login")
	})
	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
	//r.Run(":8081")
	// 查询数据用GET
	// 提交数据用POST REstfule风格
}
