package main

import (
	mysql_util "Go-000/Week02/dbUtil/mysql"
	"fmt"
)

func main() {
	db, err := mysql_util.NewMysqlConn("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		fmt.Println(err)
		return
	}
	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("连接成功")
}
