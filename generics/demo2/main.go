package main

import (
	"fmt"
	"math/rand"
)

type GenInface[T any] interface {
}

func main() {
	a := []int{1, 2, 3}
	fmt.Printf("%p,\n", &a)
	a = append([]int{1}, a...)
	fmt.Printf("%p,\n", &a)
	rand.Shuffle(1, func(i, j int) {

	})
}
