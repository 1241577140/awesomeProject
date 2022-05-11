package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.Default()
	g.POST("/api", ginFormTag)
	g.Run(":8080")
}

type Gin struct {
	User string `json:"user" form:"user" binding:"required"`
}

func ginFormTag(ctx *gin.Context) {
	g := &Gin{}
	if err := ctx.BindJSON(g); err != nil {
		return
	}
	ctx.JSON(http.StatusOK, g)
}
