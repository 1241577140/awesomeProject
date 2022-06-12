package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Person struct {
	Name string
}

func main() {
	var x interface{} = nil
	var y *int = nil
	interfaceIsNil(x)
	interfaceIsNil(y)
	断言()
}

func interfaceIsNil(x interface{}) {
	if x == nil {
		fmt.Println("empty interface")
		return
	}
	fmt.Println("non-empty interface")
}

func 断言() {
	var a = map[string]interface{}{}

	err := json.Unmarshal([]byte(`{"aa":{"Name":"s"}}`), &a)
	if err != nil {
		log.Println(err)
		return
	}
	person, ok := a["aa"].(*Person)
	fmt.Println(ok, person)
}
