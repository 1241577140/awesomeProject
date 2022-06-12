package main

import "fmt"

func main() {
	var nu1 []interface{}
	nu2 := []int{1, 3, 4}
	nu3 := append(nu1, nu2)
	fmt.Println(len(nu3))
}
