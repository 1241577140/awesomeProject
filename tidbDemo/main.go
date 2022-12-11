package main

import (
	"crypto/tls"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func main() {
	mysql.RegisterTLSConfig("tidb", &tls.Config{
		MinVersion: tls.VersionTLS12,
		ServerName: "gateway01.us-west-2.prod.aws.tidbcloud.com",
	})

	db, err := sql.Open("mysql", "3GgMXbuDSredxff.root:sunmeng0727@tcp(gateway01.us-west-2.prod.aws.tidbcloud.com:4000)/test?tls=tidb")
	if err != nil {
		panic(err)
	}
	err = db.Ping()
	fmt.Println(err)
}
