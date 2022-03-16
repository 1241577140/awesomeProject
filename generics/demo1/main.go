package main

import "fmt"

type Demo1 struct {
	dem1 int
}
type Demo2 struct {
	demo2 int
}

func (d Demo1) Print() {
	fmt.Println(1)
}
func (d Demo2) Print() {
	fmt.Println(2)
}

type Generic interface {
	Demo1 | Demo2
	Print()
}

func Gen[T Generic](a T) {
	fmt.Println(a)
}

func main() {
	var d1 = Demo1{1}
	var d2 = Demo2{2}
	Gen(d1)
	Gen(d2)
}

type Add interface {
	int | int8 | int16 | int32 | int64
}

func add[T Add](a ...T) int {
	amount := 0
	for _, v := range a {
		amount += int(v)
	}
	return amount
}

//  以下的语法现阶段还不支持，make  方法无法识别到类型
//  for range 方法同上
/*
func MakeChan[T chan bool| chan int](c T){
	_=make(T)
	_=new(T)
	_=len(c)
}
func ForEach[T []string | map[int]string](c T, f func(int, string)) {
	for i, v := range c {
	 f(i, v)
	}
   } */
