package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println(test())
}

func test() error {
	var err error
	defer func() {
		if r := recover(); r != nil {
			err = errors.New("asa")
		}
	}()
	pac()
	return err
}
func pac() {
	panic("faf")
}
