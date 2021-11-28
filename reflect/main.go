package main

import (
	"fmt"
	"reflect"
)

func printMsg() string {
	return " qqq"
}

type VV interface {
	QQ()
}
type A struct {
	aa int
}

func (a *A) QQ() {

}
func main() {
	a := &A{aa: 1}
	fmt.Println(a == nil)
	b := new(*A)
	s1 := []*A{a}
	s2 := make([]*A, 3)
	copy(s2, s1)
	// Dst and src each must have kind Slice or Array, and
	// dst and src must have the same element type.
	reflect.Copy(reflect.ValueOf(b), reflect.ValueOf(a))
	a.aa = 0
	fmt.Printf("%d,%d\n", s1[0], s2[0])
	fmt.Printf("%d,%d\n", b, a)
	fmt.Println(reflect.TypeOf(printMsg))
	str := reflect.ValueOf(printMsg).Call([]reflect.Value{})
	fmt.Println(str[0])
}
