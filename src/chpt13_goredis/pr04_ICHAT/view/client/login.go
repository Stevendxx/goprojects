package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
)

func Login(name, pwd *string) (err error) {
	fmt.Print("name -> ")
	fmt.Scanln(name)
	fmt.Print("pwd  -> ")
	fmt.Scanln(pwd)
	err = LoginVerfiy(*name, *pwd)
	if err == nil {
		fmt.Println("登录成功...")
	} else {
		fmt.Println("登录失败...")
		fmt.Println(err)
	}
	return
}

func LoginVerfiy(name, pwd string) (err error) {
	err = nil
	// 1. 请求连接及延时关闭
	network := "tcp"
	address := "localhost:8991"
	conn, errDial := net.Dial(network, address)
	if errDial != nil {
		fmt.Println("[错误]:请求连接错误,请检查网络方式和主机IP及其端口.")
		fmt.Println(errDial)
		err = errDial
		return
	}
	defer conn.Close()

	// 2. 登录信息
	var loginMsg = model.LoginMessage{
		UserId:   0,
		UserName: name,
		UserPwd:  pwd,
	}

	// 3. 登录信息序列化
	data, errJson := json.Marshal(loginMsg)
	if errJson != nil {
		fmt.Println("[错误]: 登录信息序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	// 4. 通过序列化后的登录信息构建报文
	var msg = model.Message{
		Type:   model.LoginMsgType,
		Data:   string(data),
		Length: len(data),
	}

	// 5. 发送报文至服务器
	data, errJson = json.Marshal(msg)
	if errJson != nil {
		fmt.Println("[错误]:发送报文序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	_, errWrtie := conn.Write([]byte(data))
	if errWrtie != nil {
		fmt.Println("[错误]:发送报文出现错误.")
		fmt.Println(errWrtie)
		err = errWrtie
		return
	}

	// 读取服务器返回报文
	_, err = readMsg(conn)
	return

}

// 读取服务器返回报文
func readMsg(conn net.Conn) (msg model.Message, err error) {
	err = nil

	var b = make([]byte, 10240)
	cnt, errRead := conn.Read(b)
	if errRead != nil {
		fmt.Println("[错误]:连接读取错误.")
		fmt.Println(errRead)
		err = errRead
		return
	}

	errUnjson := json.Unmarshal(b[:cnt], &msg)
	if errUnjson != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(errUnjson)
		err = errUnjson
		return
	}

	err = HandleMsg(conn, msg)
	return
}

// 处理报文
func HandleMsg(conn net.Conn, msg model.Message) (err error) {

	switch msg.Type {
	case model.LoginCodeType:
		// 处理的登录
		err = HandleReturnMsg(conn, msg)
	}
	return
}

// 处理登录报文
func HandleReturnMsg(conn net.Conn, msg model.Message) (err error) {

	var loginCode model.LoginCode
	err = json.Unmarshal([]byte(msg.Data), &loginCode)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}

	if loginCode.Code != 200 {
		err = errors.New(loginCode.Message)
		return
	}
	return
}
