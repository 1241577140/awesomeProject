package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan bool)
	close(ch)
	fmt.Println(<-ch)
	pool := sync.Pool{New: func() interface{} {
		return "dvs"
	}}
	for i := 0; i < 10; i++ {
		v := pool.Get()
		fmt.Println(v)
		pool.Put(i)
	}
}
