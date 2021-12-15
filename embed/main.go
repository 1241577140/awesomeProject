package main

import (
	"embed"
	_ "embed"
	"github.com/gin-gonic/gin"
	"net/http"
)

//go:embed preview.gif
var content []byte

/*//go:embed preview.text
var content string
  //支持 string  或者 []byte 来存储简单的文件
*/
//go:embed /embed/*
var fs embed.FS

func main() {
	engine := gin.Default()
	engine.GET("/api", func(ctx *gin.Context) {
		ctx.Data(http.StatusOK, "image/gif", content)
	})
	engine.Run(":8080")
}
