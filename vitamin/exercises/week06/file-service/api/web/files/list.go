package files

import (
	"file-service/internal/storage"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type ListArgs struct {
	PageNum  int `json:"page_num" binding:"omitempty"`
	PageSize int `json:"page_size" binding:"omitempty"`
}

func fillDefaultArgs(args *ListArgs) {
	args.PageNum = 1
	args.PageSize = 10000
}

func ListHandler(c *gin.Context) {
	//session := sessions.Default(c)
	//history := session.Get("upload_history")
	//if history == nil {
	//	history = []string{}
	//}
	var args ListArgs
	if err := c.ShouldBindQuery(&args); err != nil {
		c.JSON(http.StatusBadRequest, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("invalid argument: %s", err.Error()),
		})
	}
	fillDefaultArgs(&args)
	items, err := storage.ListItem(args.PageNum, args.PageSize)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Reply{
			Code: -1,
			Msg:  fmt.Sprintf("DB failed to list items: %s", err.Error()),
		})
		return
	}
	c.JSON(http.StatusOK, Reply{
		Data: items,
	})
}
