package main

import (
	"fmt"
	"github.com/prometheus/common/log"
	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

type Per struct {
	Uid  int    `gorm:"column:id"`   // 只有 主键是 id 才能返回
	Name string `gorm:"column:name"` // 这里需要 加上 column 字段才能匹配到数据库
	Pwd  string `gorm:"column:pwd"`
}

func (p Per) TableName() string {
	return "t_user"
}
func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	p := Per{
		Name: "11",
		Pwd:  "zz",
	}
	if db.Debug().Create(&p).Error != nil {
		log.Errorln(p)
	}
	fmt.Println(p.Uid) // 只有 主键字段名是 id 才能返回
	fmt.Println(p.Name)
}
