package main

import "fmt"

func main() {
	mm := make(map[int]int, 8)
	for i := 0; i < 5; i++ {
		mm[i] = i
	}
	fmt.Println(len(mm))
	mm[5] = 5
	fmt.Println(len(mm))
	mm[6] = 6
	fmt.Println(len(mm))
	mm[7] = 7
	fmt.Println(len(mm))
	for k := range mm {
		delete(mm, k)
	}
	fmt.Println(len(mm))
}
