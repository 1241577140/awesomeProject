package main

import (
	"fmt"
	"regexp"
	"runtime"
)

func main() {
	compile, err := regexp.Compile("http://www.zhenai.com/zhenghun/([a-z]+)/([a-z]+)")
	if err != nil {
		return
	}
	strings := compile.FindStringSubmatch("http://www.zhenai.com/zhenghun/aaa/bbb")
	for _, s := range strings {
		fmt.Println(s)
	}
	runtime.GC()
}
