package main

import (
	"fmt"
	"reflect"
	"unsafe"
)

func main() {
	slice := []int8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	s1 := slice[2:5]
	s2 := s1[2:7]
	fmt.Printf("len=%-4d cap=%-4d slice=%-1v \n", len(slice), cap(slice), slice)
	fmt.Printf("len=%-4d cap=%-4d s1=%-1v \n", len(s1), cap(s1), s1)
	fmt.Printf("len=%-4d cap=%-4d s2=%-1v \n", len(s2), cap(s2), s2)

	s2 = append(s2, 100)
	fmt.Println("-------s2 append 100-------")
	fmt.Printf("len=%-4d cap=%-4d slice=%-1v \n", len(slice), cap(slice), slice)
	fmt.Printf("len=%-4d cap=%-4d s1=%-1v \n", len(s1), cap(s1), s1)
	fmt.Printf("len=%-4d cap=%-4d s2=%-1v \n", len(s2), cap(s2), s2)

	sliceHeader0 := (*reflect.SliceHeader)(unsafe.Pointer(&slice))
	sliceHeader1 := (*reflect.SliceHeader)(unsafe.Pointer(&s1))
	sliceHeader2 := (*reflect.SliceHeader)(unsafe.Pointer(&s2))
	fmt.Println(sliceHeader0)
	fmt.Println(sliceHeader1)
	fmt.Println(sliceHeader2)

	s2 = append(s2, 101)
	fmt.Printf("len=%-4d cap=%-4d slice=%-1v \n", len(slice), cap(slice), slice)
	fmt.Printf("len=%-4d cap=%-4d s1=%-1v \n", len(s1), cap(s1), s1)
	fmt.Printf("len=%-4d cap=%-4d s2=%-1v \n", len(s2), cap(s2), s2)
}
