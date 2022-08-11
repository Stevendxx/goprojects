package service

import (
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/dao"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
)

// 判断 user 是否存在
func UserIsExist(name string) (isExist bool, err error) {
	return dao.UserIsExist(dao.ConnPool.Get(), name)
}

// 获取 user
func GetUserByName(name string) (user *model.User, err error) {
	return dao.GetUserByName(dao.ConnPool.Get(), name)
}

// 注册 user
func AddUser(user model.User) (err error) {
	return dao.SetUser(dao.ConnPool.Get(), user)
}
