package main

import (
	"log"
	"user_manage/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/users", handlers.GetUsers)
	router.POST("/users", handlers.CreateUser)
	router.PUT("/users", handlers.UpdateUser)
	router.DELETE("/users/:email", handlers.DeleteUser)
	log.Println("服务器启动，监听端口 :8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
