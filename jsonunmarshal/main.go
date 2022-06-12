package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/structs"
)

type Per struct {
	Name string
	Age  int
}

func main() {
	var per = Per{
		Name: "aaa",
		Age:  14,
	}
	m := structs.Map(&per)
	for s, i := range m {
		fmt.Printf("%T %T \n", s, i)
	}
	mapD := make(map[string]interface{})
	marshal, _ := json.Marshal(&per)
	json.Unmarshal(marshal, &mapD)
	for s, i := range mapD {
		fmt.Printf("%T %T \n", s, i)
	}
}
