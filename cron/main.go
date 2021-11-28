package main

import (
	"fmt"
	"github.com/robfig/cron/v3"
)

func main() {
	c := cron.New(cron.WithSeconds()) //精确到秒

	spec1 := "05 30 20 * * 0" //cron表达式，每5秒一次
	c.AddFunc(spec1, func() {
		fmt.Println("22222")
	})
	c.Start()
	select {} //阻塞主线程停止
}
