package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice1 := make([]int, 1, 2)
	fmt.Println("slice1  append前   len ", len(slice1), "cap", cap(slice1), "数组地址", (*reflect.SliceHeader)(unsafe.Pointer(&slice1)).Data)
	// 正常append 未超出容量
	slice1 = append(slice1, 2)
	// append后地址
	fmt.Println("slice1 append一个后 len ", len(slice1), "cap:", cap(slice1), "数组地址", (*reflect.SliceHeader)(unsafe.Pointer(&slice1)).Data)
	slice1 = append(slice1, 3, 4, 5)
	fmt.Println("slice1 扩容后 cap:", cap(slice1), "数组地址", (*reflect.SliceHeader)(unsafe.Pointer(&slice1)).Data)

	slice2 := make([]int, 1024)
	slice2 = append(slice2, 1)
	fmt.Println("slice2 扩容超过1024 len:", len(slice2), "cap:", cap(slice2))

}
