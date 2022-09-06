package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < 26; i++ {
			<-ch
			fmt.Println(string('A' + i))
			ch <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < 26; i++ {
			<-ch
			fmt.Println(i + 1)
			ch <- struct{}{}
		}
	}()
	ch <- struct{}{}
	time.Sleep(time.Second)
}
