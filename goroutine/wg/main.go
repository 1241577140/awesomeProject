package main

import "fmt"

func main() {
	var a = 1
	a = aaa(a)
	if a == 1 {
		return
	}
	defer fmt.Println(a)
}

func aaa(a int) int {
	return a
}
