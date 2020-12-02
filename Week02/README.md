err wrap 实现：

```java

package mysql_util

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql"
)

func NewMysqlConn(driveName string, dataSourceName string) (db *sql.DB, err error) {
	db, err = sql.Open(driveName, dataSourceName)
	if err != nil {
		return db, errors.New("连接数据库异常！")
	}
	return db, err
}

```

```java
package entity

type User struct {}
```

```java
package dao

import (
	mysql_util "Go-000/Week02/dbUtil/mysql"
	"Go-000/Week02/entity"
	"github.com/pkg/errors"
)

type UserDao struct{}

func (d *UserDao) FindUserById(userId int64) (user *entity.User, err error) {
	conn, err := mysql_util.NewMysqlConn("mysql", "root:12345678@tcp(127.0.0.1:3306)/go_test?charset=utf8")
	if err != nil {
		return nil, errors.Wrap(err, "UserDao 获取MySQL数据连接异常")
	}
	defer func() {conn.Close()}()
	row := conn.QueryRow("select  * from t_user_go where id = ?", userId)
	if row.Err() != nil {
		return nil, errors.Wrap(row.Err(), "UserDao 根据用户ID获取数据异常")
	}
	err = row.Scan(user)
	if err != nil {
		return nil, errors.Wrap(row.Err(), "UserDao 根据用户ID获取数据解析异常")
	}
	return user, nil
}

```

```java

package service

import (
	"Go-000/Week02/dao"
	"Go-000/Week02/entity"
	"github.com/pkg/errors"
)

type UserService struct {}

func (userService *UserService) FindUserById(userId int64)(user *entity.User, err error) {
	userDao := dao.UserDao{}
	user, err = userDao.FindUserById(userId)
	if err != nil {
		return nil, errors.Wrap(err,"UserService 根据用户ID获取用户信息异常")
	}
	return user,nil
}

```