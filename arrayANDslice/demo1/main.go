package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	array1 := [2]int{1, 2}
	array2 := array1
	array2[1] = 1
	fmt.Println("array1 kind:", reflect.TypeOf(array1).Kind(), ":", array1)
	fmt.Println("array1 kind:", reflect.TypeOf(array2).Kind(), ":", array2)
	slice1 := []int{1, 2}
	slice2 := slice1
	slice2[1] = 1
	fmt.Println("slice1 kind:", reflect.TypeOf(slice1).Kind(), ":", slice1)
	fmt.Println("slice2 kind:", reflect.TypeOf(slice2).Kind(), ":", slice2)

	fmt.Println("array1 size:", unsafe.Sizeof(array1))
	fmt.Println("slice1 size:", unsafe.Sizeof(slice1))

}
