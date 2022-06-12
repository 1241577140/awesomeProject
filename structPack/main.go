package main

import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string
}

func main() {
	var p = Person{"a"}
	fmt.Println(reflect.TypeOf(&p))
	fmt.Println(&p)
	for i := 0; i < 100000000; i++ {
		var a = float64(1)
		if int(a) != 1 {
			fmt.Println("false")
		}
	}
}
