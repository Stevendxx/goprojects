package main

import (
	"errors"
	"fmt"
	"time"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/view/client/process"
)

// 登录界面
func LoginMenu() (err error) {
	var name, pwd string
	fmt.Println("登录中...")
	fmt.Print("*name -> ")
	fmt.Scanln(&name)
	fmt.Print("*pwd  -> ")
	fmt.Scanln(&pwd)

	// 登录验证
	conn, msg, errVerfiy := process.LoginVerfiy(name, pwd)
	if errVerfiy != nil {
		fmt.Println("[错误]:登录信息验证发生错误.")
		fmt.Println(errVerfiy)
		err = errVerfiy
		return
	}
	defer conn.Close()

	err = HandleMsg(msg)
	if err != nil {
		fmt.Println("登录失败...")
		fmt.Println(err)
		return
	}
	fmt.Println("登录成功...")
	time.Sleep(time.Second * 3)
	utils.ClearScreen()

	process.UserMain(conn, name, pwd)
	utils.ClearScreen()
	return
}

// 注册界面
func SignUpMenu() (err error) {
	var user = &model.User{}
	fmt.Println("注册中...")
	fmt.Print("*name  -> ")
	fmt.Scanln(&user.Name)
	fmt.Print("*pwd   -> ")
	fmt.Scanln(&user.Pwd)
	fmt.Print("age    -> ")
	fmt.Scanln(&user.Age)
	fmt.Print("gender -> ")
	fmt.Scanln(&user.Gender)
	fmt.Print("eamil  -> ")
	fmt.Scanln(&user.Email)

	//
	msg, errVerfiy := process.SignUpVerfiy(*user)
	if errVerfiy != nil {
		fmt.Println("[错误]:登录信息验证发生错误.")
		err = errVerfiy
		return
	}

	err = HandleMsg(msg)
	if err != nil {
		fmt.Println("注册失败...")
		fmt.Println(err)
		return
	}
	time.Sleep(time.Second * 3)
	utils.ClearScreen()
	fmt.Println("注册成功, 请登录...")

	return
}

// 处理报文
func HandleMsg(msg model.Message) (err error) {

	switch msg.Type {
	case model.LoginCodeType:
		// 处理的登录
		err = process.HandleLoginReturnMsg(msg)
	case model.SignUpCodeType:
		// 处理注册
		err = process.HandleSignUpReturnMsg(msg)
	default:
		err = errors.New("[错误]:报文类型错误, " + msg.Type + "不存在.")
	}
	return
}
