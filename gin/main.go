package main

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"io/ioutil"
	"net/http"
)

func main() {
	g := gin.Default()
	g.POST("/api", middle, use)
	g.POST("/api2", middle2, use2, use3)
	g.Run(":8080")
}

func middle(ctx *gin.Context) {
	var info map[string]string
	if err := ctx.ShouldBindBodyWith(&info, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.Next()
}
func use(ctx *gin.Context) {
	var info map[string]string
	if err := ctx.ShouldBindBodyWith(&info, binding.JSON); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, info)
}
func middle2(ctx *gin.Context) {
	data, err := ctx.GetRawData()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("data: %v\n", string(data))

	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点
	ctx.Next()
}
func use2(ctx *gin.Context) {
	var info map[string]string
	if err := ctx.BindJSON(&info); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, info)
	ctx.Next()
}
func use3(ctx *gin.Context) {
	var info map[string]string
	if err := ctx.BindJSON(&info); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	ctx.JSON(http.StatusOK, info)
}
