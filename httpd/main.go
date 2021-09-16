package main

import (
	"github.com/gin-gonic/gin"
	"source/httpd/handler"
)

func main() {
	r := gin.Default()
	r.GET("/get",handler.GetUser)
	r.POST("/post",handler.PostUser)
	r.Run()
}