package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	var slice = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("slice", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))
	slice = append(slice[:3], slice[4:]...)
	fmt.Println("slice", (*reflect.SliceHeader)(unsafe.Pointer(&slice)))

	var slice1 = []int{}
	var slice2 []int
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
	fmt.Println(len(slice1) == 0)
	fmt.Println(len(slice2) == 0)
}
