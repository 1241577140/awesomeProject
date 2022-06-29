package main

import (
	"fmt"
	"sync"
)

func main() {
	str := "hello,world"

	ch := make(chan struct{})
	wg := sync.WaitGroup{}
	list := make(chan rune, len(str))
	for _, b := range str {
		list <- b
	}
	close(list)
	wg.Add(2)
	go func() {
		defer wg.Done()
		for {
			ball, ok := <-ch
			if ok {
				s, ok := <-list
				if ok {
					fmt.Println(string(s))
				} else {
					close(ch)
					return
				}
				ch <- ball
			} else {
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		for {
			ball, ok := <-ch
			if ok {
				s, ok := <-list
				if ok {
					fmt.Println(string(s))
				} else {
					close(ch)
					return
				}
				ch <- ball
			} else {
				return
			}
		}
	}()
	ch <- struct{}{}
	wg.Wait()

	chaa := make(chan int, 4)
	for i := 0; i < 4; i++ {
		chaa <- i
	}
	close(chaa)
	for i := 0; i < 19; i++ {
		g, ok := <-chaa
		if !ok {
			fmt.Println("123132")
		}
		fmt.Println(g)
	}
}
