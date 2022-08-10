package main

import (
	"encoding/json"
	"fmt"
	"net"
	"os"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
)

func main() {

	network := "tcp"
	address := "localhost:8991"
	fmt.Println("[服务器]:开启8991端口监听...")
	listener, errListen := net.Listen(network, address)
	if errListen != nil {
		fmt.Println("[错误]:开启监听错误,请检查网络方式和主机IP及端口.")
		fmt.Println(errListen)
		os.Exit(1)
	}
	defer listener.Close()

	for {
		conn, errAccept := listener.Accept()
		if errAccept != nil {
			fmt.Println("[错误]:接受请求错误.")
			fmt.Println(errAccept)
			continue
		}

		go handleConn(conn)
	}
}

// 客户端处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("[服务器]:客户端<%v>已连接...\n", conn.RemoteAddr())
	go readMsg(conn)

	for {
		// 10KB
	}
}

// 读取报文
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
	case model.LoginMsgType:
		// 处理的登录
		err = HandleLoginMsg(conn, msg)
	}
	return
}

// 处理登录报文
func HandleLoginMsg(conn net.Conn, msg model.Message) (err error) {

	var loginMsg model.LoginMessage
	err = json.Unmarshal([]byte(msg.Data), &loginMsg)
	if err != nil {
		fmt.Println("[错误]:报文解析错误.")
		fmt.Println(err)
		return
	}

	var loginCode model.LoginCode
	// root 123456
	if loginMsg.UserName == "root" && loginMsg.UserPwd == "123456" {
		loginCode.Code = 200
		loginCode.Message = "ok"
	} else {
		loginCode.Code = 500
		loginCode.Message = "no exist"
	}

	data, errJson := json.Marshal(loginCode)
	if errJson != nil {
		fmt.Println("[错误]: 返回代码序列化错误.")
		fmt.Println(errJson)
		err = errJson
		return
	}

	var returnMsg = model.Message{
		Type:   model.LoginCodeType,
		Data:   string(data),
		Length: len(data),
	}

	data, errJson = json.Marshal(returnMsg)
	if errJson != nil {
		fmt.Println("[错误]: 返回报文序列化错误.")
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
	return
}
