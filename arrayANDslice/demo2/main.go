package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := make([]int, 2)
	slice = append(slice, 1)
	fmt.Println(slice, len(slice), cap(slice))

	s1 := make([]int8, 0, 2)
	sliceHeader := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	fmt.Println(sliceHeader)

}
