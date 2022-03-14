package main

import (
	"fmt"
)

func main() {
	var a = "123"
	b := []byte(a)
	fmt.Println(b)
	b[0] = 95
	fmt.Println(string(b))
}
