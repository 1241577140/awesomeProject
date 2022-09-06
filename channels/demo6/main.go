package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	after := strings.TrimLeft("   +0 123", " ")
	fmt.Println(after)
	//print1()
	//print2()
}

// 依次打印猫狗鱼100次 三个携程

func print1() {
	cat := make(chan struct{})
	dog := make(chan struct{})
	fish := make(chan struct{})
	sum := 100
	go func() {
		for i := 0; i < sum; i++ {
			<-cat
			fmt.Println("cat")
			dog <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < sum; i++ {
			<-dog
			fmt.Println("dog")
			fish <- struct{}{}
		}
	}()
	go func() {
		for i := 0; i < sum; i++ {
			<-fish
			fmt.Println("fish")
			cat <- struct{}{}
		}
	}()
	cat <- struct{}{}
	time.Sleep(time.Second * 2)

}

// 依次打印猫狗鱼100次 三个携程
func print2() {
	cat := make(chan struct{})
	dog := make(chan struct{})
	fish := make(chan struct{})
	sum := 100
	for i := 0; i < sum; i++ {
		go func() {

			<-cat
			fmt.Println("cat")
			dog <- struct{}{}
		}()
	}
	for i := 0; i < sum; i++ {
		go func() {
			<-dog
			fmt.Println("dog")
			fish <- struct{}{}
		}()
	}
	for i := 0; i < sum; i++ {
		go func() {
			<-fish
			fmt.Println("fish")
			cat <- struct{}{}

		}()
	}
	cat <- struct{}{}
	time.Sleep(time.Second * 2)

}
