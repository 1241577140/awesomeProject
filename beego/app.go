package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context"
)

func main() {
	app := beego.NewApp()
	app.Handlers.Get("/api", func(c *context.Context) {

		c.WriteString("123")
	})
	app.Run()
}
