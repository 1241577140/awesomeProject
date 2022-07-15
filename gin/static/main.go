package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.Static("/static", "./resource")

	// 监听并在 0.0.0.0:8080 上启动服务
	router.Run(":8080")
}
