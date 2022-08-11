package model

import (
	"errors"
)

var (
	ErrorUserNameNil    = errors.New("用户名为空")
	ErrorUserPwdNil     = errors.New("用户密码为空")
	ErrorUserNamePwdNil = errors.New("用户名和密码都为空")
	ErrorUserNotExists  = errors.New("用户不存在")
	ErrorUserExists     = errors.New("用户已存在")
	ErrorUserPwd        = errors.New("用户密码错误")
	ErrorUserdbNil      = errors.New("数据不存在")
)
