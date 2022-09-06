package main

import (
	"fmt"
	"time"
)

func main() {
	SafeGo(func() {
		onPanic()
	})
	time.Sleep(time.Second)
}

func SafeGo(f func()) {
	defer onRecover()
	func() {
		f()
	}()
}

func onPanic() {
	panic("panic")
}

func onRecover() {
	if r := recover(); r != nil {
		fmt.Println(r)
	}
}
