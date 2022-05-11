package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	currTest1()
	currTest2()
	lock()
}

func currTest1() {
	slice1 := make([]int, 0)
	start := time.Now().Nanosecond()
	gos := 1000000
	var wg sync.WaitGroup
	wg.Add(gos)
	for i := 0; i < gos; i++ {
		go func(i int) {
			slice1 = append(slice1, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(slice1))
	fmt.Println("时间", time.Now().Nanosecond()-start)
}

func currTest2() {
	slice2 := make([]int, 1)
	gos := 1000
	var wg2 sync.WaitGroup
	wg2.Add(gos)
	for i := 0; i < gos; i++ {
		go func() {
			slice2[0] += 1
			wg2.Done()
		}()
	}
	wg2.Wait()
	fmt.Println(slice2[0])
}

func lock() {
	slice1 := make([]int, 0)
	start := time.Now().Nanosecond()

	gos := 1000000
	var wg sync.WaitGroup
	var lock sync.Mutex
	wg.Add(gos)
	for i := 0; i < gos; i++ {
		go func(i int) {
			lock.Lock()
			defer lock.Unlock()
			slice1 = append(slice1, i)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(len(slice1))
	fmt.Println("时间", time.Now().Nanosecond()-start)

}
