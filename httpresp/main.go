package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
)

func main() {
	//httpWithoutCloseButRead()
	//time.Sleep(1 * time.Minute)
	//fmt.Println("过了一会 ", runtime.NumGoroutine())
	iota1()
}

func httpWithoutCloseButRead() {
	for i := 0; i < 20; i++ {
		resp, err := http.Get("http://www.baidu.com")
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = ioutil.ReadAll(resp.Body)
	}
	fmt.Println("goroutine num is ", runtime.NumGoroutine())
}

const (
	i1 = iota
	i2 = iota
	i3 = iota
)

func iota1() {
	fmt.Println(i1, i2, i3)
}
