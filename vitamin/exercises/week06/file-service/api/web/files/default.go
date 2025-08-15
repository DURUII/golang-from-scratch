package files

import "github.com/gin-gonic/gin"

func DefaultHandler(c *gin.Context) {
	c.HTML(200, "upload.html", gin.H{
		"title": "文件管理系统",
	})
}
