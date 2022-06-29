package main

import (
	"fmt"
	"sync"
)

// 起300 个goroutine 按顺序打印 cat dog fish
func main() {
	//p2()
	//fmt.Println(runtime.NumCPU())

}
func p1() {
	dogCh := make(chan struct{})
	catCh := make(chan struct{})
	fishCh := make(chan struct{})
	for i := 0; i < 100; i++ {
		go cat(catCh, dogCh)
		go dog(dogCh, fishCh)
		go fish(fishCh, catCh)
		catCh <- struct{}{}
		select {
		case <-catCh:
			continue
		}
	}
}

// 用channel进行通信
func p2() {
	var wg sync.WaitGroup
	wg.Add(100)
	catCh := make(chan struct{}, 1) // 要有缓存队列
	dogCh := make(chan struct{}, 1)
	fishCh := make(chan struct{}, 1)
	catCh <- struct{}{} // 否则到这里就停止了
	for i := 0; i < 100; i++ {
		go func() {
			cat(catCh, dogCh)
		}()
		go func() {
			dog(dogCh, fishCh)
		}()
		go func() {
			fish(fishCh, catCh)
			wg.Done()
		}()
	}
	wg.Wait()
}

func cat(cat <-chan struct{}, dog chan<- struct{}) {
	<-cat
	fmt.Println("cat")
	dog <- struct{}{}

}
func dog(dog <-chan struct{}, fish chan<- struct{}) {
	<-dog
	fmt.Println("dog")
	fish <- struct{}{}

}
func fish(fish <-chan struct{}, cat chan<- struct{}) {
	<-fish
	fmt.Println("fish")
	cat <- struct{}{}
}
