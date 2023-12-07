package main

import (
	"fmt"
)

func main() {
	var a interface{}
	var b interface{}
	fmt.Println(a == b)

	var tm = TM{}
	tm.tma()
	tm.tmb()
	fc := FC(tm)
	_ = fc
	var e *int = nil
	fmt.Println(e == nil)
	fmt.Println(e == a)

}

type TM struct {
	A string
}

type FC struct {
	A string
}

func (t TM) tma() {
	fmt.Println("tma")
}

func (t *TM) tmb() {
	fmt.Println("tmb")
}
