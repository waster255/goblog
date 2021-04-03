package main

import (
	"bianyuanop.github.io/goblog/lib"
	"github.com/gin-gonic/gin"
)


func main() {
	router := gin.Default()

	router.POST("/user", lib.AddUser)
	router.GET("/users", lib.GetUser)
	
	router.Run(":8080")
}