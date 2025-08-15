package files

import (
	"file-service/config"
	"file-service/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strings"
)

type Reply struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func UploadHandler(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		c.JSON(http.StatusBadRequest, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("upload form err: %s", err.Error()),
		})
		return
	}
	files := form.File["files"]
	
	for _, file := range files {
		// 1. check file size
		if file.Size > config.MaxFileSize {
			c.JSON(http.StatusBadRequest, Reply{
				Code: -1,
				Msg:  fmt.Sprintf("File %s is too large. Max size is 5MB.", file.Filename),
			})
			return
		}

		// 2. check file type
		ext := strings.ToLower(filepath.Ext(file.Filename))
		if _, ok := config.AllowedExtensions[ext]; !ok {
			c.JSON(http.StatusBadRequest, Reply{
				Code: -1,
				Msg:  fmt.Sprintf("Extension %s not allowed", ext),
			})
			return
		}

		// 3. save the file
		newPath := filepath.Join("uploads", file.Filename)
		err := c.SaveUploadedFile(file, newPath)
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{
				Code: -1,
				Msg:  fmt.Sprintf("Failed to save the file %s", file.Filename),
			})
			return
		}
		// 4. insert into database
		err = storage.InsertItem(storage.Item{
			Type:     ext,
			FilePath: newPath,
			FileSize: file.Size,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, Reply{
				Code: -1,
				Msg:  fmt.Sprintf("Error saving file %s into DB", file.Filename),
			})
			return
		}
	}

	// 记录上传历史到 session
	//session := sessions.Default(c)
	//history := session.Get("upload_history")
	//if history == nil {
	//	history = []string{}
	//}
	//history = append(history.([]string), uploadedFiles...)
	//session.Set("upload_history", history)
	//session.Save()

	c.JSON(http.StatusOK, Reply{
		Data: files,
	})
}
