package main

import (
	"fmt"
	"reflect"
)

type Per struct {
	Say string
}

func main() {
	var p = new(Per)
	elem := reflect.TypeOf(p).Elem()
	fmt.Println(elem)
	fmt.Println(elem.Field(0).Tag)
}
