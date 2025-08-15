package files

import (
	"file-service/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

type DownloadArgs struct {
	UUID string `json:"uuid" form:"uuid"`
}

func DownloadHandler(c *gin.Context) {
	var args DownloadArgs
	if err := c.ShouldBind(&args); err != nil {
		c.JSON(http.StatusBadRequest, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("invalid argument: %s", err.Error()),
		})
		return
	}
	// Find item by UUID
	item, err := storage.FindItemByUUID(args.UUID)
	if err != nil {
		c.JSON(http.StatusBadRequest, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("DB failed to find the item by UUID %s: %s", args.UUID, err.Error()),
		})
		return
	}
	// 这里要去除前缀文件夹，留下最后的文件名+后缀名
	filename := filepath.Base(item.FilePath)
	if filename == "" {
		c.JSON(http.StatusBadRequest, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("Invalid file path for UUID %s", args.UUID),
		})
		return
	}

	// Set response headers for download
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))

	// Send the file to the client
	filePath := filepath.Join("./uploads", filename) // Safe construction of the file path
	c.File(filePath)
}
