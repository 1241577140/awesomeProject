package main

import (
	"fmt"

	"github.com/gogf/gf/v2/container/gmap"
)

type name struct {
	size string
	h    string
}

func main() {
	hashMap := gmap.HashMap{}
	fmt.Println(hashMap)
	var aa = make(map[interface{}]interface{})
	var bb []interface{}
	bb = append(bb, 1)
	var na = name{size: "12"}
	bb = append(bb, na)
	aa[1] = 2
	aa[na] = 4
	fmt.Println(aa, bb)
}
