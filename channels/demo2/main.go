package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	go func() {
		c()
	}()
	time.Sleep(time.Minute)
}

func c() {
	ch := make(chan int, 2)
	ch <- 1
	wg := sync.WaitGroup{}
	wg.Add(10)
	go func() {
		for i := 0; i < 10; i++ {
			ch <- 1
			wg.Done()
			time.Sleep(time.Second)
		}
	}()
	wg.Wait()
	close(ch)
	for v := range ch {
		fmt.Println(v)
	}
	fmt.Println("end")
}
