package main

import (
	"file-service/api/web/files"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	// 设置模板，设置静态文件目录
	r.LoadHTMLGlob("templates/*")
	r.Static("/uploads", "./uploads")
	// 设置 session
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("filesession", store))
	r.GET("/", files.DefaultHandler)          // 前端页面
	r.POST("/upload", files.UploadHandler)    // 上传
	r.GET("/list", files.ListHandler)         // 历史
	r.GET("/download", files.DownloadHandler) // 下载
	if err := r.Run(":8080"); err != nil {
		panic(err)
	}
}
