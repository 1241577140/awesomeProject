package main

import (
	"fmt"
	"github.com/samber/lo"
	"gorm.io/gorm/utils"
	"strconv"
)

type Adder interface {
	~int | ~string
}

func add[T Adder](collection ...T) {
	str := ""
	for _, t := range collection {
		str += utils.ToString(t)
	}
	fmt.Println(str)
}
func main() {
	res := lo.Map[int64, string]([]int64{1, 2, 3}, func(x int64, _ int) string {
		return strconv.FormatInt(x, 10)
	})
	fmt.Println(res)
	aa, bo := lo.Find[any]([]interface{}{1, 3, "2", "boo"}, func(t interface{}) bool {
		if t == "bool" {
			return true
		}
		return false
	})
	fmt.Println(aa, "--", bo)
	add("1", "2", "3", "123")
}
