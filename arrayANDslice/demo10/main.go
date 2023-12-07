package main

import (
	"fmt"
)

func main() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
	}
	fmt.Println(a)
	main2()
	arr := []int{1, 2, 3}
	target := 2
	index := search(arr, target)
	fmt.Println(index)
}
func main2() {
	a := [3]int{1, 2, 3}
	for k, v := range a {
		if k == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[k] = 100 + v
	}
	fmt.Println(a)
}

func search(arr []int, target int) int {

	left := 0
	right := len(arr) - 1
	for left < right {
		mid := (left + right) / 2 // left + (right-left)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}
