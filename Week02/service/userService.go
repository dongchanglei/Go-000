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
