package main

import (
	"fmt"
	"time"
)

/*
假设有一个超长的切片，切片的元素类型为int，切片中的元素为乱序排列。限时5秒，使用多个goroutine查找切片中是否存在给确定的价值，在查找目标价值或超过时间后立刻结束所有goroutine的执行。
比方说，切片[23,32,78,43,76,65,345,762,......915,86]，查找目标值为345，如果切片中存在，则目标值输出"Found it!"并立即取消仍在执行查询任务的goroutine。
如果在超时时间未查到目标值程序，则输出"Timeout！Not Found"，同时立即取消仍在执行的查询任务的goroutine。
*/
func main() {
	search([]int{23, 32, 78, 43, 76, 65, 345, 762}, 23)
}

func search(arr []int, target int) {
	ch := make(chan bool)
	timeOut := make(chan bool)
	step := len(arr) / 2
	time.AfterFunc(time.Second*5, func() {
		fmt.Println("Timeout！Not Found")
		close(timeOut)
	})
	for i := 0; i < 2; i++ {
		go func(i int) {
			for {
				select {
				case <-ch:
					return
				case <-timeOut:
					return
				default:
					temp := arr[i*step : (i+1)*step]
					find := false
					for _, v := range temp {
						if v == target {
							find = true
							fmt.Println("Found it!")
							close(ch)
						}
					}
					if !find {
						return
					}
				}
			}
		}(i)
	}
	time.Sleep(time.Second * 6)
}
