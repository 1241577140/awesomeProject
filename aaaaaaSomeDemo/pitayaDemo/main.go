package main

import (
	"fmt"
	"reflect"
)

type AAA interface {
	fmt.Stringer
}
type BBB struct {
}

func (b *BBB) String() string {
	return ""
}
func main() {
	var a AAA
	var b *BBB
	fmt.Println(b == nil)
	fmt.Println(a == nil)
	switch a.(type) {
	default:
		fmt.Println("neng")
	}
	fmt.Println(reflect.TypeOf(a).String())
}
