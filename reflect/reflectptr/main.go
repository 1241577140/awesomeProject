package main

import (
	"fmt"
	"reflect"
)

func main() {
	strVar := "aab"
	t := reflect.TypeOf(&strVar)
	kind := t.Elem().Kind()
	fmt.Println(kind)
}
