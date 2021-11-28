package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go func() {
		for i := 0; i < 3; i++ {
			ch <- 1

			time.Sleep(time.Second)
		}
		close(ch)
	}()
	for {
		select {
		case <-ch:
			fmt.Println(<-ch)
			if <-ch == 0 {
				goto jump
			}
		}
	}
jump:
	fmt.Println("over")
}
