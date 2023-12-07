package main

import "fmt"

func main() {
	l := &List{
		Next: &List{
			Next: &List{
				Next: nil,
				Val:  3,
			},
			Val: 2,
		},
		Val: 1,
	}
	l = reverse(l)
	fmt.Println(l)
}

type List struct {
	Next *List
	Val  interface{}
}

func reverse(node *List) *List {
	head := node
	pre := new(List)
	for node.Next != nil {
		pre = node
		node = node.Next
		node.Next = pre
	}
	pre.Next = head
	return pre
}
