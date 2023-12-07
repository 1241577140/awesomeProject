package main

import (
	"github.com/jmoiron/sqlx"
	"log"
)

type Per struct {
	Uid  int    `gorm:"column:id"`   // 只有 主键是 id 才能返回
	Name string `gorm:"column:name"` // 这里需要 加上 column 字段才能匹配到数据库
	Pwd  string `gorm:"column:pwd"`
}

func main() {
	dsn := "root:123456@tcp(127.0.0.1:3306)/my_test?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	p := Per{
		Name: "aaa",
		Pwd:  "111",
	}
	exec, err := db.Exec("insert into t_user(name,pwd) values (?,?)", p.Name, p.Pwd)
	exec.LastInsertId()
}
