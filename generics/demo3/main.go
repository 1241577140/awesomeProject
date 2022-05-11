package main

import (
	"fmt"
)

// 泛型语法规则：实际上就是需要将每个基础类型考虑进去
// 蜡印
type addable interface {
	~int | ~float64
}
type Number = int

type name struct {
}

func main() {
	fmt.Println(max(1.0, 2))
}
func add(a, b Number) int {
	return a + b
}
func max[T addable](a, b T) int {
	if a > b {
		return int(a)
	}
	return int(b)
}
