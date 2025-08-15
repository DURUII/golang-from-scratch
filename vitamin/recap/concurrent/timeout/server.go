package timeout

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data,omitempty"`
}

func NewRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()
	r.GET("/work", func(c *gin.Context) {
		t, _ := strconv.Atoi(c.DefaultQuery("t", "0"))
		time.Sleep(time.Duration(t) * time.Millisecond)
		c.JSON(http.StatusOK, Response{Msg: "success"})
	})
	return r
}
