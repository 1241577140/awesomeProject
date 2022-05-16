package main

import (
	"encoding/json"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	compile, err := regexp.Compile("http://www.zhenai.com/zhenghun/([a-z]+)/([a-z]+)")
	if err != nil {
		return
	}
	strs := compile.FindStringSubmatch("http://www.zhenai.com/zhenghun/aaa/bbb")
	for _, s := range strs {
		fmt.Println(s)
	}

	var str = "s1/s2/s3.json"
	fmt.Println(strings.TrimPrefix(str, "/"))
	var s = "1,2,3"
	sl := make([]int, 0)
	err = json.Unmarshal([]byte(s), &sl)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(s)
}
