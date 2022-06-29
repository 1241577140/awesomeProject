package main

import "fmt"

func main() {
	t1()
	t2()
}

func t1() {
	a := []int{1, 2, 3}
	for i, v := range a {
		if i == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[i] = 100 + v
	}
	fmt.Println(a)
}
func t2() {
	a := [3]int{1, 2, 3}
	for i, v := range a {
		if i == 0 {
			a[0], a[1] = 100, 200
			fmt.Println(a)
		}
		a[i] = 100 + v
	}
	fmt.Println(a)
}
