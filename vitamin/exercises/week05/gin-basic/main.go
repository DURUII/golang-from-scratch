package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// SumArgs 规定接受格式，避免手动解析/判断/类型转换 Query
type SumArgs struct {
	X int `form:"x" binding:"required" json:"x"`
	Y int `form:"y" binding:"required" json:"y"`
}

// SumReply 规定返回格式
type SumReply struct {
	Code    int     `json:"code"`
	Message string  `json:"message"`
	ReqData SumArgs `json:"reqData"`
	Data    int     `json:"data"`
}

func main() {
	// 创建默认的路由引擎（组）
	r := gin.Default()
	api := r.Group("/api")
	// c.Get 只在复杂中间件链中使用，这里用 ShouldBind
	api.GET("/sum", func(c *gin.Context) {
		var args SumArgs
		// 在这里 handlers 有 3 个，包含Logger 和 Recovery
		// fmt.Printf("%+v\n\n", c)
		if err := c.ShouldBind(&args); err != nil {
			// H 是 map[string]interface{} 的别名
			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		sum := args.X + args.Y
		c.JSON(http.StatusOK, SumReply{
			Code:    0,
			Message: "success",
			ReqData: args,
			Data:    sum,
		})

	})
	// 启动服务器
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
