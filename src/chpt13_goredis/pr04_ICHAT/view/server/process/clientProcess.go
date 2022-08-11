package process

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/service"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
)

// 客户端请求

// 处理登录报文
func HandleLoginMsg(conn net.Conn, msg model.Message) (err error) {

	var loginMsg model.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}
	// root 123456
	// 登录信息验证
	loginCode, _ := LoginMsgVerfiy(loginMsg)

	returnMsg, err := utils.CreateMsg(model.LoginCodeType, loginCode)
	if err != nil {
		return
	}

	// 写
	errWrite := utils.WrtieMsg(conn, returnMsg)
	if errWrite != nil {
		err = errWrite
	}
	return
}

// 登陆报文验证
func LoginMsgVerfiy(loginMsg model.LoginMessage) (loginCode model.LoginCode, err error) {

	user, errGet := service.GetUserByName(loginMsg.UserName)
	if errGet != nil {
		fmt.Printf("[错误]:用户 %s 获取失败.\n", loginMsg.UserName)
		fmt.Println(errGet)
		loginCode.Code = 500
		loginCode.Message = "User does not exist."
		err = model.ErrorUserNotExists
		return
	}

	if loginMsg.UserName == user.Name && loginMsg.UserPwd == user.Pwd {
		loginCode.Code = 200
		loginCode.Message = "ok"
	} else {
		loginCode.Code = 123
		loginCode.Message = "Password is incorrect."
		err = model.ErrorUserPwd
	}

	return
}

// 处理登录报文
func HandleSginUpMsg(conn net.Conn, msg model.Message) (err error) {

	var signUpMsg model.SignUpMessage
	err = json.Unmarshal([]byte(msg.Data), &signUpMsg)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}

	// 登录信息验证
	signUpCode, _ := SignUpMsgVerfiy(signUpMsg)

	returnMsg, err := utils.CreateMsg(model.SignUpCodeType, signUpCode)
	if err != nil {
		return
	}

	// 写
	errWrite := utils.WrtieMsg(conn, returnMsg)
	if errWrite != nil {
		err = errWrite
	}
	return
}

// 注册报文验证
func SignUpMsgVerfiy(signUpMsg model.SignUpMessage) (signUpCode model.SignUpCode, err error) {
	isExist, _ := service.UserIsExist(signUpMsg.SignUpUser.Name)

	if isExist {
		signUpCode.Code = 500
		signUpCode.Message = "User already exists."
		err = model.ErrorUserExists
	} else {
		err = service.AddUser(signUpMsg.SignUpUser)
		if err != nil {
			signUpCode.Code = 400
			signUpCode.Message = "User write-save failed."
			return
		}
		signUpCode.Code = 200
		signUpCode.Message = "ok"
	}

	return
}
