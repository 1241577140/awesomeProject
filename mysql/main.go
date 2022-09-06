package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:123456@tcp(192.144.231.147:3306)/system?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {

	}
	db.Debug()
	var a = uint8(8)
	var b = uint8(10)
	fmt.Println(a - b)
}
