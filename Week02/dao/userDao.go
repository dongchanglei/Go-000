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
