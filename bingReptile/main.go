package main

import (
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	get, err := http.Get("https://bing.ioliu.cn/photo/MistyTor_ZH-CN7520952555?force=home_18")
	if err != nil {
		return
	}
	all, err := ioutil.ReadAll(get.Body)
	if err != nil {
		return
	}
	err = ioutil.WriteFile("a.jpg", all, os.ModePerm)
	if err != nil {
		return
	}
}
