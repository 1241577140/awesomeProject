package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var mm = map[string]interface{}{}

	err := json.Unmarshal([]byte(`{"name":"1","val":0}`), &mm)
	if err != nil {
		return
	}
	for _, v := range mm {
		kind := reflect.TypeOf(v).Kind()
		fmt.Println(kind)
		ni := IsDefault(v)
		fmt.Println(ni)
	}
}

func IsDefault(in interface{}) bool {
	defer func() {
		recover()
	}()
	return reflect.ValueOf(in).IsZero()
}
