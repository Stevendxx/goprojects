package main

import (
	"errors"
	"fmt"
	"io"
	"net"

	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/model"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/utils"
	"github.com/gomodule/redigo/redis/src/chpt13_goredis/pr04_ICHAT/view/server/process"
)

// 服务端处理器

// 处理连接
func handleConn(conn net.Conn) {
	defer conn.Close()
	fmt.Printf("[服务器]:客户端<%v>已连接...\n", conn.RemoteAddr())

	// 循环等待读取报文
	for {
		msg, err := utils.ReadMsg(conn)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("[服务器]:客户端<%v>已离线...\n", conn.RemoteAddr())
			}
			return
		}
		err = HandleMsg(conn, msg)
		if err != nil {
			if err == io.EOF {
				fmt.Printf("[服务器]:客户端<%v>已离线...\n", conn.RemoteAddr())
				return
			}
			fmt.Println("[错误]:处理报文时发生错误.")
			fmt.Println(err)
			continue
		}
	}
}

// 报文转发
func HandleMsg(conn net.Conn, msg model.Message) (err error) {

	switch msg.Type {
	case model.LoginMsgType:
		// 处理登录报文
		err = process.HandleLoginMsg(conn, msg)
	case model.SignUpMsgType:
		// 处理注册报文
		err = process.HandleSginUpMsg(conn, msg)
	default:
		err = errors.New("[错误]:报文类型错误, " + msg.Type + "不存在.")
	}
	return
}
