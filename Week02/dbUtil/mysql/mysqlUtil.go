package mysql_util

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

//driveName:mysql;dataSourceName:root:123456@tcp(127.0.0.1:3306)/test?charset=utf8
func NewMysqlConn(driveName string, dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driveName, dataSourceName)
	if err != nil {
		return db, errors.New("连接数据库异常！")
	}
	return db, err
}
